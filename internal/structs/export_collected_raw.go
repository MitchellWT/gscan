package structs

// ExportCollectedRaw stores the information that will be exported when
// a command is called
type ExportCollectedRaw struct {
	UnixStartTime int64
	UnixEndTime   int64
	RootDir       string
	TotalDiff     map[int64]int64
}
