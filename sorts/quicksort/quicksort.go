package quicksort

// Разбиение: перераспределение элементов в массиве таким образом,
// что элементы, меньшие опорного (pivot), помещаются перед ним,
// а большие или равные - после.
// возвращаем индекс pivot'a
// разбиение Ломуто
func partition(arr []int) int {
    last := len(arr) - 1
    pivot := arr[last]
    i := -1

    // последний элем не берем, потому что он - pivot
    for j := 0; j < last; j++ {
        if arr[j] <= pivot {
            i++
            // сдвигаем элем меньшие pivot влево
            arr[i], arr[j] = arr[j], arr[i]
        }
    }

    // ставим pivot на границу элементов меньше pivot с большими чем он
    arr[i+1], arr[last] = arr[last], arr[i+1]
    return i + 1
}

func QuickSort(arr []int) {
    if len(arr) <= 1 {
        return
    }

    // Выбрать элемент из массива. Назовём его опорным.
	q := partition(arr)

    // Рекурсивно применить первые два шага к двум
    // подмассивам слева и справа от опорного элемента.
	QuickSort(arr[:q])
	QuickSort(arr[q+1:])
}
