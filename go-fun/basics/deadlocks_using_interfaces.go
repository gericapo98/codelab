package deadlocks_using_interfaces

import (
  "container/list"
  "errors"
  "fmt"
  "strings"
)

type State interface {
  ID() String
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

type deadlock_checker bool {
  
}
func main(){
}
