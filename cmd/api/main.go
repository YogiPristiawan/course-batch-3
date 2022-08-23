package main

import (
	"course/app/auth"
	"course/app/exercise"
	"course/interface/http/api"
	"course/pkg/databases"
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

	// initialize use case
	authUseCase := auth.NewAuthUseCase(userRepo, tokenize.GenerateAccessToken)
	exerciseUseCase := exercise.NewExerciseUseCase(exerciseRepository)

	// initialize validator
	authValidator := restValidator.NewAuthValidator(validator)
	exerciseValidator := restValidator.NewExerciseValidator(validator)

	// instance routes
	api.NewAuthRoute(router, authUseCase, authValidator)
	api.NewExerciseRoute(router, exerciseUseCase, exerciseValidator)

	// db := database.NewDabataseConn()
	// exerciseUcs := usecase.NewExerciseUsecase(db)
	// userUcs := userUc.NewUserUsecase(db)
	// r.GET("/hello", func(c *gin.Context) {
	// 	c.JSON(200, map[string]string{
	// 		"message": "hello world",
	// 	})
	// })
	// exercise
	// r.GET("/exercises/:id", middleware.WithAuthentication(userUcs), exerciseUcs.GetExercise)
	// r.GET("/exercises/:id/scores", middleware.WithAuthentication(userUcs), exerciseUcs.CalculateScore)

	// // user
	// r.POST("/register", userUcs.Register)
	// r.POST("/login", userUcs.Login)
	router.Run(":1234")
}
