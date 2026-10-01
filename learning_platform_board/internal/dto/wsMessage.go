package dto

import (
	"learning-platform/board/internal/models/enum"
	mongoModels "learning-platform/board/internal/models/mongo"
)

type Envelope struct {
	Type    enum.ActionType `json:"type"`
	BoardID int64           `json:"board_id"`
	Version int64           `json:"version"`
}

type SnapshotMessage struct {
	Envelope Envelope                `json:"envelope"`
	Items    []mongoModels.BoardItem `json:"items"`
}
type AddMessage struct {
	Envelope Envelope              `json:"envelope"`
	Item     mongoModels.BoardItem `json:"item"`
}
type UpdateMessage struct {
	Envelope Envelope              `json:"envelope"`
	Item     mongoModels.BoardItem `json:"item"`
}
type DeleteMessage struct {
	Envelope  Envelope `json:"envelope"`
	ObjectIDs []string `json:"object_ids"`
}
