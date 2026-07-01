//! Edmonds' Blossom algorithm for maximum cardinality matching in a general
//! (not necessarily bipartite) graph.
//!
//! The difficulty of the algorithm comes from "blossoms": odd-length cycles
//! that appear while searching for an augmenting path. A naive alternating-path
//! search (which works fine for bipartite graphs) can get stuck on these odd
//! cycles, so Edmonds' insight was to contract each blossom into a single
//! super-vertex, continue the search on the contracted graph, and later lift
//! the contraction when reconstructing the augmenting path.
//!
//! This is the classic O(V^3) adjacency-list implementation built around an
//! alternating BFS forest. Run `cargo run` for a demo and `cargo test` for the
//! test suite (which includes a randomized comparison against a brute-force
//! oracle).

use std::collections::VecDeque;

/// An undirected graph on vertices `0..n`.
pub struct Graph {
    n: usize,
    adj: Vec<Vec<usize>>,
}

impl Graph {
    /// Creates a graph with `n` vertices and no edges.
    pub fn new(n: usize) -> Self {
        Graph {
            n,
            adj: vec![Vec::new(); n],
        }
    }

    /// Adds an undirected edge between `u` and `v`. Self-loops, out-of-range
    /// endpoints, and duplicate edges are ignored, so it is safe to call
    /// repeatedly.
    pub fn add_edge(&mut self, u: usize, v: usize) {
        if u == v || u >= self.n || v >= self.n {
            return;
        }
        if self.adj[u].contains(&v) {
            return;
        }
        self.adj[u].push(v);
        self.adj[v].push(u);
    }

    /// Computes a maximum cardinality matching.
    ///
    /// Returns a vector `m` where `m[v]` is the vertex matched with `v`, or
    /// `None` if `v` is unmatched.
    pub fn max_matching(&self) -> Vec<Option<usize>> {
        let mut s = State::new(self);
        for v in 0..self.n {
            if s.matched[v].is_none() {
                if let Some(u) = s.find_augmenting_path(v) {
                    s.augment(u);
                }
            }
        }
        s.matched
    }
}

/// Mutable scratch data used by a single maximum-matching run.
struct State<'a> {
    g: &'a Graph,
    matched: Vec<Option<usize>>,
    parent: Vec<Option<usize>>,
    base: Vec<usize>,
    used: Vec<bool>,
    blossom: Vec<bool>,
    queue: VecDeque<usize>,
}

impl<'a> State<'a> {
    fn new(g: &'a Graph) -> Self {
        State {
            g,
            matched: vec![None; g.n],
            parent: vec![None; g.n],
            base: (0..g.n).collect(),
            used: vec![false; g.n],
            blossom: vec![false; g.n],
            queue: VecDeque::new(),
        }
    }

    /// Finds the lowest common ancestor of `a` and `b` in the alternating
    /// forest, walking up through blossom bases. Used to locate the base of a
    /// newly discovered blossom.
    fn lca(&self, mut a: usize, mut b: usize) -> usize {
        let mut seen = vec![false; self.g.n];
        loop {
            a = self.base[a];
            seen[a] = true;
            match self.matched[a] {
                None => break,
                Some(m) => a = self.parent[m].unwrap(),
            }
        }
        loop {
            b = self.base[b];
            if seen[b] {
                return b;
            }
            b = self.parent[self.matched[b].unwrap()].unwrap();
        }
    }

    /// Walks from `v` up to the blossom base `b`, flagging every base on the way
    /// as part of the blossom and rerouting parent pointers so the contracted
    /// odd cycle can be traversed in both directions.
    fn mark_path(&mut self, mut v: usize, b: usize, mut child: usize) {
        while self.base[v] != b {
            self.blossom[self.base[v]] = true;
            let mv = self.matched[v].unwrap();
            self.blossom[self.base[mv]] = true;
            self.parent[v] = Some(child);
            child = mv;
            v = self.parent[mv].unwrap();
        }
    }

    /// Runs an alternating BFS rooted at the unmatched vertex `root`,
    /// contracting blossoms as they are found. Returns the far endpoint of an
    /// augmenting path if one exists.
    fn find_augmenting_path(&mut self, root: usize) -> Option<usize> {
        for i in 0..self.g.n {
            self.used[i] = false;
            self.parent[i] = None;
            self.base[i] = i;
        }
        self.used[root] = true;
        self.queue.clear();
        self.queue.push_back(root);

        while let Some(v) = self.queue.pop_front() {
            for idx in 0..self.g.adj[v].len() {
                let to = self.g.adj[v][idx];
                // Ignore edges inside the same blossom and the matched edge.
                if self.base[v] == self.base[to] || self.matched[v] == Some(to) {
                    continue;
                }
                let is_outer =
                    to == root || self.matched[to].is_some_and(|m| self.parent[m].is_some());
                if is_outer {
                    // Found a blossom: v and to are both outer vertices.
                    let cur_base = self.lca(v, to);
                    for b in self.blossom.iter_mut() {
                        *b = false;
                    }
                    self.mark_path(v, cur_base, to);
                    self.mark_path(to, cur_base, v);
                    for i in 0..self.g.n {
                        if self.blossom[self.base[i]] {
                            self.base[i] = cur_base;
                            if !self.used[i] {
                                self.used[i] = true;
                                self.queue.push_back(i);
                            }
                        }
                    }
                } else if self.parent[to].is_none() {
                    self.parent[to] = Some(v);
                    match self.matched[to] {
                        None => return Some(to), // augmenting path ends here
                        Some(m) => {
                            self.used[m] = true;
                            self.queue.push_back(m);
                        }
                    }
                }
            }
        }
        None
    }

