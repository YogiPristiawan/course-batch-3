package exercise

import (
	"course/domain"
	"strconv"
	"strings"
	"sync"
)

type exerciseUseCase struct {
	exerciseRepository domain.ExerciseRepository
	questionRepository domain.QuestionRepository
}

func NewExerciseUseCase(
	exerciseRepository domain.ExerciseRepository,
	questionRepository domain.QuestionRepository,
) domain.ExerciseUseCase {
	return &exerciseUseCase{
		exerciseRepository: exerciseRepository,
		questionRepository: questionRepository,
	}
}

func (e *exerciseUseCase) CreateExercise(in *domain.ExerciseCreateRequest) (out domain.ExerciseCreateResponse) {
	// create exercise
	exercise := &domain.ExerciseModel{
		Title:       in.Title,
		Description: in.Description,
	}

	if err := e.exerciseRepository.Create(exercise); err != nil {
		out.SetError(500, err.Error())
		return
	}

	out.ID = exercise.ID
	out.Title = exercise.Title
	out.Description = exercise.Description
	return
}

func (e *exerciseUseCase) GetById(in *domain.ExerciseGetByIdRequest) (out domain.ExerciseGetByIdResponse) {
	// get exercise
	exercise, err := e.exerciseRepository.GetById(in.ID)
	domain.HandleHttpError(err, &out.CommonResult)

	// find questions
	questions, err := e.questionRepository.FindByExerciseId(exercise.ID)
	domain.HandleHttpError(err, &out.CommonResult)

	out.ID = exercise.ID
	out.Title = exercise.Title
	out.Description = exercise.Description

	if len(questions) > 0 {
		for _, val := range questions {
			question := make(map[string]interface{})
			question["id"] = val.ID
			question["body"] = val.Body
			question["option_a"] = val.OptionA
			question["option_b"] = val.OptionB
			question["option_c"] = val.OptionC
			question["option_d"] = val.OptionD
			question["score"] = val.Score
			question["created_at"] = val.CreatedAt
			question["updated_at"] = val.UpdatedAt

			out.Questions = append(out.Questions, question)
		}
	} else {
		out.Questions = []map[string]interface{}{}
	}

	return
}

func (e *exerciseUseCase) GetExerciseScore(in *domain.ExerciseScoreRequest) (out domain.ExerciseScoreResponse) {
	// verify if exercise exists
	_, err := e.exerciseRepository.GetById(in.ID)
	domain.HandleHttpError(err, &out.CommonResult)

	exercises, err := e.exerciseRepository.FindUserQuestionAnswer(in.ID, in.AuthUserId)
	domain.HandleHttpError(err, &out.CommonResult)

	// calculate score
	var wg sync.WaitGroup
	var m sync.Mutex
	var score int

	for _, val := range exercises {
		wg.Add(1)
		go func(val map[string]interface{}) {
			m.Lock()
			defer m.Unlock()
			defer wg.Done()

			if strings.EqualFold(val["correct_answer"].(string), val["user_answer"].(string)) {
				score += int(val["score"].(int32))
			}
		}(val)
	}
	wg.Wait()

	out.Score = strconv.Itoa(score)
	return
}
