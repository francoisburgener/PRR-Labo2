package main

import (
	"PRR-Labo2/labo2/processus"
	"bufio"
	"flag"
	"fmt"
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
		fmt.Println("1 - Read critical value")
		fmt.Println("2 - Update critical value")
		fmt.Println("3 - Quit")
		fmt.Print("> ")

		scanner.Scan()
		choice :=  scanner.Text()

		switch choice {
		case "1":
			a := p.Mut.GetResource()
			fmt.Printf("The resource is %d \n", a)
		case "2":
			p.Mut.Ask()
			fmt.Println("Process is asking for the resource")
			p.Mut.Wait()
			fmt.Println("Other Processes gave us permission")
			fmt.Print("You can now enter your value:  ")

			scanner.Scan()
			a :=  scanner.Text()

			b, err := strconv.Atoi(a)

			if err != nil {
				fmt.Println("Enter a number please")
				continue
			}

			fmt.Println(" value is", b)

			p.Mut.Update(uint32(b))
			p.Mut.End()
		case "3":
			os.Exit(0)
		default:
			fmt.Println("Choose 1 or 2")
		}

		test++
	}

}
