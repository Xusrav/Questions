package fifthquestion

import "log"

// FromNumbersToLetters предоставляет нужные буквы из полученной цифры.
func FromNumbersToLetters(numberToGenerate int) string {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"	
	var index int
	var arrey [10000]int
	for numberToGenerate > 0 {
		arrey[index] = numberToGenerate % len(alphabet)
		numberToGenerate = numberToGenerate / len(alphabet)
		index++
	}

	for i := 0; i < index-1; i++ {
		if arrey[i] <= 0 {
			arrey[i] += len(alphabet)
			arrey[i+1] = arrey[i+1] - 1
		}
	}

	result := ""
	for i := index; i >= 0; i-- {
		if arrey[i] > 0 {
			result += string(rune(arrey[i] - 1 + 65))
		}
	}

	log.Println(result)
	return result
}
