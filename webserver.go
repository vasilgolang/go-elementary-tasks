package main

import (
	_ "github.com/vasilgolang/go-elementary-tasks/task1"
	_ "github.com/vasilgolang/go-elementary-tasks/task2"
	"github.com/vasilgolang/go-elementary-tasks/webserver"
)

func main() {
	webserver.Run()
}
