package utils

import (
	"PRR-Labo2/labo2/config"
	"encoding/binary"
	"strconv"
)


/************************************************
*                UTILS METHODE       		    *
*************************************************/

/**
 * Method to convert an unint32 in byte array
 * @param value we want to update
 * @return the value converted in byte array
 */
func uint32ToByteArray(i uint32) []byte{
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, i)
	return buf
}

/**
 * Method to convert an unint16 in byte array
 * @param value we want to update
 * @return the value converted in byte array
 */
func uint16ToByteArray(i uint16) []byte{
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, i)
	return buf
}

/**
 * Method to init a message(REQ or OK). TYPE + Stamp + id
 * @param message to convert
 * @param message in byte array
 */
func InitMessage(stamp uint32,id uint16, _type []byte) []byte{
	var buf []byte

	_id := uint16ToByteArray(id)
	_stamp := uint32ToByteArray(stamp)

	buf = append(buf, _type...)
	buf = append(buf, _stamp...)
	buf = append(buf, _id...)
	buf = append(buf, '\n')
	return buf
}

/**
 * Method to init an update message (UPD + VALUE)
 * @param buf we want to convert
 * @return update message in byte array
 */
func InitMessageUpdate(value uint32, _type []byte) []byte{
	var buf []byte

	_value := uint32ToByteArray(value)

	buf = append(buf, _type...)
	buf = append(buf, _value...)
	buf = append(buf, '\n')

	return buf
}

/**
 * Method to convert a byte array to unint32
 * @param buf we want to convert
 * @return the value in uint32
 */
func ConverByteArrayToUint32(buf []byte) uint32{
	return binary.LittleEndian.Uint32(buf)
}

/**
 * Method to convert a byte array to unint16
 * @param buf we want to convert
 * @return the value in uint16
 */
func ConverByteArrayToUint16(buf []byte) uint16{
	return binary.LittleEndian.Uint16(buf)
}

/**
 * Method to get the adress:port of the processus by id
 * @param id of the processus
 * @return address:port in string
 */
func AddressByID(id uint16) string{
	port := config.PORT + id
	return config.ADDR + ":" + strconv.Itoa(int(port))
}

