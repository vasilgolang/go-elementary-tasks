package webserver

import (
	"net/http"
	"log"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"github.com/vasilgolang/go-elementary-tasks/taskmanager"
	"strings"
)

type TaskResult struct {
	Err    string
	Result string
}

func handlerMainPage(w http.ResponseWriter, r *http.Request) {
	// redirect to static index.html
	http.Redirect(w, r, `/static/index.html`, http.StatusSeeOther)
}

func handlerAllTasks(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Errorf("ioutil.ReadAll(r.Body) Error:%s", err)
	}
	//fmt.Println("body:", string(body))
	defer r.Body.Close()

	params := map[string]json.RawMessage{}
	err = json.Unmarshal([]byte(body), &params)
	if err != nil {
		fmt.Errorf("Unmarshal Error:", err)
	}
	results := map[string]TaskResult{}
	for k, v := range params {
		sNumber := strings.Replace(k, `task`, ``, 1)
		if i, err := strconv.Atoi(sNumber); err != nil {
			fmt.Println("Error:", err)
			continue
		} else {
			result, err := taskmanager.RunTask(i, string(v))
			results[k] = TaskResult{
				Result: result,
				Err:    ErrToString(err),
			}
		}

	}

	fmt.Printf("Results:\r\n%#v\r\n", results)
	b, _:= json.Marshal(results)
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
		Err:    ErrToString(err),
		Result: result,
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
