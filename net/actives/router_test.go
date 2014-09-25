package router

import "testing"

func TestRouting(t *testing.T) {

	var pass bool

	pass = Routing(1, 2)

	if pass != true {
		t.Error("Expected true, got ", pass)
	}

}
