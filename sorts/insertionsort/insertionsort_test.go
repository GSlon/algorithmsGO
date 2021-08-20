package insertionsort

import "testing"

func compareSlices(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for pos := range a {
		if a[pos] != b[pos] {
			return false
		}
	}

	return true
}

func TestInsertionSort(t *testing.T) {
    want := []int{1,2,3,4,5,6,7,8,9}
    got := []int{1,6,2,3,8,9,4,7,5}
    InsertionSort(got)
    if !compareSlices(got, want) {
        t.Errorf("sort -> %q, want -> %q", got, want)
    }
}
