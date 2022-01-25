package gscan

import "log"

// Used to log errors and exit out of application
func ErrorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
