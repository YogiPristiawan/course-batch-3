package exercise

import (
	"course/domain"
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
