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

