package repositories

import (
	"context"
	"fmt"
	"notes-api-golang/framework/mongo/schemas"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type NoteRepository struct {
	mongoDatabase *mongo.Database
}

func NewNoteRepository(mongoDatabase *mongo.Database) *NoteRepository {
	return &NoteRepository{
		mongoDatabase,
	}
}

func (repository *NoteRepository) Create(note schemas.Note) (schemas.Note, error) {
	collection := repository.mongoDatabase.Collection("notes")

	insertResult, errInsert := collection.InsertOne(context.Background(), note)
	if errInsert != nil {
		return schemas.Note{}, errInsert
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	fetchedNote, errFetch := repository.FetchNoteById(insertResult.InsertedID, note.CreatedBy)

	return fetchedNote, errFetch
}

func (repository *NoteRepository) FetchNoteById(id interface{}, userId string) (schemas.Note, error) {
	var note schemas.Note
	collection := repository.mongoDatabase.Collection("notes")
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&note)
	if err != nil {
		return note, err
	}

	return note, nil
}

func (repository *NoteRepository) FetchAllNotes(userId string) ([]schemas.Note, error) {
	var notes []schemas.Note
	collection := repository.mongoDatabase.Collection("notes")
	cursor, err := collection.Find(context.Background(), bson.M{"created_by": userId})
	if err != nil {
		return notes, err
	}
	if err = cursor.All(context.Background(), &notes); err != nil {
		return notes, err
	}
	return notes, nil
}