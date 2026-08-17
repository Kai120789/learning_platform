package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"io"
	"learning-platform/api-gateway/internal/dto/enum"
	"learning-platform/api-gateway/internal/dto/materialDto"
	"learning-platform/api-gateway/internal/dto/mediaDto"
	"net/http"
	"strconv"
)

type MaterialHandler struct {
	service MaterialService
	logger  *zap.Logger
}

type MaterialService interface {
	CreateFolder(folder materialDto.CreateFolder) (*materialDto.MaterialFolder, error)
	MoveFolders(folderIDs []int64, parentFolderID *int64) error
	RenameFolder(folderID int64, newTitle string) error
	DeleteOneFolder(folderID int64) error
	DeleteFolders(folderIDs []int64) error
	GetMaterials(userID int64, folderID *int64, role enum.UserRole) (*materialDto.FoldersAndMaterials, error)
	CreateMaterials(tutorID int64, folderID *int64, files []mediaDto.FileDataType) ([]materialDto.Material, error)
	MoveMaterials(materialIDs []int64, folderID *int64) error
	RenameMaterial(materialID int64, newTitle string) error
	DeleteOneMaterial(materialID int64) error
	DeleteMaterials(materialIDs []int64) error
	UpdateUsersMaterialsAccess(userIDs, materialIDs []int64) error
	UpdateUsersFoldersAccess(userIDs, folderIDs []int64) error
}

func NewMaterialHandler(service MaterialService, logger *zap.Logger) *MaterialHandler {
	return &MaterialHandler{
		service: service,
		logger:  logger,
	}
}

