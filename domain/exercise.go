package domain

type ExerciseModel struct {
	ID          int `gorm:"primaryKey"`
	Title       string
	Description string
}

func (e *ExerciseModel) TableName() string {
	return "exercises"
}

type ExerciseCreateRequest struct {
	Title       string `json:"title" validate:"required,max=200"`
	Description string `json:"description" validate:"required,max=2000"`
}

type ExerciseValidator interface {
	ValidateCreateExercisePayload(*ExerciseCreateRequest) error
}

type ExerciseCreateResponse struct {
	CommonResult
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ExerciseUseCase interface {
	CreateExercise(*ExerciseCreateRequest) ExerciseCreateResponse
}

type ExerciseRepository interface {
	Create(*ExerciseModel) error
}
