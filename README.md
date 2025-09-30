# FSM - Generic Finite State Machine Library

A type-safe, generic Go library for building and executing Finite State Machines (FSMs).

[![Go Reference](https://pkg.go.dev/badge/github.com/atilaromero/fsm.svg)](https://pkg.go.dev/github.com/atilaromero/fsm)

## Features

- **Type-safe generics** - Define custom state and alphabet types with compile-time safety
- **Simple API** - Create FSMs with just a transition function
- **Input validation** - Built-in alphabet validation during processing
- **Zero dependencies** - Pure Go standard library implementation
- **Well-documented** - Comprehensive examples and godoc comments

## Installation

```bash
go get github.com/atilaromero/fsm
```

## Quick Start

```go
package main

import (
    "fmt"
    fsm "github.com/atilaromero/fsm"
)

// Define your state type
type State int
const (
    S0 State = 0
    S1 State = 1
    S2 State = 2
)

// Define your alphabet type
type Alphabet rune
const (
    Zero Alphabet = '0'
    One  Alphabet = '1'
)

func main() {
    // Define transition function
    tf := func(q State, a Alphabet) State {
        switch q {
        case S0:
            if a == Zero { return S0 }
            return S1
        case S1:
            if a == Zero { return S2 }
            return S0
        case S2:
            if a == Zero { return S1 }
            return S2
        }
        return q
    }

    // Create FSM
    machine := fsm.NewFSM(S0, []State{S0, S1, S2}, tf)

    // Validate input
    isValid := func(c rune) bool {
        return c == '0' || c == '1'
    }

    // Process input
    finalState, err := machine.ProcessInput("101", isValid)
    if err != nil {
        panic(err)
    }
    fmt.Println(finalState) // Output: 2
}
```

## Theory

A finite automaton (FA) is formally defined as a 5-tuple **(Q, A, q₀, F, δ)** where:

- **Q** - Finite set of states
- **A** - Finite input alphabet
- **q₀ ∈ Q** - Initial state
- **F ⊆ Q** - Set of accepting/final states
- **δ: Q × A → Q** - Transition function

For any state q ∈ Q and symbol a ∈ A, we interpret δ(q, a) as the state to which the FSM transitions when in state q and receiving input a.

## API

### Types

```go
// QStates constraint for state types (must be ~int)
type QStates interface { ~int }

// Alphabet constraint for input types (must be ~rune)
type Alphabet interface { ~rune }

// TransitionFunction defines state transitions
type TransitionFunction[Q QStates, A Alphabet] func(Q, A) Q

// FSM represents a finite state machine
type FSM[Q QStates, A Alphabet] struct {
    Q           Q                        // Current state
    Q0          Q                        // Initial state
    ValidFinalQ []Q                      // Valid accepting states
    tf          TransitionFunction[Q, A] // Transition function
}
```

### Functions

```go
// NewFSM creates a new finite state machine
func NewFSM[Q QStates, A Alphabet](
    q0 Q,
    validFinalQ []Q,
    tf TransitionFunction[Q, A],
) *FSM[Q, A]

// ProcessInput processes a string through the FSM
func (f FSM[Q, A]) ProcessInput(
    input string,
    isValid func(rune) bool,
) (Q, error)
```

## Examples

### Modulo 3 Calculator

A complete example that calculates `n mod 3` for binary strings is available in [`examples/mod3/`](examples/mod3/).

```go
import "github.com/atilaromero/fsm/examples/mod3"

machine := mod3.NewMod3FSM()
result, _ := machine.ProcessInput("101", isValid) // 5 in binary
fmt.Println(result) // Output: 2 (5 mod 3 = 2)
```

Run the example:
```bash
go test ./examples/mod3 -v
```

## Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run examples
go test -run Example
```

## Documentation

View documentation locally:
```bash
go doc github.com/atilaromero/fsm
```

Or browse online at [pkg.go.dev](https://pkg.go.dev/github.com/atilaromero/fsm).

## Use Cases

- **Lexical analysis** - Tokenizing input streams
- **Pattern matching** - Regular expression engines
- **Protocol validation** - Network protocol state machines
- **Game logic** - Character or game state management
- **Input parsing** - Validating structured input

## License

MIT

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.