package heapsort

// i -> это индекс рута поддерева, а не самого дерева
// To heapify a SUBtree (arr) rooted with node i which is
// an index in arr[]. Size for 'deleting' prev root
// i -> индекс root'a (не max, a того что с ним поменяли)
func heapify(arr []int, size, i int) {
	l, r := 2*i+1, 2*i+2 // дочерние элементы к нашему root
	max := i

	// если l или r больше size, то нет потомков
	// если дочерний элемент больше, то должны поменять их с root
	if l < size && arr[l] > arr[max] {
		max = l
	}
	if r < size && arr[r] > arr[max] {
		max = r
	}

	// дочерний элемент оказался больше
	if max != i {
        // меняем дочерний с root
		arr[i], arr[max] = arr[max], arr[i]

		// продолжаем пока root не станет листом
		heapify(arr, size, max)
	}
}


func HeapSort(arr []int) {
    size := len(arr)    // можно воспринимать уже как дерево

    // Build maxheap (rearrange array)
    // применяем к верхним узлам (нижним некуда опускаться)
    // половина будет листами -> смысла их опускать нет
    for i := size / 2 - 1; i >= 0; i-- { //-1 ->так как нумерация с 0
		heapify(arr, size, i)
	}

    for i := size - 1; i > 0; i-- {
        // 0й элемент -> максимальный. Ставим его в конец
        arr[0], arr[i] = arr[i], arr[0]

		// теперь наверху не макс элем -> нужно снова делать
        // maxHeap из него
		heapify(arr, i, 0) // массив без очередного элемента в конце
	}
}
