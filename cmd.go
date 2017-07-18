package main

import (
	_ "github.com/vasilgolang/go-elementary-tasks/task1"
	_ "github.com/vasilgolang/go-elementary-tasks/task2"
	"fmt"
	"flag"
	"os"
	"io/ioutil"
	"encoding/json"
	"github.com/vasilgolang/go-elementary-tasks/taskmanager"
)

func inputFile() (file string) {
	for {
		fmt.Print("Please enter JSON file name:")
		fmt.Scanln(&file)
		fileInfo, err := os.Stat(file)
		if err == nil && fileInfo.IsDir() {
			fmt.Println("You should enter _file_ name")
			continue
		}
		if !os.IsNotExist(err) {
			break
		}
		fmt.Println("File doesn't exists")

	}
	return
}

func main() {
	var (
		task int    // task number
		file string // JSON file name

	)
	flag.IntVar(&task, "task", 0, "Task number")
	flag.StringVar(&file, "file", "", "JSON file name")
	flag.Parse()
	if file == "" {
		file = inputFile()
	}
	fmt.Println("task", task)
	fmt.Println("file", file)

	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %s:%s", file, err)
		return
	}
	fmt.Println(string(content))

	params := map[int]json.RawMessage{}
	err = json.Unmarshal(content, &params)
	if err != nil {
		fmt.Errorf("Unmarshal Error:", err)
	}
	for task, jsonParams := range params {
		fmt.Printf("Running task #%d\n", task)
		result, err := taskmanager.RunTask(task, string(jsonParams))
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(result)
		}
	}
}
