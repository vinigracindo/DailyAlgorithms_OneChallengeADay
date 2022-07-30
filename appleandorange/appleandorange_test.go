package appleandorange

import "testing"

type input struct {
	s       int
	t       int
	a       int
	b       int
	apples  []int
	oranges []int
}

type output struct {
	apple  int
	orange int
}

func TestMain_AppleAndOrange(t *testing.T) {
	tests := []struct {
		input  input
		output output
	}{
		{
			input: input{
				s:       7,
				t:       10,
				a:       4,
				b:       12,
				apples:  []int{2, 3, -4},
				oranges: []int{3, -2, -4},
			},
			output: output{
				apple:  1,
				orange: 2,
			},
		},
		{
			input: input{
				s:       7,
				t:       11,
				a:       5,
				b:       15,
				apples:  []int{-2, 2, 1},
				oranges: []int{5, -6},
			},
			output: output{
				apple:  1,
				orange: 1,
			},
		},
	}

	for _, test := range tests {
		apple, orange := CountApplesAndOranges(test.input.s, test.input.t, test.input.a, test.input.b, test.input.apples, test.input.oranges)
		if apple != test.output.apple || orange != test.output.orange {
			t.Errorf("Expected %d and %d, got %d and %d", test.output.apple, test.output.orange, apple, orange)
		}
	}
}
