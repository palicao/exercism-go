// Package binarysearchtree implements a binary search tree!
package binarysearchtree

type SearchTreeData struct {
	data  int
	left  *SearchTreeData
	right *SearchTreeData
}

// Bst creates a new SearchTreeData given an int
func Bst(data int) *SearchTreeData {
	return &SearchTreeData{data: data}
}

// Insert inserts a new element in the tree
func (s *SearchTreeData) Insert(data int) {
	node := &s.left
	if data > s.data {
		node = &s.right
	}
	if *node == nil {
		*node = Bst(data)
	} else {
		(*node).Insert(data)
	}
}

// MapString returns the ordered tree as array of strings applying a func to each element
func (s *SearchTreeData) MapString(f func(int) string) []string {
	ints := s.MapInt(func(i int) int { return i })
	out := make([]string, 0)
	for _, i := range ints {
		out = append(out, f(i))
	}
	return out
}

// MapInt returns the ordered tree as array of ints applying a func to each element
func (s *SearchTreeData) MapInt(f func(int) int) []int {
	if s == nil {
		return make([]int, 0)
	}
	out := []int{f(s.data)}
	out = append(s.left.MapInt(f), out...)
	return append(out, s.right.MapInt(f)...)
}
