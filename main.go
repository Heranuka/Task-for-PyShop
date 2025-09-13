package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const TIMESTAMPS_COUNT = 500
const PROBABILITY_SCORE_CHANGED = 0.01
const PROBABILITY_HOME_SCORE = 0.45
const OFFSET_MAX_STEP = 3

type Score struct {
	Home int
	Away int
}

type ScoreStamp struct {
	Offset int
	Score  Score
}

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	var stamps = fillScores()

	for _, stamp := range stamps {
		fmt.Printf("%v: %v -- %v\n", stamp.Offset, stamp.Score.Home, stamp.Score.Away)
	}

	offset := 100
	score := getScore(stamps, offset)
	fmt.Printf("Score at offset %d: %v -- %v\n", offset, score.Home, score.Away)

	offset = 0
	score = getScore(stamps, offset)
	fmt.Printf("Score at offset %d: %v -- %v\n", offset, score.Home, score.Away)

	offset = 50000
	score = getScore(stamps, offset)
	fmt.Printf("Score at offset %d: %v -- %v\n", offset, score.Home, score.Away)

	offset = -1
	score = getScore(stamps, offset)
	fmt.Printf("Score at offset %d: %v -- %v\n", offset, score.Home, score.Away)
}

func generateStamp(previousValue ScoreStamp) ScoreStamp {
	scoreChanged := random.Float32() > 1-PROBABILITY_SCORE_CHANGED
	homeScoreChange := 0
	if scoreChanged && random.Float32() > 1-PROBABILITY_HOME_SCORE {
		homeScoreChange = 1
	}

	awayScoreChange := 0
	if scoreChanged && homeScoreChange == 0 {
		awayScoreChange = 1
	}

	offsetChange := int(math.Floor(random.Float64()*OFFSET_MAX_STEP)) + 1

	return ScoreStamp{
		Offset: previousValue.Offset + offsetChange,
		Score: Score{
			Home: previousValue.Score.Home + homeScoreChange,
			Away: previousValue.Score.Away + awayScoreChange,
		},
	}
}

func fillScores() []ScoreStamp {
	scores := make([]ScoreStamp, TIMESTAMPS_COUNT)
	prevScore := ScoreStamp{
		Offset: 0,
		Score:  Score{Home: 0, Away: 0},
	}
	scores[0] = prevScore

	for i := 1; i < TIMESTAMPS_COUNT; i++ {
		scores[i] = generateStamp(prevScore)
		prevScore = scores[i]
	}

	return scores
}

func getScore(gameStamps []ScoreStamp, offset int) Score {
	if len(gameStamps) == 0 {
		return Score{Home: 0, Away: 0}
	}

	if offset < gameStamps[0].Offset {
		return Score{Home: 0, Away: 0}
	}
	if offset >= gameStamps[len(gameStamps)-1].Offset {
		return gameStamps[len(gameStamps)-1].Score
	}

	low := 0
	high := len(gameStamps) - 1

	for low <= high {
		mid := (low + high) / 2
		if gameStamps[mid].Offset == offset {
			return gameStamps[mid].Score
		} else if gameStamps[mid].Offset < offset {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if high >= 0 {
		return gameStamps[high].Score
	}
	return Score{Home: 0, Away: 0}
}
