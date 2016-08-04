// Package school is a small archiving program that stores students' names along with the grade that they are in.
package school

import "sort"

type School map[int]Grade

type Grade struct {
	grade    int
	students []string
}

// New creates a new School
func New() *School {
	return &School{}
}

// Add adds a student in a grade
func (s School) Add(student string, grade int) {
	g, ok := s[grade]
	if !ok {
		g = Grade{grade, make([]string, 0)}
	}
	g.students = append(g.students, student)
	sort.Strings(g.students)
	s[grade] = g
}

// Grade returns a specific grade
func (s School) Grade(g int) []string {
	return s[g].students
}

// Enrollment returns an ordered list of grades
func (s School) Enrollment() []Grade {
	var keys []int
	var grades []Grade
	for i := range s {
		keys = append(keys, i)
	}
	sort.Ints(keys)
	for _, i := range keys {
		grades = append(grades, s[i])
	}
	return grades
}
