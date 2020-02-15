package gdengine

import (
	"errors"
	"fmt"
	"strings"
)

func CreateGDEngine() GDEngine {

	var eng GDEngine

	OptionsManager := CreateGDEngineOptionsManager()
	eng.OptionsManager = OptionsManager

	return eng
}

// Parse the Header
func (gdpe GDEngine) ParseHeader(ColumnIndexes []GDColumnIndex, line []string) []GDColumnIndex {

	var parsedColumnIndexes []GDColumnIndex

	// let's walk through the column indexes passed in
	for j := 0; j < len(ColumnIndexes); j++ {

		var ci GDColumnIndex

		// make local copy of index
		ci = ColumnIndexes[j]

		added := false

		// is the name value set?
		// if it is  set then we shall find the data
		if ci.Name != "" {

			columnname := strings.ToLower(ci.Name)

			gdpe.Logger.LogTrace(fmt.Sprintf("Searching for column index based on %s", columnname))

			// walk through the header row
			for p := 0; p < len(line); p++ {

				// grab the data and make lower case
				datalower := strings.ToLower(line[p])

				if columnname == datalower {
					ci.Index = p
					gdpe.Logger.LogTrace(fmt.Sprintf("Matched on Name - read - column index: %s - %+v", columnname, ci))

					parsedRelatedIndexes := gdpe.ParseHeader(ci.RelatedIndexes, line)
					if len(parsedRelatedIndexes) > 0 {
						gdpe.Logger.LogTrace(fmt.Sprintf("Found %d related columns", len(parsedRelatedIndexes)))
						ci.RelatedIndexes = parsedRelatedIndexes
					}

					parsedColumnIndexes = append(parsedColumnIndexes, ci)
					added = true
					break
				} else {
					gdpe.Logger.LogTrace(fmt.Sprintf("Did not match column index based on %s - data was %s ", columnname, datalower))
				}

			}

		}

		if !added {

			// walk through the header and
			for p := 0; p < len(line); p++ {
				// use the index
				if p == ci.Index {
					ci.Name = line[p]
					//ci := ColumnIndex{Index: p, Name: strings.ToLower(line[p])}
					gdpe.Logger.LogTrace(fmt.Sprintf("Matched on column index - read - column index: %d - %+v", p, ci))

					parsedRelatedIndexes := gdpe.ParseHeader(ci.RelatedIndexes, line)

					if len(parsedRelatedIndexes) > 0 {
						gdpe.Logger.LogTrace(fmt.Sprintf("Found %d related columns", len(parsedRelatedIndexes)))
						ci.RelatedIndexes = parsedRelatedIndexes
					}

					parsedColumnIndexes = append(parsedColumnIndexes, ci)
					added = true
					break
				}

			}
		}

	}
	return parsedColumnIndexes
}

func (gdpe GDEngine) MapRowData(ColumnIndexes []GDColumnIndex, line []string, linecount int, pass int) GDMappedRowData {

	opts := gdpe.OptionsManager.GetCurrentOptions()
	rowdata := GDMappedRowData{Index: linecount, Options: &opts, Pass: pass}

	// let's go through the indexes
	for j := 0; j < len(ColumnIndexes); j++ {

		var celldata GDCellData

		celldata.Index = ColumnIndexes[j]
		celldata.Data = line[ColumnIndexes[j].Index]
		celldata.RelatedData = make(map[*GDColumnIndex]string)

		for r := 0; r < len(ColumnIndexes[j].RelatedIndexes); r++ {
			relatedindex := ColumnIndexes[j].RelatedIndexes[r]
			celldata.RelatedData[&relatedindex] = line[relatedindex.Index]
		}

		rowdata.Cells = append(rowdata.Cells, celldata)
	}

	return rowdata
}

func (gdpe GDEngine) Execute(Rowparser GDRowParser) error {

	gdpe.Logger.LogTrace("Started Executing the GD Engine")
	gdpe.Logger.LogTrace("Parsing options of the GD Engine")
	succes, err := gdpe.OptionsManager.ParseOptions(gdpe.Logger)

	if !succes {
		return err
	}

	if gdpe.InputFilePath == "" {
		gdpe.Logger.LogTrace("no filename")
		return errors.New("no filename")
	} else {
		gdpe.Logger.LogTrace(fmt.Sprintf("Executing a parser against file: %s", gdpe.InputFilePath))
		return gdpe.Parser.ParseFile(Rowparser, gdpe.InputFilePath, gdpe.Logger, gdpe.OptionsManager.GetCurrentOptions())
	}

	return nil
}
