package gdengine

import (
	"fmt"
)

func CreateRColumnIndexbyIndex(Index int, rcolumns []GDColumnIndex) GDColumnIndex {
	return CreateColumnIndex(Index, "", rcolumns)
}

func CreateRColumnIndexbyName(Name string, rcolumns []GDColumnIndex) GDColumnIndex {
	return CreateColumnIndex(-1, Name, rcolumns)
}

func CreateColumnIndexbyIndex(Index int) GDColumnIndex {
	return CreateColumnIndex(Index, "", nil)
}

func CreateColumnIndexbyName(Name string) GDColumnIndex {
	return CreateColumnIndex(-1, Name, nil)
}

// Create Column Index
func CreateColumnIndex(Index int, Name string, rcolumns []GDColumnIndex) GDColumnIndex {
	ci := GDColumnIndex{Index, Name, rcolumns}
	return ci
}

//Print ColumnIndex
func (ci GDColumnIndex) String() string {

	if ci.Name != "" {
		if ci.Index < 0 {
			return fmt.Sprintf("[%s]", ci.Name)
		} else {
			return fmt.Sprintf("[%s(%d)]", ci.Name, ci.Index)
		}
	} else {
		return fmt.Sprintf("[%d]", ci.Index)
	}

}
