package utils

import (
	"PRR-Labo2/labo2/message"
	"encoding/binary"
	"strconv"
)


/************************************************
*                UTILS METHODE       		    *
*************************************************/
func InitMessage(stamp uint32, id uint16, _type []byte) message.Message{
	return message.Message{id, stamp, _type}
}

func uint32ToByteArray(i uint32) []byte{
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, i)
	return buf
}

func uint16ToByteArray(i uint16) []byte{
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, i)
	return buf
}

/**
 * Method to convert a Message type to an array of byte
 * @param message to convert
 */
func ConvertMessageToBytes(msg message.Message) []byte{
	buf := []byte{}

	id := uint16ToByteArray(msg.Id)
	stamp := uint32ToByteArray(msg.Stamp)

	buf = append(buf, stamp...)
	buf = append(buf, id...)
	buf = append(buf, msg.Type...)
	return buf
}

func ConverByteArrayToUint32(buf []byte) uint32{
	return binary.LittleEndian.Uint32(buf)
}

func ConverByteArrayToUint16(buf []byte) uint16{
	return binary.LittleEndian.Uint16(buf)
}

/**
 * Method to get the adress:port of the processus by id
 * @param id of the processus
 */
func AddressByID(id uint16) string{
	i := 3000 + id
	return "127.0.0.1:" + strconv.Itoa(int(i))
}

