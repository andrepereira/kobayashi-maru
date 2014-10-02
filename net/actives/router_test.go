package router

import "testing"

//Unit test to verify the comportament of Routing function
//Give a IP and receive true if this value exists at route table
//Take a false value if the IP don't exist  at route table

func TestRouting(t *testing.T) {

	var pass bool

	pass = Routing("192.168.1.67")

	if pass != true {
		t.Error("Expected true, got ", pass)
	}

}
