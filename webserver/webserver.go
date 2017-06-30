package webserver

import (
	"net/http"
	"log"
	"fmt"
	"github.com/vasilgolang/go-elementary-tasks/task1"
	"encoding/json"
	"io/ioutil"
	"unicode/utf8"
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
	fmt.Println("body:", string(body))
	var t task1.Params

	err = json.Unmarshal(body, &t)
	if err != nil {
		fmt.Errorf("json.Unmarshal(body, t) Error:", err)
	}
	defer r.Body.Close()
	fmt.Printf("Values: %#v\r\n", t)
	//fmt.Println("handlerTask1", r.URL.Path)
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("body:", string(body))
	symbol, _ := utf8.DecodeRuneInString(t.Symbol) // symbol contains the first rune of the string
	board, err := task1.ChessBoard(t.Width, t.Height, symbol)
	fmt.Println(board)
	w.Write([]byte(board))
}

func Run() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlerMainPage)   // set router
	http.HandleFunc("/task/1", handlerTask1) // set router

	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}