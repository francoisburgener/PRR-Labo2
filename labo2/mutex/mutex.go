package mutex

import (
	"fmt"
)


const (
	REST = false
	WAITING = true
)

const (
	NOT_IN_CS = false
	CRITICAL_SECTION = true
)

type void struct{}

type Mutex struct {
	_N uint16				// Total number of processes
	_me uint16				// The id of the Process
	_stamp uint32			// The logic clock
	_askedReq bool			// True if there is a current request
	_isCS bool				// True if this P is in critical section
	_stampAsk uint32		// Stamp of the submitted request
	_pDiff map[uint16]void	// set of the processes we differed the OK
	_pWait map[uint16]void	// set of the processes we must wait a permission
	askChan chan bool
	waitChan chan bool
}

func (m *Mutex) Init(id uint16, n uint16) {
	m.askChan = make(chan bool)
	m.waitChan = make(chan bool)
	m._stamp = 0
	m._me = id
	m._N = n
	m._askedReq = REST

	go func() {
		for {
			select {
				case m._askedReq = <- m.askChan:
					fmt.Println("Someone asked")
				case m.waitChan <- true:
					fmt.Println("ohy")
			}
		}
	}()
}

func (m *Mutex) Ask() {
	m.askChan <- true
}

func (m *Mutex) Wait() {
	<- m.waitChan
	fmt.Println("I waited")
}

func (m *Mutex) End() {
	m.askChan <- false
}