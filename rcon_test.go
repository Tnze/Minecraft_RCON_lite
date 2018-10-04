package mcrcon

import "testing"

func TestRCON(t *testing.T) {
	var r RCON
	err := r.Connect("localhost:25575", "123456")
	if err != nil {
		t.Error(err)
	}
	r.Do("stop")
}
