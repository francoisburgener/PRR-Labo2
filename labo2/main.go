package main

import (
	"PRR-Labo2/labo2/processus"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main(){
	id,N := argValue()
	p := processus.Processus{}
	p.Init(uint16(id),N)
	fmt.Println("\nInitialisation done\n")
	console(&p, uint16(id))
}

func argValue() (uint16, int) {
	var proc string
	var procN string
	flag.StringVar(&proc, "proc", "", "Usage")
	flag.StringVar(&procN, "N", "", "Usage")
	flag.Parse()
	id,err :=strconv.Atoi(proc)
	N,err :=strconv.Atoi(procN)
	if err != nil {
		log.Print("Veuillez mettre un chiffre")
	}

	return uint16(id),N

}

func console(p *processus.Processus,id uint16) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Choix (nombre)")
	fmt.Println("---------------------")
	for{
		fmt.Println("1 - Lire la valeur critique")
		fmt.Println("2 - Modifier la valeur critique")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		switch text {
		case "1":
			//TODO
		case "2":
			//TODO
		default:
			fmt.Println("Choose 1 or 2")
		}
	}

}
