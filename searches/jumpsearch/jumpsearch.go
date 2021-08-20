package jumpsearch

import "math"

// return index of target or -1 if does not exist
// for ordered list
func JumpSearch(arr []int, target int) int {
    bound := 0
    jump := int(math.Sqrt(float64(len(arr))))

    // -1 чтобы выйти сразу после того как Min вернет
    // второй вариант
    for bound < len(arr)-1 && arr[bound] <= target {
        bound = int(math.Min(float64(bound + jump), float64(len(arr)-1)))
    }

    if (bound >= len(arr)) {
            return -1
    }

    // последней итерацией один раз перепрыгнули target
    // значит надо вычесть jump для левой границы,
    // а правой послужит текущий left
    // это обычный линейный поиск
    for i := bound - jump; i <= bound; i++ {
        if arr[i] == target {
            return i
        }
    }

    return -1
}
