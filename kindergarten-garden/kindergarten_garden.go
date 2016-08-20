// Package kindergarten tells you which plants each child in the kindergarten class is responsible for
package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

type Garden map[string][]string

var plantNames = map[rune]string{
	'V': "violets",
	'R': "radishes",
	'G': "grass",
	'C': "clover",
}

// NewGarden constructs a new Garden given a diagram and a list of children
func NewGarden(diagram string, children []string) (*Garden, error) {
	numChildren := len(children)

	diagramRows, err := parseDiagram(diagram, numChildren)
	if err != nil {
		return nil, err
	}

	sortedChildren := sortStrings(children)

	g := make(map[string][]string, numChildren)
	for i := 0; i < numChildren; i++ {
		plants := make([]string, 4)

		for j := 0; j < 4; j++ {
			var ok bool
			plants[j], ok = getPlantName(diagramRows, i, j)
			if !ok {
				return nil, errors.New("Unkown plant")
			}
		}
		if _, ok := g[sortedChildren[i]]; ok {
			return nil, errors.New("Duplicated children")
		}
		g[sortedChildren[i]] = plants
	}
	garden := Garden(g)
	return &garden, nil
}

// Plants tells you which plants a child is responsible for
func (g *Garden) Plants(child string) ([]string, bool) {
	gmap := map[string][]string(*g)
	p, ok := gmap[child]
	return p, ok
}

func parseDiagram(diagram string, numChildren int) ([]string, error) {
	diagramRows := strings.Split(diagram, "\n")
	if len(diagramRows) != 3 {
		return nil, errors.New("Incorrect diagram")
	}

	// discard first row
	diagramRows = diagramRows[1:]

	diagramRowLen := numChildren * 2
	if len(diagramRows[0]) != diagramRowLen || len(diagramRows[1]) != diagramRowLen {
		return nil, errors.New("Mismatched rows")
	}

	return diagramRows, nil
}

func sortStrings(a []string) []string {
	c := make([]string, len(a))
	copy(c, a)
	sort.Strings(c)
	return c
}

func getPlantName(diagramRows []string, childId, plantId int) (string, bool) {
	row := plantId / 2
	col := childId*2 + plantId%2
	plantCode := rune(diagramRows[row][col])
	name, ok := plantNames[plantCode]
	return name, ok
}
