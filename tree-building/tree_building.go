// Package tree builds a tree of records
package tree

import (
	"errors"
	"sort"
)

const testVersion = 3

/*
Benchmark:

- Original version:
BenchmarkTwoTree-2    	       1	20123595538 ns/op
BenchmarkTenTree-2    	       3	 470625948 ns/op
BenchmarkShallowTree-2	       1	1817715881 ns/op
ok  	_/Users/palicao/exercism/go/tree-building	24.803s

- Modified version:
BenchmarkTwoTree-2    	      50	  26985974 ns/op
BenchmarkTenTree-2    	     500	   3318511 ns/op
BenchmarkShallowTree-2	     500	   2824984 ns/op
ok  	_/Users/palicao/exercism/go/tree-building	5.972s

*/

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

// Sorting records by ID utils
type Records []Record

// Len calculates the length Records
func (r Records) Len() int { return len(r) }

// Swap swaps 2 elements
func (r Records) Swap(i, j int) { r[i], r[j] = r[j], r[i] }

// Less compares 2 elements
func (r Records) Less(i, j int) bool { return r[i].ID < r[j].ID }

// Build builds the tree of records and returns a pointer to the root node
func Build(records []Record) (*Node, error) {
	l := len(records)
	if l == 0 {
		return nil, nil
	}

	sort.Sort(Records(records))
	nodes := make([]Node, l)

	for i, r := range records {
		if r.ID >= l {
			return nil, errors.New("The ID must be between 0 (inclusive) and the length of the list (exclusive)")
		}

		if i == 0 && (r.ID != 0 || r.Parent != 0) {
			return nil, errors.New("Root node must have Parent = 0 and ID = 0")
		}

		if i != 0 && r.ID <= r.Parent {
			return nil, errors.New("ID of a non-root node cannot be <= Parent")
		}

		nodes[i] = Node{ID: r.ID}
		if r.ID != 0 {
			nodes[r.Parent].Children = append(nodes[r.Parent].Children, &nodes[i])
		}
	}
	return &nodes[0], nil
}
