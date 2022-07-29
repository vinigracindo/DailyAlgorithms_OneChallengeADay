package encryption

import (
	"testing"
)

func TestMain_Encrytion(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  "if man was meant to stay on the ground god would have given us roots",
			output: "imtgdvs fearwer mayoogo anouuio ntnnlvt wttddes aohghn sseoau",
		},
		{
			input:  "haveaniceday",
			output: "hae and via ecy",
		},
		{
			input:  "feedthedog",
			output: "fto ehg ee dd",
		},
		{
			input:  "chillout",
			output: "clu hlt io",
		},
	}

	for _, test := range tests {
		if result := Encryption(test.input); result != test.output {
			t.Errorf("Encrypt(%q) = %q, want %q", test.input, result, test.output)
		}
	}
}
