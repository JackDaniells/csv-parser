package commons

import "strings"

func TrimSpacesFromElementsInArray(arr []string) []string {
	out := []string{}
	for _, item := range arr {
		out = append(out, strings.TrimSpace(item))
	}
	return out
}

func FindInStringArray(arr []string, field string) bool {
	arr = TrimSpacesFromElementsInArray(arr)
	for _, item := range arr {
		if item == field {
			return true
		}
	}
	return false
}

func FindInIntArray(arr []int, field int) bool {
	for _, item := range arr {
		if item == field {
			return true
		}
	}
	return false
}

func FindIndexInArray(arr []string, field string) (bool, int) {
	arr = TrimSpacesFromElementsInArray(arr)
	for index, item := range arr {
		if item == field {
			return true, index
		}
	}
	return false, -1
}

func HasDuplicatedElementsInArray(arr []string) bool {
	visited := make(map[string]bool)
	for i := range arr {
		if visited[arr[i]] {
			return true
		}
		visited[arr[i]] = true
	}
	return false
}

func GetDuplicatedElementsIndexesInArray(arr []string) [][]int {
	visited := make(map[string][]int)
	for index, element := range arr {
		visited[element] = append(visited[element], index)
	}

	duplicated := [][]int{}
	for _, indexes := range visited {
		if len(indexes) > 1 {
			duplicated = append(duplicated, indexes)
		}
	}
	return duplicated
}

func RemoveDuplicatedElementsInArray(arr []int) (output []int) {
	keys := make(map[int]bool)
	for _, item := range arr {
		if _, value := keys[item]; !value {
			keys[item] = true
			output = append(output, item)
		}
	}
	return output
}

func ConvertMatrixToArray(matrix [][]int) []int {
	out := []int{}
	for _, row := range matrix {
		for _, cell := range row {
			out = append(out, cell)
		}
	}
	return out
}
