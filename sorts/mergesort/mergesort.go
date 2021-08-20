package mergesort

func merge(a []int, b []int) []int {
    res := make([]int, len(a) + len(b))
    i := 0
    j := 0

    for i < len(a) && j < len(b) {
        if a[i] <= b[j] {
            res[i+j] = a[i]
            i++
        } else {
            res[i+j] = b[j]
            j++
        }
    }

    for i < len(a) {
        res[i+j] = a[i]
        i++
    }

    for j < len(b) {
        res[i+j] = b[j]
        j++
    }

    return res
}

func MergeSort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }

    // делим пополам, пока не останется
    var middle = len(arr) / 2
    var a = MergeSort(arr[:middle])
    var b = MergeSort(arr[middle:])

    return merge(a, b)
}
