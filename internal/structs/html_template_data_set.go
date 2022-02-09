package structs

// HTMLTemplateDataSet stores one data set for the HTML export. This data
// directly maps to the data required for Chart.js
type HTMLTemplateDataSet struct {
	Label      string
	LineColour string
	Data       []float32
}
