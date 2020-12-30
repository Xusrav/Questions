package secondquestion

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"

	// "io/ioutil"
	"log"
	"os"
	// "path/filepath"
)

type fakeMap struct {
	candidateID   int
	voteQount int
}

// Top топ кандидатов с наибольшим кол-ом голоса
type Top []fakeMap

var topCandidates Top

func (t Top) Len() int {
	return len(t)
}

func (t Top) Less(i, j int) bool {
	return t[i].voteQount < t[j].voteQount
}

func (t Top) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// TopCandidatesComputation вычисление топ кандидатов по голосам
func TopCandidatesComputation(topX int) Top {

	allCandidates := make(map[int]int)

	var x []string
	src, err := os.Open("../questions/secondquestion/votes.md")
	if err != nil {
		log.Print(err)
		return nil
	}
	defer func() {
		if cerr := src.Close(); cerr != nil {
			log.Print(cerr)
		}
	}()
	reader := bufio.NewReader(src)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			log.Print(line)
			break
		}
		if err != nil {
			log.Print(err)
			return nil
		}
		line = strings.TrimSuffix(line, "\n")
		x = append(x, line)
		log.Print(line)
	}
	log.Print(x)

	for _, value := range x {
		mySlice := strings.Split(value, ",")
		log.Print(mySlice)
		voterID, err := strconv.Atoi(mySlice[0])
		if err != nil {
			log.Print(err)
			return nil
		}
		candidateID, err := strconv.Atoi(mySlice[1])
		if err != nil {
			log.Print(err)
			return nil
		}

		if allCandidates[voterID] != 0 {
			log.Println("Жулик пойман", voterID)
			continue
		}
		allCandidates[voterID] = candidateID
	}

	log.Print(allCandidates)

	candidatesRating := make(map[int]int)

	for _, candidateID := range allCandidates {
		candidatesRating[candidateID]++
	}

	for key, value := range candidatesRating {
		topCandidates = append(topCandidates, fakeMap{
			candidateID: key,
			voteQount: value,
		})
	}

	sort.Sort(sort.Reverse(topCandidates))
	return topCandidates[:topX]

}
