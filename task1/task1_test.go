package task1

import (
	"testing"
	"unicode/utf8"
)

var tests = []struct {
	input        Params
	passValidate bool
	board        string
}{
	{
		input:        Params{Width: 3, Height: 3, Symbol: "*"},
		passValidate: true,
		board:        "* *\r\n * \r\n* *\r\n",
	},
	{
		input:        Params{Width: 3, Height: 3, Symbol: "ab"},
		passValidate: true,
		board:        "a a\r\n a \r\na a\r\n",
	},
	{
		input:        Params{Width: -1, Height: 10, Symbol: "*"},
		passValidate: false,
	},
	{
		input:        Params{Width: 5, Height: -10, Symbol: "*"},
		passValidate: false,
	},
}

func TestChessBoard(t *testing.T) {
	for _, test := range tests {
		symbol, _ := utf8.DecodeRuneInString(test.input.Symbol) // symbol contains the first rune of the string
		if board, err := ChessBoard(test.input.Width, test.input.Height, symbol); (err == nil) && board != test.board {
			t.Errorf("Wait board:\r\n%s\r\nGet board:\r\n%s", test.board, board)
		}
	}
}
