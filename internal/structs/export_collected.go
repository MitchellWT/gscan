package structs

import sorting "github.com/MitchellWT/gscan/internal/sorting"

type TotalDiff map[int64]int64

// ExportCollectedRaw stores the information that will be exported when
// a command is called
type ExportCollected struct {
	UnixStartTime int64
	UnixEndTime   int64
	RootDir       string
	TotalDiff     TotalDiff
}

func (td TotalDiff) GetSortedKeys() []int64 {
	diffKeys := make([]int64, 0)
	for key, _ := range td {
		diffKeys = append(diffKeys, key)
	}

	return sorting.MergeSort(diffKeys)
}
