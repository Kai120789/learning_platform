package mongoModels

import "learning-platform/board/internal/models/enum"

type BoardEvent struct {
	BoardID   int64           `bson:"board_id"`
	AuthorID  int64           `bson:"author_id"`
	Version   int64           `bson:"version"`
	Type      enum.ActionType `bson:"type"`
	Item      *BoardItem      `bson:"item,omitempty"`
	ObjectIDs []string        `bson:"object_ids,omitempty"`
}
