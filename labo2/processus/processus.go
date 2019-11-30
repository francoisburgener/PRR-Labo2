package processus

import (
	"PRR-Labo2/labo2/mutex"
	"PRR-Labo2/labo2/network"
	"math/rand"
)

type Process struct{
	Id uint16
	Net network.Network
	Mut mutex.Mutex
	N uint16
}

func (p *Process) Init(id uint16, N uint16){
	p.Id = id
	p.N = N
	p.Net = network.Network{}
	p.Mut = mutex.Mutex{
		Debug: true,
	}

	const max = 100
	const min = 1

	initStamp := uint32(rand.Intn(max - min + 1) + min)

	p.Mut.Init(p.Id, initStamp, N, &p.Net)
	p.Net.Init(p.Id, N, &p.Mut)
}
