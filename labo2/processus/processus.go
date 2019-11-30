package processus

import (
	"PRR-Labo2/labo2/mutex"
	"PRR-Labo2/labo2/network"
)

type Processus struct{
	Id uint16
	Net network.Network
	Mut mutex.Mutex
	N uint16
}

func (p *Processus) Init(id uint16, N uint16){
	p.Id = id
	p.N = N
	p.Net = network.Network{}
	p.Mut = mutex.Mutex{}

	//TODO init random stamp
	p.Mut.Init(p.Id,0,N,&p.Net)
	p.Net.Init(p.Id,N,&p.Mut)


}
