package task

import (
	"errors"
	"log"
	"time"

	"github.com/nikitasmall/audio-challenge/util"
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

// Query returns raw query method
func (pz *PizzaTask) Query() string {
	return pz.RawQuery
}

// defineTask defines type of a task by RawQuery field
func (task *BaseTask) defineTask() (Tasker, error) {
	taskType := task.determinate()

	name, addr, date := task.getQueryParams()

	log.Println(name, addr, date)

	switch taskType {
	case "pizza":
		return &PizzaTask{
			RawQuery: task.RawQuery,
			Status:   task.Status,
			command:  "pizza",
			time:     date,
			orderDetails: orderDetails{
				userName: name,
				address:  addr,
			},
		}, nil
	default:
		return nil, errors.New("Cannot determinate task")
	}
}

func (task *BaseTask) getQueryParams() (string, string, time.Time) {
	jsonParams := paramsRequest(task.RawQuery)
	params := util.ParseJSON(jsonParams)

	name := parseName(params)
	addr := parseAddr(params)
	date := parseTime(params)

	log.Println(addr)
	return name, addr, date
}
