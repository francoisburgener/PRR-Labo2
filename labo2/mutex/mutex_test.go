package mutex

import "testing"

func TestMutex_ask(t *testing.T) {
	type fields struct {
		asked bool
	}
	tests := []struct {
		name string
		fields fields
	}{
		{"Should know when someone asked", fields{false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mutex{false}

			m.ask()

			if m.asked != true {
				t.Errorf("Asked is %v, want %v", m.asked, true)
			}
		})
	}
}

func TestMutex_end(t *testing.T) {
	type fields struct {
		asked bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Should know when someone released", fields{true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mutex{
				asked: tt.fields.asked,
			}

			m.end()

			if m.asked != false {
				t.Errorf("Asked is %v, want %v", m.asked, false)
			}
		})
	}
}