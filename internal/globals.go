package gscan

import "log"

// Used to log errors and exit out of application
func ErrorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// DataDir is a file path for the data directory
var DataDir = "/var/lib/gscan/data/"
