package task

import (
	"os"
	"strings"

	"github.com/nikitasmall/audio-challenge/util"
)

func (task BaseTask) determinate() string {
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
