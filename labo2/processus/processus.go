package processus

import (
	"PRR-Labo2/labo2/mutex"
	"PRR-Labo2/labo2/network"
	"math/rand"
)

const (
	stampMax = 100
	stampMin = 0
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
	p.Net = network.Network{
		Debug: false,
	}
	p.Mut = mutex.Mutex{
		Debug: false,
	}

	// Ensures everyone has a different seed
	rand.Seed(int64(id + N))

	initStamp := uint32(rand.Intn(stampMax - stampMin + 1) + stampMin)

	p.Mut.Init(p.Id, initStamp, N, &p.Net)
	p.Net.Init(p.Id, N, &p.Mut)
}
