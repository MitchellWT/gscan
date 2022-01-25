package enums

import "strings"

type FileFormat int
type UndefinedFileFormatError struct{}

func (uffe UndefinedFileFormatError) Error() string {
	return "Error: Undefined file format provided!"
}

const (
	FileFormatUndefined FileFormat = -1
	JSON                FileFormat = 0
	HTML                FileFormat = 1
	CSV                 FileFormat = 2
)

func ToFileFormat(s string) (FileFormat, error) {
	s = strings.ToLower(s)
	switch s {
	case "json":
		return JSON, nil
	case "html":
		return HTML, nil
	case "csv":
		return CSV, nil
	default:
		return FileFormatUndefined, UndefinedFileFormatError{}
	}
}

func (ff FileFormat) String() string {
	switch ff {
	case JSON:
		return "json"
	case HTML:
		return "html"
	case CSV:
		return "csv"
	default:
		return "nil"
	}
}
