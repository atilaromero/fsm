package modthree

import (
	"fmt"
	"testing"
)

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
		if got := ModThree(tt); got != want {
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
		got := ModThree(intToBString(tt))
		want := tt % 3
		if got != want {
			t.Errorf("ModThree(%q) = %d, want %d", tt, got, want)
		}
	}
	for tt := 0; tt < 100; tt++ {
		got := ModThree(intToBString(tt))
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
