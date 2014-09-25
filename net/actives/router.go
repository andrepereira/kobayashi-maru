package router

//Authorize the route between a from and a to
//This make the function of a router device
//Uses a from string with a IP address and a to string with a IP address
//Return a boolean with authorization ornot to routing (if the route
//exists at routing table
func Routing(from string, to string) bool {

	routingTable := make(map[string]bool)
	routingTable["192.168.1.65"] = true

	if from != "" && routingTable[to] == true {

		return true
	} else {
		return false
	}

}
