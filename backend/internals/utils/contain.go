package utils

func Contains[T comparable](targetArray []T, element T) bool {
	for _, target := range targetArray {
		if target == element {
			return true
		}
	}
	return false
}
