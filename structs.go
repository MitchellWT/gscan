package gscan

import "log"

// ScanFile stores the minimal file information used in the applicaiton
type ScanFile struct {
	Path string
	Size int64
}

// ScanData stores the information that will be saved in /var/lib/scan/data
type ScanData struct {
	DateTime  string
	RootDir   string
	ScanFiles []ScanFile
}

// Used to log errors and exit out of application
func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
