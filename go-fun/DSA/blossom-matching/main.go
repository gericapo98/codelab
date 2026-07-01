// Command blossom-matching demonstrates Edmonds' Blossom algorithm for
// computing a maximum cardinality matching in a general (not necessarily
// bipartite) graph.
//
// The difficulty of the algorithm comes from "blossoms": odd-length cycles that
// appear while searching for an augmenting path. A naive alternating-path search
// (which works fine for bipartite graphs) can get stuck on these odd cycles, so
// Edmonds' insight was to contract each blossom into a single super-vertex,
// continue the search on the contracted graph, and later lift the contraction
// when reconstructing the augmenting path.
//
// This is the classic O(V^3) adjacency-list implementation built around an
// alternating BFS forest.
package main

import "fmt"

// Graph is an undirected graph on vertices [0, N).
type Graph struct {
	n   int
	adj [][]int
}

// NewGraph returns a graph with n vertices and no edges.
func NewGraph(n int) *Graph {
	return &Graph{n: n, adj: make([][]int, n)}
}

// AddEdge adds an undirected edge between u and v. Self-loops and duplicate
// edges are ignored, so it is safe to call repeatedly.
func (g *Graph) AddEdge(u, v int) {
	if u == v || u < 0 || v < 0 || u >= g.n || v >= g.n {
		return
	}
	for _, w := range g.adj[u] {
		if w == v {
			return
		}
	}
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}

// state holds the mutable scratch data used by a single maximum-matching run.
type state struct {
	g       *Graph
	match   []int  // match[v] = vertex matched to v, or -1
	parent  []int  // parent in the alternating forest (for outer vertices)
	base    []int  // base[v] = base of the blossom currently containing v
	used    []bool // used[v] = v is an outer vertex in the forest
	blossom []bool // scratch flag: v lies on a freshly contracted blossom
	queue   []int  // BFS queue of outer vertices
}

// MaxMatching returns a maximum cardinality matching. The result is a slice
// where result[v] is the vertex matched with v, or -1 if v is unmatched. The
// number of matched edges is (count of non-negative entries) / 2.
func (g *Graph) MaxMatching() []int {
	s := &state{
		g:       g,
		match:   make([]int, g.n),
		parent:  make([]int, g.n),
		base:    make([]int, g.n),
		used:    make([]bool, g.n),
		blossom: make([]bool, g.n),
	}
	for i := range s.match {
		s.match[i] = -1
	}

	for v := 0; v < g.n; v++ {
		if s.match[v] == -1 {
			if u := s.findAugmentingPath(v); u != -1 {
				s.augment(u)
			}
		}
	}

	result := make([]int, g.n)
	copy(result, s.match)
	return result
}

// lca finds the lowest common ancestor of a and b in the alternating forest,
// walking up through blossom bases. It is used to locate the base of a newly
// discovered blossom.
func (s *state) lca(a, b int) int {
	seen := make([]bool, s.g.n)
	for {
		a = s.base[a]
		seen[a] = true
		if s.match[a] == -1 {
			break
		}
		a = s.parent[s.match[a]]
	}
	for {
		b = s.base[b]
		if seen[b] {
			return b
		}
		b = s.parent[s.match[b]]
	}
}

// markPath walks from v up to the blossom base b, flagging every base on the
// way as part of the blossom and rerouting parent pointers so that the
// contracted odd cycle can be traversed in both directions.
func (s *state) markPath(v, b, child int) {
	for s.base[v] != b {
		s.blossom[s.base[v]] = true
		s.blossom[s.base[s.match[v]]] = true
		s.parent[v] = child
		child = s.match[v]
		v = s.parent[s.match[v]]
	}
}

// findAugmentingPath runs an alternating BFS rooted at an unmatched vertex root,
// contracting blossoms as they are found. It returns the far endpoint of an
// augmenting path if one exists, or -1 otherwise.
func (s *state) findAugmentingPath(root int) int {
	g := s.g
	for i := 0; i < g.n; i++ {
		s.used[i] = false
		s.parent[i] = -1
		s.base[i] = i
	}
	s.used[root] = true
	s.queue = s.queue[:0]
	s.queue = append(s.queue, root)

	for len(s.queue) > 0 {
		v := s.queue[0]
		s.queue = s.queue[1:]

		for _, to := range g.adj[v] {
			// Ignore edges inside the same blossom and the matched edge itself.
			if s.base[v] == s.base[to] || s.match[v] == to {
				continue
			}
			if to == root || (s.match[to] != -1 && s.parent[s.match[to]] != -1) {
				// Found a blossom: v and to are both outer vertices.
				curBase := s.lca(v, to)
				for i := range s.blossom {
					s.blossom[i] = false
				}
				s.markPath(v, curBase, to)
				s.markPath(to, curBase, v)
				for i := 0; i < g.n; i++ {
					if s.blossom[s.base[i]] {
						s.base[i] = curBase
						if !s.used[i] {
							s.used[i] = true
							s.queue = append(s.queue, i)
						}
					}
				}
			} else if s.parent[to] == -1 {
				s.parent[to] = v
				if s.match[to] == -1 {
					// Augmenting path found, ending at the free vertex `to`.
					return to
				}
				// Extend the alternating tree through the matched edge.
				s.used[s.match[to]] = true
				s.queue = append(s.queue, s.match[to])
			}
		}
	}
	return -1
}

// augment flips matched and unmatched edges along the augmenting path ending
// at u, increasing the matching size by one.
func (s *state) augment(u int) {
	for u != -1 {
		pv := s.parent[u]
		ppv := s.match[pv]
		s.match[u] = pv
		s.match[pv] = u
		u = ppv
	}
}

// Size returns the number of edges in the matching described by m (as returned
// by MaxMatching).
func Size(m []int) int {
	count := 0
	for _, v := range m {
		if v != -1 {
			count++
		}
	}
	return count / 2
}

func main() {
	// Demo: the Petersen graph, which has a perfect matching of 5 edges.
	g := NewGraph(10)
	for i := 0; i < 5; i++ {
		g.AddEdge(i, (i+1)%5)     // outer 5-cycle
		g.AddEdge(i, i+5)         // spokes
		g.AddEdge(5+i, 5+(i+2)%5) // inner pentagram
	}
	m := g.MaxMatching()
	fmt.Printf("Petersen graph maximum matching: %d edges\n", Size(m))
	for v := 0; v < len(m); v++ {
		if w := m[v]; w > v {
			fmt.Printf("  %d - %d\n", v, w)
		}
	}
}
