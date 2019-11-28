package main

import (
	"PRR-Labo2/labo2/network"
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main(){
	/*var proc string
	var procN string
	flag.StringVar(&proc, "proc", "", "Usage")
	flag.StringVar(&procN, "N", "", "Usage")
	flag.Parse()
	id,err :=strconv.Atoi(proc)
	N,err :=strconv.Atoi(procN)
	if err != nil {
		log.Print("Veuillez mettre un chiffre")
	}
	n := network.Network{}
	n.Init(uint16(id),N)

	fmt.Println("\nInitialisation done\n")

	console(&n, uint16(id))*/
}

func console(n *network.Network,id uint16) {
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
			n.REQ(0,1)
		case "2":
			n.OK(0,id)
		default:
			fmt.Println("Choose 1 or 2")
		}
	}

}
