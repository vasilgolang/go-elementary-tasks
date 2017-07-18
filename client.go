package main

import (
	"fmt"
	"flag"
	"os"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"bytes"
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

func taskRunner(task int, jsonParams []byte) {
	url := fmt.Sprintf("http://localhost:9090/task/%d", task)
	fmt.Println("URL:>", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonParams))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
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
		taskRunner(task, jsonParams)
		//result, err := taskmanager.RunTask(task, string(jsonParams))
		//if err != nil {
		//	fmt.Println("Error:", err)
		//} else {
		//	fmt.Println(result)
		//}
	}
}
