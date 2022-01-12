package gscan

// ScanFile stores the minimal file information used in the applicaiton
type ScanFile struct {
	Path string
	Size int64
}

// ScanData stores the information that will be saved in /var/lib/scan/data
type ScanData struct {
	UnixTime  int64
	RootDir   string
	ScanFiles []ScanFile
}

// ExportRaw stores the information that will be exported when a command
// is called
type ExportRaw struct {
	UnixStartTime int64
	UnixEndTime   int64
	RootDir       string
	ScanFiles     map[int64][]ScanFile
}

// ExportCollectedRaw stores the information that will be exported when
// a command is called
type ExportCollectedRaw struct {
	UnixStartTime int64
	UnixEndTime   int64
	RootDir       string
	TotalDiff     map[int64]int64
}
