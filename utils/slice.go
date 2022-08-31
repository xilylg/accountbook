package utils

func deleteK(slice []int, k int) []int {
	return append(slice[:k], slice[k+1:]...)
}