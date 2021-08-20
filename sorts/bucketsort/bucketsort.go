package bucketsort

import (
    q "github.com/GSlon/algorithmsGO/sorts/quicksort"
)

func BucketSort(arr []int) []int {
    if len(arr) == 0 {
        return arr
    }

    size := len(arr)
    min := arr[0]
    max := arr[0]
    // ищем min и max
    for _, elem := range arr {
        if elem < min {
            min = elem
        } else if elem > max {
            max = elem
        }
    }

    // кол-во buckets
    count := (max - min) / size + 1

    var buckets [][]int
    for i := 0; i < count; i++ {
        buckets = append(buckets, []int{})
    }

    for i := 0; i < len(arr); i++ {
        // выбираем в какую корзину добавить
        index := (arr[i] - min) / size
        buckets[index] = append(buckets[index], arr[i])
    }

    var res []int
    // сортируем элементы в каждом buckets и кладем в res
    for i := 0; i < len(buckets); i++ {
        q.QuickSort(buckets[i])

        for j := 0; j < len(buckets[i]); j++ {
            res = append(res, buckets[i][j])
        }
    }

    return res
}
