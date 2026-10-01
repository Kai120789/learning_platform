package storage

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.mongodb.org/mongo-driver/mongo"
)

type Storage struct {
	PostgresStorage *PostgresStorage
	MongoStorage    *MongoStorage
}

func New(conn *pgxpool.Pool, client *mongo.Client) *Storage {
	return &Storage{
		PostgresStorage: NewPostgresStorage(conn),
		MongoStorage:    NewMongoStorage(client),
	}
}
