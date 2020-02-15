package gdengine

import (
	geinterfaces "github.com/eshu0/gdengine/interfaces"
)

type GDCellData struct {
	// Index for the Cell
	Index GDColumnIndex
	// data in the cell
	Data string
	// any other cell data that should be related
	RelatedData map[*GDColumnIndex]string
}

type GDRowData struct {
	geinterfaces.IGDRowdata
	// row index
	Index int
	// raw data
	RawData []string
	// mapped data
	Pass int
	// reference to parser options
	//Options *GDEngineOptions
}

type GDMappedRowData struct {
	geinterfaces.IGDMappedRowData
	// row index
	Index int
	// row pass
	Pass int
	// reference to parser options
	Options *GDEngineOptions
	// Parsed Cell Data
	Cells []GDCellData
}
