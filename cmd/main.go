package main

import (
	"log"

	"github.com/Xusrav/Questions/bonus/fourthquestion"
	"github.com/Xusrav/Questions/bonus/thirdquestion"
	"github.com/Xusrav/Questions/questions/firstquestion"
	"github.com/Xusrav/Questions/questions/secondquestion"
)

func main() {
	// Run first question
	firstquestion.AddUnit("inch", "foot", 12, 1)
	result := firstquestion.Convert("inch", 24, "foot")
	log.Print(result)

	// Run third question
	testArrey := [5]int{1,3,4,4,5}
	total := thirdquestion.SearchFixPoint(testArrey)
	log.Print(total)

	// Run second question
	top3 := secondquestion.TopCandidatesComputation(3)
	log.Print(top3)

	// 
	digit1, digit2 := 4, 13
	fourthquestion.EgyptVersionCalculate(digit1, digit2)

	
	

}
