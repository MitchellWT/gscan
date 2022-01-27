package structs

import sorting "github.com/MitchellWT/gscan/internal/sorting"

type ScanFileMap map[int64][]ScanFile

// ExportRaw stores the information that will be exported when a command
// is called
type ExportRaw struct {
	UnixStartTime int64
	UnixEndTime   int64
	RootDir       string
	ScanFileMap   ScanFileMap
}

func (sfm ScanFileMap) GetSortedKeys() []int64 {
	scanKeys := make([]int64, 0)
	for key, _ := range sfm {
		scanKeys = append(scanKeys, key)
	}

	return sorting.MergeSort(scanKeys)
}
