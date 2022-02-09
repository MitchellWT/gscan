package structs

// HTMLTemplateData stores the information used in HTML exporting
type HTMLTemplateData struct {
	Title       string
	GraphLabels []string
	DataSets    []HTMLTemplateDataSet
}
