package routes

import (
	"quiz-api/controller"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, c *controller.Controllers) {
	// Healthcheck endpoint
	e.GET("/api/ping", controller.Ping())

	// Question Endpoints
	e.GET("/api/questions", c.QuestionController.GetQuestions())
}
