package mod3_test

import (
	"fmt"
	"testing"

	"github.com/atilaromero/mod_three/examples/mod3"
)

func ModThree(s string) (int, error) {
	isValidRune := func(c rune) bool {
		return c == '0' || c == '1'
	}
	var q, err = mod3.NewMod3FSM().ProcessInput(s, isValidRune)
	if err != nil {
		return 0, err
	}
	return int(q), nil
}

func TestModThreeError(t *testing.T) {
	table := []string{
		"10101019",
	}
	for _, tt := range table {
		_, err := ModThree(tt)
		if err == nil {
			t.Errorf("ModThree(%q) = nil, want error", tt)
		}
	}
}

func TestModThreeString(t *testing.T) {
	table := []string{
		"0",
		"1",
		"10",
		"11",
		"101",
		"10101",
	}
	for _, tt := range table {
		want := bStringToInt(tt) % 3
		got, err := ModThree(tt)
		if err != nil {
			t.Errorf("ModThree(%q) failed with error %v", tt, err)
		}
		if got != want {
			t.Errorf("ModThree(%q) = %d, want %d", tt, got, want)
		}
	}
}

func TestModThreeInt(t *testing.T) {
	table := []int{
		0,
		1,
		2,
		3,
		5,
	}
	for _, tt := range table {
		got, err := ModThree(intToBString(tt))
		if err != nil {
			t.Errorf("ModThree(%q) failed with error %v", tt, err)
		}
		want := tt % 3
		if got != want {
			t.Errorf("ModThree(%q) = %d, want %d", tt, got, want)
		}
	}
	for tt := 0; tt < 100; tt++ {
		got, err := ModThree(intToBString(tt))
		if err != nil {
			t.Errorf("ModThree(%q) failed with error %v", tt, err)
		}
		want := tt % 3
		if got != want {
			t.Errorf("ModThree(%q) = %d, want %d", tt, got, want)
		}
	}
}

func intToBString(n int) string {
	binaryStr := fmt.Sprintf("%b", n)
	return binaryStr
}

// bStringToInt converts a string of 0s and 1s to an integer.
// Any non '0' or '1' characters are skipped.
func bStringToInt(s string) int {
	var n int
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '0' {
			n = n << 1
		} else if c == '1' {
			n = (n << 1) + 1
		} else {
			// ignore any non-binary character
			continue
		}
	}
	return n
}

func TestBStringToInt(t *testing.T) {
	table := []struct {
		in   string
		want int
	}{
		{"0", 0},
		{"", 0},
		{"1", 1},
		{"10", 2},
		{"11", 3},
		{"101", 5},
		{"10101", 21},
	}
	for _, tt := range table {
		if got := bStringToInt(tt.in); got != tt.want {
			t.Errorf("bStringToInt(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}
