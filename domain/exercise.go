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
	CreateExerciseQuestion(*ExerciseQuestionCreateRequest) ExerciseQuestionCreateResponse
	CreateExerciseAnswer(*ExerciseAnswerCreateRequest) ExerciseAnswerCreateResponse
}

type ExerciseRepository interface {
	Create(*ExerciseModel) error
	GetById(int) (*ExerciseModel, HttpError)

	FindUserQuestionAnswer(exerciseId int, userId int) ([]map[string]interface{}, HttpError)
}

type ExerciseValidator interface {
	ValidateCreateExercisePayload(*ExerciseCreateRequest) error
}

// create exercise
type ExerciseCreateRequest struct {
	Title       string `json:"title" validate:"required,max=200"`
	Description string `json:"description" validate:"required,max=2000"`
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

// create exercise question
type ExerciseQuestionCreateRequest struct {
	RequestMetadata
	ExerciseId    int    `json:"-"`
	Body          string `json:"body" validate:"required"`
	OptionA       string `json:"option_a" validate:"required"`
	OptionB       string `json:"option_b" validate:"required"`
	OptionC       string `json:"option_c" validate:"required"`
	OptionD       string `json:"option_d" validae:"required"`
	Score         int    `json:"score" validate:"required"`
	CorrectAnswer string `json:"correct_answer" validate:"required,oneof=a b c d"`
}

type ExerciseQuestionCreateResponse struct {
	CommonResult
	Message string `json:"message"`
}

// create answer
type ExerciseAnswerCreateRequest struct {
	RequestMetadata
	ExerciseId int
	QuestionId int
	Answer     string `json:"answer" validate:"required,oneof=a b c d"`
}

type ExerciseAnswerCreateResponse struct {
	CommonResult
	Message string `json:"message"`
}
