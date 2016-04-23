/*
 * This package holds ways to get and setup environment and configs.
 * This file contains function to get the db connect to work with.
 */
package config

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

// Connect tries to connect to mongoDB and returns an actual session.
func Connect() *mgo.Session {
	uri := os.Getenv("MONGO_CONNECTION_URL")

	session, err := mgo.Dial(uri)
	if err != nil {
		log.Panic("Error on db connection! ", err.Error())
	}

	return session
}
