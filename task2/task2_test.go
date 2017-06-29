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

func TestValidate(t *testing.T) {
	for _, test := range tests {
		if err := Validate(test.input); (err == nil) != test.passError {
			t.Errorf("passError: %v != %v", err == nil, test.passError)
		}
	}
}

func TestRun(t *testing.T) {
	for _, test := range tests {
		if canEnclose, minEnvelope, err := Run(test.input); (err == nil) != test.passError ||
			canEnclose != test.canEnclose ||
			minEnvelope != test.minEnvelope {
			t.Errorf("Wrong results. We wait canEnclose %v instead %v, %d instead  %d\r\n",
				canEnclose, test.canEnclose, minEnvelope, test.minEnvelope)
		}
	}
}
func TestCanEncloseEnvelopes(t *testing.T) {
	for _, test := range tests {
		if canEnclose, minEnvelope := canEncloseEnvelopes(test.input.Envelope1, test.input.Envelope2); test.passError &&
			(canEnclose != test.canEnclose || minEnvelope != test.minEnvelope) {
			t.Errorf("Wrong results. We wait canEnclose %v instead %v, minEnvelope %d instead  %d\r\n",
				canEnclose, test.canEnclose, minEnvelope, test.minEnvelope)
		}
	}
}
