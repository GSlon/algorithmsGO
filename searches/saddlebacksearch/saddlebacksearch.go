package saddlebacksearch

// для отсортированных по row и col матриц
// идем по правому столбцу -> это max элементы строки
// если target больше -> опускаемся на строку ниже
// если меньше переходим на столбец влево
func SaddlebackSearch(matrix [][]int, target int) []int {
    row := len(matrix)
    col := len(matrix[0])

    i := 0
    j := col - 1

    for i < row && j >= 0 {
        elem := matrix[i][j]

        if elem == target {
            return []int{i, j}
        } else if elem < target {
            i++
        } else {
            j--
        }
    }

    return []int{-1, -1}
}
