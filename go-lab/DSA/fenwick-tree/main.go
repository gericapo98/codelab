package main

import "fmt"


type FenwickTree struct {
	tree []int
}


func NewFenwickTree(n int) *FenwickTree {
	return &FenwickTree{
		tree: make([]int, n+1),
	}
}


func (ft *FenwickTree) Update(idx, val int) {
	for idx < len(ft.tree) {
		ft.tree[idx] += val
		idx += idx & -idx
	}
}


func (ft *FenwickTree) Query(idx int) int {
	sum := 0
	for idx > 0 {
		sum += ft.tree[idx]
		idx -= idx & -idx
	}
	return sum
}


func (ft *FenwickTree) RangeQuery(left, right int) int {
	return ft.Query(right) - ft.Query(left-1)
}

func main() {

	n := 10
	ft := NewFenwickTree(n)


	ft.Update(1, 5)
	ft.Update(2, 3)
	ft.Update(3, 7)
	ft.Update(4, 6)
	ft.Update(5, 5)


	fmt.Println("Sum of first 5 elements:", ft.Query(5))  // Output: 26
	fmt.Println("Sum of first 3 elements:", ft.Query(3))  // Output: 15

	
	fmt.Println("Sum from index 2 to 4:", ft.RangeQuery(2, 4)) 
}

