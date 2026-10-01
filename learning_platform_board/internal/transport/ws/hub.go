package ws

import (
	"context"
	"encoding/json"
	"github.com/coder/websocket"
	"learning-platform/board/internal/dto"
	"learning-platform/board/internal/models/enum"
	mongoModels "learning-platform/board/internal/models/mongo"
	"strconv"
	"sync"
)

type Hub struct {
	mu    sync.Mutex
	rooms map[string]map[*websocket.Conn]struct{}
	live  LiveService
}

type LiveService interface {
	AddBoardEvent(id, userID int64, msg dto.AddMessage) error
	UpdateBoardEvent(id, userID int64, msg dto.UpdateMessage) error
	DeleteBoardEvents(id, userID int64, msg dto.DeleteMessage) error
	CancelLastBoardEvent(eventID string) error
	GetSnapshot(boardID int64) (*mongoModels.BoardSnapshot, error)
}

func NewHub(live LiveService) *Hub {
	return &Hub{
		rooms: make(map[string]map[*websocket.Conn]struct{}),
		live:  live,
	}
}

func (h *Hub) Join(boardID string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.rooms[boardID] == nil {
		h.rooms[boardID] = make(map[*websocket.Conn]struct{})
	}
	h.rooms[boardID][conn] = struct{}{}
}

func (h *Hub) Leave(boardID string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.rooms[boardID], conn)
	if len(h.rooms[boardID]) == 0 {
		delete(h.rooms, boardID)
	}
}

func (h *Hub) Snapshot(boardID int64) ([]byte, error) {
	snapshot, err := h.live.GetSnapshot(boardID)
	if err != nil {
		return nil, err
	}

	snap, err := json.Marshal(dto.SnapshotMessage{
		Envelope: dto.Envelope{
			Type:    enum.GetSnapshot,
			BoardID: snapshot.BoardID,
			Version: snapshot.Version,
		},
		Items: snapshot.Items,
	})
	if err != nil {
		return nil, err
	}

	return snap, nil
}

func (h *Hub) Broadcast(
	ctx context.Context,
	boardID int64,
	userID int64,
	from *websocket.Conn,
	payload []byte,
) {
	var head struct {
		Envelope dto.Envelope `json:"envelope"`
	}
	if err := json.Unmarshal(payload, &head); err != nil {
		return
	}

	switch head.Envelope.Type {
	case enum.AddItem:
		var msg dto.AddMessage
		if err := json.Unmarshal(payload, &msg); err != nil {
			return
		}
		if err := h.live.AddBoardEvent(boardID, userID, msg); err != nil {
			return
		}
	case enum.UpdateItem:
		var msg dto.UpdateMessage
		if err := json.Unmarshal(payload, &msg); err != nil {
			return
		}
		if err := h.live.UpdateBoardEvent(boardID, userID, msg); err != nil {
			return
		}
	case enum.DeleteItems:
		var msg dto.DeleteMessage
		if err := json.Unmarshal(payload, &msg); err != nil {
			return
		}
		if err := h.live.DeleteBoardEvents(boardID, userID, msg); err != nil {
			return
		}
	default:
		return
	}

	h.mu.Lock()
	conns := make([]*websocket.Conn, 0, len(h.rooms[strconv.FormatInt(boardID, 10)]))
	for c := range h.rooms[strconv.FormatInt(boardID, 10)] {
		if c != from {
			conns = append(conns, c)
		}
	}
	h.mu.Unlock()

	for _, c := range conns {
		_ = c.Write(ctx, websocket.MessageText, payload)
	}
}