func (m *MaterialHandler) CreateFolder(w http.ResponseWriter, r *http.Request) {
	var folder materialDto.CreateFolder
	err := json.NewDecoder(r.Body).Decode(&folder)
	if err != nil {
		m.logger.Error("invalid folder dto", zap.Error(err))
		http.Error(w, "invalid folder dto", http.StatusBadRequest)
		return
	}

	res, err := m.service.CreateFolder(folder)
	if err != nil {
		m.logger.Error("failed to create folder", zap.Error(err))
		http.Error(w, "failed to create folder", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (m *MaterialHandler) MoveFolders(w http.ResponseWriter, r *http.Request) {
	var moveFoldersDTO materialDto.MoveFolders
	err := json.NewDecoder(r.Body).Decode(&moveFoldersDTO)
	if err != nil {
		m.logger.Error("invalid move folders dto", zap.Error(err))
		http.Error(w, "invalid move folders dto", http.StatusBadRequest)
		return
	}

	err = m.service.MoveFolders(moveFoldersDTO.FolderIDs, moveFoldersDTO.ParentFolderID)
	if err != nil {
		m.logger.Error("failed to move folders", zap.Error(err))
		http.Error(w, "failed to move folders", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (m *MaterialHandler) RenameFolder(w http.ResponseWriter, r *http.Request) {
	strFolderID := chi.URLParam(r, "folderID")
	folderID, err := strconv.Atoi(strFolderID)
	if err != nil {
		m.logger.Error("invalid param folder id", zap.Error(err))
		http.Error(w, "invalid param folder id", http.StatusBadRequest)
		return
	}
	title := r.URL.Query().Get("title")

	err = m.service.RenameFolder(int64(folderID), title)
	if err != nil {
		m.logger.Error("failed to rename folder", zap.Error(err))
		http.Error(w, "failed to rename folder", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (m *MaterialHandler) DeleteOneFolder(w http.ResponseWriter, r *http.Request) {
	strFolderID := chi.URLParam(r, "folderID")
	folderID, err := strconv.Atoi(strFolderID)
	if err != nil {
		m.logger.Error("invalid param folder id", zap.Error(err))
		http.Error(w, "invalid param folder id", http.StatusBadRequest)
		return
	}

	err = m.service.DeleteOneFolder(int64(folderID))
	if err != nil {
		m.logger.Error("failed to delete one folder", zap.Error(err))
		http.Error(w, "failed to delete one folder", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (m *MaterialHandler) DeleteFolders(w http.ResponseWriter, r *http.Request) {
	type deleteFoldersDTO struct {
		FolderIDs []int64 `json:"folder_ids"`
	}
	var deleteFolders deleteFoldersDTO
	err := json.NewDecoder(r.Body).Decode(&deleteFolders)
	if err != nil {
		m.logger.Error("invalid delete folders dto", zap.Error(err))
		http.Error(w, "invalid delete folders dto", http.StatusBadRequest)
		return
	}

	err = m.service.DeleteFolders(deleteFolders.FolderIDs)
	if err != nil {
		m.logger.Error("failed to delete folders", zap.Error(err))
		http.Error(w, "failed to delete folders", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (m *MaterialHandler) GetMaterials(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		m.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	role, ok := r.Context().Value("role").(enum.UserRole)
	if !ok {
		m.logger.Error(
			"user unauthorized",
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	strFolderID := chi.URLParam(r, "folderID")
	folderID, err := strconv.Atoi(strFolderID)
	if err != nil {
		m.logger.Error("invalid param folder id", zap.Error(err))
		http.Error(w, "invalid param folder id", http.StatusBadRequest)
		return
	}

	fid := int64(folderID)

	var folderIDPtr *int64
	if folderID != 0 {
		folderIDPtr = &fid
	}

	res, err := m.service.GetMaterials(userID, folderIDPtr, role)
	if err != nil {
		m.logger.Error("failed to get materials and folders", zap.Error(err))
		http.Error(w, "failed to get materials and folders", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (m *MaterialHandler) CreateMaterials(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		m.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	strFolderID := chi.URLParam(r, "folderID")
	folderID, err := strconv.Atoi(strFolderID)
	if err != nil {
		m.logger.Error("invalid param folder id", zap.Error(err))
		http.Error(w, "invalid param folder id", http.StatusBadRequest)
		return
	}

	fid := int64(folderID)

	var folderIDPtr *int64
	if folderID != 0 {
		folderIDPtr = &fid
	}

	if err := r.ParseMultipartForm(100 << 20); err != nil {
		m.logger.Error("failed to parse multipart form", zap.Error(err))
		http.Error(w, "invalid multipart form", http.StatusBadRequest)
		return
	}

	fileHeaders := r.MultipartForm.File["files"]
	if len(fileHeaders) == 0 {
		http.Error(w, "files are required", http.StatusBadRequest)
		return
	}

	files := make([]mediaDto.FileDataType, 0, len(fileHeaders))

	for _, header := range fileHeaders {
		file, err := header.Open()
		if err != nil {
			m.logger.Error(
				"failed to open uploaded file",
				zap.String("fileName", header.Filename),
				zap.Error(err),
			)
			http.Error(w, "failed to open file", http.StatusBadRequest)
			return
		}

		data, err := io.ReadAll(file)
		_ = file.Close()

		if err != nil {
			m.logger.Error(
				"failed to read uploaded file",
				zap.String("fileName", header.Filename),
				zap.Error(err),
			)
			http.Error(w, "failed to read file", http.StatusBadRequest)
			return
		}

		contentType := header.Header.Get("Content-Type")
		if contentType == "" {
			contentType = http.DetectContentType(data)
		}

		files = append(files, mediaDto.FileDataType{
			FileMetadata: mediaDto.FileMetadata{
				FileName:    header.Filename,
				ContentType: contentType,
				Size:        uint64(len(data)),
			},
			Data: data,
		})
	}

	res, err := m.service.CreateMaterials(userID, folderIDPtr, files)
	if err != nil {
		m.logger.Error("failed to create materials", zap.Error(err))
		http.Error(w, "failed to create materials", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (m *MaterialHandler) MoveMaterials(w http.ResponseWriter, r *http.Request) {
	var moveMaterials materialDto.MoveMaterials
	err := json.NewDecoder(r.Body).Decode(&moveMaterials)
	if err != nil {
		m.logger.Error("invalid move materials dto", zap.Error(err))
		http.Error(w, "invalid move materials dto", http.StatusBadRequest)
		return
	}

	err = m.service.MoveMaterials(moveMaterials.MaterialIDs, moveMaterials.FolderID)
	if err != nil {
		m.logger.Error("failed to move materials", zap.Error(err))
		http.Error(w, "failed to move materials", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (m *MaterialHandler) RenameMaterial(w http.ResponseWriter, r *http.Request) {
	strMaterialID := chi.URLParam(r, "materialID")
	materialID, err := strconv.Atoi(strMaterialID)
	if err != nil {
		m.logger.Error("invalid param material id", zap.Error(err))
		http.Error(w, "invalid param material id", http.StatusBadRequest)
		return
	}

	title := r.URL.Query().Get("title")

	err = m.service.RenameMaterial(int64(materialID), title)
	if err != nil {
		m.logger.Error("failed to rename material", zap.Error(err))
		http.Error(w, "failed to rename material", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (m *MaterialHandler) DeleteOneMaterial(w http.ResponseWriter, r *http.Request) {
	strMaterialID := chi.URLParam(r, "materialID")
	materialID, err := strconv.Atoi(strMaterialID)
	if err != nil {
		m.logger.Error("invalid param material id", zap.Error(err))
		http.Error(w, "invalid param material id", http.StatusBadRequest)
		return
	}

	err = m.service.DeleteOneMaterial(int64(materialID))
	if err != nil {
		m.logger.Error("failed to delete one material", zap.Error(err))
		http.Error(w, "failed to delete one material", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (m *MaterialHandler) DeleteMaterials(w http.ResponseWriter, r *http.Request) {
	type deleteMaterialsDTO struct {
		MaterialIDs []int64 `json:"material_ids"`
	}
	var deleteMaterials deleteMaterialsDTO
	err := json.NewDecoder(r.Body).Decode(&deleteMaterials)
	if err != nil {
		m.logger.Error("invalid delete materials dto", zap.Error(err))
		http.Error(w, "invalid delete materials dto", http.StatusBadRequest)
		return
	}

	err = m.service.DeleteMaterials(deleteMaterials.MaterialIDs)
	if err != nil {
		m.logger.Error("failed to delete materials", zap.Error(err))
		http.Error(w, "failed to delete materials", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (m *MaterialHandler) UpdateUsersMaterialsAccess(w http.ResponseWriter, r *http.Request) {
	var materialsAccess materialDto.MaterialsAccess
	err := json.NewDecoder(r.Body).Decode(&materialsAccess)
	if err != nil {
		m.logger.Error("invalid update users materials access dto", zap.Error(err))
		http.Error(w, "invalid update users materials access dto", http.StatusBadRequest)
		return
	}

	err = m.service.UpdateUsersMaterialsAccess(materialsAccess.UserIDs, materialsAccess.MaterialIDs)
	if err != nil {
		m.logger.Error("failed to update users materials access", zap.Error(err))
		http.Error(w, "failed to update users materials access", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (m *MaterialHandler) UpdateUsersFoldersAccess(w http.ResponseWriter, r *http.Request) {
	var foldersAccess materialDto.FoldersAccess
	err := json.NewDecoder(r.Body).Decode(&foldersAccess)
	if err != nil {
		m.logger.Error("invalid update users folders access dto", zap.Error(err))
		http.Error(w, "invalid update users folders access dto", http.StatusBadRequest)
		return
	}

	err = m.service.UpdateUsersFoldersAccess(foldersAccess.UserIDs, foldersAccess.FolderIDs)
	if err != nil {
		m.logger.Error("failed to update users folders access", zap.Error(err))
		http.Error(w, "failed to update users folders access", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
