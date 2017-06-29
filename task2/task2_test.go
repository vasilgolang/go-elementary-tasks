package task2

import "testing"

var tests = []struct {
	input       Params
	canEnclose  bool
	minEnvelope uint8
	passError   bool
}{
	{
		input: Params{
			Envelope1: Envelope{Width: 5, Height: 5},
			Envelope2: Envelope{Width: 5.1, Height: 5.1},
		},
		canEnclose:  true,
		minEnvelope: 1,
		passError:   true,
	},
}

func TestCanEncloseEnvelopes(t *testing.T) {
	for _, test := range tests {
		canEnclose, minEnvelope, err := CanEncloseEnvelopes(test.input.Envelope1, test.input.Envelope2)
		if (err == nil) && test.passError && (canEnclose != test.canEnclose || minEnvelope != test.minEnvelope) {
			t.Errorf("Wrong results. We wait canEnclose %v instead %v, minEnvelope %d instead  %d\r\n",
				canEnclose, test.canEnclose, minEnvelope, test.minEnvelope)
		}
	}
}
