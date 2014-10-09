package definition

import "testing"

//Unit test to verify if the hardware was created or altered with sucess
func TestSetHardware(t *testing.T) {

	var pass bool
	processRuning := [20][2]string{{"apache2", "2.2.14"}, {"mysqld", "5.5.6"}}

	_, pass = SetHardware("192.168.1.67", "CISCO", "X300", "LINUX", 1200, 1650, 550, 320, 10000, 10000, 1000, 1000, "janaina", "andre", "janaina2", processRuning)

	if pass != true {
		t.Error("Expected true, got ", pass)
	}

}
