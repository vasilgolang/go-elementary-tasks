package webserver

import (
	"net/http"
	"log"
	"fmt"
	"github.com/vasilgolang/go-elementary-tasks/task1"
	"encoding/json"
	"io/ioutil"
	"unicode/utf8"
	"github.com/vasilgolang/go-elementary-tasks/task2"
	"strconv"
	"github.com/vasilgolang/go-elementary-tasks/taskmanager"
)

func handlerMainPage(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()       // parse arguments, you have to call this by yourself
	//fmt.Println(r.Form) // print form information in server side
	fmt.Println("handlerMainPage", r.URL.Path)
	//fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	//for k, v := range r.Form {
	//	fmt.Println("key:", k)
	//	fmt.Println("val:", strings.Join(v, ""))
	//}
	//fmt.Fprintf(w, "Hello astaxie!") // send data to client side
}

func handlerTask1(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Errorf("ioutil.ReadAll(r.Body) Error:", err)
	}
	defer r.Body.Close()
	fmt.Println("body:", string(body))
	var params task1.Params

	err = json.Unmarshal(body, &params)
	if err != nil {
		fmt.Errorf("json.Unmarshal(body, t) Error:", err)
	}

	symbol, _ := utf8.DecodeRuneInString(params.Symbol) // symbol contains the first rune of the string
	board, err := task1.ChessBoard2(params.Width, params.Height, symbol)

	jsonData, _ := json.Marshal(struct {
		Err   string
		Board string
	}{
		Err:   ErrToString(err),
		Board: board,
	})
	w.Write(jsonData)
}

func handlerTask2(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Errorf("ioutil.ReadAll(r.Body) Error:", err)
	}
	fmt.Println("body:", string(body))
	var t task2.Params

	err = json.Unmarshal(body, &t)
	if err != nil {
		fmt.Errorf("json.Unmarshal(body, t) Error:", err)
	}
	defer r.Body.Close()
	fmt.Printf("Values: %#v\r\n", t)
	res, minEnvelope, err := task2.CanEncloseEnvelopes(t.Envelope1, t.Envelope2)
	fmt.Println(res, minEnvelope, err)
	w.Write([]byte(fmt.Sprintf("res:%v, minEnvelope:%v, err:%v", res, minEnvelope, err)))
}

func handlerAllTasks(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Errorf("ioutil.ReadAll(r.Body) Error:", err)
	}
	fmt.Println("body:", string(body))
	var t task2.Params

	err = json.Unmarshal(body, &t)
	if err != nil {
		fmt.Errorf("json.Unmarshal(body, t) Error:", err)
	}
	defer r.Body.Close()
	fmt.Printf("Values: %#v\r\n", t)
	res, minEnvelope, err := task2.CanEncloseEnvelopes(t.Envelope1, t.Envelope2)
	fmt.Println(res, minEnvelope, err)
	w.Write([]byte(fmt.Sprintf("res:%v, minEnvelope:%v, err:%v", res, minEnvelope, err)))
}

func handlerTask(w http.ResponseWriter, r *http.Request) {
	taskNumber, err := strconv.Atoi(r.RequestURI[len(`/task/`):])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You should use this route like /task/N - where N is number"))
	}

	fmt.Println("TaskNumber:", taskNumber)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Errorf("ioutil.ReadAll(r.Body) Error:", err).Error()))

	}

	result, err := taskmanager.RunTask(taskNumber, string(body))

	b, _ := json.Marshal(struct {
		Err    string
		Result string
	}{
		Err:    ErrToString(err),
		Result: result,
	})

	w.Write(b)
}

func Run() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlerMainPage) // set router
	//http.HandleFunc("/task/1", handlerTask1)       // set router
	//http.HandleFunc("/task/2", handlerTask2)       // set router
	//http.HandleFunc("/tasks/all", handlerAllTasks) // set router
	http.HandleFunc("/task/", handlerTask) // set router

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
