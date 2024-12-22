package deadlocks_using_interfaces

import (
  "container/list"
  "errors"
  "fmt"
  "strings"
)

type State interface {
  ID() string
}

type Transitions interface {
  Source() State 
  Label() string
  Destination() State
}

type LTS interface {
  States() []State 
  Transitions() []Transition
  InitialStates() []State 
  AddState(State)
  AddTransition(Transition)
}

type BisimulationCheck interface {
  CheckBisimulation(LTS,LTS) bool
}
// https://dev.to/envitab/function-signatures-in-go-38ja
type ltsImpl struct {
	states        []State
	transitions   []Transition
	initialStates []State
}

func NewLTS(initial ...State) LTS {
	return &ltsImpl{
		states:        []State{},
		transitions:   []Transition{},
		initialStates: initial,
	}
}

type deadlock_checker bool {
  
}


