package storage

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mongoModels "learning-platform/board/internal/models/mongo"
)

type MongoStorage struct {
	client *mongo.Client
}

func NewMongoStorage(client *mongo.Client) *MongoStorage {
	return &MongoStorage{
		client: client,
	}
}

func MongoConnection(connectStr string) (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectStr))
	if err != nil {
		return nil, fmt.Errorf("unable connect to mongo db %w", err)
	}

	return client, nil
}

func (m *MongoStorage) InsertBoardEvent(event mongoModels.BoardEvent) error {
	boardEvents := m.client.Database("board").Collection("board-events")

	_, err := boardEvents.InsertOne(context.Background(), event)
	if err != nil {
		return fmt.Errorf("add board item: %w", err)
	}

	return nil
}

func (m *MongoStorage) CancelBoardEvent(eventID string) error {
	items := m.client.Database("board").Collection("board-events")

	objectID, err := primitive.ObjectIDFromHex(eventID)
	if err != nil {
		return fmt.Errorf("cancel board event (get object id): %w", err)
	}

	filter := bson.D{
		{
			Key:   "_id",
			Value: objectID,
		},
	}

	_, err = items.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("cancel board event: %w", err)
	}

	return nil
}

func (m *MongoStorage) GetAllBoardEvents(boardID int64) ([]mongoModels.BoardEvent, error) {
	events := m.client.Database("board").Collection("board-events")

	filter := bson.D{
		{
			Key:   "board_id",
			Value: boardID,
		},
	}

	var resBoardEvents []mongoModels.BoardEvent
	cur, err := events.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("get all board events: %w", err)
	}

	if err := cur.All(context.Background(), &resBoardEvents); err != nil {
		return nil, err
	}

	return resBoardEvents, nil
}

func (m *MongoStorage) GetBoardEventsByVersion(boardID, version int64) ([]mongoModels.BoardEvent, error) {
	events := m.client.Database("board").Collection("board-events")

	filter := bson.D{
		{
			Key:   "board_id",
			Value: boardID,
		},
		{
			Key: "version",
			Value: bson.D{
				{Key: "$gt", Value: version},
			},
		},
	}
	opts := options.Find().SetSort(
		bson.D{
			{
				Key:   "version",
				Value: 1,
			},
		},
	)

	var resBoardEvents []mongoModels.BoardEvent
	cur, err := events.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, fmt.Errorf("get board events by version: %w", err)
	}

	if err := cur.All(context.Background(), &resBoardEvents); err != nil {
		return nil, err
	}

	return resBoardEvents, nil
}

func (m *MongoStorage) GetSnapshot(boardID int64) (*mongoModels.BoardSnapshot, error) {
	items := m.client.Database("board").Collection("board-snapshots")

	filter := bson.D{
		{
			Key:   "board_id",
			Value: boardID,
		},
	}

	var snapshot mongoModels.BoardSnapshot
	err := items.FindOne(context.Background(), filter).Decode(&snapshot)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, fmt.Errorf("get snapshot: %w", err)
	}

	return &snapshot, nil
}

func (m *MongoStorage) CreateSnapshot(boardItems []mongoModels.BoardItem) error {
	return nil
}

func (m *MongoStorage) UpdateSnapshot(boardItems []mongoModels.BoardItem) error {
	return nil
}
