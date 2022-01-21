package enums

import "strings"

type ExportType int
type UndefinedExportTypeError struct{}

func (uete UndefinedExportTypeError) Error() string {
	return "Error: Undefined export type provided!"
}

const (
	ExportTypeUndefined ExportType = -1
	Raw                 ExportType = 0
	Total               ExportType = 1
)

func ToExportType(s string) (ExportType, error) {
	s = strings.ToLower(s)
	switch s {
	case "total":
		return Total, nil
	case "raw":
		return Raw, nil
	default:
		return ExportTypeUndefined, UndefinedExportTypeError{}
	}
}

func (et ExportType) String() string {
	switch et {
	case Total:
		return "total"
	case Raw:
		return "raw"
	default:
		return "nil"
	}
}
