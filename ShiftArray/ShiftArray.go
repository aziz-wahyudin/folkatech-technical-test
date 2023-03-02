package main

import "fmt"

func ShiftArray(array []int, i int) []int {
	// change 1d array to 2d array
	// since the original array consist of 9 elements
	// hence i will change it to 3x3 matrix
	rows := 3
	cols := 3
	arr2D := make([][]int, rows)
	for j := range arr2D {
		arr2D[j] = make([]int, cols)
	}

	// fill 2d array from original array
	for k := 0; k < rows; k++ {
		for l := 0; l < cols; l++ {
			index := k*cols + l
			if index < len(array) {
				arr2D[k][l] = array[index]
			}
		}
	}

	// rotating the values on the outer ring of a matrix of size HxW
	height := len(arr2D)
	width := len(arr2D[0])
	matrixMap := mapMatrix(height, width)
	i %= len(matrixMap)
	rotatedMap := append(matrixMap[len(matrixMap)-i:], matrixMap[:len(matrixMap)-i]...)
	newMatrix := make(map[[2]int]int)
	for _, el := range matrixMap {
		newMatrix[el] = arr2D[el[0]][el[1]]
	}
	for i, el := range rotatedMap {
		arr2D[matrixMap[i][0]][matrixMap[i][1]] = newMatrix[el]
	}

	// return 2d array to its original form; 1d array
	flattened := make([]int, 0)
	for _, row := range arr2D {
		for _, val := range row {
			flattened = append(flattened, val)
		}
	}
	return flattened
}

func mapMatrix(h, w int) [][2]int {
	matrix := make([][2]int, 0)
	for i := 0; i < w; i++ {
		matrix = append(matrix, [2]int{0, i})
	}
	for i := 1; i < h; i++ {
		matrix = append(matrix, [2]int{i, w - 1})
	}
	for i := w - 2; i >= 0; i-- {
		matrix = append(matrix, [2]int{h - 1, i})
	}
	for i := h - 2; i > 0; i-- {
		matrix = append(matrix, [2]int{i, 0})
	}
	return matrix
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(ShiftArray(arr, 1))
	fmt.Println(ShiftArray(arr, 2))
}
