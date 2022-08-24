package api

import (
	"course/domain"
	"course/pkg/middleware"
	"course/presentation"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewExerciseRoute(
	router *gin.Engine,
	useCase domain.ExerciseUseCase,
	exerciseValidator domain.ExerciseValidator,
	questionValidator domain.QuestionValidator,
) {
	handler := &exerciseHandler{
		exerciseValidator: exerciseValidator,
		useCase:           useCase,
		questionValidator: questionValidator,
	}

	router.POST("/exercises", middleware.AuthMiddleware(), handler.handleCreateExercise)
	router.GET("/exercises/:exerciseId", middleware.AuthMiddleware(), handler.handleGetExerciseById)
	router.GET("/exercises/:exerciseId/score", middleware.AuthMiddleware(), handler.handleGetExerciseScore)
	router.POST("/exercises/:exerciseId/questions", middleware.AuthMiddleware(), handler.handleCreateExerciseQuestion)
	router.POST("/exercises/:exerciseId/questions/:questionId/answer", middleware.AuthMiddleware(), handler.handleCreateExerciseAnswer)
}

// Handler
type exerciseHandler struct {
	exerciseValidator domain.ExerciseValidator
	questionValidator domain.QuestionValidator
	useCase           domain.ExerciseUseCase
}

func (e *exerciseHandler) handleCreateExercise(c *gin.Context) {
	in := domain.ExerciseCreateRequest{}

	if !presentation.ReadRestIn(c, &in) {
		return
	}

	if err := e.exerciseValidator.ValidateCreateExercisePayload(&in); err != nil {
		out := struct {
			CommonResult domain.CommonResult `json:"-"`
		}{
			CommonResult: domain.CommonResult{
				ResErrorCode:    400,
				ResErrorMessage: err.Error(),
			},
		}
		presentation.WriteRestOut(c, out, &out.CommonResult)
		return
	}

	out := e.useCase.CreateExercise(&in)
	presentation.WriteRestOut(c, out, &out.CommonResult)
}

func (e *exerciseHandler) handleGetExerciseById(c *gin.Context) {
	exerciseId, err := strconv.Atoi(c.Param("exerciseId"))
	// if param not number
	if err != nil {
		out := struct {
			CommonResult domain.CommonResult
		}{
			CommonResult: domain.CommonResult{
				ResErrorCode:    400,
				ResErrorMessage: "parameter harus berupa angka",
			},
		}

		presentation.WriteRestOut(c, out, &out.CommonResult)
		return
	}

	in := domain.ExerciseGetByIdRequest{
		ID: exerciseId,
	}

	// call use case
	out := e.useCase.GetById(&in)
	presentation.WriteRestOut(c, out, &out.CommonResult)
}

func (e *exerciseHandler) handleGetExerciseScore(c *gin.Context) {
	exerciseId, err := strconv.Atoi(c.Param("exerciseId"))
	if err != nil {
		out := struct {
			CommonResult domain.CommonResult
		}{
			CommonResult: domain.CommonResult{
				ResErrorCode:    400,
				ResErrorMessage: "parameter harus berupa angka",
			},
		}

		presentation.WriteRestOut(c, out, &out.CommonResult)
		return
	}

	authUserId, _ := c.Get("user_id")
	in := domain.ExerciseScoreRequest{
		RequestMetadata: domain.RequestMetadata{
			AuthUserId: int(authUserId.(float64)),
		},
		ID: exerciseId,
	}

	// call use case
	out := e.useCase.GetExerciseScore(&in)
	presentation.WriteRestOut(c, out, &out.CommonResult)
}

func (e *exerciseHandler) handleCreateExerciseQuestion(c *gin.Context) {
	exerciseId, err := strconv.Atoi(c.Param("exerciseId"))
	if err != nil {
		out := struct {
			CommonResult domain.CommonResult
		}{
			CommonResult: domain.CommonResult{
				ResErrorCode:    400,
				ResErrorMessage: "parameter harus berupa angka",
			},
		}

		presentation.WriteRestOut(c, out, &out.CommonResult)
		return
	}

	authUserId, _ := c.Get("user_id")
	in := domain.ExerciseQuestionCreateRequest{
		RequestMetadata: domain.RequestMetadata{
			AuthUserId: int(authUserId.(float64)),
		},
		ExerciseId: exerciseId,
	}

	if !presentation.ReadRestIn(c, &in) {
		return
	}

	if err = e.questionValidator.ValidateCreateQuestionPayload(&in); err != nil {
		out := struct {
			CommonResult domain.CommonResult
		}{
			CommonResult: domain.CommonResult{
				ResErrorCode:    400,
				ResErrorMessage: err.Error(),
			},
		}

		presentation.WriteRestOut(c, out, &out.CommonResult)
		return
	}

	out := e.useCase.CreateExerciseQuestion(&in)
	presentation.WriteRestOut(c, out, &out.CommonResult)
}

func (e *exerciseHandler) handleCreateExerciseAnswer(c *gin.Context) {
	exerciseId, err1 := strconv.Atoi(c.Param("exerciseId"))
	questionId, err2 := strconv.Atoi(c.Param("questionId"))
	authUserId, _ := c.Get("user_id")

	if err1 != nil || err2 != nil {
		out := struct {
			CommonResult domain.CommonResult
		}{
			CommonResult: domain.CommonResult{
				ResErrorCode:    400,
				ResErrorMessage: "parameter harus berupa angka",
			},
		}
		presentation.WriteRestOut(c, out, &out.CommonResult)
		return
	}

	in := domain.ExerciseAnswerCreateRequest{
		RequestMetadata: domain.RequestMetadata{
			AuthUserId: int(authUserId.(float64)),
		},
		ExerciseId: exerciseId,
		QuestionId: questionId,
	}

	if !presentation.ReadRestIn(c, &in) {
		return
	}

	if err := e.questionValidator.ValidateCreateAnswerPayload(&in); err != nil {
		out := struct {
			CommonResult domain.CommonResult
		}{
			CommonResult: domain.CommonResult{
				ResErrorCode:    400,
				ResErrorMessage: err.Error(),
			},
		}
		presentation.WriteRestOut(c, out, &out.CommonResult)
		return
	}

	out := e.useCase.CreateExerciseAnswer(&in)
	presentation.WriteRestOut(c, out, &out.CommonResult)
}
