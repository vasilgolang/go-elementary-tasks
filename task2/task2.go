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
)

type Params struct {
	Envelope1 Envelope `json:"envelope1"`
	Envelope2 Envelope `json:"envelope2"`
}

// Returns error when params can't pass validation
func Validate(params Params) (err error) {
	if !(params.Envelope1.Validate() && params.Envelope2.Validate() ) {
		return errors.New("Width and height of both envelopes must be positive numbers")
	}
	return
}

func Demo(params []Params) {
	for _, param := range params {
		fmt.Printf("Received params:\r\nEnvelope1: %#v\r\nEnvelope2: %#v\r\n", param.Envelope1, param.Envelope2)
		if CanEnclose, minEnvelope, err := Run(param); err != nil {
			fmt.Println("Error:", err)
		} else {
			if CanEnclose {
				fmt.Println("Envelopes can be enclosed. The smallest envelope is", minEnvelope)
			} else {
				fmt.Println("Envelopes can't be enclosed.")
			}
		}
	}
}

func Run(params Params) (CanEnclose bool, minEnvelope uint8, err error) {
	if err = Validate(params); err != nil {
		return
	}
	CanEnclose, minEnvelope = canEncloseEnvelopes(params.Envelope1, params.Envelope2)
	return
}

type Envelope struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

func (e *Envelope) Validate() bool {
	if e.Width <= 0 || e.Height <= 0 {
		return false
	}
	return true
}

func canEncloseEnvelopes(e1, e2 Envelope) (res bool, minEnvelope uint8) {

	minE1 := math.Min(e1.Width, e1.Height)
	minE2 := math.Min(e2.Width, e2.Height)
	maxE1 := math.Max(e1.Width, e1.Height)
	maxE2 := math.Max(e2.Width, e2.Height)

	switch {
	case minE1 == minE2 || maxE1 == maxE2:
		return
	case minE1 < minE2 && maxE1 < maxE2:
		return true, 1
	case minE1 > minE2 && maxE1 > maxE2:
		return true, 2
	}
	return
}
