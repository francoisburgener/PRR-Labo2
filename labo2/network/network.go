package network

import (
	"PRR-Labo2/labo2/utils"
	"fmt"
	"log"
	"net"
	"strconv"
)

/************************************************
*                   INTERFACE       		    *
*************************************************/
type Mutex interface {
	REQ(stamp uint32, id uint16)
	OK(stamp uint32, id uint16)
	UPDATE(value uint)
}

/************************************************
*                   STRUCTURE       		    *
*************************************************/
type Network struct {
	id uint16
	nProc int
	directory map[uint16]net.Conn
	Done chan string
	mutex Mutex
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
	msg := utils.InitMessage(stamp,id,[]byte("REQ"))
	buf := utils.ConvertMessageToBytes(msg);
	n.directory[id].Write(buf)
}

/**
 * Method of Network to send a REQ message
 * @param stamp (logic clock) of the processus
 * @param id of the processus
 */
func (n *Network) OK(stamp uint32, id uint16){
	msg := utils.InitMessage(stamp,id,[]byte("OK_"))
	buf := utils.ConvertMessageToBytes(msg);
	n.directory[id].Write(buf)
}

/**
 * Method of Network to send a UPDATE message
 * @param value to update
 * @param id of the processus
 */
func (n *Network) UPDATE(value uint){
	for i:=0; i < len(n.directory) + 1; i++{
		if i != int(n.id){
			n.directory[uint16(i)].Write([]byte("UDP" + strconv.Itoa(int(value))))
		}
	}
}


/**
 * Method to init a new Network
 */
func (n *Network) initServ(id uint16){
	addr := utils.AddressByID(id)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	for {

		if len(n.directory) == n.nProc-1{
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
		fmt.Println("Serv Connection between P" + strconv.Itoa(int(id)) + " and P" + strconv.Itoa(idConn))
		n.directory[uint16(idConn)] = conn

		go n.handleConn(conn)
	}
}

/**
 * Method to init all dial connection
 * @param n reference of network
 * @param id of the processus to connect
 * @param N number of processus
 */
func (n *Network) initAllConn(id uint16) {
	for i:=0 ; i < n.nProc; i++ {
		if i != int(id) {
			n.initConn(i,id)
		}
	}
}

/**
 * Method to init a dial connection
 * @param n reference of network
 * @param i id of the processus we want to connect
 * @param id of our processus
 */
func (n *Network)initConn(i int,id uint16) {
	addr := utils.AddressByID(uint16(i))
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Printf("Connection refused with P%d",i)
	}else{
		n.directory[uint16(i)] = conn
		conn.Write([]byte(strconv.Itoa(int(id))))
		fmt.Println("Dial Connection between P" + strconv.Itoa(int(id)) + " and P" + strconv.Itoa(i))
	}
}

/**
 * Method to init the server and get all connection between processus
 * @param id of the processus
 * @param N number of processus
 */
func (n *Network) Init(id uint16,N int, mutex Mutex) {
	n.directory = make(map[uint16]net.Conn,N)
	n.Done = make(chan string)
	n.mutex = mutex;
	n.id = id
	n.nProc = N

	go func() {
		n.initAllConn(id)
		n.initServ(id)
	}()

	<- n.Done

}

/**
 * Method to ...
 */
func (n *Network)handleConn(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 32)

	// Read the incoming connection into the buffer.
	l, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	n.decodeMessage(buf,l)
}

func (n *Network) decodeMessage(bytes []byte,l int) {

	_type := string(bytes[0:3])
	var stamp uint32
	var id uint16
	var value uint

	if(_type == "UDP"){
		tmp, err := strconv.Atoi(string(bytes[3:l]))
		if err != nil{
			log.Fatal(err)
		}
		value = uint(tmp)
	}else{
		stamp = utils.ConverByteArrayToUint32(bytes[3:7])
		id = utils.ConverByteArrayToUint16(bytes[7:l])
	}

	switch _type {
	case "REQ":
		n.mutex.REQ(stamp,id)
	case "OK_":
		n.mutex.OK(stamp,id)
	case "UPD":
		n.mutex.UPDATE(value)
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