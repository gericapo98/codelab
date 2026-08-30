package main

import (
	"math/rand"
	"testing"
)

// validMatching checks that m is a symmetric, non-overlapping matching that
// only uses edges present in g.
func validMatching(t *testing.T, g *Graph, m []int) {
	t.Helper()
	for v := 0; v < g.n; v++ {
		w := m[v]
		if w == -1 {
			continue
		}
		if m[w] != v {
			t.Fatalf("matching not symmetric: m[%d]=%d but m[%d]=%d", v, w, w, m[w])
		}
		found := false
		for _, x := range g.adj[v] {
			if x == w {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("matched pair (%d,%d) is not an edge", v, w)
		}
	}
}

func TestEmpty(t *testing.T) {
	g := NewGraph(0)
	if got := Size(g.MaxMatching()); got != 0 {
		t.Fatalf("empty graph: got %d, want 0", got)
	}
}

func TestSingleEdge(t *testing.T) {
	g := NewGraph(2)
	g.AddEdge(0, 1)
	m := g.MaxMatching()
	validMatching(t, g, m)
	if got := Size(m); got != 1 {
		t.Fatalf("single edge: got %d, want 1", got)
	}
}

// TestTriangle is the classic case that breaks a naive bipartite-style search:
// an odd cycle (blossom). Maximum matching is 1.
func TestTriangle(t *testing.T) {
	g := NewGraph(3)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	m := g.MaxMatching()
	validMatching(t, g, m)
	if got := Size(m); got != 1 {
		t.Fatalf("triangle: got %d, want 1", got)
	}
}

// TestPetersen: the Petersen graph has a perfect matching (5 edges).
func TestPetersen(t *testing.T) {
	g := NewGraph(10)
	// Outer 5-cycle.
	for i := 0; i < 5; i++ {
		g.AddEdge(i, (i+1)%5)
	}
	// Spokes.
	for i := 0; i < 5; i++ {
		g.AddEdge(i, i+5)
	}
	// Inner pentagram.
	for i := 0; i < 5; i++ {
		g.AddEdge(5+i, 5+(i+2)%5)
	}
	m := g.MaxMatching()
	validMatching(t, g, m)
	if got := Size(m); got != 5 {
		t.Fatalf("Petersen: got %d, want 5 (perfect matching)", got)
	}
}

// TestBlossomChain builds a graph with adjacent odd cycles connected by an
// edge, exercising repeated blossom contraction.
func TestBlossomChain(t *testing.T) {
	// Two triangles 0-1-2 and 3-4-5 joined by edge 2-3.
	g := NewGraph(6)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(5, 3)
	g.AddEdge(2, 3)
	m := g.MaxMatching()
	validMatching(t, g, m)
	// A perfect matching exists: (0,1),(2,3),(4,5).
	if got := Size(m); got != 3 {
		t.Fatalf("blossom chain: got %d, want 3", got)
	}
}

// bruteForce computes the maximum matching size by exhaustive search. Only
// usable for small graphs; used as an oracle for randomized testing.
func bruteForce(g *Graph) int {
	type edge struct{ u, v int }
	var edges []edge
	for u := 0; u < g.n; u++ {
		for _, v := range g.adj[u] {
			if u < v {
				edges = append(edges, edge{u, v})
			}
		}
	}
	best := 0
	used := make([]bool, g.n)
	var rec func(i, count int)
	rec = func(i, count int) {
		if count > best {
			best = count
		}
		for j := i; j < len(edges); j++ {
			e := edges[j]
			if !used[e.u] && !used[e.v] {
				used[e.u], used[e.v] = true, true
				rec(j+1, count+1)
				used[e.u], used[e.v] = false, false
			}
		}
	}
	rec(0, 0)
	return best
}

// TestRandomAgainstBruteForce compares the algorithm against an exhaustive
// oracle on many small random graphs.
func TestRandomAgainstBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	for trial := 0; trial < 500; trial++ {
		n := 1 + rng.Intn(9) // 1..9 vertices
		g := NewGraph(n)
		for u := 0; u < n; u++ {
			for v := u + 1; v < n; v++ {
				if rng.Float64() < 0.4 {
					g.AddEdge(u, v)
				}
			}
		}
		m := g.MaxMatching()
		validMatching(t, g, m)
		got := Size(m)
		want := bruteForce(g)
		if got != want {
			t.Fatalf("trial %d (n=%d): got %d, want %d", trial, n, got, want)
		}
	}
}
