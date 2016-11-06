// Package binarysearch implements... binary search!
package binarysearch

import "fmt"

// SearchInts imitates the std lib's sort.SearchInts using binary search
func SearchInts(slice []int, key int) int {
	return searchInSlice(slice, 0, len(slice)-1, key)
}

func searchInSlice(slice []int, min, max, key int) int {
	if min > max {
		// return the index of the first value greater than the search key
		return min
	}
	mid := (min + max) / 2
	if slice[mid] == key {
		// find the first occurrence
		i := mid
		for ; i > 0 && slice[i-1] == slice[mid]; i-- {
		}
		return i
	}
	if slice[mid] < key {
		return searchInSlice(slice, mid+1, max, key)
	}
	return searchInSlice(slice, min, mid-1, key)
}

// Message returns a human-readable version of search result
func Message(slice []int, key int) string {

	l := len(slice)
	if l == 0 {
		return "slice has no values"
	}

	i := SearchInts(slice, key)

	switch {
	case i > l-1:
		return fmt.Sprintf("%d > all %d values", key, l)
	case slice[i] != key && i == 0:
		return fmt.Sprintf("%d < all values", key)
	case slice[i] != key:
		return fmt.Sprintf("%d > %d at index %d, < %d at index %d", key, slice[i-1], i-1, slice[i], i)
	case i == 0:
		return fmt.Sprintf("%d found at beginning of slice", key)
	case i == l-1:
		return fmt.Sprintf("%d found at end of slice", key)
	default:
		return fmt.Sprintf("%d found at index %d", key, i)
	}
}
