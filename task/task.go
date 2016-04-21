package task

import (
	"errors"
	"log"
	"time"
)

// Tasker is an interface that all the Tasks must implement
type Tasker interface {
	process() error
	setStatus(status bool)
}

// BaseTask is a struct to hold rawQuery string and to determinate the task inside the query
type BaseTask struct {
	RawQuery string
	Status   bool
}

type orderDetails struct {
	phone    string
	userName string
	address  string
}

// PizzaTask is a struct to perform pizza requests
type PizzaTask struct {
	RawQuery string
	command  string

	orderDetails orderDetails
	pizzaName    string
	pizzeriaName string

	time   time.Time
	Status bool
}

// Process does the PizzaTask work: goes for a pizza
func (pz *PizzaTask) process() error {
	return nil
}

func (pz *PizzaTask) setStatus(status bool) {
	pz.Status = status
}

// DefineTask defines type of a task by RawQuery field
func (task *BaseTask) DefineTask() (Tasker, error) {
	taskType := task.determinate()
	log.Println(taskType)

	switch taskType {
	case "pizza":
		return &PizzaTask{}, nil
	default:
		return nil, errors.New("Cannot determinate task")
	}
}
