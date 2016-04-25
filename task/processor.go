/*
 * This package holds the main logic. Here tasks are parsed, created, processed.
 * This file contains main processor functions and struct.
 */
package task

import (
	"errors"
	"log"
	"os"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/nikitasmall/audio-challenge/config"
	"github.com/nikitasmall/audio-challenge/util"
)

// Processor is a struct which take part in task background processing.
// It gets available tasks that not ready yet and just run their `process` method
type Processor struct {
	taskTypes []string // tasks types available to work with
}

// NewProcessor function ask for legal task list and put it inside of Processor struct.
// A pointer to new Processor struct is returned.
func NewProcessor() *Processor {
	var types []string

	for t := range util.FillTaskList(os.Getenv("TASK_LIST_FILE")) {
		types = append(types, t)
	}

	return &Processor{
		taskTypes: types,
	}
}

// Start function runs an endless loop to process all the undone tasks
// while the app is working.
func (processor *Processor) Start() {
	for {
		for _, taskType := range processor.taskTypes {
			task, err := getUndoneTask(taskType)
			if err != nil {
				// this spams a lot
				if err.Error() == "not found" {
					// log.Printf("Uncomleted task not found for %s task type, continue.", taskType)
					continue
				}

				// shows only unusual errors
				log.Println(err)
				continue
			}

			task.process()
		}
	}
}

// getUndoneTask returns single undone task which taskProcessor is triyng to process
// and complete.
func getUndoneTask(taskType string) (Tasker, error) {
	tasksCollection := config.DB.DB(os.Getenv("MONGO_DB_NAME")).C(collectionName)

	switch taskType {
	case "pizza":
		var task PizzaTask

		err := tasksCollection.Find(bson.M{"status": false, "time": bson.M{"$lte": time.Now()}}).One(&task)
		return &task, err
	default:
		return nil, errors.New("Wrong task type")
	}
}
