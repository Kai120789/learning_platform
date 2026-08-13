package dto

import "learning-platform/materials/internal/models"

type FoldersAndMaterials struct {
	Folders   []models.MaterialFolder `json:"folders"`
	Materials []models.Material       `json:"materials"`
}
