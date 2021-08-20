package insertionsort

func InsertionSort(arr []int) {
    // проверяем элем слева, поэтому нет смысла начинать с 0го элем
    for i := 1; i < len(arr); i++ {
        temp := arr[i]
        var j int // сохраняем индекс, в который положим temp
        for j = i; j > 0 && arr[j-1] >= temp; j-- {
            arr[j] = arr[j-1] // двигаем вправо
        }
        // в цикле мы сдвигали вправо, теперь у нас есть
        // пустая ячейка -> кладем туда наш элем
        arr[j] = temp
    }
}
