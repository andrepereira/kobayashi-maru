package mongo

import "gopkg.in/mgo.v2"

//Non terminated package to make centralized connection to MongoDB
//Strange bug. "Session already closed [recovered]"
//Non integrated yet to tested software
func Connect(server string, mode bool, db string, collection string) *mgo.Collection {

	session, err := mgo.Dial(server)
	if err != nil {
		panic(err)
	}
	//	defer session.Close() // Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, mode)
	c := session.DB(db).C(collection)

	return c

}
