package exercise_struct

type CountMemberDoExercise struct {
	ExerciseId int `gorm:"column:exercise_id"`
	Count      int `gorm:"count"`
}
