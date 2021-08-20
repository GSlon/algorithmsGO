package fibsearch

func FibSearch(arr []int, target int) int {
    // генерируем числа Фиббоначи, пока они не превысят
    // длину arr
    fibs := []int{0,1}
    for fibs[len(fibs)-1] < len(arr) {
        fibs = append(fibs, fibs[len(fibs)-1] + fibs[len(fibs)-2])
    }

    // если последний элем вылез за пределы arr,
    // то берем предпоследний
    k := len(fibs)-1
    if fibs[k] > len(arr)-1 {
        k--
    }

    lastindex := 0
    for k > 0 {
        // сумма чисел Фиббоначи используется как индекс в arr
        index := lastindex + fibs[k]

        // если значение по индексу больше target, то
        // берем число Фиббоначи поменьше
        if index >= len(arr) || target < arr[index] {
            k--
        } else if target > arr[index] {
            lastindex = index
            k--
        } else {
            return index
        }
    }

    return -1
}
