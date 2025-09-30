// Package fsm_generator provides a generic framework for building and executing
// Finite State Machines (FSMs) with type-safe states and alphabets.
// A finite automaton (FA) is a 5-tuple (Q,A,q0,validFinalQ, tf) where,
//
// Q is a finite set of states;
// A is a finite input alphabet;
// q0 ∈ Q is the initial state;
// validFinalQ ⊆ Q is the set of accepting/final states; and
// tf:Q×A→Q is the transition function.
// For any element q of Q and any symbol a∈A, we interpret tf(q,a) as the state to which the FA
// moves, if it is in state q and receives the input a.
package fsm

import "fmt"

// QStates is a constraint for state types in a finite state machine.
// States must have an underlying type of int.
type QStates interface {
	~int
}

// Alphabet is a constraint for input alphabet types in a finite state machine.
// Alphabets must have an underlying type of rune and provide validation.
type Alphabet interface {
	~rune
}

// TransitionFunction defines the state transition logic for a finite state machine.
// It takes the current state Q and an input symbol A, returning the next state Q.
type TransitionFunction[Q QStates, A Alphabet] func(Q, A) Q

// FSM represents a generic Finite State Machine with parameterized state and alphabet types.
type FSM[Q QStates, A Alphabet] struct {
	Q           Q                        // Current state
	Q0          Q                        // Initial state
	ValidFinalQ []Q                      // Set of valid accepting/final states
	tf          TransitionFunction[Q, A] // State transition function
}

// NewFSM creates and initializes a new finite state machine.
// Parameters:
//   - q0: the initial state
//   - validFinalQ: slice of valid accepting states
//   - tf: the transition function that defines state changes
func NewFSM[Q QStates, A Alphabet](q0 Q, validFinalQ []Q, tf TransitionFunction[Q, A]) *FSM[Q, A] {
	return &FSM[Q, A]{
		Q0:          q0,
		ValidFinalQ: validFinalQ,
		tf:          tf,
	}
}

// ProcessInput processes a string input through the FSM, transitioning states
// according to the transition function. Returns the final state and any error
// encountered during processing (e.g., invalid input character).
func (f FSM[Q, A]) ProcessInput(input string, isValid func(rune) bool) (Q, error) {
	for _, c := range input {
		// Validate each character against the alphabet
		if !isValid(c) {
			return f.Q, fmt.Errorf("FSM.ProcessInput: invalid alphabet %c", c)
		}
		// Convert rune to alphabet type and transition
		var a A = A(c)
		f.Q = f.tf(f.Q, a)
	}
	return f.Q, nil
}
