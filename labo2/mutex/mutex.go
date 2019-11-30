package mutex

import (
	"log"
)

/**
 * ENUM declaration of the states
 */
const (
	REST = iota
	WAITING
	CRITICAL
)

/**
 * Interface wanted for the Network
 */
type Network interface {
	REQ(stamp uint32, id uint16)
	OK(stamp uint32, id uint16)
	UPDATE(value uint)
}

/**
 * Passing stamp and id through channels
 */
type Message struct {
	stamp uint32
	id uint16
}

/**
 * Hides the values used by the mutex to handle his internal state
 */
type mutexPrivate struct {
	N uint16				// Total number of processes
	me uint16				// The id of the Process
	stamp uint32			// The logic clock
	state uint8			// Rest, Waiting or CS
	stampAsk uint32		// Stamp of the submitted request
	pDiff map[uint16]bool	// set of the processes we differed the OK
	pWait map[uint16]bool	// set of the processes we must wait a permission
	netWorker Network
}

/**
 * Hides the communication channels used by the mutex
 */
type mutexChans struct {
	reqChan chan Message
	okChan chan Message
	endChan chan bool
	updateChan chan uint
	askChan chan bool
	waitChan chan bool
}

/**
 * This is the class you may want to export
 */
type Mutex struct {
	private  mutexPrivate
	channels mutexChans
	resource uint
}

/**
 * Constructor
 * This method is responsible to initialize everything in order.
 * ALWAYS CALL IT BEFORE DOING ANYTHING ELSE
 */
func (m *Mutex) Init(id uint16, initialStamp uint32, numberOfProcess uint16, netWorker Network) {

	m.private = mutexPrivate{
		N:         numberOfProcess,
		me:        id,
		stamp:     initialStamp,
		state:     REST,
		stampAsk:  0,
		pDiff:     make(map[uint16]bool),
		pWait:     make(map[uint16]bool),
		netWorker: netWorker,
	}

	m.channels = mutexChans{
		reqChan:    make(chan Message),
		okChan:     make(chan Message),
		endChan:    make(chan bool),
		updateChan: make(chan uint),
		askChan:     make(chan bool),
		waitChan:    make(chan bool),
	}

	// We start with some tokens already
	m.initpWait()

	// Here the manager starts
	go m.manager()
}

/**
 * This function runs in a goroutine
 * It is the main handler of the mutex
 * Every method called passes through here
 */
func (m *Mutex) manager() {
	for {
		select {
			// ASK: Client called Ask()
			case <- m.channels.askChan:
				log.Printf("Mutex: Client asked me the CS")

				if m.private.state == REST {
					m.incrementClock(0)
					m.private.state = WAITING
					m.private.stampAsk = m.private.stamp
					m.reqAll() // Sending req to the Ps to ask token
				}

			// END: Client released SC
			case <- m.channels.endChan:
				log.Printf("Mutex: Client released the CS")
				m.incrementClock(0)
				m.private.state = REST		// Leaving SC
				m.private.netWorker.UPDATE(m.resource)
				m.okAll() // Sending ok to the differed Ps


			// P asked a token
			case message := <- m.channels.reqChan:
				log.Printf("Mutex: Req received from %d", message.id)
				m.incrementClock(message.stamp) // Increment, max between mine and P

				if m.private.state == CRITICAL ||
					m.private.state == WAITING &&
					m.private.stampAsk < message.stamp ||
					(m.private.stampAsk == message.stamp && m.private.me < message.id) {

					m.private.pDiff[message.id] = true // We have to differ the obtain from other P
				} else {

					m.private.pWait[message.id] = true // Adding to waiting set
					m.private.netWorker.OK(m.private.stamp, message.id) //Sending the signal
				}

			// P sent Ok
			case message:= <- m.channels.okChan:
				log.Printf("Mutex: Ok received from %d", message.id)
				m.incrementClock(message.stamp) // Increment, max between mine and P
				delete(m.private.pWait, message.id) // removing wait from here

			// Network told us to update
			case val := <- m.channels.updateChan:
				log.Printf("Mutex: Network told us to update")
				if m.private.state != CRITICAL {
					m.resource = val
				}

			default:
				// If we need to enter CS and don't wait on anyone
				if m.private.state == WAITING && len(m.private.pWait) == 0 {
					m.private.state = CRITICAL
					m.channels.waitChan <- true // we release our client
				}
		}
	}
}

// CLIENT SIDE METHODS --------------------------------

/**
 * Call this to ask the network for a future usage of the SC
 */
func (m *Mutex) Ask() {
	m.channels.askChan <- true
}

/**
 * Block until the SC is ready
 */
func (m *Mutex) Wait() {
	<- m.channels.waitChan
}

/**
 * Release the SC
 */
func (m *Mutex) End() {
	m.channels.endChan <- true
}

// SEVER SIDE METHOD ---------------------------------

/**
 * Pass an incoming network REQ here
 */
func (m *Mutex) Req(stamp uint32, id uint16) {
	message := Message{
		stamp: stamp,
		id:    id,
	}
	m.channels.reqChan <- message
}

/**
 * Pass an incoming network OK here
 */
func (m *Mutex) Ok(stamp uint32, id uint16) {
	message := Message{
		stamp: stamp,
		id:    id,
	}
	m.channels.okChan <- message
}

/**
 * SETTER: call this if you want to change the SC val
 * Never call it without being in SC (ask, wait, update, end)
 */
func (m *Mutex) Update(value uint) {
	m.channels.updateChan <- value
}

// PRIVATE utils methods ---------------------------------

/**
 * Sends ok to all differed P in network
 */
func (m *Mutex) okAll() {
	for key, _ := range m.private.pDiff {
		m.private.netWorker.OK(m.private.stamp, key)
	}

	// Clean the structure
	m.private.pDiff =  make(map[uint16]bool)
}

/**
 * Sends req to all P in network you're waiting
 */
func (m *Mutex) reqAll() {
	for key, _ := range m.private.pWait  {
		log.Printf("Sending req to %d\n", key)
		m.private.netWorker.REQ(m.private.stamp, key)
	}
}

/**
 * Takes max and increments the stamp
 * value uint32 -  the value of the other stamp
 */
func (m *Mutex) incrementClock(value uint32){
	if value > m.private.stamp {
		m.private.stamp = value
	}

	m.private.stamp += 1
}

/**
 * Initialize the tokens this P has over the others
 */
func (m *Mutex) initpWait(){
	for i := m.private.me + 1; i < m.private.N; i++ {
		m.private.pWait[i] = true
	}
}