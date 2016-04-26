/*Package router holds router and hadnlers for its routes.
 * In this file session functions and helpers are declared.
 */
package router

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/nikitasmall/audio-challenge/user"
)

// local storage for session
var localStorage = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

// helper function that hides all the operations with session
// and only saves username to session
func setSessionUser(username string, phone string, c *gin.Context) {
	session, err := localStorage.Get(c.Request, "auth")
	if err != nil {
		log.Println("Cannot obtain session. ", err.Error())
	}

	session.Values["username"] = username
	session.Values["phone"] = phone
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		log.Println("Cannot save session. ", err.Error())
	}
}

// helper function that hides all the operations with session
// and only returns phone from session.
// also returns error if the phone is not set in session
func getSessionPhone(c *gin.Context) (string, error) {
	session, err := localStorage.Get(c.Request, "auth")
	if err != nil {
		log.Println("Cannot obtain session. ", err.Error())
	}

	phone := session.Values["phone"]
	if phone == nil {
		return "", errors.New("No users stored in session")
	}

	return phone.(string), nil
}

// helper function that hides all the operations with session
// and only returns username from session.
// also returns error if the username is not set in session
func getSessionUser(c *gin.Context) (string, error) {
	session, err := localStorage.Get(c.Request, "auth")
	if err != nil {
		log.Println("Cannot obtain session. ", err.Error())
	}

	username := session.Values["username"]
	if username == nil {
		return "", errors.New("No users stored in session")
	}

	return username.(string), nil
}

// function binds json input from request to passed user struct
func bindUser(user *user.User, c *gin.Context) {
	err := c.BindJSON(user)
	if err != nil {
		log.Panic("Error on user binding from json. ", err.Error())
	}
}

// function that returns a name if user is set in session,
// returns nil otherwise
func checkUser(c *gin.Context) {
	username, err := getSessionUser(c)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"username": nil})
	} else {
		c.JSON(http.StatusOK, gin.H{"username": username})
	}
}

// handler function that makes attempt to register provided user
// set session for username in case of success
// and returns error with description if failed
func register(c *gin.Context) {
	u := user.CreateUser()
	bindUser(u, c)

	err := u.Register()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		setSessionUser(u.Username, u.Phone, c)
		c.JSON(http.StatusOK, gin.H{"message": "Hello, " + u.Username + "!"})
	}
}

// handler function that makes attempt to login provided user
// set session for username in case of success
// and returns error with description if failed
func login(c *gin.Context) {
	u := user.CreateUser()
	bindUser(u, c)

	err := u.CheckUser()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		setSessionUser(u.Username, u.Phone, c)
		c.JSON(http.StatusOK, gin.H{"message": "Hello, " + u.Username + "!"})
	}
}

// handler function that remove user's session from local storage
func logout(c *gin.Context) {
	session, err := localStorage.Get(c.Request, "auth")
	if err != nil {
		log.Println("Cannot obtain session. ", err.Error())
	}

	delete(session.Values, "username")
	delete(session.Values, "phone")

	err = session.Save(c.Request, c.Writer)
	if err != nil {
		log.Panic("Cannot save session. ", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logout is successfull."})
}
