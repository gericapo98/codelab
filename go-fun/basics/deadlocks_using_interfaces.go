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
  AddStates(...State) error 
  AddTransitions(...Transition) error 
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
		states:        make([]State, 0),
		transitions:   make([]Transition,0),
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
* We need to be able to add states passed as variadic parameter 
* into the array of State which should be empty in the beginning.
*/

// a function that adds states to a lts should anticipate duplication and,
// skip any such occurrence
//
// Update: implement this in two functions. 
func (l *ltsImpl) AddState(s State) error {
	for _, existing := range l.states {
		if existing.ID() == s.ID() {
			return errors.New("duplicate state ID: " + s.ID())
		}
	}
	l.states = append(l.states, s)
	return nil
}
// adding multiple states using variadic parameter
func (l *ltsImpl) AddStates(sts ...State) error {
	for _, s := range sts {
		if err := l.AddState(s); err != nil {
			return err
		}
	}
	return nil
}
// only passing the value of the transition interface,
// not sure yet where to point for methods arguments
func (l *ltsImpl) AddTransition(t Transition) error {
	l.transitions = append(l.transitions, t)
	return nil
}
// be aware 
func (l *ltsImpl) AddTransitions(ts ...Transition) error {
	for _, t := range ts {
		if err := l.AddTransition(t); err != nil {
			return err
		}
	}
	return nil
}


type deadlock_checker bool {
  
}


