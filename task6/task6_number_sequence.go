package task6

import (
	"math"
	"os"
	"strconv"
	"io/ioutil"
	"strings"
	"github.com/pkg/errors"
	"fmt"
)

/*
Вывести в файл через запятую ряд длиной n, состоящий из натуральных чисел,
квадрат которых не меньше заданного m.
Входные параметры : длина и значение минимального квадрата
Выход : nil если сохранение удалось и err в противном случае
*/

type Params struct {
	Length int `json:"length"`
	Square int `json:"square"`
}

// Returns error when params can't pass validation
func Validate(params Params) (err error) {
	if params.Length <= 0 {
		return errors.New("Length must be more than 0")
	}
	if params.Square <= 0 {
		return errors.New("Square must be more than 0")
	}
	return nil
}

func Run(params Params) (err error) {
	if err = Validate(params); err != nil {
		return
	}
	return WriteNumbers(params.Length, params.Square)
}

func Demo(params []Params) {
	for _, param := range params {
		fmt.Printf("Received Length:%d, Square:%d\r\n", param.Length, param.Square)
		if err := Run(param); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Please see result in file numbers.txt")
		}
	}
}

func WriteNumbers(length int, square int) error {
	if square < 0 {
		return errors.New("Square can't be less than 0")
	}

	f, err := os.Create(`numbers.txt`)
	if err != nil {
		return err
	}
	defer f.Close()

	var numsS []string
	i := int(math.Sqrt(float64(square))) // маленькая оптимизация начального значения i :)
	for j := 0; j <= length; j, i = j+1, i+1 {
		if i*i < square {
			continue
		}

		numsS = append(numsS, strconv.Itoa(i))
	}

	return ioutil.WriteFile(`numbers.txt`, []byte(strings.Join(numsS, `,`)), 0777)
}
