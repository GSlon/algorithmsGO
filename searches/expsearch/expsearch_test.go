package expsearch

import "testing"

func TestExpSearch(t *testing.T) {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    want := 2
    got := arr[ExpSearch(arr, want)]

    if got != want {
        t.Errorf("sort -> %q, want -> %q", got, want)
    }
}
