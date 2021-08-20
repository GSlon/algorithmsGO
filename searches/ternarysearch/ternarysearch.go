package ternarysearch

func TernarySearch(arr []int, target int) int {
    size := len(arr) - 1
    high := size
    low  := 0

    for high >= low {
        mid1 := low + (high - low) / 3    // первая треть
        mid2 := high - (high - low) / 3   // последняя треть

        if arr[mid1] == target {
            return mid1
        }

        if arr[mid2] ==  target {
            return mid2
        }

        // выбираем в какой трети будет находиться target
        if target < arr[mid1] {
            high = mid1 - 1
        } else if target > arr[mid2] {
            low = mid2 + 1
        } else {    // вторая треть
            low = mid1 + 1
            high = mid2 - 1
        }
    }

    return -1
}

