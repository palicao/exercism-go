// Package pascal allows you to build Pascal's Triangles
package pascal

// Triangle builds Pascal's Triangle of size n
func Triangle(n int) [][]int {
	triangle := [][]int{}
	for i := 0; i < n; i++ {
		tmp := []int{}
		for j := 0; j <= i; j++ {
			tmp = append(tmp, pascal(j, i))
		}
		triangle = append(triangle, tmp)
	}
	return triangle
}

// pascal returns the value of the triangle for a given row and column
func pascal(r, c int) int {
	if r == c || r == 0 || c == 0 {
		return 1
	}
	return pascal(r-1, c-1) + pascal(r, c-1)
}
