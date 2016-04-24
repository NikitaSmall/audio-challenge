/*
 * This package holds the main logic. Here tasks are parsed, created, processed.
 * This file holds task definition method and some task-specific functions,
 * such as pizzeria name recognition.
 */
package task

import (
	"os"
	"regexp"
	"strings"

	"github.com/nikitasmall/audio-challenge/util"
)

var cashRegexp = regexp.MustCompile("налич|при доставке|по доставке|курьер|cash")

// determinateTask works with help of outer file which contains
// task name and its keywords.
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

// determinatePizzeria checks for reserved keywords in the provided text
// and returns the found result (pizzeria name).
func (task BaseTask) determinatePizzeria() string {
	pizzeriaList := util.FillList(os.Getenv("PIZZERIA_LIST_FILE"))

	for _, pizzeria := range pizzeriaList {

		if strings.Contains(task.RawQuery, pizzeria) {
			return pizzeria
		}
	}

	return ""
}

// determinateFood checks for reserved keywords in the provided text
// and returns the found result (joined food name array).
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

// determinatePaymentType checks for key words about payment type
// inside the provided message and returns it.
func (task BaseTask) determinatePaymentType() string {
	if cashRegexp.FindStringIndex(task.RawQuery) != nil {
		return "cash"
	}

	return "terminal"
}
