package src

import (
	"fmt"
	"strings"
)

type Response struct {
	Data         struct{ Results string }
	ErrorCode    int
	ErrorMessage string
	ServerTime   int
}

type Score struct {
	Math           float32
	Literature     float32
	Physics        float32
	Chemistry      float32
	Biology        float32
	NaturalScience float32
	History        float32
	Geography      float32
	Civic          float32
	SocialScience  float32
	Language       float32
}

type Student struct {
	SBD string
	Score
}

type StudentChannel struct {
	id   string
	data *Student
}

type BoundChannel struct {
	areaCode int
	bound    int
}

func (score *Score) String() string {
	str := []string{
		formatScore(score.Math),
		formatScore(score.Literature),
		formatScore(score.Physics),
		formatScore(score.Chemistry),
		formatScore(score.Biology),
		formatScore(score.History),
		formatScore(score.Geography),
		formatScore(score.Civic),
		formatScore(score.Language),
	}
	return strings.Join(str, ",")
}

func (std *Student) String() string {
	return fmt.Sprintf("%s,%s", std.SBD, std.Score.String())
}
