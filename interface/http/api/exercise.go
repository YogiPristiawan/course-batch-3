package api

import (
	"course/domain"
	"course/pkg/middleware"
	"course/presentation"

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
