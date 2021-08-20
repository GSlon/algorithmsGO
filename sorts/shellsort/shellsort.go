package shellsort

func ShellSort(arr []int) {
    size := len(arr)

    for d := int(size / 2); d > 0; d /= 2 {    // расстояние
        // для каждого элем делаем сравнение с элем на расстоянии d слева
        for i := d; i < size; i++ {
            for j := i - d; j >= 0 && arr[j] > arr[j+d]; j -= d {
                arr[j], arr[j+d] = arr[j+d], arr[j]
            }
        }
    }
}
