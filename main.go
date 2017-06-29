package main

import (
	"fmt"
	"github.com/vasilgolang/go-elementary-tasks/task1"
	"github.com/vasilgolang/go-elementary-tasks/task2"
	"github.com/vasilgolang/go-elementary-tasks/task3"
	"github.com/vasilgolang/go-elementary-tasks/task4"
	"github.com/vasilgolang/go-elementary-tasks/task5"
	"github.com/vasilgolang/go-elementary-tasks/task6"
	"github.com/vasilgolang/go-elementary-tasks/task7"
	"encoding/xml"
	"os"
	"strings"
	"path/filepath"
	"io/ioutil"
	"log"
	"encoding/json"
)

type Params struct {
	Params1 []task1.Params `json:"task1params" xml:"task1params"`
	Params2 []task2.Params `json:"task2params" xml:"task2params"`
	Params3 []task3.Params `json:"task3params" xml:"task3params"`
	Params4 []task4.Params `json:"task4params" xml:"task4params"`
	Params5 []task5.Params `json:"task5params" xml:"task5params"`
	Params6 []task6.Params `json:"task6params" xml:"task6params"`
	Params7 []task7.Params `json:"task7params" xml:"task7params"`
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Runs tasks\r\n", os.Args[0], "filename\r\n  filename - must contain JSON format")
		return
	}

	fileName := os.Args[1]
	extension := strings.ToLower(filepath.Ext(fileName))

	fmt.Println("FILE:", fileName)

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Fatal error: ", err)
	}

	var params Params
	switch extension {
	case ".json":
		err = json.Unmarshal(content, &params)
		if err != nil {
			log.Fatal("JSON error: ", err)
		}
	case ".xml":
		err = xml.Unmarshal(content, &params)
		if err != nil {
			log.Fatal("XML error: ", err)
		}
	default:
		fmt.Println("Unknown file extension")
		return
	}

	fmt.Printf("%#v\r\n", params)

	task1.Demo(params.Params1)
	task2.Demo(params.Params2)
	return
	task3.Demo(params.Params3)
	task4.Demo(params.Params4)
	task5.Demo(params.Params5)
	task6.Demo(params.Params6)
	task7.Demo(params.Params7)

}
