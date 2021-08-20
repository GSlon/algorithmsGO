package intersearch

// return index of target or -1 if does not exist
func InterSearch(arr []int, target int) int {
    startIndex := 0
    endIndex := len(arr) - 1

    for startIndex <= endIndex {
        mid := startIndex + (((endIndex - startIndex) / (arr[endIndex] - arr[startIndex])) * (target - arr[startIndex]))

        if arr[mid] > target {
            endIndex = mid - 1
        } else if arr[mid] < target {
            startIndex = mid + 1
        } else {
            return mid
        }

    }

    return -1
}
