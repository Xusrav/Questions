package thirdquestion

// SearchFixPoint находит число, которое равно индексу массива
func SearchFixPoint(arrey [5]int) bool {
	count := len(arrey)
	for i := 0; i < count; i++ {
		if arrey[i] == i {
			return true
		}
	}
	return false
}
