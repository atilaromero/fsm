package fsm_test

import (
	"fmt"

	fsm "github.com/atilaromero/fsm"
)

func ExampleFSM() {

	isValidRune := func(c rune) bool {
		return c == '0' || c == '1'
	}

	type MyState int
	const (
		MyState0 = 0
		MyState1 = 1
		MyState2 = 2
	)

	type MyAlphabet rune
	const (
		MyAlphabet0 = '0'
		MyAlphabet1 = '1'
	)

	tf := func(q MyState, a MyAlphabet) MyState {
		switch q {
		case MyState0:
			if a == MyAlphabet0 {
				return MyState0
			}
			return MyState1
		case MyState1:
			if a == MyAlphabet0 {
				return MyState2
			}
			return MyState0
		case MyState2:
			if a == MyAlphabet0 {
				return MyState1
			}
			return MyState2
		}
		return q
	}

	mod3FSM := fsm.NewFSM(MyState0, []MyState{0, 1, 2}, tf)

	fmt.Println(mod3FSM.ProcessInput("x", isValidRune)) // checking invalid input
	fmt.Println(mod3FSM.ProcessInput("0", isValidRune))
	fmt.Println(mod3FSM.ProcessInput("1", isValidRune))
	fmt.Println(mod3FSM.ProcessInput("10", isValidRune))
	fmt.Println(mod3FSM.ProcessInput("11", isValidRune))
	fmt.Println(mod3FSM.ProcessInput("100", isValidRune))
	fmt.Println(mod3FSM.ProcessInput("101", isValidRune))
	fmt.Println(mod3FSM.ProcessInput("110", isValidRune))

	// Output:
	// 0 FSM.ProcessInput: invalid alphabet x
	// 0 <nil>
	// 1 <nil>
	// 2 <nil>
	// 0 <nil>
	// 1 <nil>
	// 2 <nil>
	// 0 <nil>
}
