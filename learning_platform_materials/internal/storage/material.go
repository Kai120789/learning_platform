package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"learning-platform/materials/internal/dto"
	"learning-platform/materials/internal/models"
	"strings"
)

type MaterialStorage struct {
	conn *pgxpool.Pool
}

func NewMaterialStorage(conn *pgxpool.Pool) *MaterialStorage {
	return &MaterialStorage{
		conn: conn,
	}
}

func (m *MaterialStorage) GetStudentMaterials(studentID int64, folderID *int64) ([]models.Material, error) {
	var resMaterials []models.Material
	query := `
		SELECT m.id, m.title, m.size, m.folder_id, m.tutor_id, m.mime_type, m.media_object_id
		FROM materials AS m
		JOIN material_users AS mu
		ON mu.material_id = m.id
		WHERE mu.user_id = $1 AND m.folder_id IS NOT DISTINCT FROM $2::bigint
	`

	rows, err := m.conn.Query(
		context.Background(),
		query,
		studentID,
		folderID,
	)
	if err != nil {
		return nil, fmt.Errorf("get student materials: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var oneMaterial models.Material
		err := rows.Scan(
			&oneMaterial.ID,
			&oneMaterial.Title,
			&oneMaterial.Size,
			&oneMaterial.FolderID,
			&oneMaterial.TutorID,
			&oneMaterial.MimeType,
			&oneMaterial.MediaObjectID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan one student material: %w", err)
		}

		resMaterials = append(resMaterials, oneMaterial)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate student materials: %w", err)
	}

	return resMaterials, nil
}

func (m *MaterialStorage) GetTutorMaterials(tutorID int64, folderID *int64) ([]models.Material, error) {
	var resMaterials []models.Material
	query := `
		SELECT id, title, size, folder_id, tutor_id, mime_type, media_object_id
		FROM materials
		WHERE tutor_id = $1 AND folder_id IS NOT DISTINCT FROM $2::bigint
	`

	rows, err := m.conn.Query(
		context.Background(),
		query,
		tutorID,
		folderID,
	)
	if err != nil {
		return nil, fmt.Errorf("get tutor materials: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var oneMaterial models.Material
		err := rows.Scan(
			&oneMaterial.ID,
			&oneMaterial.Title,
			&oneMaterial.Size,
			&oneMaterial.FolderID,
			&oneMaterial.TutorID,
			&oneMaterial.MimeType,
			&oneMaterial.MediaObjectID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan one tutor material: %w", err)
		}

		resMaterials = append(resMaterials, oneMaterial)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate tutor materials: %w", err)
	}

	return resMaterials, nil
}

func (m *MaterialStorage) CreateMaterials(materials []dto.CreateMaterial) ([]models.Material, error) {
	args := make([]any, 0, len(materials)*6)
	values := make([]string, 0, len(materials))

	for i, material := range materials {
		offset := i * 6

		values = append(values,
			fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)",
				offset+1,
				offset+2,
				offset+3,
				offset+4,
				offset+5,
				offset+6,
			),
		)

		args = append(args,
			material.Title,
			material.Size,
			material.FolderID,
			material.TutorID,
			material.MimeType,
			material.MediaObjectID,
		)
	}

	var resMaterials []models.Material
	query := fmt.Sprintf(`
		INSERT INTO materials (title, size, folder_id, tutor_id, mime_type, media_object_id)
		VALUES %s
		RETURNING id, title, size, folder_id, tutor_id, mime_type, media_object_id
	`, strings.Join(values, ", "))

	rows, err := m.conn.Query(
		context.Background(),
		query,
		args...,
	)
	if err != nil {
		return nil, fmt.Errorf("create materials: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var oneMaterial models.Material
		err := rows.Scan(
			&oneMaterial.ID,
			&oneMaterial.Title,
			&oneMaterial.Size,
			&oneMaterial.FolderID,
			&oneMaterial.TutorID,
			&oneMaterial.MimeType,
			&oneMaterial.MediaObjectID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan one material: %w", err)
		}

		resMaterials = append(resMaterials, oneMaterial)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate materials: %w", err)
	}

	return resMaterials, nil
}

func (m *MaterialStorage) MoveMaterials(materialIDs []int64, folderID *int64) error {
	query := `
		UPDATE materials
		SET folder_id = $2::bigint
		WHERE id = ANY($1)
	`

	_, err := m.conn.Exec(
		context.Background(),
		query,
		materialIDs,
		folderID,
	)
	if err != nil {
		return fmt.Errorf("move materials: %w", err)
	}

	return nil
}

func (m *MaterialStorage) RenameMaterial(materialID int64, newTitle string) error {
	query := `
		UPDATE materials
		SET title = $2
		WHERE id = $1
	`

	_, err := m.conn.Exec(
		context.Background(),
		query,
		materialID,
		newTitle,
	)
	if err != nil {
		return fmt.Errorf("rename material: %w", err)
	}

	return nil
}

func (m *MaterialStorage) DeleteOneMaterial(materialID int64) error {
	query := `
		DELETE FROM materials
		WHERE id = $1
	`

	_, err := m.conn.Exec(
		context.Background(),
		query,
		materialID,
	)
	if err != nil {
		return fmt.Errorf("delete one material: %w", err)
	}

	return nil
}

func (m *MaterialStorage) DeleteMaterials(materialIDs []int64) error {
	query := `
		DELETE FROM materials
		WHERE id = ANY($1)
	`

	_, err := m.conn.Exec(
		context.Background(),
		query,
		materialIDs,
	)
	if err != nil {
		return fmt.Errorf("delete materials: %w", err)
	}

	return nil
}
