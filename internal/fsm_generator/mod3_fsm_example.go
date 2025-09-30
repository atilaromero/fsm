package fsm_generator

type Mod3State int

const (
	S0 Mod3State = 0
	S1 Mod3State = 1
	S2 Mod3State = 2
)

type Mod3Alphabet rune

const (
	A0 Mod3Alphabet = '0'
	A1 Mod3Alphabet = '1'
)

func (a Mod3Alphabet) isValid(c rune) bool {
	return c == '0' || c == '1'
}

func tf(q Mod3State, a Mod3Alphabet) Mod3State {
	switch q {
	case S0:
		if a == A0 {
			return S0
		}
		return S1
	case S1:
		if a == A0 {
			return S2
		}
		return S0
	case S2:
		if a == A0 {
			return S1
		}
		return S2
	}
	return q
}

func NewMod3FSM() *FSM[Mod3State, Mod3Alphabet] {
	return NewFSM(S0, []Mod3State{S0, S1, S2}, tf)
}
