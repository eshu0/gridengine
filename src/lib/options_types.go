package gdengine

import (
	geinterfaces "github.com/eshu0/gdengine/interfaces"
)

//This struct represents the options for parsing the rows
// designed to be loaded from file as JSON or be a in memory object
type GDRowParsingOptions struct {
	// how many times do we read the row?
	NumOfRowPasses int `json:"numofrowpasses"`
	// do we skip the first row?
	// the header row will always be read so column indexes can be built

	//SkipHeaderRow bool `json:"skipheaderrow"`
	HeaderRowIndex int `json:"headerrowindex"`

	//Max rows to read
	MaxRows int `json:"maxrows"`

	//what row to start at
	StartRowIndex int `json:"startrowindex"`
}

//These are the Parse Options
type GDEngineOptions struct {
	/// dump column mappings to file
	DumpColumnMapping         bool
	DumpColumnMappingFilePath string
	// dump row options to a file
	DumpRowOptions         bool
	DumpRowOptionsFilePath string

	/// Don't do any parsing
	DumpOnly bool

	// input file path
	//InputFilePath string

	// Row parsing options
	RowParsingoptions GDRowParsingOptions

	// are we loading the options from a JSON file?
	// do not set if inmmeory options should be used
	RowParsingOptionsFilepath string

	// Column Mappings
	Columnmapping GDColumMapping

	// are we loading the mapping from a JSON file?
	// do not set if inmmeory options should be used
	ColumMappingFilepath string
}

type GDEngineOptionsManager struct {
	// inherit from interface to get methods
	geinterfaces.IGDEngineOptionsManager

	// Options for the engine
	Options GDEngineOptions
}
