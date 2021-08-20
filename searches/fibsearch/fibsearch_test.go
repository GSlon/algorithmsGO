package fibsearch

import "testing"

func TestFibSearch(t *testing.T) {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    want := 4
    got := arr[FibSearch(arr, want)]

    if got != want {
        t.Errorf("sort -> %q, want -> %q", got, want)
    }
}
