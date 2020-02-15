package geinterfaces

import (
	slinterfaces "github.com/eshu0/simplelogger/interfaces"
)

type IGDEngineOptionsManager interface {
	SaveRowOptions(Logger slinterfaces.ISimpleLogger) error

	GetRowParsingOptions(Logger slinterfaces.ISimpleLogger) (bool, IGDRowParsingOptions)

	GetColumMapping(Logger slinterfaces.ISimpleLogger) (bool, IGDColumMapping)

	SaveColumnmapping(Logger slinterfaces.ISimpleLogger) error

	ParseOptions(Logger slinterfaces.ISimpleLogger) (bool, error)

	// Current options// returns the memeber of the Engine Opts
	GetCurrentOptions() IGDEngineOptions

	// Current options// returns the memeber of the Engine Opts
	SetCurrentOptions(opts IGDEngineOptions) IGDEngineOptionsManager
}
