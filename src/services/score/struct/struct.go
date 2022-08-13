package score_struct

type GetScoreOutput struct {
	Id      int                   `json:"id"`
	Name    string                `json:"name"`
	Avatar  string                `json:"avatar"`
	Scores  []ExerciseScoreOutput `json:"scores"`
	Average float64               `json:"average"`
}

type ExerciseScoreOutput struct {
	Id   int     `json:"id"`
	Mark float64 `json:"mark"`
}
