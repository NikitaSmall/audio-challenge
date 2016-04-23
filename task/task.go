/*
 * This package holds the main logic. Here tasks are parsed, created, processed.
 * This file holds main declarations and some main task methods.
 */
package task

import (
	"errors"
	"io"
	"os"
	"time"

	"github.com/nikitasmall/audio-challenge/config"
	"gopkg.in/mgo.v2/bson"
)

// Config variables for handy connection.
var collectionName = "tasks"

// Tasker is an interface that all the Tasks must implement
type Tasker interface {
	id() string
	process() error
	setStatus(status bool)

	Query() string
}

// BaseTask is a struct to hold rawQuery string and to determinate the task inside the query
type BaseTask struct {
	RawQuery string
	Status   bool
}

// OrderDetails contains additional general information about order tasks
type OrderDetails struct {
	Phone    string
	UserName string
	Address  string
}

// PizzaTask is a struct to perform pizza requests
type PizzaTask struct {
	Id string `json:"id" bson:"_id,omitempty"`

	RawQuery string
	Command  string

	OrderDetails OrderDetails
	OrderList    string
	PizzeriaName string

	Time   time.Time `json:"time"`
	Status bool      `json:"status"`
}

// Process does the PizzaTask work: goes for a pizza
func (pz *PizzaTask) process() error {
	return nil
}

func (pz *PizzaTask) setStatus(status bool) {
	pz.Status = status
}

// Query returns raw query method
func (pz *PizzaTask) Query() string {
	return pz.RawQuery
}

func (pz *PizzaTask) id() string {
	return pz.Id
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
			Id:           bson.NewObjectId().Hex(),
			RawQuery:     task.RawQuery,
			Status:       task.Status,
			OrderList:    task.determinateFood(),
			PizzeriaName: task.determinatePizzeria(),
			Command:      "pizza",
			Time:         date,
			OrderDetails: OrderDetails{
				UserName: name,
				Address:  addr,
			},
		}, nil
	default:
		return nil, errors.New("Cannot determinate task")
	}
}

// saveTask stores sucessfully parsed task to mongo collection
func saveTask(t Tasker) (Tasker, error) {
	session := config.Connect()
	defer session.Close()

	tasksCollection := session.DB(os.Getenv("MONGO_DB_NAME")).C(collectionName)
	return t, tasksCollection.Insert(t)
}

// TaskList returns list of all the possible tasks of some special type
func TaskList(kind string) (interface{}, error) {
	session := config.Connect()
	defer session.Close()

	tasksCollection := session.DB(os.Getenv("MONGO_DB_NAME")).C(collectionName)

	switch kind {
	case "pizza":
		var tasks []PizzaTask

		err := tasksCollection.Find(nil).All(&tasks)
		return tasks, err
	default:
		return nil, errors.New("Wrong task type")
	}
}
