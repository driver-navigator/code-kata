package matrix

import (
	"fmt"
)

func Clone(matrix [][]bool) [][]bool {
	duplicate := make([][]bool, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]bool, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}

func AreEqual(matrix1, matrix2 [][]bool) bool {
	if len(matrix1) != len(matrix2) {
		return false
	}

	for i := 0; i < len(matrix1); i++ {
		if len(matrix1[i]) != len(matrix2[i]) {
			return false
		}
		for j := 0; j < len(matrix1[i]); j++ {
			if matrix1[i][j] != matrix2[i][j] {
				return false
			}
		}
	}

	return true
}

func Exists(matrix [][]bool, ele bool) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == ele {
				return true
			}
		}
	}

	return false
}

func Print(matrix [][]bool) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] {
				fmt.Print("■")
			} else {
				fmt.Print("□")
			}
		}
		fmt.Println("")
	}
}
