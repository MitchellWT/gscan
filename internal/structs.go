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
