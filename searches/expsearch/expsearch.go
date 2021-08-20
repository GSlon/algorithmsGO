package expsearch

import (
    "math"
    b "github.com/GSlon/algorithmsGO/searches/binarysearch"
)

func ExpSearch(arr []int, target int) int {
    bound := 1
    for bound < len(arr) && arr[bound] <= target {
        bound *= 2
    }

    if bound >= len(arr) {
        return -1
    }

    // чтобы не выйти за пределы arr
    rbound := int(math.Min(float64(bound), float64(len(arr)-1)))
    return (bound/2) + b.BinarySearch(arr[bound / 2: rbound], target)
}
