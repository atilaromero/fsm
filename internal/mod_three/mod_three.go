package modthree

import "fmt"

type State int

const (
	S0 State = 0
	S1 State = 1
	S2 State = 2
)

func ModThree(s string) (int, error) {
	state := S0
	for _, c := range s {
		if c != '0' && c != '1' {
			return 0, fmt.Errorf("ModThree: invalid character %c", c)
		}
		state = stateTransition(state, c)
	}
	return int(state), nil
}

func stateTransition(state State, c rune) State {
	switch state {
	case S0:
		if c == '0' {
			return S0
		}
		return S1
	case S1:
		if c == '0' {
			return S2
		}
		return S0
	case S2:
		if c == '0' {
			return S1
		}
		return S2
	}
	return state
}
