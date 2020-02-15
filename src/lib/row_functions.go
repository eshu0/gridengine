package gdengine

// Default Parse Row behaviour
func OnParseRow(Rowparser GDRowParser, rowdata GDRowData) int {
	return Rowparser.ParseRow(rowdata)
}

// Default parse Mapped Data row behaviours
func OnParseMappedRow(Rowparser GDRowParser, data GDMappedRowData) int {
	return Rowparser.ParseMappedRow(data)
}
