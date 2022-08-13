package score_service

import (
	"context"
	"ggclass_go/src/enums"
	score_struct "ggclass_go/src/services/score/struct"
)

func (s *service) GetScore(ctx context.Context, classId int) (*score_struct.GetScoreOutput, error) {

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

	return nil, nil
}
