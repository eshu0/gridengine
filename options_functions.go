package gdengine

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func CreateGDEngineOptionsManager() GDEngineOptionsManager {

	var OptionsManager GDEngineOptionsManager

	// template
	opts := GDEngineOptions{
		// dump the mapping to a file
		DumpColumnMapping: false,
		// dump the row template to a file
		DumpRowOptions: false,
		// the input file
		//InputFilePath: "",
		// default dump mapping file path
		DumpColumnMappingFilePath: "",
		// defaul dum row options file path
		DumpRowOptionsFilePath: "",

		RowParsingOptionsFilepath: "",

		ColumMappingFilepath: "",
	}

	// template
	opts.RowParsingoptions = GDRowParsingOptions{

		NumOfRowPasses: 1,

		//SkipHeaderRow: true,

		HeaderRowIndex: 1,

		StartRowIndex: -1,
	}

	OptionsManager.Options = opts

	return OptionsManager
}

func (gdpe GDEngineOptionsManager) ParseOptions(Logger GDLogger) (bool, error) {

	if gdpe.Options.RowParsingOptionsFilepath != "" {

		Logger.LogTrace(fmt.Sprintf("Loading Row options file %s", gdpe.Options.RowParsingOptionsFilepath))

		var sucess, rtp = gdpe.GetRowParsingOptions(Logger)

		if sucess {
			Logger.LogTrace("loaded custom template")
			gdpe.Options.RowParsingoptions = rtp
		} else {
			Logger.LogError("failed to load row parsing options")
			return false, errors.New("failed to load row parsing options")
		}
	} else {
		Logger.LogTrace("Did not Load Row options from file")
	}

	if gdpe.Options.ColumMappingFilepath != "" {

		Logger.LogTrace(fmt.Sprintf("Loading custom column mapping file %s", gdpe.Options.ColumMappingFilepath))

		var sucess, rtp = gdpe.GetColumMapping(Logger)
		if sucess {
			Logger.LogTrace("loaded custom Column Mapping")
			gdpe.Options.Columnmapping = rtp
		} else {
			Logger.LogError("failed to load column mapping")
			return false, errors.New("failed to load column mapping")
		}
	} else {
		Logger.LogTrace("Did not Load custom column mapping from file")
	}

	// are we dumping the column mappings?
	if gdpe.Options.DumpColumnMapping {

		Logger.LogTrace("Dumping Column Mapping")

		if gdpe.Options.DumpColumnMappingFilePath == "" {
			Logger.LogTrace(fmt.Sprintf("dump column mapping file path is empty - setting to default filepath: %s ", "./dump_RowTempleParser.json"))
			gdpe.Options.DumpColumnMappingFilePath = "./dump_ColumMapping.json"
		} else {
			Logger.LogTrace(fmt.Sprintf("dumping column mapping to file path: %s ", gdpe.Options.DumpColumnMappingFilePath))
		}

		err := gdpe.SaveColumnmapping(Logger)
		if err != nil {
			return false, err
		}

	} else {
		Logger.LogTrace("Not Savong column map to file")
	}

	if gdpe.Options.DumpRowOptions {

		Logger.LogTrace("Dumping Row Options")

		if gdpe.Options.DumpRowOptionsFilePath == "" {
			Logger.LogTrace(fmt.Sprintf("dump row options file path is empty - setting to default filepath: %s ", "./dump_RowTempleParser.json"))
			gdpe.Options.DumpRowOptionsFilePath = "./dump_RowTempleParser.json"
		} else {
			Logger.LogTrace(fmt.Sprintf("dumping row options to file path: %s ", gdpe.Options.DumpRowOptionsFilePath))
		}

		err := gdpe.SaveRowOptions(Logger)
		if err != nil {
			return false, err
		}

	} else {
		Logger.LogTrace("Not Saving row options to file")
	}

	// we are only dumping the options
	if gdpe.Options.DumpOnly {
		return false, nil
	}

	return true, nil
}

func (parser GDEngineOptionsManager) GetRowParsingOptions(Logger GDLogger) (bool, GDRowParsingOptions) {

	var c GDRowParsingOptions

	if CheckFileExists(parser.Options.RowParsingOptionsFilepath) {

		raw, err := ioutil.ReadFile(parser.Options.RowParsingOptionsFilepath)
		if err != nil {
			Logger.LogFatal(err)
			return false, c
		}

		json.Unmarshal(raw, &c)
		return true, c

	} else {
		return false, c
	}

}

func (parser GDEngineOptionsManager) GetColumMapping(Logger GDLogger) (bool, GDColumMapping) {

	var c GDColumMapping

	if CheckFileExists(parser.Options.ColumMappingFilepath) {

		raw, err := ioutil.ReadFile(parser.Options.ColumMappingFilepath)
		if err != nil {
			Logger.LogFatal(err) //.Error())
			return false, c
		}

		json.Unmarshal(raw, &c)
		return true, c

	} else {
		return false, c
	}

}

func (parser GDEngineOptionsManager) GetCurrentOptions() GDEngineOptions {
	return parser.Options
}

func (parser GDEngineOptionsManager) SetCurrentOptions(opts GDEngineOptions) GDEngineOptionsManager {
	parser.Options = opts
	return parser
}

// Exists reports whether the named file or directory exists.
func CheckFileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
