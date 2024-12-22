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
// 
type ltsImpl struct {
	states        []State
	transitions   []Transition
	initialStates []State
}
// https://www.digitalocean.com/community/tutorials/how-to-use-variadic-functions-in-go
// The initial states are designated as possible traversing points after creating the lts
// Will be changed accordingly...
func NewLTS(initial ...State) LTS {
	return &ltsImpl{
		states:        []State{},
		transitions:   []Transition{},
		initialStates: initial,
	}
}
func (l *ltsImpl) States() []State {
	return l.states
}

func (l *ltsImpl) Transitions() []Transition {
	return l.transitions
}

/*
* TODO::
* Now that we have a framework for lts.
* We need to be able to add states passed variadic parameter 
* into the array of State which should be empty in the beginning.
*/

type deadlock_checker bool {
  
}


