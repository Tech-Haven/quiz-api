package repository

import (
	"quiz-api/config"
	"quiz-api/repository/questionrepo"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	QuestionRepo *questionrepo.QuestionRepo
}

func InitRepositories(db *mongo.Client, config *config.Config) *Repositories {
	questionrepo := questionrepo.NewQuestionRepo(db, config)
	return &Repositories{QuestionRepo: questionrepo}
}
