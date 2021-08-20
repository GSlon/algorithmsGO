package bucketsort

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

func TestBucketSort(t *testing.T) {
    want := []int{3,9,21,25,29,37,43,49} // wiki example
    got := []int{29,25,3,49,9,37,21,43}
    got = BucketSort(got)
    if !compareSlices(got, want) {
        t.Errorf("sort -> %q, want -> %q", got, want)
    }
}
