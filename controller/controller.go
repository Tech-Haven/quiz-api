package controller

import (
	"net/http"
	"quiz-api/controller/question"
	"quiz-api/repository"
	"quiz-api/responses"

	"github.com/labstack/echo/v4"
)

type Controllers struct {
	QuestionController *question.Controller
}

func InitControllers(repos *repository.Repositories) *Controllers {
	return &Controllers{
		QuestionController: question.InitController(repos.QuestionRepo),
	}
}

func Ping() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, responses.HTTPResponse{Status: http.StatusOK, Message: responses.Success, Data: &echo.Map{"data": "Pong!"}})
	}
}
