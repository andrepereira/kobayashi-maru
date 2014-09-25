package router

import "testing"

func TestRouting(t *testing.T) {

	var pass bool

	pass = Routing("192.168.1.5", "192.168.1.65")

	if pass != true {
		t.Error("Expected true, got ", pass)
	}

}
