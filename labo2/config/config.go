/*
 -----------------------------------------------------------------------------------
 Lab 		 : 02
 File    	 : config.go
 Authors   	 : François Burgener - Tiago P. Quinteiro
 Date        : 03.12.2019

 Goal        : Config file for the network layer
 -----------------------------------------------------------------------------------
*/

package config

const (
	ADDR = "127.0.0.1"
	PORT = 3000
	MessageREQ    = "REQ"
	MessageOK     = "OK_"
	MessageUPDATE = "UPD"
	Debug_Mutex = false
	Debug_Network = false
)
