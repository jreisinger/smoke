package helper

import (
	"testing"
)

func TestCountNonEmptyLines(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		input []byte
		want  int
	}{
		{input: []byte("line 1\nline 2\nline 3"), want: 3},
		{input: []byte("line 1\nline 2\nline 3\n"), want: 3},
		{input: []byte(""), want: 0},
		{input: []byte("\n"), want: 0},
	}

	for i, tc := range testCases {
		got := CountNonEmptyLines(tc.input)
		if tc.want != got {
			t.Errorf("tc %d: want %d, got %d", i, tc.want, got)
		}
	}
}
