package router

import "github.com/andrepereira/kobayashi-maru/net/controls"

var model string //Model of emulated device

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
