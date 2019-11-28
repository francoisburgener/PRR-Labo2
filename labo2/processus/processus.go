package processus

import (
	"PRR-Labo2/labo2/mutex"
	"PRR-Labo2/labo2/network"
)

type Processus struct{
	Id uint16
	Net network.Network
	Mut mutex.Mutex
	N int
}

func (p *Processus) Init(id uint16, N int){
	p.Id = id
	p.N = N
	p.Net = network.Network{}
	p.Mut = mutex.Mutex{}

	p.Net.Init(id,N,&p.Mut)

	//TODO init mutex

}
