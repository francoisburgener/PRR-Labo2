package mutex

import (
	"fmt"
	"testing"
	"time"
)

type NetworkMock struct {}

func (n NetworkMock) Req(stamp uint32, id uint16){}
func (n NetworkMock) Ok(stamp uint32, id uint16){}
func (n NetworkMock) Update(value uint){}

func TestMutex_Ask_Increment(t *testing.T) {
	type fields struct {
		private  mutexPrivate
		channels mutexChans
		resource uint
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "Should increment clock when someone stamp", fields: struct {
			private  mutexPrivate
			channels mutexChans
			resource uint
		}{private: mutexPrivate{} , channels: mutexChans{}, resource: 10 } },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mutex{
				private:  tt.fields.private,
				channels: tt.fields.channels,
				resource: tt.fields.resource,
			}
			// Start params
			stamp := uint32(12)
			n := NetworkMock{}

			// Executes
			m.Init(1,stamp, n)
			m.Ask()

			time.Sleep(time.Duration(time.Second * 1))

			// Asserts
			if m.private.stamp != stamp + 1 {
				t.Errorf("The stamps hasn't been incremented")
			}

		})
	}
}

func TestMutex_Ask_STATE(t *testing.T) {
	type fields struct {
		private  mutexPrivate
		channels mutexChans
		resource uint
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "Should change state when asked first time", fields: struct {
			private  mutexPrivate
			channels mutexChans
			resource uint
		}{private: mutexPrivate{} , channels: mutexChans{}, resource: 10 } },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mutex{
				private:  tt.fields.private,
				channels: tt.fields.channels,
				resource: tt.fields.resource,
			}
			// Start params
			stamp := uint32(12)
			n := NetworkMock{}

			// Executes
			m.Init(1,stamp, n)
			m.Req(uint32(2), uint16(2))
			m.Ask()

			time.Sleep(time.Duration(time.Second * 1))

			fmt.Println(m.private.state)

			// Asserts
			if m.private.state != WAITING {
				t.Errorf("State should be waiting. It is now %d", m.private.state)
			}
		})
	}
}