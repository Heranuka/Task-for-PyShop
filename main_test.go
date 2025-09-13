package main

import (
	"math/rand"
	"testing"
)

func createScoreStamps(offsets []int, scores []Score) []ScoreStamp {
	if len(offsets) != len(scores) {
		panic("offsets and scores must have the same length")
	}

	stamps := make([]ScoreStamp, len(offsets))
	for i := 0; i < len(offsets); i++ {
		stamps[i] = ScoreStamp{Offset: offsets[i], Score: scores[i]}
	}
	return stamps
}

func TestGetScore(t *testing.T) {
	simpleOffsets := []int{10, 20, 30, 40, 50}
	simpleScores := []Score{
		{Home: 1, Away: 0},
		{Home: 1, Away: 1},
		{Home: 2, Away: 1},
		{Home: 2, Away: 2},
		{Home: 3, Away: 2},
	}
	simpleStamps := createScoreStamps(simpleOffsets, simpleScores)

	t.Run("OffsetBeforeFirstStamp", func(t *testing.T) {
		score := getScore(simpleStamps, 5)
		if score != (Score{Home: 0, Away: 0}) {
			t.Errorf("Expected (0, 0), got (%v, %v)", score.Home, score.Away)
		}
	})

	t.Run("OffsetEqualToFirstStamp", func(t *testing.T) {
		score := getScore(simpleStamps, 10)
		if score != (Score{Home: 1, Away: 0}) {
			t.Errorf("Expected (1, 0), got (%v, %v)", score.Home, score.Away)
		}
	})

	t.Run("OffsetExactMatch", func(t *testing.T) {
		score := getScore(simpleStamps, 30)
		if score != (Score{Home: 2, Away: 1}) {
			t.Errorf("Expected (2, 1), got (%v, %v)", score.Home, score.Away)
		}
	})

	t.Run("OffsetWithinRangeNoMatch", func(t *testing.T) {
		score := getScore(simpleStamps, 25)
		if score != (Score{Home: 1, Away: 1}) {
			t.Errorf("Expected (1, 1), got (%v, %v)", score.Home, score.Away)
		}
	})

	t.Run("OffsetAfterLastStamp", func(t *testing.T) {
		score := getScore(simpleStamps, 60)
		if score != (Score{Home: 3, Away: 2}) {
			t.Errorf("Expected (3, 2), got (%v, %v)", score.Home, score.Away)
		}
	})

	t.Run("EmptyGameStamps", func(t *testing.T) {
		emptyStamps := []ScoreStamp{}
		score := getScore(emptyStamps, 50)
		if score != (Score{Home: 0, Away: 0}) {
			t.Errorf("Expected (0, 0), got (%v, %v)", score.Home, score.Away)
		}
	})

	t.Run("OneElementInList", func(t *testing.T) {
		oneElementStamps := createScoreStamps([]int{50}, []Score{{Home: 5, Away: 5}})
		score := getScore(oneElementStamps, 50)
		if score != (Score{Home: 5, Away: 5}) {
			t.Errorf("Expected (5, 5), got (%v, %v)", score.Home, score.Away)
		}
		score = getScore(oneElementStamps, 40)
		if score != (Score{Home: 0, Away: 0}) {
			t.Errorf("Expected (0, 0), got (%v, %v)", score.Home, score.Away)
		}
		score = getScore(oneElementStamps, 60)
		if score != (Score{Home: 5, Away: 5}) {
			t.Errorf("Expected (5, 5), got (%v, %v)", score.Home, score.Away)
		}
	})

	t.Run("ScoreReturnZero", func(t *testing.T) {
		empty := []ScoreStamp{}
		zeroScore := getScore(empty, 10)
		if zeroScore.Home != 0 || zeroScore.Away != 0 {
			t.Errorf("Expected zero score for empty slice; got %v", zeroScore)
		}

		stamps := []ScoreStamp{
			{Offset: 5, Score: Score{Home: 1, Away: 2}},
		}
		zeroScore = getScore(stamps, 0)
		if zeroScore.Home != 0 || zeroScore.Away != 0 {
			t.Errorf("Expected zero score for offset less than first; got %v", zeroScore)
		}
	})

	gameStamps := fillScores()
	randomIndex := rand.Intn(len(gameStamps))
	offset := gameStamps[randomIndex].Offset
	expectedScore := gameStamps[randomIndex].Score

	score := getScore(gameStamps, offset)
	if score != expectedScore {
		t.Errorf("Test Case 8: Random - expected %v, got %v", expectedScore, score)
	}
}
