package gdengine

import (
	geinterfaces "github.com/eshu0/gdengine/interfaces"
)

//This structs has the mapping of the column indexs
type GDColumMapping struct {
	ColumnIndexes []geinterfaces.IGDColumnIndex `json:"columnindexes,omitempty"`
}

// This struct is the column index
// it can be based around the name or an integer index
// the name will be read by the header row
type GDColumnIndex struct {
	geinterfaces.IGDColumnIndex

	Index int `json:"index"`

	Name string `json:"name"`
	// do we need to associate other columns with this column?
	// for example: column 1 is the first name and column 2 is the second name they should be associated
	RelatedIndexes []geinterfaces.IGDColumnIndex `json:"relatedindexes,omitempty"`
}
