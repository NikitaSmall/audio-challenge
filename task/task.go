package task

import (
	"time"
)

// Tasker is an interface that all the Tasks must implement
type Tasker interface {
	process() error
	setStatus(status bool)
}

// BaseTask is an example struct to implement Taskter interface
type BaseTask struct {
	rawQuery string
	command  string

	time   time.Time // time to complete the task
	status bool
}

type orderDetails struct {
	phone    string
	userName string
	address  string
}

// PizzaTask is a struct to perform pizza requests
type PizzaTask struct {
	rawQuery string
	command  string

	orderDetails orderDetails
	pizzaName    string
	pizzeriaName string

	time   time.Time
	status bool
}
