package mongoModels

import "time"

type BoardSnapshot struct {
	BoardID   int64       `json:"board_id" bson:"board_id"`
	Version   int64       `json:"version" bson:"version"`
	Items     []BoardItem `json:"items" bson:"items"`
	UpdatedAt time.Time   `json:"updated_at" bson:"updated_at"`
}
