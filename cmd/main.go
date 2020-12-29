package main

import "log"

// UnitsInStruct структура с мапами
type UnitsInStruct struct {
	Units map[string]int
}

var total = &UnitsInStruct{
	Units: map[string]int{
		"inch": 1,
	},
}

func main() {
	AddUnit("inch", "foot", 12, 1)
	log.Print(total)
	
	result := Convert("inch", 24, "foot")
	log.Print(result)
}

// AddUnit Добавляет единицы с структуру
func AddUnit(first string, needInsert string, forFirst int, forInsert int) {
	result := forFirst / forInsert
	if first != "inch" {
		result = result * total.Units[first]
	}
	total.Units[needInsert] = result
}

// Convert конвертирует одну единицу в другую
func Convert(param string, qty int, needConvert string) (result float64)  {
	result = (float64(qty) * float64(total.Units[param])) / float64(total.Units[needConvert])
	return result
}