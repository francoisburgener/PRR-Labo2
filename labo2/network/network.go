package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

/************************************************
*                   STRUCTURE       		    *
*************************************************/
type Network struct {
	directory map[uint16]net.Conn
}

type Message struct {
	_id    uint16
	_stamp uint32
	_type  []byte
}

/************************************************
*                UTILS METHODE       		    *
*************************************************/
func initMessage(stamp uint32, id uint16, _type []byte) Message{
	return Message{id, stamp, _type}
}

/**
 * Method to convert a Message type to an array of byte
 * @param message to convert
 */
func convertMessageToBytes(msg Message) []byte{
	buf := make([]byte,128)
	buf = append(buf, byte(msg._stamp))
	buf = append(buf, byte(msg._id))
	buf = append(buf, msg._type...)
	return buf
}

/**
 * Method to get the adress:port of the processus by id
 * @param id of the processus
 */
func addressByID(id uint16) string{
	i := 3000 + id
	return "127.0.0.1:" + strconv.Itoa(int(i))
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
	msg := initMessage(stamp,id,[]byte("REQ"))
	buf := convertMessageToBytes(msg);
	n.directory[id].Write(buf)
}

/**
 * Method of Network to send a REQ message
 * @param stamp (logic clock) of the processus
 * @param id of the processus
 */
func (n *Network) OK(stamp uint32, id uint16){
	msg := initMessage(stamp,id,[]byte("OK"))
	buf := convertMessageToBytes(msg);
	n.directory[id].Write(buf)
}


/**
 * Method to init a new Network
 */
func (n *Network) initServ(id uint16,N int){
	addr := addressByID(id)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
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
		fmt.Println("Reading id : " + strconv.Itoa(idConn))
		fmt.Println("Serv Connection between P" + strconv.Itoa(int(id)) + " and P" + strconv.Itoa(idConn))
		n.directory[uint16(idConn)] = conn


		go n.handleConn(conn)
	}
}

func (n *Network)initAllConn(id uint16, N int) {
	for i:=0 ; i < N; i++ {
		if i != int(id) {
			n.initConn(i,id)
		}
	}
}

func (n *Network) initConn(i int,id uint16) {
	addr := addressByID(uint16(i))
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Printf("Connection refused with P%d",i)
	}else{
		n.directory[uint16(i)] = conn
		conn.Write([]byte(strconv.Itoa(int(id))))
		fmt.Println("Writing id : " + strconv.Itoa(int(id)))
		fmt.Println("Dial Connection between P" + strconv.Itoa(int(id)) + " and P" + strconv.Itoa(i))

	}
}

func (n *Network) Init(id uint16,N int) {
	n.directory = make(map[uint16]net.Conn,N)
	n.initAllConn(id,N)
	go n.initServ(id,N)
}

/**
 * Method to ...
 */
func (n *Network) handleConn(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)

	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	//TODO send response to chan
	fmt.Println(buf)
}


func main(){
	//network.Init(uint16(i),N);
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	fmt.Print(" Processus : ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	idConn, err := strconv.Atoi(text)
	if err != nil {
		log.Print(err)
	}
	network := Network{}
	network.Init(uint16(idConn),5)
	fmt.Println(network.directory)
	select {

	}
}