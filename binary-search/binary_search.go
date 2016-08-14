// Package binarysearch implements... binary search!
package binarysearch

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
