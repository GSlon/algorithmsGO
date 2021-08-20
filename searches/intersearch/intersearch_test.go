package intersearch

import "testing"

func TestInterSearch(t *testing.T) {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    want := 8
    got := arr[InterSearch(arr, want)]

    if got != want {
        t.Errorf("sort -> %q, want -> %q", got, want)
    }
}
