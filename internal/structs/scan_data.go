package structs

// ScanData stores the information that will be saved in /var/lib/scan/data
type ScanData struct {
	UnixTime  int64
	RootDir   string
	ScanFiles []ScanFile
}
