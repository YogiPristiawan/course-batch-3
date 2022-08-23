package api

import (
	"course/domain"
	"course/presentation"

	"github.com/gin-gonic/gin"
)

func NewAuthRoute(
	router *gin.Engine,
	useCase domain.AuthUseCase,
	validator domain.AuthValidator,
) {
	handler := &authHandler{
		router:    router,
		useCase:   useCase,
		validator: validator,
	}

	router.POST("/register", handler.handleRegister)
}

// Handler
type authHandler struct {
	router    *gin.Engine
	useCase   domain.AuthUseCase
	validator domain.AuthValidator
}

func (a *authHandler) handleRegister(c *gin.Context) {
	in := domain.AuthRegisterRequest{}

	if !presentation.ReadRestIn(c, &in) {
		return
	}

	err := a.validator.ValidateRegisterRequest(&in)
	if err != nil {
		out := struct {
			CommonResult domain.CommonResult `json:"-"`
			Message      string              `json:"message"`
		}{
			CommonResult: domain.CommonResult{
				ResErrorCode:    400,
				ResErrorMessage: err.Error(),
			},
		}
		presentation.WriteRestOut(c, out, &out.CommonResult)
		return
	}

	out := a.useCase.Register(&in)
	presentation.WriteRestOut(c, out, &out.CommonResult)
	return
}
