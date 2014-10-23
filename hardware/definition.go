package definition

import "github.com/andrepereira/kobayashi-maru/net/actives"

type id struct {
	IP            string        //IP address
	Router        string        //IP of router to connect this hardware
	Vendor        string        //Vendor of emulated hardware
	Model         string        //Model of emulated hardware
	OS            string        //Operating systemof emulated hardware
	TX            int64         //Actual uplink speed in kbps
	RX            int64         //Actual downlink speed in kbps
	TXPPS         int64         //Actual uplink speed in pps
	RXPPS         int64         //Actual downlink speed in pps
	TXM           int64         //Maximum uplink in kbps
	RXM           int64         //Maximum downlink in kbps
	TXPPSM        int64         //Maximum uplink in pps
	RXPPSM        int64         //Maximum downlink in pps
	RootPassword  string        //Password of administrative user
	User          string        //Username of a non privileged user at active
	UserPassword  string        //Password of a non privileged user at system
	Powered       bool          //Hardware ON = true or OFF = false
	ProcessRuning [20][3]string //Table 20x3 of runing process with process name, your version and the network listener port(if used)

}

//The function define the set of details permanents and variables to a instance of hardware
func SetHardware(ip string, router string, vendor string, model string, os string, tx int64, rx int64, txpps int64, rxpps int64, txm int64, rxm int64, txppsm int64, rxppsm int64, rootPassword string, user string, userPassword string, powered bool, processRuning [20][3]string) (id, bool) {

	var hw id
	hw.IP = ip
	hw.Router = router
	hw.Vendor = vendor
	hw.Model = model
	hw.OS = os
	hw.TX = tx
	hw.RX = rx
	hw.TXPPS = txpps
	hw.RXPPS = rxpps
	hw.TXM = txm
	hw.RXM = rxm
	hw.TXPPSM = txppsm
	hw.RXPPSM = rxppsm
	hw.RootPassword = rootPassword
	hw.User = user
	hw.UserPassword = userPassword
	hw.Powered = powered
	hw.ProcessRuning = processRuning

	return hw, true

}

//Make acheckto verify if exists a routeto a ping be delivered
//ToDo: Verify if the destinyis powered and if it make a pong back
func Ping(ip string) bool {

	sucess := router.Routing(ip)

	return sucess

}
