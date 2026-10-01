package ws

import (
	"github.com/coder/websocket"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func NewWSRouter(hub *Hub) http.Handler {
	r := chi.NewRouter()
	r.Get("/ws/{boardID}", func(w http.ResponseWriter, r *http.Request) {
		handleBoardWS(w, r, hub)
	})
	return r
}

func handleBoardWS(w http.ResponseWriter, r *http.Request, hub *Hub) {
	strBoardID := chi.URLParam(r, "boardID")
	strUserID := r.Header.Get("X-User-Id")
	if strUserID == "" {
		http.Error(w, "missing user", http.StatusUnauthorized)
		return
	}

	boardID, err := strconv.ParseInt(strBoardID, 10, 64)
	if err != nil {
		return
	}

	userID, err := strconv.ParseInt(strUserID, 10, 64)
	if err != nil {
		return
	}

	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return
	}
	defer conn.CloseNow()

	hub.Join(strBoardID, conn)
	defer hub.Leave(strBoardID, conn)

	snapshot, err := hub.Snapshot(boardID)
	if err != nil {
		conn.Close(websocket.StatusInternalError, "snapshot")
		return
	}
	if err := conn.Write(r.Context(), websocket.MessageText, snapshot); err != nil {
		return
	}

	for {
		_, payload, err := conn.Read(r.Context())
		if err != nil {
			return
		}
		hub.Broadcast(r.Context(), boardID, userID, conn, payload)
	}
}
