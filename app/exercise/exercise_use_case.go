package exercise

import "course/domain"

type exerciseUseCase struct {
	exerciseRepository domain.ExerciseRepository
}

func NewExerciseUseCase(
	exerciseRepository domain.ExerciseRepository,
) domain.ExerciseUseCase {
	return &exerciseUseCase{
		exerciseRepository: exerciseRepository,
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
