/*
 -----------------------------------------------------------------------------------
 Lab 		 : 02
 File    	 : config.go
 Authors   	 : François Burgener - Tiago P. Quinteiro
 Date        : 03.12.2019

 Goal        : Fichier de config avec l'address réseau ainsi que le port, ainsi que
               les trois type de message (REQ,OK et UPDATE)
 -----------------------------------------------------------------------------------
*/

package config

const (
	ADDR = "127.0.0.1"
	PORT = 3000
	MessageREQ    = "REQ"
	MessageOK     = "OK_"
	MessageUPDATE = "UPD"
)
