package routertbl

import "testing"

//Unit test to verify if the routes are been inserted correctly in
//the routing table
func TestInsertRoute(t *testing.T) {

	var pass bool

	pass = InsertRoute("192.168.1.67")

	if pass != true {
		t.Error("Expected true, got ", pass)
	}

}

//Unit test to verify if the funtion that get a route in the
//routing table are been running correctly
func TestGetRoute(t *testing.T) {

	var pass bool

	pass = GetRoute("192.168.1.67")

	if pass != true {
		t.Error("Expected true, got ", pass)
	}

}
