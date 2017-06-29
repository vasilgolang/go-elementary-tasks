package task7

import (
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
	"os"
	"errors"
)

/*
7. Ряд Фибоначчи для диапазона
Вывести все числа Фибоначчи, которые удовлетворяют переданному в функцию
ограничению: находятся в указанном диапазоне, либо имеют указанную длину.
Входные параметры : файл context со значениями min и max, либо с полем length
Выход : срез сгенерированных чисел
*/

type Params struct {
	Context string `json:"context"` // context file name
}

// Returns error when params can't pass validation
func Validate(params Params) (err error) {
	if _, err := os.Stat(params.Context); os.IsNotExist(err) {
		return errors.New("File" + params.Context + "doesn't exist")
	}

	return nil
}

func Demo(params []Params) {
	for _, param := range params {
		fmt.Printf("Received file name:%s\r\n", param.Context)
		if numbers, err := Run(param); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Numbers:", numbers)
		}
	}
}

func Run(params Params) (numbers []int, err error) {
	restriction, err := ParseContext(params.Context)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	return Fib(restriction), nil
}

type restriction struct {
	min, max, length int
}

func parseLine(s string) (result int, err error) {
	intVal, err := strconv.Atoi(s)
	if err != nil {
		return -1, errors.New("Restriction must be an integer")
	}
	return intVal, nil
}

func ParseContext(filename string) (r restriction, err error) {
	r = restriction{-1, -1, -1}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")
	if len(lines) > 2 {
		return r, errors.New("File " + filename + " must contain 1 or 2 strings")
	}
	for k, v := range lines {
		lines[k] = strings.TrimSpace(v)
	}
	if len(lines) == 2 {
		if r.min, err = parseLine(lines[0]); err != nil {
			return r, err
		}
		if r.max, err = parseLine(lines[1]); err != nil {
			return r, err
		}
	} else {
		if r.length, err = parseLine(lines[0]); err != nil {
			return r, err
		}
	}

	return
}

func Fib(r restriction) (res []int) {
	fmt.Printf("restriction:%#v\r\n", r)
	a, b := 0, 1
	for i := 0; ; i++ {
		a, b = b, a+b

		if r.length == -1 {
			if a < r.min {
				continue
			}
			if a > r.max {
				break
			}

		} else {
			l := len(strconv.Itoa(a))
			if l < r.length {
				continue
			}
			if l > r.length {
				break
			}

		}
		res = append(res, a)

	}
	return
}
