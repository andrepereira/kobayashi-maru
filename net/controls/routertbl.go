package routertbl

import (
	"log"

	"github.com/andrepereira/kobayashi-maru/config"
	"gopkg.in/mgo.v2/bson"
)

type Route struct {
	To     string //Destiny of this route in routing table
	Result bool   //Result is true if route exist and is enabled
}

//This function insert a new route in the MongoDB
//The argument is a IP to this route
//ToDo: verify if the route exist before the insertion to notmake duplicates
//ToDo: make the connection to database a package outside (SOLID)
func InsertRoute(to string) bool {

	//Legacyi block. Veryfing if is gonna be allright
	//	session, err := mgo.Dial("127.0.0.1")
	//	if err != nil {
	//		panic(err)
	//	}
	//	defer session.Close() // Optional. Switch the session to a monotonic behavior.
	//	session.SetMode(mgo.Monotonic, true)
	//	c := session.DB("test").C("routes")

	c := mongo.Connect("127.0.0.1", true, "test", "routes")

	err := c.Insert(&Route{to, true})

	if err != nil {
		log.Fatal(err)
		return false
	} else {
		return true
	}

}

//This function take a route if it exists in the routing table
//The argument is aIP to this route
//ToDo: make the connection to database a package outside (SOLID)
func GetRoute(to string) bool {

	c := mongo.Connect("127.0.0.1", true, "test", "routes")

	result := Route{}
	err := c.Find(bson.M{"to": to}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	if result.Result == true {

		return true
	} else {
		return false
	}

}
