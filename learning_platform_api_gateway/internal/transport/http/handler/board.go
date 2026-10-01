package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"learning-platform/api-gateway/internal/dto/boardDto"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

type BoardHandler struct {
	service      BoardService
	logger       *zap.Logger
	boardHTTPURL string
}

type BoardService interface {
	GetBoard(boardID int64) (*boardDto.Board, error)
	CreateBoard(title string, tutorID int64) (*boardDto.Board, error)
	GetTutorBoards(tutorID int64) ([]boardDto.Board, error)
	DeleteBoard(boardID, tutorID int64) error
	UpdateBoard(title string, tutorID, boardID int64) (*boardDto.Board, error)
}

func NewBoardHandler(service BoardService, logger *zap.Logger, boardHTTPURL string) *BoardHandler {
	return &BoardHandler{
		service:      service,
		logger:       logger,
		boardHTTPURL: boardHTTPURL,
	}
}

func (b *BoardHandler) GetBoard(w http.ResponseWriter, r *http.Request) {
	strBoardID := chi.URLParam(r, "boardID")

	boardID, err := strconv.Atoi(strBoardID)
	if err != nil {
		b.logger.Error("invalid param board id", zap.Error(err))
		http.Error(w, "invalid param board id", http.StatusBadRequest)
		return
	}

	res, err := b.service.GetBoard(int64(boardID))
	if err != nil {
		b.logger.Error("failed to get board", zap.Error(err))
		http.Error(w, "failed to get board", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (b *BoardHandler) CreateBoard(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		b.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var req boardDto.BoardTitleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid board dto", http.StatusBadRequest)
		return
	}

	res, err := b.service.CreateBoard(req.Title, userID)
	if err != nil {
		b.logger.Error("failed to create board", zap.Error(err))
		http.Error(w, "failed to create board", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (b *BoardHandler) GetTutorBoards(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		b.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	res, err := b.service.GetTutorBoards(userID)
	if err != nil {
		b.logger.Error("failed to get tutor boards", zap.Error(err))
		http.Error(w, "failed to get tutor boards", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (b *BoardHandler) DeleteBoard(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		b.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	strBoardID := chi.URLParam(r, "boardID")

	boardID, err := strconv.Atoi(strBoardID)
	if err != nil {
		b.logger.Error("invalid param board id", zap.Error(err))
		http.Error(w, "invalid param board id", http.StatusBadRequest)
		return
	}

	err = b.service.DeleteBoard(int64(boardID), userID)
	if err != nil {
		b.logger.Error("failed to delete board", zap.Error(err))
		http.Error(w, "failed to delete board", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (b *BoardHandler) UpdateBoard(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		b.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	strBoardID := chi.URLParam(r, "boardID")

	boardID, err := strconv.Atoi(strBoardID)
	if err != nil {
		b.logger.Error("invalid param board id", zap.Error(err))
		http.Error(w, "invalid param board id", http.StatusBadRequest)
		return
	}

	var req boardDto.BoardTitleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid board dto", http.StatusBadRequest)
		return
	}

	res, err := b.service.UpdateBoard(req.Title, userID, int64(boardID))
	if err != nil {
		b.logger.Error("failed to update board", zap.Error(err))
		http.Error(w, "failed to update board", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (b *BoardHandler) ProxyBoardWS(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		b.logger.Error("user unauthorized", zap.Int64("userID", userID))
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	strBoardID := chi.URLParam(r, "boardID")
	boardID, err := strconv.ParseInt(strBoardID, 10, 64)
	if err != nil || boardID <= 0 {
		b.logger.Error("invalid param board id", zap.Error(err))
		http.Error(w, "invalid param board id", http.StatusBadRequest)
		return
	}

	target, err := url.Parse(b.boardHTTPURL)
	if err != nil || target.Host == "" {
		b.logger.Error("invalid board http url", zap.String("url", b.boardHTTPURL), zap.Error(err))
		http.Error(w, "board ws proxy", http.StatusInternalServerError)
		return
	}

	proxy := &httputil.ReverseProxy{
		Rewrite: func(pr *httputil.ProxyRequest) {
			pr.SetURL(target)
			pr.Out.URL.Path = "/ws/" + strconv.FormatInt(boardID, 10)
			pr.Out.URL.RawPath = ""
			pr.Out.Header.Set("X-User-Id", strconv.FormatInt(userID, 10))
		},
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			b.logger.Error("board ws proxy", zap.Error(err))
			http.Error(w, "board ws proxy", http.StatusBadGateway)
		},
	}

	proxy.ServeHTTP(w, r)
}
