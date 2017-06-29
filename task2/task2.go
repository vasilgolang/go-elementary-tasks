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
func validate(e1, e2 Envelope) (err error) {
	if !(e1.Validate() && e2.Validate() ) {
		return errors.New("Width and height of both envelopes must be positive numbers")
	}
	return
}

func Demo(params []Params) {
	for _, param := range params {
		fmt.Printf("Received params:\r\nEnvelope1: %#v\r\nEnvelope2: %#v\r\n", param.Envelope1, param.Envelope2)
		if canEnclose, minEnvelope, err := CanEncloseEnvelopes(param.Envelope1, param.Envelope2); err != nil {
			fmt.Println("Error:", err)
		} else {
			if canEnclose {
				fmt.Println("Envelopes can be enclosed. The smallest envelope is", minEnvelope)
			} else {
				fmt.Println("Envelopes can't be enclosed.")
			}
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
