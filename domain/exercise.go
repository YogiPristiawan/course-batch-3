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

// type exerciseGetByIdQuestion struct {
// 	QuestionId int
// 	Body       string `json:"body"`
// 	OptionA    string `json:"option_a"`
// 	OptionB    string `json:"option_b"`
// 	OptionC    string `json:"option_c"`
// 	OptionD    string `json:"option_d"`
// 	Score      string `json:"score"`
// 	CreatedAt  string `json:"created_at"`
// 	UpdatedAt  string `json:"updated_at"`
// }

type ExerciseUseCase interface {
	CreateExercise(*ExerciseCreateRequest) ExerciseCreateResponse
	GetById(*ExerciseGetByIdRequest) ExerciseGetByIdResponse
}

type ExerciseRepository interface {
	Create(*ExerciseModel) error
	GetById(int) (*ExerciseModel, HttpError)
}
