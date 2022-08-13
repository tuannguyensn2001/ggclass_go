package score_service

import (
	"context"
	"ggclass_go/src/enums"
	"ggclass_go/src/models"
	slicepkg "ggclass_go/src/packages/slice"
	score_struct "ggclass_go/src/services/score/struct"
	"sort"
)

func (s *service) GetScore(ctx context.Context, classId int) ([]score_struct.GetScoreOutput, error) {

	exercises, err := s.exerciseService.GetByClassId(ctx, classId)
	if err != nil {
		return nil, err
	}

	modeGetMarkForFirstTimeToDo := make([]int, 0)
	modeGetMarkForTimeNewest := make([]int, 0)
	modeGetHighestMark := make([]int, 0)

	for _, item := range exercises {
		if item.Mode == enums.GetMarkForFirstTimeToDo {
			modeGetMarkForFirstTimeToDo = append(modeGetMarkForFirstTimeToDo, item.Id)
		} else if item.Mode == enums.GetHighestMark {
			modeGetHighestMark = append(modeGetHighestMark, item.Id)
		} else if item.Mode == enums.GetMarkForTimeNewest {
			modeGetMarkForTimeNewest = append(modeGetMarkForTimeNewest, item.Id)
		}
	}

	assignmentGetMarkForFirstTimeToDo, err := s.repository.GetMarkForFirstTimeToDo(ctx, modeGetMarkForFirstTimeToDo)
	if err != nil {
		return nil, err
	}

	assignmentGetMarkForHighest, err := s.repository.GetMarkForHighest(ctx, modeGetHighestMark)
	if err != nil {
		return nil, err
	}

	assignmentGetMarkForTimeNewest, err := s.repository.GetMarkForNewest(ctx, modeGetMarkForTimeNewest)
	if err != nil {
		return nil, err
	}

	mapUser := make(map[int][]score_struct.ExerciseScoreOutput)

	for _, item := range assignmentGetMarkForFirstTimeToDo {
		_, ok := mapUser[item.UserId]
		if !ok {
			mapUser[item.UserId] = make([]score_struct.ExerciseScoreOutput, 0)
		}
		mapUser[item.UserId] = append(mapUser[item.UserId], score_struct.ExerciseScoreOutput{
			Id:   item.ExerciseId,
			Mark: item.Mark,
		})
	}

	for _, item := range assignmentGetMarkForHighest {
		_, ok := mapUser[item.UserId]
		if !ok {
			mapUser[item.UserId] = make([]score_struct.ExerciseScoreOutput, 0)
		}
		mapUser[item.UserId] = append(mapUser[item.UserId], score_struct.ExerciseScoreOutput{
			Id:   item.ExerciseId,
			Mark: item.Mark,
		})
	}

	for _, item := range assignmentGetMarkForTimeNewest {
		_, ok := mapUser[item.UserId]
		if !ok {
			mapUser[item.UserId] = make([]score_struct.ExerciseScoreOutput, 0)
		}
		mapUser[item.UserId] = append(mapUser[item.UserId], score_struct.ExerciseScoreOutput{
			Id:   item.ExerciseId,
			Mark: item.Mark,
		})
	}

	userIds := make([]int, 0)

	for k, _ := range mapUser {
		userIds = append(userIds, k)
	}

	users, err := s.userService.GetUsersByIds(ctx, userIds)
	if err != nil {
		return nil, err
	}

	userResults := make(map[int]models.User)

	for _, item := range users {
		userResults[item.Id] = item
	}

	var result []score_struct.GetScoreOutput

	for k, item := range mapUser {
		user, _ := userResults[k]

		sort.SliceStable(item, func(i, j int) bool {
			return item[i].Mark > item[j].Mark
		})

		result = append(result, score_struct.GetScoreOutput{
			Name:    user.Username,
			Avatar:  user.Profile.Avatar,
			Scores:  item,
			Average: calculateAverageScoreOutput(item),
			Id:      user.Id,
		})

	}

	return result, nil
}

func calculateAverageScoreOutput(array []score_struct.ExerciseScoreOutput) float64 {
	scores := make([]float64, len(array))

	for index, item := range array {
		scores[index] = item.Mark
	}

	return slicepkg.Average(scores)

}
