package score_struct

type GetScoreOutput struct {
	Name   string                `json:"name"`
	Avatar string                `json:"avatar"`
	Scores []ExerciseScoreOutput `json:"scores"`
}

type ExerciseScoreOutput struct {
	Id   int     `json:"id"`
	Mark float64 `json:"mark"`
}
