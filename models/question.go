package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Tags           []string           `bson:"tags,omitempty"`
	Question       string             `bson:"question,omitempty"`
	Answers        []string           `bson:"answers,omitempty"`
	CorrectAnswers []string           `bson:"correct_answers,omitempty"`
}
