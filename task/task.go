/*
 * This package holds the main logic. Here tasks are parsed, created, processed.
 * This file holds main declarations and some main task methods.
 */
package task

import (
	"errors"
	"io"
	"time"
)

// Tasker is an interface that all the Tasks must implement
type Tasker interface {
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
	RawQuery string
	Command  string

	OrderDetails OrderDetails
	OrderList    string
	PizzeriaName string

	Time   time.Time
	Status bool
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

// ProcessMessage sends the message file to the Yandex API and returns parsed task
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

	return task, nil
}

// defineTask defines type of a task by RawQuery field
func (task *BaseTask) defineTask() (Tasker, error) {
	taskType := task.determinateTask()

	name, addr, date := task.getQueryParams()

	switch taskType {
	case "pizza":
		return &PizzaTask{
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
