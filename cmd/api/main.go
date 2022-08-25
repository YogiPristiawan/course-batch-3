package main

import (
	"course/app/auth"
	"course/app/exercise"
	"course/interface/http/api"
	"course/pkg/databases"
	"course/pkg/helpers"
	"course/pkg/repositories"
	"course/pkg/tokenize"
	restValidator "course/pkg/validator/rest"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	// initialize database
	db := databases.NewDatabaseConn()

	// initialize pkg
	validator := validator.New()

	// initialize repository
	userRepo := repositories.NewUserRepository(db)
	exerciseRepository := repositories.NewExerciseRepository(db)
	questionRepository := repositories.NewQuestionRepository(db)

	// initialize use case
	authUseCase := auth.NewAuthUseCase(userRepo, tokenize.GenerateAccessToken, helpers.CompareHashAndPassword)
	exerciseUseCase := exercise.NewExerciseUseCase(exerciseRepository, questionRepository)

	// initialize validator
	authValidator := restValidator.NewAuthValidator(validator)
	exerciseValidator := restValidator.NewExerciseValidator(validator)
	questionValidator := restValidator.NewQuestionValidator(validator)

	// instance routes
	api.NewAuthRoute(router, authUseCase, authValidator)
	api.NewExerciseRoute(router, exerciseUseCase, exerciseValidator, questionValidator)

	router.Run(":1234")
}
