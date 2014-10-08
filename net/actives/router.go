package router

import "github.com/andrepereira/kobayashi-maru/net/controls"

type Id struct {
	IP     string //IP address
	Vendor string //Vendor of emulated hardware
	Model  string //Model of emulated hardware
	OS     string //Operating systemof emulated hardware
}

//Authorize the route to a destiny "to"
//This make the function of a router device
//Uses a to string with a IP address
//Return a boolean with authorization or not to routing (if the route
//exists at routing table
func Routing(to string) bool {

	if routertbl.GetRoute(to) {

		return true
	} else {
		return false
	}

}
