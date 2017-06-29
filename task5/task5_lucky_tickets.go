package task5

import (
	"fmt"
	"errors"
)

/*
Есть 2 способа подсчёта счастливых билетов:
1. Простой — если на билете напечатано
шестизначное число, и сумма первых трёх цифр
равна сумме последних трёх, то этот билет
считается счастливым.
2. Сложный — если сумма чётных цифр билета равна
сумме нечётных цифр билета, то билет считается
счастливым.
Определить программно какой вариант подсчёта счастливых билетов даст их
большее количество на заданном интервале.
Входные параметры : границы min и max
Выход : элемент структурного типа с информацией о победившем методе и
количестве счастливых билетов для каждого способа подсчёта.
*/

type Params struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type Result struct {
	Method int
	Count  int
}

// Returns error when params can't pass validation
func Validate(params Params) (err error) {
	if params.Min > params.Max {
		return errors.New(fmt.Sprintf("Min (%d) must be less than Max (%d)", params.Min, params.Max))
	}
	if params.Min > 999999 || params.Min < 0 || params.Max > 999999 {
		return errors.New(fmt.Sprintf("Min (%d) and Max (%d) must be in range from 0 to 999999", params.Min, params.Max))
	}
	return nil
}

func Run(params Params) (result Result, err error) {
	if err = Validate(params); err != nil {
		return
	}
	return BestCountingSuccessTickets(params.Min, params.Max), nil
}

func Demo(params []Params) {
	for _, param := range params {
		fmt.Printf("Received range. Min:%d, Max:%d\r\n", param.Min, param.Max)
		if result, err := Run(param); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Best method is %d, found %d lucky tickets\r\n", result.Method, result.Count)
		}
	}
}

func firstMethod(d [6]uint8) bool {
	return d[0]+d[1]+d[2] == d[3]+d[4]+d[5]
}

func secondMethod(d [6]uint8) bool {
	//return d[0]+d[2]+d[4] == d[1]+d[3]+d[5] // wrong understood task
	sumOdd, sumEven := 0, 0
	for _, v := range d {
		if (v % 2) == 0 {
			sumEven += int(v)
		} else {
			sumOdd += int(v)
		}
	}
	return sumOdd == sumEven
}

func BestCountingSuccessTickets(min, max int) (r Result) {
	firstMethodCounter := 0
	secondMethodCounter := 0
	for i := min; i <= max; i++ {
		ticket := num2Ticket(i)

		if firstMethod(ticket) {
			firstMethodCounter++
		}

		if secondMethod(ticket) {
			secondMethodCounter++
		}
	}
	if firstMethodCounter > secondMethodCounter {
		return Result{
			Method: 1,
			Count:  firstMethodCounter,
		}
	} else {
		return Result{
			Method: 2,
			Count:  secondMethodCounter,
		}
	}

}

func num2Ticket(num int) (ticket [6]uint8) {
	for i := 5; i >= 0; i-- {
		ticket[i] = uint8(num % 10)
		num /= 10
	}
	return
}

