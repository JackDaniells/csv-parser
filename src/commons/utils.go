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
