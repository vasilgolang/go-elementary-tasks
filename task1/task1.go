package task1

/*
1. Шахматная доска
Вывести шахматную доску с заданными размерами высоты и ширины, по принципу:
* * * * * *
* * * * * *
* * * * * *
* * * * * *
Входные параметры : длина, ширина, символ для отображения.
Выход : строка с представлением шахматной доски
*/

import (
	"errors"
	"unicode/utf8"
	"fmt"
)

type Params struct {
	Width  int `json:"width" xml:"width"`       // chess board width
	Height int `json:"height" xml:"height"`     // chess board height
	Symbol string  `json:"symbol" xml:"symbol"` // chess board symbol for white fields
}

func Demo(params []Params) {
	for _, param := range params {
		fmt.Printf("Received params:\r\nWidth: %d\r\nHeight: %d\r\nSymbol: %s\r\n", param.Width, param.Height, param.Symbol)
		symbol, _ := utf8.DecodeRuneInString(param.Symbol) // symbol contains the first rune of the string
		if result, err := ChessBoard(param.Width, param.Height, symbol); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Result:\r\n", result)
		}
	}
}

// Returns error when params can't pass validation
func validate(width, height int) (err error) {
	// Check if width and height are positive numbers
	if width < 0 || height < 0 {
		return errors.New("Width and height must be more than 0")
	}
	return nil
}

// Returns text plain chess board
func ChessBoard(width, height int, symbol rune) (board string, err error) {
	if err := validate(width, height); err != nil {
		return "", err
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			// Detection "white" or "black" field of the chess board
			if (i+j)%2 == 0 {
				board += string(symbol)
			} else {
				board += ` `
			}
		}
		board += "\r\n"
	}
	return
}