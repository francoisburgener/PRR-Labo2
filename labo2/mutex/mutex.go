package mutex

type Mutex struct {
	asked bool
}

func (m *Mutex) ask() {
	m.asked = true
}

func (m *Mutex) wait() {
}

func (m *Mutex) end() {
	m.asked = false
}