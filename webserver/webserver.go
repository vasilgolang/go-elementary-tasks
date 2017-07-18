package webserver

import (
	"net/http"
	"log"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"github.com/vasilgolang/go-elementary-tasks/taskmanager"
)

type TaskResult struct {
	Task   int
	Reason string
	Resp   string
}

func handlerMainPage(w http.ResponseWriter, r *http.Request) {
	// redirect to static index.html
	//http.Redirect(w, r, `/static/index.html`, http.StatusSeeOther)
	if content, err := ioutil.ReadFile("./static/index.html"); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
	} else {
		w.Write(content)
	}
}

func handlerAllTasks(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Errorf("ioutil.ReadAll(r.Body) Error:%s", err)
	}
	defer r.Body.Close()
	fmt.Println("body:", string(body))

	params := map[int]json.RawMessage{}
	err = json.Unmarshal(body, &params)
	if err != nil {
		fmt.Errorf("Unmarshal Error:", err)
	}

	results := []TaskResult{}
	for task, jsonParams := range params {
		result, err := taskmanager.RunTask(task, string(jsonParams))
		results = append(results, TaskResult{
			Task:   task,
			Resp:   result,
			Reason: ErrToString(err),
		})

	}

	fmt.Printf("Results:\r\n%#v\r\n", results)
	b, _ := json.Marshal(results)
	//fmt.Println(err, string(b))
	w.Write(b)
}

func handlerTask(w http.ResponseWriter, r *http.Request) {
	taskNumber, err := strconv.Atoi(r.RequestURI[len(`/task/`):])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You should use this route like /task/N - where N is number"))
	}

	fmt.Println("TaskNumber:", taskNumber) // debun in console

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Errorf("ioutil.ReadAll(r.Body) Error: %v", err).Error()))
	}

	result, err := taskmanager.RunTask(taskNumber, string(body))

	b, _ := json.Marshal(TaskResult{
		Task:   taskNumber,
		Reason: ErrToString(err),
		Resp:   result,
	})

	w.Write(b)
}

func Run() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlerMainPage)       // set router
	http.HandleFunc("/task/", handlerTask)      // set router
	http.HandleFunc("/tasks/", handlerAllTasks) // set router

	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func ErrToString(err error) (msg string) {
	if err == nil {
		return ""
	}
	return err.Error()
}
