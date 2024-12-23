package main
import (
    "fmt"
    "lts_lib"
)

type MyState struct {
    name string
}
func (m MyState) ID() string {
    return m.name
}

type MyTransition struct {
    src lts_lib.State
    lbl string
    dst lts_lib.State
}

func (t MyTransition) Source() lts_lib.State       { return t.src }
func (t MyTransition) Label() string             { return t.lbl }
func (t MyTransition) Destination() lts_lib.State { return t.dst }

func main() {
    lts := lts_lib.NewLTS()
    s1 := MyState{name: "S1"}
    s2 := MyState{name: "S2"}
    s3 := MyState{name: "S3"}
    _ = lts.AddStates(s1, s2, s3)

    t1 := MyTransition{src: s1, lbl: "go", dst: s2}
    t2 := MyTransition{src: s2, lbl: "proceed", dst: s3}

    _ = lts.AddTransitions(t1, t2)
    lts_lib.TraverseAndPrintBFS(lts, s1)

    fmt.Println("Main done.")
}

