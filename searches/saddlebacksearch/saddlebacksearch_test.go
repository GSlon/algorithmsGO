package saddlebacksearch

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

func TestSaddlebackSearch(t *testing.T) {
    arr := [][]int {
        []int{1, 2, 3, 4, 5},
        []int{6, 7, 8, 9, 11},
        []int{12,13,14,15,16},
    }

    // row, col
    want := []int{1,2}
    got := SaddlebackSearch(arr, 8)

    if !compareSlices(got, want) {
        t.Errorf("sort -> %q, want -> %q", got, want)
    }
}
