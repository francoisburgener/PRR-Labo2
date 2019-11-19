package main

import "fmt"
import "prr-labo2/labo2/mutex"

func main() {
	fmt.Println("Hello words")

	mutex := mutex.Mutex{}

	mutex.Init()

	mutex.Ask()

	mutex.Wait()

}
