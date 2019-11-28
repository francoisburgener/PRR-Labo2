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

	//TODO init random stamp
	p.Mut.Init(p.Id,0,&p.Net)
	p.Net.Init(p.Id,N,&p.Mut)


}
