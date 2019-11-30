package main

import (
	"PRR-Labo2/labo2/processus"
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

func main(){
	id,N := argValue()
	p := processus.Process{}
	p.Init(id,N)
	log.Println("\nClient: Initialisation done\n")
	console(&p)
}

func argValue() (uint16, uint16) {
	var proc string
	var procN string
	flag.StringVar(&proc, "proc", "", "Usage")
	flag.StringVar(&procN, "N", "", "Usage")
	flag.Parse()
	id,err :=strconv.Atoi(proc)
	N,err :=strconv.Atoi(procN)
	if err != nil {
		log.Fatal("Client: Please put a number !")
	}

	return uint16(id),uint16(N)

}

func console(p *processus.Process) {
	reader := bufio.NewReader(os.Stdin)
	log.Println("Client: Choice (number)")
	log.Println("Client: ---------------------")
	for{
		log.Println("Client: 1 - Read critical value")
		log.Println("Client: 2 - Update critical value")
		log.Println("Client: 3 - Quit")
		log.Print("Client: > ")


		choice, _ := reader.ReadString('\n')
		choice = strings.Replace(choice, "\n", "", -1)

		switch choice {
		case "1":
			a := p.Mut.GetResource()
			log.Printf("Client: The resource is %d \n", a)
		case "2":
			p.Mut.Ask()
			log.Println("Client: Process is asking for the resource")
			p.Mut.Wait()
			log.Println("Client: Other Processes gave us permission")
			p.Mut.Update(42)
			p.Mut.End()
		case "3":
			os.Exit(0)
		default:
			log.Println("Client: Choose 1 or 2")
		}
	}

}
