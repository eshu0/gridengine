package geinterfaces

// Interface for Grid data Parser engine
type IGDEngine interface {

	// Parse the grid, with a row parser specified
	Execute(Rowparser IGDRowParser) error

	// how to handle the header row
	ParseHeader(ColumnIndexes []IGDColumnIndex, line []string) []IGDColumnIndex

	// Map the data in the row that was read
	MapRowData(ColumnIndexes []IGDColumnIndex, line []string, linecount int, pass int) IGDMappedRowData
}
