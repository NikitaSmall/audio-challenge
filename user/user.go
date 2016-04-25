/*Package user holds main auth logic.
 * Here users are created, checked in db and so on.
 */
package user

import (
	"crypto/md5"
	"errors"
	"log"
	"os"

	"github.com/nikitasmall/audio-challenge/config"
	"gopkg.in/mgo.v2/bson"
)

// configs for connection
var collectionName = "users"

// Authenticator is an interface that determinates behavior of users
type Authenticator interface {
	Register() error
	CheckUser() error

	hasSameUsername() bool
	encryptPassword()
}

// User is basic user struct to hold some info
type User struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

// CreateUser returns user with generated id
func CreateUser() *User {
	return &User{
		ID: bson.NewObjectId().Hex(),
	}
}

// Register makes attempt to register user.
// In case of success returns nil,
// in other returns error with reason
func (user *User) Register() error {
	if user.hasSameUsername() {
		return errors.New("This username already taken. Choose another one.")
	}

	user.encryptPassword()
	usersCollection := config.DB.DB(os.Getenv("MONGO_DB_NAME")).C(collectionName)

	err := usersCollection.Insert(user)
	if err != nil {
		log.Print("Error on user inserting. ", err.Error())
		return err
	}

	return nil
}

// CheckUser check the user's existence
// (for login or similar actions),
// returns nil if the user exists,
// otherwise returns error with description
func (user *User) CheckUser() error {
	usersCollection := config.DB.DB(os.Getenv("MONGO_DB_NAME")).C(collectionName)
	user.encryptPassword()

	u := CreateUser()
	err := usersCollection.Find(bson.M{"username": user.Username, "password": user.Password}).One(u)

	if err != nil {
		if err.Error() == "not found" {
			return errors.New("Cannot find user with such username or password!")
		}

		// if error other than "not found" we log this
		log.Print("Cannot establish uniquiness of username. ", err.Error())
		return err
	}

	return nil
}

// function returns true in case of user with same name existence,
// otherwise (if username is new) return false
func (user *User) hasSameUsername() bool {
	usersCollection := config.DB.DB(os.Getenv("MONGO_DB_NAME")).C(collectionName)
	var users []User

	err := usersCollection.Find(bson.M{"username": user.Username}).All(&users)
	if err != nil {
		log.Panic("Cannot set uniquiness of username. ", err.Error())
	}

	return len(users) > 0
}

// function encrypt the password of user.
// 'true' password is not stored, only hash result of checksum
func (user *User) encryptPassword() {
	checksum := md5.Sum([]byte(user.Password))
	user.Password = string(checksum[:])
}
