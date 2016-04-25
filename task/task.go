/*Package task holds the main logic. Here tasks are parsed, created, processed.
 * This file holds main declarations and some main task methods.
 */
package task

import (
	"errors"
	"io"
	"log"
	"os"
	"time"

	"github.com/nikitasmall/audio-challenge/config"
	"github.com/nikitasmall/audio-challenge/socket"
	"gopkg.in/mgo.v2/bson"
)

// Config variables for handy connection.
var collectionName = "tasks"

// Tasker is an interface that all the Tasks must implement
type Tasker interface {
	id() string
	process() error
	changeStatus(status bool)

	Query() string
}

// BaseTask is a struct to hold rawQuery string and to determinate the task inside the query
type BaseTask struct {
	RawQuery string
	Status   bool
}

// OrderDetails contains additional general information about order tasks
type OrderDetails struct {
	Phone       string `json:"phone"`
	UserName    string `json:"username"`
	Address     string `json:"address"`
	PaymentType string `json:"paymenttype"`
}

// PizzaTask is a struct to perform pizza requests
type PizzaTask struct {
	ID string `json:"id" bson:"_id,omitempty"`

	RawQuery string
	Command  string `json:"command"`

	OrderDetails OrderDetails `json:"orderdetails"`
	OrderList    string       `json:"orderlist"`
	PizzeriaName string       `json:"pizzerianame"`

	Time   time.Time `json:"time"`
	Status bool      `json:"status"`
}

// Process does the PizzaTask work: goes for a pizza
func (pz *PizzaTask) process() error {
	pz.changeStatus(true)
	return nil
}

func (pz *PizzaTask) changeStatus(status bool) {
	tasksCollection := config.DB.DB(os.Getenv("MONGO_DB_NAME")).C(collectionName)
	err := tasksCollection.Update(bson.M{"_id": pz.ID}, bson.M{"$set": bson.M{"status": status}})
	if err != nil {
		log.Printf("Error on changing tasks status: %s", err.Error())
	} else {
		pz.Status = true
		socket.MainHub.SendMessage(socket.TaskComplete, pz)
	}
}

// Query returns raw query method
func (pz *PizzaTask) Query() string {
	return pz.RawQuery
}

func (pz *PizzaTask) id() string {
	return pz.ID
}

// ProcessMessage sends the message file to the Yandex API and returns parsed
// task as an interface type Tasker.
func ProcessMessage(message io.Reader) (Tasker, error) {
	parsedResult := messageRequest(message)

	t, err := newTask(parsedResult)
	if err != nil {
		return nil, err
	}

	task, err := t.defineTask()
	if err != nil {
		return nil, err
	}

	return saveTask(task)
}

// defineTask defines type of a task by RawQuery field.
func (task *BaseTask) defineTask() (Tasker, error) {
	taskType := task.determinateTask()

	name, addr, date := task.getQueryParams()

	switch taskType {
	case "pizza":
		return &PizzaTask{
			ID:           bson.NewObjectId().Hex(),
			RawQuery:     task.RawQuery,
			Status:       task.Status,
			OrderList:    task.determinateFood(),
			PizzeriaName: task.determinatePizzeria(),
			Command:      "pizza",
			Time:         date,
			OrderDetails: OrderDetails{
				UserName:    name,
				Address:     addr,
				PaymentType: task.determinatePaymentType(),
			},
		}, nil
	default:
		return nil, errors.New("Cannot determinate task")
	}
}

// saveTask stores sucessfully parsed task to mongo collection
func saveTask(t Tasker) (Tasker, error) {
	tasksCollection := config.DB.DB(os.Getenv("MONGO_DB_NAME")).C(collectionName)
	return t, tasksCollection.Insert(t)
}

// TaskList returns list of all the possible tasks.
// interface returning value is used for getting mutability and
// easy way to get different types of tasks.
func TaskList() (interface{}, error) {
	tasksCollection := config.DB.DB(os.Getenv("MONGO_DB_NAME")).C(collectionName)

	var tasks []interface{}

	err := tasksCollection.Find(nil).All(&tasks)
	if err != nil {
		log.Print(err)
	}

	return tasks, err
}
