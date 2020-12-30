package thirdquestion

// SearchFixPoint находит число, которое равно индексу массива
func SearchFixPoint(arrey [5]int) int {
	count := len(arrey)
	for i := 0; i < count; i++ {
		if arrey[i] == i {
			return i
		}
	}
	return -1
}
