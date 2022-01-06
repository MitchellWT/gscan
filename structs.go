package gscan

// ScanFile is the singular file scanned from the application
type ScanFile struct {
	Path string
	Size int64
}

// ScanData is the data that will be stored in json form in the /var/lib/gscan/data dir
type ScanData struct {
	DateTime  string
	RootDir   string
	ScanFiles []ScanFile
}
