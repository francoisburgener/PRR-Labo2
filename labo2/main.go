package main

import (
	"PRR-Labo2/labo2/processus"
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
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

	log.Println("Client: Choice (number)")
	log.Println("Client: ---------------------")

	scanner := bufio.NewScanner(os.Stdin)

	test := uint16(1)

	for{
		log.Println("Client: 1 - Read critical value")
		log.Println("Client: 2 - Update critical value")
		log.Println("Client: 3 - Quit")
		log.Print("Client: > ")

		scanner.Scan()
		choice :=  scanner.Text()

		switch choice {
		case "1":
			a := p.Mut.GetResource()
			log.Printf("Client: The resource is %d \n", a)
		case "2":
			p.Mut.Ask()
			log.Println("Client: Process is asking for the resource")
			p.Mut.Wait()
			log.Println("Client: Other Processes gave us permission")
			log.Println("Client: You can now enter your value")

			scanner.Scan()
			a :=  scanner.Text()

			b, err := strconv.Atoi(a)

			if err != nil {
				log.Println("Enter a number please")
				continue
			}

			p.Mut.Update(uint(b))
			p.Mut.End()
		case "3":
			os.Exit(0)
		default:
			log.Println("Client: Choose 1 or 2")
		}

		test++
	}

}
