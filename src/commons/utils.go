package commons

import "strings"

func TrimSpacesFromArray(arr []string) []string {
	out := []string{}
	for _, item := range arr {
		out = append(out, strings.TrimSpace(item))
	}
	return out
}

func FindInArray(arr []string, field string) bool {
	for _, item := range arr {
		if item == field {
			return true
		}
	}
	return false
}

func FindFieldIndexInArray(arr []string, field string) int {
	for index, item := range arr {
		if item == field {
			return index
		}
	}
	return -1
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
