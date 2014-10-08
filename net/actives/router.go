package router

import "github.com/andrepereira/kobayashi-maru/net/controls"

type Id struct {
	IP     string //IP address
	Vendor string //Vendor of emulated hardware
	Model  string //Model of emulated hardware
	OS     string //Operating systemof emulated hardware
	TX     int64  //Actual uplink speed in kbps
	RX     int64  //Actual downlink speed in kbps
	TXPPS  int64  //Actual uplink speed in pps
	RXPPS  int64  //Actual downlink speed in pps
	TXM    int64  //Maximum uplink in kbps
	RXM    int64  //Maximum downlink in kbps
	TXPPSM int64  //Maximum uplink in pps
	RXPPSM int64  //Maximum downlink in pps

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
