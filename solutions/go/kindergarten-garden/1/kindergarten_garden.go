package kindergarten

import (
	"errors"
	"slices"
	"strings"
)

type Plant struct {
	symbol rune
	owner string
}

// Define the Garden type here.
type Garden struct {
	windows string
	firstRow []Plant
	secondRow []Plant
}
// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`


func NewGarden(diagram string, children []string) (*Garden, error) {
	var firstRow []Plant
	var secondRow []Plant
	childrenCopy := slices.Clone(children)
	slices.Sort(childrenCopy)
	validPlants := []string{"G", "C", "R", "V"}
	diagramSlice := strings.Split(diagram, "\n")
	if len(diagramSlice) != 3 || len(diagramSlice[1]) != len(diagramSlice[2]) || len(diagramSlice[1]) % 2 != 0 {
		return nil, errors.New("Wrong diagram format.")
	}
	if len(slices.Compact(childrenCopy)) != len(childrenCopy) {
		return nil, errors.New("Can't have duplicate names.")
	}
	plantsByChild := 0
	childIndex := 0
	for i := 0; i < len(diagramSlice[1]); i++ {
		if plantsByChild == 2 {
			plantsByChild = 0
			childIndex++
		}
		if !slices.Contains(validPlants, string(diagramSlice[1][i])) || !slices.Contains(validPlants, string(diagramSlice[2][i])) {
			return nil, errors.New("Invalid plant.")
		}
		firstRow = append(firstRow, Plant{rune(diagramSlice[1][i]), childrenCopy[childIndex]})
		secondRow = append(secondRow, Plant{rune(diagramSlice[2][i]), childrenCopy[childIndex]})
		plantsByChild++
	}
	resGarden := Garden{"[window]", firstRow, secondRow}
	return &resGarden, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	var res []string
	plantType := map[rune]string{'G': "grass", 'C': "clover", 'R': "radishes", 'V': "violets"}
	for i := 0; i < len(g.firstRow); i++ {
		if g.firstRow[i].owner == child {
			res = append(res, plantType[g.firstRow[i].symbol])
		}
	}
	for i := 0; i < len(g.secondRow); i++ {
		if g.secondRow[i].owner == child {
			res = append(res, plantType[g.secondRow[i].symbol])
		}
	}
	if res == nil {
		return nil, false
	}
	return res, true
}