    /// Flips matched and unmatched edges along the augmenting path ending at
    /// `u`, increasing the matching size by one.
    fn augment(&mut self, mut u: usize) {
        loop {
            let pv = self.parent[u].unwrap();
            let ppv = self.matched[pv];
            self.matched[u] = Some(pv);
            self.matched[pv] = Some(u);
            match ppv {
                Some(next) => u = next,
                None => break,
            }
        }
    }
}

/// Returns the number of edges in the matching described by `m` (as returned by
/// [`Graph::max_matching`]).
pub fn size(m: &[Option<usize>]) -> usize {
    m.iter().filter(|x| x.is_some()).count() / 2
}

fn main() {
    // Demo: the Petersen graph, which has a perfect matching of 5 edges.
    let mut g = Graph::new(10);
    for i in 0..5 {
        g.add_edge(i, (i + 1) % 5); // outer 5-cycle
        g.add_edge(i, i + 5); // spokes
        g.add_edge(5 + i, 5 + (i + 2) % 5); // inner pentagram
    }
    let m = g.max_matching();
    println!("Petersen graph maximum matching: {} edges", size(&m));
    for v in 0..m.len() {
        if let Some(w) = m[v] {
            if w > v {
                println!("  {v} - {w}");
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    /// Checks that `m` is a symmetric, non-overlapping matching that only uses
    /// edges present in `g`.
    fn assert_valid(g: &Graph, m: &[Option<usize>]) {
        for v in 0..g.n {
            if let Some(w) = m[v] {
                assert_eq!(m[w], Some(v), "matching not symmetric at {v}");
                assert!(
                    g.adj[v].contains(&w),
                    "matched pair ({v},{w}) is not an edge"
                );
            }
        }
    }

    #[test]
    fn empty() {
        let g = Graph::new(0);
        assert_eq!(size(&g.max_matching()), 0);
    }

    #[test]
    fn single_edge() {
        let mut g = Graph::new(2);
        g.add_edge(0, 1);
        let m = g.max_matching();
        assert_valid(&g, &m);
        assert_eq!(size(&m), 1);
    }

    /// The classic case that breaks a naive bipartite-style search: an odd
    /// cycle (blossom). Maximum matching is 1.
    #[test]
    fn triangle() {
        let mut g = Graph::new(3);
        g.add_edge(0, 1);
        g.add_edge(1, 2);
        g.add_edge(2, 0);
        let m = g.max_matching();
        assert_valid(&g, &m);
        assert_eq!(size(&m), 1);
    }

    /// The Petersen graph has a perfect matching (5 edges).
    #[test]
    fn petersen() {
        let mut g = Graph::new(10);
        for i in 0..5 {
            g.add_edge(i, (i + 1) % 5); // outer 5-cycle
            g.add_edge(i, i + 5); // spokes
            g.add_edge(5 + i, 5 + (i + 2) % 5); // inner pentagram
        }
        let m = g.max_matching();
        assert_valid(&g, &m);
        assert_eq!(size(&m), 5);
    }

    /// Two triangles joined by an edge; a perfect matching of size 3 exists.
    #[test]
    fn blossom_chain() {
        let mut g = Graph::new(6);
        g.add_edge(0, 1);
        g.add_edge(1, 2);
        g.add_edge(2, 0);
        g.add_edge(3, 4);
        g.add_edge(4, 5);
        g.add_edge(5, 3);
        g.add_edge(2, 3);
        let m = g.max_matching();
        assert_valid(&g, &m);
        assert_eq!(size(&m), 3);
    }

    /// Exhaustive maximum-matching oracle for small graphs.
    fn brute_force(g: &Graph) -> usize {
        let mut edges = Vec::new();
        for u in 0..g.n {
            for &v in &g.adj[u] {
                if u < v {
                    edges.push((u, v));
                }
            }
        }
        let mut used = vec![false; g.n];
        let mut best = 0;
        fn rec(
            edges: &[(usize, usize)],
            i: usize,
            count: usize,
            used: &mut [bool],
            best: &mut usize,
        ) {
            if count > *best {
                *best = count;
            }
            for j in i..edges.len() {
                let (u, v) = edges[j];
                if !used[u] && !used[v] {
                    used[u] = true;
                    used[v] = true;
                    rec(edges, j + 1, count + 1, used, best);
                    used[u] = false;
                    used[v] = false;
                }
            }
        }
        rec(&edges, 0, 0, &mut used, &mut best);
        best
    }

    /// A tiny deterministic xorshift PRNG so the test needs no dependencies.
    struct Rng(u64);
    impl Rng {
        fn next(&mut self) -> u64 {
            let mut x = self.0;
            x ^= x << 13;
            x ^= x >> 7;
            x ^= x << 17;
            self.0 = x;
            x
        }
        fn below(&mut self, n: u64) -> u64 {
            self.next() % n
        }
    }

    /// Compare against the exhaustive oracle on many small random graphs.
    #[test]
    fn random_against_brute_force() {
        let mut rng = Rng(0x9E3779B97F4A7C15);
        for _ in 0..500 {
            let n = 1 + rng.below(9) as usize; // 1..=9 vertices
            let mut g = Graph::new(n);
            for u in 0..n {
                for v in (u + 1)..n {
                    if rng.below(100) < 40 {
                        g.add_edge(u, v);
                    }
                }
            }
            let m = g.max_matching();
            assert_valid(&g, &m);
            assert_eq!(size(&m), brute_force(&g), "n={n}");
        }
    }
}
