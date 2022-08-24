package domain

type ExerciseModel struct {
	ID          int `gorm:"primaryKey"`
	Title       string
	Description string
	QuestionModel
}

func (e *ExerciseModel) TableName() string {
	return "exercises"
}

type ExerciseUseCase interface {
	CreateExercise(*ExerciseCreateRequest) ExerciseCreateResponse
	GetById(*ExerciseGetByIdRequest) ExerciseGetByIdResponse
	GetExerciseScore(*ExerciseScoreRequest) ExerciseScoreResponse
}

type ExerciseRepository interface {
	Create(*ExerciseModel) error
	GetById(int) (*ExerciseModel, HttpError)

	FindUserQuestionAnswer(exerciseId int, userId int) ([]map[string]interface{}, HttpError)
}

// create exercise
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

// get exercise by id
type ExerciseGetByIdRequest struct {
	ID int
}

type ExerciseGetByIdResponse struct {
	CommonResult
	ID          int                      `json:"id"`
	Title       string                   `json:"title"`
	Description string                   `json:"description"`
	Questions   []map[string]interface{} `json:"questions"`
}

// get exercise score
type ExerciseScoreRequest struct {
	RequestMetadata
	ID int
}

type ExerciseScoreResponse struct {
	CommonResult
	Score string `json:"score"`
}
