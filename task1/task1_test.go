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

func TestDifferentAlgorithms(t *testing.T) {
	r1, _ := ChessBoard(100, 100, 'Ы')
	r2, _ := ChessBoard2(100, 100, 'Ы')
	if r1 != r2 {
		t.Error("Algorithms are not equal")
	}
}

func BenchmarkChessBoard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChessBoard(1000, 1000, 'X')
	}
}

func BenchmarkChessBoard2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChessBoard2(1000, 1000, 'X')
	}
}
