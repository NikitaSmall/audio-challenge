package task

import (
	"os"
	"strings"

	"github.com/nikitasmall/audio-challenge/util"
)

func (task BaseTask) determinateTask() string {
	taskList := util.FillTaskList(os.Getenv("TASK_LIST_FILE"))

	for taskName, keyWords := range taskList {
		for _, keyWord := range keyWords {

			if strings.Contains(task.RawQuery, keyWord) {
				return taskName
			}
		}
	}

	return "unknown"
}

func (task BaseTask) determinateFood() string {
	var order []string
	foodList := util.FillList(os.Getenv("FOOD_LIST_FILE"))

	for _, food := range foodList {

		if strings.Contains(task.RawQuery, food) {
			order = append(order, food)
		}
	}

	return strings.Join(order, ", ")
}
