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
	validator domain.ExerciseValidator,
) {
	handler := &exerciseHandler{
		validator: validator,
		useCase:   useCase,
	}

	router.POST("/exercises", middleware.AuthMiddleware(), handler.handleCreateExercise)
	router.GET("/exercises/:id", middleware.AuthMiddleware(), handler.handleGetExerciseById)
	router.GET("/exercises/:id/score", middleware.AuthMiddleware(), handler.handleGetExerciseScore)
}

// Handler
type exerciseHandler struct {
	validator domain.ExerciseValidator
	useCase   domain.ExerciseUseCase
}

func (e *exerciseHandler) handleCreateExercise(c *gin.Context) {
	in := domain.ExerciseCreateRequest{}

	if !presentation.ReadRestIn(c, &in) {
		return
	}

	if err := e.validator.ValidateCreateExercisePayload(&in); err != nil {
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
	return
}

func (e *exerciseHandler) handleGetExerciseById(c *gin.Context) {
	exerciseId, err := strconv.Atoi(c.Param("id"))
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
	return
}

func (e *exerciseHandler) handleGetExerciseScore(c *gin.Context) {
	exerciseId, err := strconv.Atoi(c.Param("id"))
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
