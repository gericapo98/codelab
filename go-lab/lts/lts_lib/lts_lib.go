/*
* TODO: The task is to implement a library for working with labeled transition systems (LTSs).
* The tasks following are unordered and will be completed based on the priority of the tasks. 1 being the highest priority and 10 being the lowest.
* Any task that is not marked by a number will be left to the end.
* methods to check if the LTS is deterministic, and to convert a non-deterministic LTS to a deterministic one
* methods to check if the LTS is complete, and to convert a non-complete LTS to a complete one
* 1.methods that return the set of states reachable from a given state, and the set of states that can reach a given state
* methods to check if the LTS is minimal, and to convert a non-minimal LTS to a minimal one
* methods to check if the LTS is aperiodic, and to convert a periodic LTS to an aperiodic one
* methods to check if the LTS is reversible, and to convert a non-reversible LTS to a reversible one
* methods to check if the LTS is synchronizable, and to convert a non-synchronizable LTS to a synchronizable one
* 2.methods to check if the LTS is bisimilar to another LTS, and to compute the bisimulation relation between two LTSs
* methods that return the set of states that are equivalent to a given state, and the set of states that are not equivalent to a given state
* methods to check if the LTS is deadlock-free, and to convert a LTS with deadlocks to a deadlock-free one
* Another important aspect is to write test cases for each of the methods implemented.
*/


package lts_lib

import (
    "errors"
    "fmt"
)

type State interface {
    ID() string
}

type Transition interface {
    Source() State
    Label() string
    Destination() State
}

type LTS interface {
    States() []State
    Transitions() []Transition
    InitialStates() []State
    AddState(State) error
    AddTransition(Transition) error
    AddStates(...State) error
    AddTransitions(...Transition) error
    PrintStates()
    PrintTransitions()
}

type ltsImpl struct {
    states        []State
    transitions   []Transition
    initialStates []State
}

func NewLTS(initial ...State) LTS {
    return &ltsImpl{
        states:        make([]State, 0),
        transitions:   make([]Transition, 0),
        initialStates: initial,
    }
}

func (l *ltsImpl) States() []State {
    return l.states
}

func (l *ltsImpl) Transitions() []Transition {
    return l.transitions
}

func (l *ltsImpl) InitialStates() []State {
    return l.initialStates
}

func (l *ltsImpl) AddState(s State) error {
    for _, existing := range l.states {
        if existing.ID() == s.ID() {
            return errors.New("duplicate state ID: " + s.ID())
        }
    }
    l.states = append(l.states, s)
    return nil
}

func (l *ltsImpl) AddStates(sts ...State) error {
    for _, s := range sts {
        if err := l.AddState(s); err != nil {
            return err
        }
    }
    return nil
}

func (l *ltsImpl) AddTransition(t Transition) error {
    l.transitions = append(l.transitions, t)
    return nil
}

func (l *ltsImpl) AddTransitions(ts ...Transition) error {
    for _, t := range ts {
        if err := l.AddTransition(t); err != nil {
            return err
        }
    }
    return nil
}

func (l *ltsImpl) PrintStates() {
    fmt.Println("States in the LTS:")
    for _, state := range l.states {
        fmt.Printf(" - %s\n", state.ID())
    }
}

func (l *ltsImpl) PrintTransitions() {
    fmt.Println("Transitions in the LTS:")
    for _, t := range l.transitions {
        fmt.Printf(" - %s --[%s]--> %s\n", t.Source().ID(), t.Label(), t.Destination().ID())
    }
}

func TraverseAndPrintBFS(l LTS, start State) {
    visited := make(map[string]bool)
    queue := []State{start}
    visited[start.ID()] = true

    fmt.Printf("BFS traversal beginning at state: %s\n", start.ID())

    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        fmt.Printf(" - Visited state: %s\n", current.ID())

        for _, tr := range l.Transitions() {
            if tr.Source().ID() == current.ID() {
                fmt.Printf("    via '%s' -> %s\n", tr.Label(), tr.Destination().ID())

                if !visited[tr.Destination().ID()] {
                    visited[tr.Destination().ID()] = true
                    queue = append(queue, tr.Destination())
                }
            }
        }
    }
}

