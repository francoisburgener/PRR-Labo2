package utils

import (
	"encoding/binary"
	"strconv"
)


/************************************************
*                UTILS METHODE       		    *
*************************************************/

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
func InitMessage(stamp uint32,id uint16, _type []byte) []byte{
	buf := []byte{}

	_id := uint16ToByteArray(id)
	_stamp := uint32ToByteArray(stamp)

	buf = append(buf, _type...)
	buf = append(buf, _stamp...)
	buf = append(buf, _id...)
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

