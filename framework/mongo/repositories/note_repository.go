package repositories

import (
	"context"
	"errors"
	"notes-api-golang/framework/mongo/schemas"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	fetchedNote, errFetch := repository.FetchNoteById(insertResult.InsertedID, note.CreatedBy)

	return fetchedNote, errFetch
}

func (repository *NoteRepository) FetchNoteById(id interface{}, userId string) (schemas.Note, error) {
	var note schemas.Note
	collection := repository.mongoDatabase.Collection("notes")

	var objectID primitive.ObjectID
	var objectIDErr error
	switch id.(type) {
	case string:
		var err error
		objectID, objectIDErr = primitive.ObjectIDFromHex(id.(string))
		if err != nil {
			return note, errors.New("Invalid note id")
		}
	case primitive.ObjectID:
		objectID = id.(primitive.ObjectID)
	default:
		return note, errors.New("Invalid note id")
	}

	objectIDErr = collection.FindOne(context.Background(), bson.M{"_id": objectID, "created_by": userId, "deleted": bson.M{
		"$exists": false,
	}}).Decode(&note)
	if objectIDErr != nil {
		return note, errors.New("Note not found")
	}

	return note, nil
}

func (repository *NoteRepository) FetchAllNotes(filter bson.M) ([]schemas.Note, error) {
	var notes []schemas.Note
	collection := repository.mongoDatabase.Collection("notes")

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return notes, err
	}
	if err = cursor.All(context.Background(), &notes); err != nil {
		return notes, err
	}
	return notes, nil
}

func (repository *NoteRepository) SoftDelete(id interface{}, userId string) (schemas.Note, error) {
	var note schemas.Note
	collection := repository.mongoDatabase.Collection("notes")

	objectID, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return note, errors.New("Invalid note id")
	}

	update := bson.M{
		"$set": bson.M{
			"deleted":    true,
			"deleted_at": time.Now(),
			"deleted_by": userId,
		},
	}

	err = collection.FindOneAndUpdate(context.Background(), bson.M{"_id": objectID, "created_by": userId, "deleted": bson.M{
		"$exists": false,
	}}, update).Decode(&note)
	if err != nil {
		return note, errors.New("Note not found")
	}

	return note, nil
}

func (repository *NoteRepository) Update(id interface{}, note schemas.Note, userId string) (schemas.Note, error) {
	var noteResult schemas.Note
	collection := repository.mongoDatabase.Collection("notes")

	objectID, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return noteResult, errors.New("Invalid note id")
	}

	update := bson.M{
		"$set": bson.M{
			"title":      note.Title,
			"content":    note.Content,
			"updated_at": time.Now(),
			"updated_by": userId,
		},
	}

	err = collection.FindOneAndUpdate(context.Background(), bson.M{"_id": objectID, "created_by": userId, "deleted": bson.M{
		"$exists": false,
	}}, update).Decode(&noteResult)
	if err != nil {
		return noteResult, errors.New("Note not found")
	}

	return noteResult, nil
}

func (repository *NoteRepository) HardDelete(id interface{}) (schemas.Note, error) {
	var note schemas.Note
	collection := repository.mongoDatabase.Collection("notes")

	objectID, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return note, errors.New("Invalid note id")
	}

	err = collection.FindOneAndDelete(context.Background(), bson.M{"_id": objectID, "deleted": true}).Decode(&note)
	if err != nil {
		return note, errors.New("Note not found")
	}

	return note, nil
}
