package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"learning-platform/materials/internal/dto"
	"learning-platform/materials/internal/models"
)

type MaterialFolderStorage struct {
	conn *pgxpool.Pool
}

func NewMaterialFolderStorage(conn *pgxpool.Pool) *MaterialFolderStorage {
	return &MaterialFolderStorage{
		conn: conn,
	}
}

func (mf *MaterialFolderStorage) CreateFolder(folder dto.CreateFolder) (*models.MaterialFolder, error) {
	var resFolder models.MaterialFolder
	query := `
		INSERT INTO material_folders (title, parent_folder_id, tutor_id)
		VALUES ($1, $2, $3)
		RETURNING id, title, parent_folder_id, tutor_id
	`

	err := mf.conn.QueryRow(
		context.Background(),
		query,
		folder.Title,
		folder.ParentFolderID,
		folder.TutorID,
	).Scan(
		&resFolder.ID,
		&resFolder.Title,
		&resFolder.ParentFolderID,
		&resFolder.TutorID,
	)
	if err != nil {
		return nil, fmt.Errorf("create folder: %w", err)
	}

	return &resFolder, nil
}

func (mf *MaterialFolderStorage) MoveFolders(folderIDs []int64, newParentFolderID *int64) error {
	query := `
		WITH RECURSIVE chain AS (
			SELECT id, parent_folder_id
			FROM material_folders
			WHERE id = $2::bigint
			UNION
			SELECT f.id, f.parent_folder_id
			FROM chain c
			JOIN material_folders f ON f.id = c.parent_folder_id
		)
		UPDATE material_folders
		SET parent_folder_id = $2::bigint
		WHERE id = ANY($1)
		  AND NOT EXISTS (
			  SELECT 1
			  FROM chain
			  WHERE chain.id = material_folders.id
		  )
	`

	tag, err := mf.conn.Exec(context.Background(),
		query,
		folderIDs,
		newParentFolderID,
	)
	if err != nil {
		return fmt.Errorf("move folder: %w", err)
	}
	if tag.RowsAffected() != int64(len(folderIDs)) {
		return errors.New("folder not found or move would create a cycle")
	}

	return nil
}

func (mf *MaterialFolderStorage) RenameFolder(folderID int64, newFolderTitle string) error {
	query := `
		UPDATE material_folders
		SET title = $2
		WHERE id = $1
	`

	_, err := mf.conn.Exec(
		context.Background(),
		query,
		folderID,
		newFolderTitle,
	)
	if err != nil {
		return fmt.Errorf("rename folder: %w", err)
	}

	return nil
}

func (mf *MaterialFolderStorage) DeleteOneFolder(folderID int64) error {
	query := `
		DELETE FROM material_folders
		WHERE id = $1
	`

	_, err := mf.conn.Exec(
		context.Background(),
		query,
		folderID,
	)
	if err != nil {
		return fmt.Errorf("delete one folder: %w", err)
	}

	return nil
}

func (mf *MaterialFolderStorage) DeleteFolders(folderIDs []int64) error {
	query := `
		DELETE FROM material_folders
		WHERE id = ANY($1)
	`

	_, err := mf.conn.Exec(
		context.Background(),
		query,
		folderIDs,
	)
	if err != nil {
		return fmt.Errorf("delete folders: %w", err)
	}

	return nil
}

func (mf *MaterialFolderStorage) GetStudentFolders(studentID int64, folderID *int64) ([]models.MaterialFolder, error) {
	var resFolders []models.MaterialFolder
	query := `
		WITH RECURSIVE
		parent_chain AS (
			SELECT f.id, f.parent_folder_id
			FROM material_folders f
			WHERE f.id = $2::bigint
		
			UNION ALL
		
			SELECT f.id, f.parent_folder_id
			FROM parent_chain pc
			JOIN material_folders f ON f.id = pc.parent_folder_id
		),
		
		inherited AS (
			SELECT EXISTS (
				SELECT 1
				FROM folder_users fu
				JOIN parent_chain pc ON pc.id = fu.folder_id
				WHERE fu.user_id = $1
			) AS ok
		),
		
		access_points AS (
			SELECT fu.folder_id
			FROM folder_users fu
			WHERE fu.user_id = $1
		
			UNION
		
			SELECT m.folder_id
			FROM material_users mu
			JOIN materials m ON m.id = mu.material_id
			WHERE mu.user_id = $1 AND m.folder_id IS NOT NULL
		),
		
		visible_path AS (
			SELECT ap.folder_id AS id
			FROM access_points ap
		
			UNION
		
			SELECT f.parent_folder_id
			FROM visible_path vp
			JOIN material_folders f ON f.id = vp.id
			WHERE f.parent_folder_id IS NOT NULL
		)
		
		SELECT f.id, f.title, f.parent_folder_id, f.tutor_id
		FROM material_folders f
		WHERE f.parent_folder_id IS NOT DISTINCT FROM $2::bigint
		  AND (
				(SELECT ok FROM inherited)
				OR f.id IN (SELECT id FROM visible_path)
			  )
		ORDER BY f.id
	`

	rows, err := mf.conn.Query(
		context.Background(),
		query,
		studentID,
		folderID,
	)
	if err != nil {
		return nil, fmt.Errorf("get student folders: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var oneFolder models.MaterialFolder
		err := rows.Scan(
			&oneFolder.ID,
			&oneFolder.Title,
			&oneFolder.ParentFolderID,
			&oneFolder.TutorID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan one student folder: %w", err)
		}

		resFolders = append(resFolders, oneFolder)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate student folders: %w", err)
	}

	return resFolders, nil
}

func (mf *MaterialFolderStorage) GetTutorFolders(tutorID int64, folderID *int64) ([]models.MaterialFolder, error) {
	var resFolders []models.MaterialFolder
	query := `
		SELECT id, title, parent_folder_id, tutor_id
		FROM material_folders
		WHERE tutor_id = $1 AND parent_folder_id IS NOT DISTINCT FROM $2::bigint
		ORDER BY id
	`

	rows, err := mf.conn.Query(
		context.Background(),
		query,
		tutorID,
		folderID,
	)
	if err != nil {
		return nil, fmt.Errorf("get tutor folders: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var oneFolder models.MaterialFolder
		err := rows.Scan(
			&oneFolder.ID,
			&oneFolder.Title,
			&oneFolder.ParentFolderID,
			&oneFolder.TutorID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan one tutor folder: %w", err)
		}

		resFolders = append(resFolders, oneFolder)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate tutor folders: %w", err)
	}

	return resFolders, nil
}
