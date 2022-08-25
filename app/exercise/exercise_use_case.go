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

	err := e.exerciseRepository.Create(exercise)
	if domain.HandleHttpError(err, &out.CommonResult) {
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
	if domain.HandleHttpError(err, &out.CommonResult) {
		return
	}

	// find questions
	questions, err := e.questionRepository.FindByExerciseId(exercise.ID)
	if domain.HandleHttpError(err, &out.CommonResult) {
		return
	}

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
	if domain.HandleHttpError(err, &out.CommonResult) {
		return
	}

	exercises, err := e.exerciseRepository.FindUserQuestionAnswer(in.ID, in.AuthUserId)
	if domain.HandleHttpError(err, &out.CommonResult) {
		return
	}

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

func (e *exerciseUseCase) CreateExerciseQuestion(in *domain.ExerciseQuestionCreateRequest) (out domain.ExerciseQuestionCreateResponse) {
	// verify if exercise exists
	_, err := e.exerciseRepository.GetById(in.ExerciseId)
	if domain.HandleHttpError(err, &out.CommonResult) {
		return
	}

	// create qeustion
	question := domain.QuestionModel{
		Body:          in.Body,
		OptionA:       in.OptionA,
		OptionB:       in.OptionB,
		OptionC:       in.OptionC,
		OptionD:       in.OptionD,
		Score:         in.Score,
		CorrectAnswer: in.CorrectAnswer,
		CreatorId:     in.AuthUserId,
		ExerciseId:    in.ExerciseId,
	}
	err = e.questionRepository.Create(&question)
	if domain.HandleHttpError(err, &out.CommonResult) {
		return
	}

	out.Message = "berhasil menambah pertanyaaan"
	return
}

func (e *exerciseUseCase) CreateExerciseAnswer(in *domain.ExerciseAnswerCreateRequest) (out domain.ExerciseAnswerCreateResponse) {
	// verify if question exists
	err := e.questionRepository.VerifyExerciseAndQuestionId(in.ExerciseId, in.QuestionId)
	if domain.HandleHttpError(err, &out.CommonResult) {
		return
	}

	// verify if answer already created
	err = e.questionRepository.VerifyExistsAnswer(in.AuthUserId, in.ExerciseId, in.QuestionId)
	if domain.HandleHttpError(err, &out.CommonResult) {
		return
	}

	// craete answer
	answer := domain.AnswerModel{
		ExerciseId: in.ExerciseId,
		QuestionId: in.QuestionId,
		UserId:     in.AuthUserId,
		Answer:     in.Answer,
	}
	err = e.questionRepository.CreateQuestionAnswer(&answer)
	if domain.HandleHttpError(err, &out.CommonResult) {
		return
	}

	out.Message = "berhasil menjawab pertanyaan"
	return
}
