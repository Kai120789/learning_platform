package service

import (
	"learning-platform/board/internal/dto"
	"learning-platform/board/internal/models/enum"
	mongoModels "learning-platform/board/internal/models/mongo"
	"sync"
)

type LiveBoardService struct {
	mongo   MongoStorage
	rooms   map[int64]*Room
	roomsMu sync.Mutex
}

type MongoStorage interface {
	InsertBoardEvent(event mongoModels.BoardEvent) error
	CancelBoardEvent(eventID string) error
	GetSnapshot(boardID int64) (*mongoModels.BoardSnapshot, error)
	GetAllBoardEvents(boardID int64) ([]mongoModels.BoardEvent, error)
	GetBoardEventsByVersion(boardID, version int64) ([]mongoModels.BoardEvent, error)
}

func NewLiveBoardService(mongo MongoStorage) *LiveBoardService {
	return &LiveBoardService{
		mongo: mongo,
		rooms: make(map[int64]*Room),
	}
}

type Room struct {
	mu      sync.Mutex
	BoardID int64
	Version int64
	Items   []mongoModels.BoardItem
}

func (l *LiveBoardService) AddBoardEvent(id, userID int64, msg dto.AddMessage) error {
	room, err := l.ensureRoom(id)
	if err != nil {
		return err
	}

	room.mu.Lock()
	defer room.mu.Unlock()

	room.Version++
	msg.Item.AuthorID = userID
	msg.Item.ZIndex = nextZ(room.Items)
	boardEvent := mongoModels.BoardEvent{
		BoardID:  id,
		Version:  room.Version,
		Type:     msg.Envelope.Type,
		Item:     &msg.Item,
		AuthorID: userID,
	}

	err = l.mongo.InsertBoardEvent(boardEvent)
	if err != nil {
		room.Version--
		return err
	}

	room.Items = applyEvent(room.Items, boardEvent)
	return nil
}

// TODO: сделать функционал, чтобы преподаватель мог редачить и удалять все айтемы, а ученики только свои
func (l *LiveBoardService) UpdateBoardEvent(id, userID int64, msg dto.UpdateMessage) error {
	room, err := l.ensureRoom(id)
	if err != nil {
		return err
	}

	room.mu.Lock()
	defer room.mu.Unlock()

	for i := range room.Items {
		if room.Items[i].ObjectID == msg.Item.ObjectID {
			msg.Item.AuthorID = room.Items[i].AuthorID
			break
		}
	}

	room.Version++
	boardEvent := mongoModels.BoardEvent{
		BoardID:  id,
		Version:  room.Version,
		Type:     msg.Envelope.Type,
		Item:     &msg.Item,
		AuthorID: userID,
	}

	err = l.mongo.InsertBoardEvent(boardEvent)
	if err != nil {
		room.Version--
		return err
	}

	room.Items = applyEvent(room.Items, boardEvent)
	return nil
}

// TODO: сделать функционал, чтобы преподаватель мог редачить и удалять все айтемы, а ученики только свои
func (l *LiveBoardService) DeleteBoardEvents(id, userID int64, msg dto.DeleteMessage) error {
	room, err := l.ensureRoom(id)
	if err != nil {
		return err
	}

	room.mu.Lock()
	defer room.mu.Unlock()

	room.Version++
	boardEvent := mongoModels.BoardEvent{
		BoardID:   id,
		Version:   room.Version,
		Type:      msg.Envelope.Type,
		ObjectIDs: msg.ObjectIDs,
		AuthorID:  userID,
	}

	err = l.mongo.InsertBoardEvent(boardEvent)
	if err != nil {
		room.Version--
		return err
	}

	room.Items = applyEvent(room.Items, boardEvent)
	return nil
}

// TODO: Доработать логику
func (l *LiveBoardService) CancelLastBoardEvent(eventID string) error {
	err := l.mongo.CancelBoardEvent(eventID)
	if err != nil {
		return err
	}

	return nil
}

func (l *LiveBoardService) GetSnapshot(boardID int64) (*mongoModels.BoardSnapshot, error) {
	room, err := l.ensureRoom(boardID)
	if err != nil {
		return nil, err
	}

	room.mu.Lock()
	defer room.mu.Unlock()

	items := make([]mongoModels.BoardItem, len(room.Items))
	copy(items, room.Items)
	return &mongoModels.BoardSnapshot{
		BoardID: boardID,
		Version: room.Version,
		Items:   items,
	}, nil
}

func applyEvent(items []mongoModels.BoardItem, ev mongoModels.BoardEvent) []mongoModels.BoardItem {
	switch ev.Type {
	case enum.AddItem:
		if ev.Item == nil {
			return items
		}
		return append(items, *ev.Item)
	case enum.UpdateItem:
		if ev.Item == nil {
			return items
		}
		out := make([]mongoModels.BoardItem, len(items))
		copy(out, items)
		for i := range out {
			if out[i].ObjectID == ev.Item.ObjectID {
				out[i] = *ev.Item
				break
			}
		}
		return out
	case enum.DeleteItems:
		if len(ev.ObjectIDs) == 0 {
			return items
		}
		drop := make(map[string]struct{}, len(ev.ObjectIDs))
		for _, id := range ev.ObjectIDs {
			drop[id] = struct{}{}
		}
		out := make([]mongoModels.BoardItem, 0, len(items))
		for _, it := range items {
			if _, ok := drop[it.ObjectID]; !ok {
				out = append(out, it)
			}
		}
		return out
	default:
		return items
	}
}

func (l *LiveBoardService) ensureRoom(boardID int64) (*Room, error) {
	l.roomsMu.Lock()
	if room, ok := l.rooms[boardID]; ok {
		l.roomsMu.Unlock()
		return room, nil
	}
	l.roomsMu.Unlock()

	snapshot, err := l.mongo.GetSnapshot(boardID)
	if err != nil {
		return nil, err
	}

	var items []mongoModels.BoardItem
	var version int64
	if snapshot != nil {
		items = snapshot.Items
		version = snapshot.Version
	}

	events, err := l.mongo.GetBoardEventsByVersion(boardID, version)
	if err != nil {
		return nil, err
	}
	for _, ev := range events {
		items = applyEvent(items, ev)
		version = ev.Version
	}

	room := &Room{
		BoardID: boardID,
		Version: version,
		Items:   items,
	}

	l.roomsMu.Lock()
	if existing, ok := l.rooms[boardID]; ok {
		l.roomsMu.Unlock()
		return existing, nil
	}
	l.rooms[boardID] = room
	l.roomsMu.Unlock()
	return room, nil
}

func nextZ(items []mongoModels.BoardItem) int64 {
	var maxZ int64
	for i := range items {
		if items[i].ZIndex > maxZ {
			maxZ = items[i].ZIndex
		}
	}
	return maxZ + 1
}
