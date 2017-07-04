package taskmanager

import (
	"errors"
	"fmt"
)

//map[TASK_NUMBER int]func(JSON_PARAMS string) (RESULT string, ERR error)
var jsonRunners map[int]func(string) (string, error) = make(map[int]func(string) (string, error))

func RegisterJsonRunner(task int, jsonRunner func(string) (string, error)) {
	jsonRunners[task] = jsonRunner
}

func RunTask(task int, jsonParams string) (result string, err error) {
	taskFunc, ok := jsonRunners[task]
	if !ok {
		return "", errors.New(fmt.Sprintf("Task #%d didn't registered runner"))
	}
	return taskFunc(jsonParams)
}
