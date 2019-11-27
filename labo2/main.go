package main

import (
	"fmt"
	"time"
)
import "prr-labo2/labo2/mutex"

type NetworkMock struct {}

func (n NetworkMock) Req(stamp uint32, id uint16){}
func (n NetworkMock) Ok(stamp uint32, id uint16){}
func (n NetworkMock) Update(value uint32){}

func main() {
	fmt.Println("Hello words")

	mutex := mutex.Mutex{}
 	n := NetworkMock{}
	mutex.Init(1, 1, n)

	mutex.Ask()

	mutex.Req(3,3)

 	time.Sleep(time.Second * 1)

}
