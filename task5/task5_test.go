package task5

import (
	"testing"
)

func TestNum2Ticket(t *testing.T) {
	var tests = []struct {
		num    int
		digits [6]uint8
	}{
		{num: 0, digits: [...]uint8{0, 0, 0, 0, 0, 0}},
		{num: 1, digits: [...]uint8{0, 0, 0, 0, 0, 1}},
		{num: 20, digits: [...]uint8{0, 0, 0, 0, 2, 0}},
		{num: 100000, digits: [...]uint8{1, 0, 0, 0, 0, 0}},
		{num: 555555, digits: [...]uint8{5, 5, 5, 5, 5, 5}},

	}

	for _, test := range tests {
		if ticket := num2Ticket(test.num); test.digits != ticket {
			t.Errorf("num2Digits(%d) must return %v instead of %v", test.num, test.digits, ticket)
		}
	}
}

func BenchmarkBestCountingSuccessTickets(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BestCountingSuccessTickets(0, 999999)
	}
}
