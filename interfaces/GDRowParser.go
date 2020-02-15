package geinterfaces

type IGDRowParser interface {
	// Parse the row
	ParseRow(data IGDRowData) int
	// Parse the Mapped Row
	ParseMappedRow(data IGDMappedRowData) int
}
