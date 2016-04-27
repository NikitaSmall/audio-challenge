/*Package config holds ways to get and setup environment and configs.
 * This file contains function to get the db connect to work with.
 */
package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/mgo.v2"
)

// DB is main dial instance to work with database
var DB *mgo.Session

// Connect tries to connect to mongoDB and returns an actual session.
func Connect() {
	uri := os.Getenv("MONGO_CONNECTION_URL")

	session, err := mgo.Dial(uri)
	session.SetSocketTimeout(1 * time.Hour)
	if err != nil {
		log.Panic("Error on db connection! ", err.Error())
	}

	DB = session
}
