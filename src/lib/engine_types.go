package gdengine

import (
	geinterfaces "github.com/eshu0/gdengine/interfaces"
)

// Grid data Parser with sturcts
type GDEngine struct {
	// inherit from the engine interface
	geinterfaces.IGDEngine

	// Options
	// this is the options for this engine
	OptionsManager IGDEngineOptionsManager

	// input file path
	InputFilePath string

	// for logging
	Logger GDLogger

	// for parsing the file
	// this can be HTML, CSV etc
	Parser IGDParser
}
