package services

import (
	"context"
	"parapilot/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AggregatePassages(collection *mongo.Collection) ([]models.Passage, error) {
	pipeline := bson.A{
		bson.M{
			"$lookup": bson.M{
				"from":         "sentences",
				"localField":   "_id",
				"foreignField": "passageId",
				"as":           "sentences",
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "questions",
				"localField":   "_id",
				"foreignField": "passageId",
				"as":           "questions",
			},
		},
		bson.M{
			"$unwind": bson.M{
				"path":                       "$sentences",
				"preserveNullAndEmptyArrays": true,
			},
		},
		bson.M{
			"$sort": bson.M{
				"sentences.order": 1,
			},
		},
		bson.M{
			"$group": bson.M{
				"_id":       "$_id",
				"title":     bson.M{"$first": "$title"},
				"sentences": bson.M{"$push": "$sentences"},
				"questions": bson.M{"$first": "$questions"},
			},
		},
		bson.M{
			"$project": bson.M{
				"_id":       1,
				"title":     1,
				"sentences": "$sentences.text",
				"questions": 1,
			},
		},
	}

	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var passages []models.Passage
	for cursor.Next(context.Background()) {
		var passage models.Passage
		if err := cursor.Decode(&passage); err != nil {
			return nil, err
		}
		passages = append(passages, passage)
	}

	return passages, nil
}
