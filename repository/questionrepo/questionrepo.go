package questionrepo

import (
	"context"
	"quiz-api/config"
	"quiz-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuestionRepo struct {
	db     *mongo.Client
	config *config.Config
}

func NewQuestionRepo(db *mongo.Client, config *config.Config) *QuestionRepo {
	return &QuestionRepo{
		db:     db,
		config: config,
	}
}

// Returns all questions from the database
func (repo *QuestionRepo) GetQuestions(tags []string) ([]models.Question, error) {
	coll := repo.db.Database(repo.config.Mongo.Database).Collection(repo.config.Mongo.Collection)

	filter := bson.D{{}}

	if len(tags) > 0 {
		filter = bson.D{{"tags", bson.D{{"$all", tags}}}}
	}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var results []models.Question
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results, nil
}

// Returns 1 random question from the database
func (repo *QuestionRepo) GetRandomQuestion(tags []string) ([]models.Question, error) {
	coll := repo.db.Database(repo.config.Mongo.Database).Collection(repo.config.Mongo.Collection)

	selectOneStage := bson.D{{"$sample", bson.D{{"size", 1}}}}

	var cursor *mongo.Cursor
	var err error

	if len(tags) > 0 {
		filter := bson.D{{"$match", bson.D{{"tags", bson.D{{"$all", tags}}}}}}
		cursor, err = coll.Aggregate(context.TODO(), mongo.Pipeline{filter, selectOneStage})
	} else {
		cursor, err = coll.Aggregate(context.TODO(), mongo.Pipeline{selectOneStage})
	}

	if err != nil {
		panic(err)
	}

	var results []models.Question
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results, nil
}
