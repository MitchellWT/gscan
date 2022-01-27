package structs

// ExportRaw stores the information that will be exported when a command
// is called
type ExportRaw struct {
	UnixStartTime int64
	UnixEndTime   int64
	RootDir       string
	ScanFiles     map[int64][]ScanFile
}
