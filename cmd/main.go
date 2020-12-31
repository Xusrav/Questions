package main

import (
	"log"

	"github.com/Xusrav/Questions/bonus/fifthquestion"
	// "github.com/Xusrav/Questions/questions/secondquestion"
	// "github.com/Xusrav/Questions/bonus/thirdquestion"
	// "github.com/Xusrav/Questions/questions/firstquestion"
	// "log"
	// "fmt"
	// "github.com/Xusrav/Questions/bonus/fourfthquestion"
)

func main() {
	// // Testing first question
	// firstquestion.AddUnit("inch", "foot", 12, 1)
	// result := firstquestion.Convert("inch", 24, "foot")
	// log.Print(result)

	// // Testing third question
	// testArrey := [5]int{1,3,4,4,5}
	// total := thirdquestion.SearchFixPoint(testArrey)
	// log.Print(total)

	// top3 := secondquestion.TopCandidatesComputation(3)
	// log.Print(top3)

	// digit1, digit2 := 4, 13
	// fmt.Print("Egyptian Fraction Representation of ", nr, "/", dr, " is\n ")
	// fourthquestion.EgyptVersionCalculate(digit1, digit2)

	
	result := fifthquestion.FromNumbersToLetters(705)
	log.Println("Correct answer is: 'AAC', my answer is", result)
	
}
