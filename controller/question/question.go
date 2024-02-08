package question

import (
	"net/http"
	"quiz-api/models"
	"quiz-api/repository/questionrepo"
	"quiz-api/responses"
	"quiz-api/utils"

	"github.com/labstack/echo/v4"
)

type repository interface {
	GetQuestions(tags []string) ([]models.Question, error)
	GetRandomQuestion(tags []string) ([]models.Question, error)
}

type Controller struct {
	service repository
}

func InitController(questionRepo *questionrepo.QuestionRepo) *Controller {
	return &Controller{
		service: questionRepo,
	}
}

func (m *Controller) GetQuestions() echo.HandlerFunc {
	return func(c echo.Context) error {

		tags, err := utils.ParseQueryParamsTags(c)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.HTTPResponse{Status: http.StatusInternalServerError, Message: responses.Error, Data: &echo.Map{"data": err.Error()}})
		}

		questions, err := m.service.GetQuestions(tags)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.HTTPResponse{Status: http.StatusInternalServerError, Message: responses.Error, Data: &echo.Map{"data": err.Error()}})
		}

		return c.JSON(http.StatusOK, responses.HTTPResponse{Status: http.StatusOK, Message: responses.Success, Data: &echo.Map{"data": questions}})
	}
}

func (m *Controller) GetRandomQuestion() echo.HandlerFunc {
	return func(c echo.Context) error {
		tags, err := utils.ParseQueryParamsTags(c)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.HTTPResponse{Status: http.StatusInternalServerError, Message: responses.Error, Data: &echo.Map{"data": err.Error()}})
		}

		question, err := m.service.GetRandomQuestion(tags)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.HTTPResponse{Status: http.StatusInternalServerError, Message: responses.Error, Data: &echo.Map{"data": err.Error()}})
		}

		return c.JSON(http.StatusOK, responses.HTTPResponse{Status: http.StatusOK, Message: responses.Success, Data: &echo.Map{"data": question}})
	}
}
