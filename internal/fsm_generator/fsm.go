package fsm_generator

import "fmt"

type QStates interface {
	~int
}

type Alphabet interface {
	~rune
	isValid(c rune) bool
}

type TransitionFunction[Q QStates, A Alphabet] func(Q, A) Q

type FSM[Q QStates, A Alphabet] struct {
	Q           Q
	A           A
	Q0          Q
	ValidFinalQ []Q
	tf          TransitionFunction[Q, A]
}

func NewFSM[Q QStates, A Alphabet](q0 Q, validFinalQ []Q, tf TransitionFunction[Q, A]) *FSM[Q, A] {
	return &FSM[Q, A]{
		Q0:          q0,
		ValidFinalQ: validFinalQ,
		tf:          tf,
	}
}

func (f FSM[Q, A]) ProcessInput(input string) (Q, error) {
	for _, c := range input {
		if !f.A.isValid(c) {
			return f.Q, fmt.Errorf("FSM.ProcessInput: invalid alphabet %c", c)
		}
		var a A = A(c)
		f.Q = f.tf(f.Q, a)
	}
	return f.Q, nil
}
