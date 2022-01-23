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
)

func ToFileFormat(s string) (FileFormat, error) {
	s = strings.ToLower(s)
	switch s {
	case "json":
		return JSON, nil
	case "html":
		return HTML, nil
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
	default:
		return "nil"
	}
}
