package network

import (
	"PRR-Labo2/labo2/utils"
	"log"
	"net"
	"strconv"
)

/************************************************
*                   INTERFACE       		    *
*************************************************/
type Mutex interface {
	Req(stamp uint32, id uint16)
	Ok(stamp uint32, id uint16)
	Update(value uint)
}

/************************************************
*                   STRUCTURE       		    *
*************************************************/
type Network struct {
	id uint16 //id of our processus
	nProc uint16 // Number of processus
	directory map[uint16]net.Conn // map of connection
	Done chan string // channel to say if the server initialisation is done
	mutex Mutex	//Ref of our mutex
	Debug bool
}

/************************************************
*                NETWORK METHOD        		    *
*************************************************/

/**
 * Method of Network to send a REQ message
 * @param stamp (logic clock) of the processus
 * @param id of the processus
 */
func (n *Network) REQ(stamp uint32, id uint16){
	msg := utils.InitMessage(stamp,n.id,[]byte("REQ"))
	buf := utils.ConvertMessageToBytes(msg)
	_, err := n.directory[id].Write(buf)

	if err != nil{
		log.Fatal(err)
	}
	if n.Debug{
		log.Printf("Network: Send message type:%s stamp:%d id:%d \n",msg.Type,msg.Stamp,msg.Id)
	}

}

/**
 * Method of Network to send a REQ message
 * @param stamp (logic clock) of the processus
 * @param id of the processus
 */
func (n *Network) OK(stamp uint32, id uint16){
	msg := utils.InitMessage(stamp,n.id,[]byte("OK_"))
	buf := utils.ConvertMessageToBytes(msg);
	_, err := n.directory[id].Write(buf)

	if err != nil{
		log.Fatal(err)
	}

	if n.Debug{
		log.Printf("Network: Send message type:%s stamp:%d id:%d \n",msg.Type,msg.Stamp,msg.Id)
	}
}

/**
 * Method of Network to send a UPDATE message
 * @param value to update
 * @param id of the processus
 */
func (n *Network) UPDATE(value uint){
	for i:=0; i < len(n.directory) + 1; i++{
		if i != int(n.id){
			_, err := n.directory[uint16(i)].Write([]byte("UPD" + strconv.Itoa(int(value))))
			if err != nil{
				log.Fatal(err)
			}

			if n.Debug{
				log.Printf("Network: Send message Update P%d value: %d",i,value)
			}
		}
	}
}


/**
 * Method to init the server and get all connection between processus
 * @param id of the processus
 * @param N number of processus
 */
func (n *Network) Init(id uint16,N uint16, mutex Mutex,debug bool) {
	log.Printf("Network: Initialisation ")
	n.directory = make(map[uint16]net.Conn,N)
	n.Done = make(chan string)
	n.mutex = mutex
	n.id = id
	n.nProc = N
	n.Debug = debug

	go func() {
		n.initAllConn()
		n.initServ()
	}()

	<- n.Done
}

// PRIVATE methods ---------------------------------

/**
 * Method to init all dial connection
 * @param n reference of network
 * @param id of the processus to connect
 * @param N number of processus
 */
func (n *Network) initAllConn() {
	for i:=uint16(0) ; i < n.nProc; i++ {
		if i != uint16(n.id) {
			n.initConn(i)
		}
	}
}

/**
 * Method to init a dial connection
 * @param n reference of network
 * @param i id of the processus we want to connect
 * @param id of our processus
 */
func (n *Network)initConn(i uint16) {
	addr := utils.AddressByID(uint16(i))
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		log.Printf("Network error : Connection refused with P%d",i)
	}else{
		n.directory[uint16(i)] = conn
		_, err := conn.Write([]byte(strconv.Itoa(int(n.id))))
		if err != nil{
			log.Fatal(err)
		}

		if n.Debug{
			log.Printf("Network : Dial Connection between P%d and P%d\n", n.id, i)
		}

		go n.handleConn(conn)
	}
}


/**
 * Method to init a new Network
 */
func (n *Network) initServ(){
	addr := utils.AddressByID(n.id)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	for {

		if len(n.directory) == int(n.nProc-1) {
			n.Done <- "done"
		}

		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}


		tmp := make([]byte,128)
		l, err := conn.Read(tmp)
		if err != nil {
			log.Print(err)
		}
		str := string(tmp[0:l])
		idConn, err := strconv.Atoi(str)
		log.Println("Network: Serv Connection between P" + strconv.Itoa(int(n.id)) + " and P" + strconv.Itoa(idConn))
		n.directory[uint16(idConn)] = conn

		go n.handleConn(conn)
	}
}

/**
 * Method to ...
 */
func (n *Network)handleConn(conn net.Conn) {
	for {
		// Make a buffer to hold incoming data.
		buf := make([]byte, 32)

		// Read the incoming connection into the buffer.
		l, err := conn.Read(buf)
		if err != nil {
			log.Printf("Network error: Error reading:", err.Error())
		}
		n.decodeMessage(buf,l)
	}
}

func (n *Network) decodeMessage(bytes []byte,l int) {

	_type := string(bytes[0:3])
	var stamp uint32
	var id uint16
	var value uint

	if _type == "UPD"{
		tmp, err := strconv.Atoi(string(bytes[3:l]))
		if err != nil{
			log.Fatal(err)
		}
		value = uint(tmp)

		if n.Debug{
			log.Printf("Network: Decoded message type:%s value:%d",_type,value)
		}

	}else if _type == "OK_" || _type == "REQ"{
		stamp = utils.ConverByteArrayToUint32(bytes[3:7])
		id = utils.ConverByteArrayToUint16(bytes[7:l])

		if n.Debug{
			log.Printf("Network: Decoded message type:%s stamp:%d id:%d",_type,stamp,id)
		}

	}


	switch _type {
	case "REQ":
		n.mutex.Req(stamp,id)
	case "OK_":
		n.mutex.Ok(stamp,id)
	case "UPD":
		n.mutex.Update(value)
	default:
		log.Println("Network: Incorrect type message !")
	}
}



/*func main(){
	n := Network{}
	n.directory = make(map[uint16]net.Conn,2)
	n.Done = make(chan string)
	n.nProc = 2

	go n.initServ(2)
	n.initConn(2,1)
	n.REQ(50000,2)
	//n.UPDATE(42)
	select {

	}
}*/