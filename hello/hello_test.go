package hello

import "testing"

func Test_hello(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{name: "hello Adam", arg: "Adam", want: "hello Adam"},
		{name: "hello empty", arg: "", want: "hello"},
		{name: "hello 123", arg: "123", want: "hello 123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hello(tt.arg); got != tt.want {
				t.Errorf("hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
