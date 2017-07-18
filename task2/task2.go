package task2

/*
2. Анализ конвертов
Есть два конверта со сторонами (a,b) и (c,d). Требуется определить, можно ли один
конверт вложить в другой. Программа должна обрабатывать ввод чисел с плавающей
точкой.
Входные параметры : структурные типы конверт1 и конверт2
Выход : номер конверта, если вложение возможно, 0 – если вложение невозможно.
UPD: номер меньшего конверта
*/

import (
	"math"
	"errors"
	"fmt"
	"encoding/json"
	"github.com/vasilgolang/go-elementary-tasks/taskmanager"
)

type Params []Envelope

func JsonRunner(jsonData string) (result string, err error) {
	var params Params
	err = json.Unmarshal([]byte(jsonData), &params)
	if err != nil {
		return
	}
	return Demo(params)
}

func init() {
	taskmanager.RegisterJsonRunner(2, JsonRunner)
}

// Returns error when params can't pass validation
func validate(e1, e2 Envelope) (err error) {
	if !(e1.Validate() && e2.Validate() ) {
		return errors.New("Width and height of both envelopes must be positive numbers")
	}
	return
}

func Demo(param Params) (result string, err error) {
	fmt.Printf("Received params:\r\nEnvelope1: %#v\r\nEnvelope2: %#v\r\n", param[0], param[1])
	if canEnclose, minEnvelope, err := CanEncloseEnvelopes(param[0], param[1]); err != nil {
		return "", err
	} else {
		if canEnclose {
			return fmt.Sprintf("Envelopes can be enclosed. The smallest envelope is %d", minEnvelope), nil
		} else {
			return "", errors.New(fmt.Sprint("Envelopes can't be enclosed."))
		}
	}
}

type Envelope struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

func (e Envelope) Validate() bool {
	if e.Width <= 0 || e.Height <= 0 {
		return false
	}
	return true
}

func CanEncloseEnvelopes(e1, e2 Envelope) (res bool, minEnvelope uint8, err error) {
	if err = validate(e1, e2); err != nil {
		return
	}
	minE1 := math.Min(e1.Width, e1.Height)
	minE2 := math.Min(e2.Width, e2.Height)
	maxE1 := math.Max(e1.Width, e1.Height)
	maxE2 := math.Max(e2.Width, e2.Height)

	switch {
	case minE1 == minE2 || maxE1 == maxE2:
		return
	case minE1 < minE2 && maxE1 < maxE2:
		return true, 1, nil
	case minE1 > minE2 && maxE1 > maxE2:
		return true, 2, nil
	}
	return
}
