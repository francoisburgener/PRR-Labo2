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
	p := processus.Process{}
	p.Init(id,N)
	fmt.Println("\nInitialisation done\n")
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
		log.Fatal("Please put a number !")
	}

	return uint16(id),uint16(N)

}

func console(p *processus.Process) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Choice (number)")
	fmt.Println("---------------------")
	for{
		fmt.Println("1 - Read critical value")
		fmt.Println("2 - Update critical value")
		fmt.Println("3 - Quit")
		fmt.Print("> ")


		choice, _ := reader.ReadString('\n')
		choice = strings.Replace(choice, "\n", "", -1)

		switch choice {
		case "1":
			a := p.Mut.GetResource()
			fmt.Printf("Client: The resource is %d \n", a)
		case "2":
			p.Mut.Ask()
			fmt.Println("Client: Process is asking for the resource")
			p.Mut.Wait()
			fmt.Println("Client: Other Processes gave us permission")
			p.Mut.Update(42)
			p.Mut.End()
		case "3":
			os.Exit(0)
		default:
			fmt.Println("Choose 1 or 2")
		}
	}

}
