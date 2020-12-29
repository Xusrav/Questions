package main

import (
	"github.com/Xusrav/Questions/bonus/thirdquestion"
	"github.com/Xusrav/Questions/questions/firstquestion"
	"log"
)



func main() {
	// Testing first question
	firstquestion.AddUnit("inch", "foot", 12, 1)	
	result := firstquestion.Convert("inch", 24, "foot")
	log.Print(result)

	// Testing third question
	testArrey := [5]int{1,3,4,3,5}
	total := thirdquestion.SearchFixPoint(testArrey)
	log.Print(total)
}

