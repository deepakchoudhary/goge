package score

type __Score__ struct {
	Current     int
	Lowest      int
	Highest     int
	GamesPlayed int
}

func (score *__Score__) __UpdateScore__(temp *__TempScore__) {
	if score.GamesPlayed == 1 {
		score.Lowest = temp.Score
	} else {
		if score.Lowest > temp.Score {
			score.Lowest = temp.Score
		}
	}
	if score.Highest < temp.Score {
		score.Highest = temp.Score
	}
}
func (score *__Score__) Add(temp *__TempScore__) *__Score__ {
	score.Current += temp.Score
	score.GamesPlayed += 1
	score.__UpdateScore__(temp)
	temp.Score = 0
	return score
}
func (score *__Score__) Reset() *__Score__ {
	score.Current = 0
	score.Highest = 0
	score.Lowest = 0
	score.GamesPlayed = 0
	return score
}

func NewPerm() *__Score__ {
	return &__Score__{Current: 0, Lowest: 0, Highest: 0, GamesPlayed: 0}
}

type __TempScore__ struct {
	Score int
}

func (temp *__TempScore__) Add(num int) *__TempScore__ {
	temp.Score += num
	return temp
}
func (temp *__TempScore__) Subtract(num int) *__TempScore__ {
	temp.Score -= num
	return temp
}
func (temp *__TempScore__) Set(num int) *__TempScore__ {
	temp.Score = num
	return temp
}
func (temp *__TempScore__) Reset() *__TempScore__ {
	temp.Score = 0
	return temp
}

func NewTemp() *__TempScore__ {
	return &__TempScore__{Score: 0}
}
