// Package allergies lists one person's allergies
package allergies

var allergens = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

// Allergies returns the list of allergies given a person's score
func Allergies(score int) []string {
	list := make([]string, 0)
	for i, a := range allergens {
		flag := 1<<uint(i)
		if score&flag == flag {
			list = append(list, a)
		}
	}
	return list
}

// AllergicTo returns true if the person is allergic to a specific allergen
func AllergicTo(score int, allergen string) bool {
	for i, a := range allergens {
		if a == allergen {
			flag := 1<<uint(i)
			return score&flag == flag
		}
	}
	return false
}
