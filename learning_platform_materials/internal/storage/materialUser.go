package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MaterialUserStorage struct {
	conn *pgxpool.Pool
}

func NewMaterialUserStorage(conn *pgxpool.Pool) *MaterialUserStorage {
	return &MaterialUserStorage{
		conn: conn,
	}
}

func (mu *MaterialUserStorage) UpdateUsersMaterialsAccess(userIDs, materialIDs []int64) error {
	query := `
		WITH deleted AS (
			DELETE FROM material_users
			WHERE material_id = ANY($2)
		)
		INSERT INTO material_users (user_id, material_id)
		SELECT u.user_id, m.material_id
		FROM unnest($1::bigint[]) AS u(user_id)
		CROSS JOIN unnest($2::bigint[]) AS m(material_id);
	`

	_, err := mu.conn.Exec(
		context.Background(),
		query,
		userIDs,
		materialIDs,
	)
	if err != nil {
		return fmt.Errorf("update users materials access: %w", err)
	}

	return nil
}

func (mu *MaterialUserStorage) UpdateUsersFoldersAccess(userIDs, folderIDs []int64) error {
	query := `
		WITH deleted AS (
			DELETE FROM folder_users
			WHERE folder_id = ANY($2)
		)
		INSERT INTO folder_users (user_id, folder_id)
		SELECT u.user_id, f.folder_id
		FROM unnest($1::bigint[]) AS u(user_id)
		CROSS JOIN unnest($2::bigint[]) AS f(folder_id);
	`

	_, err := mu.conn.Exec(
		context.Background(),
		query,
		userIDs,
		folderIDs,
	)
	if err != nil {
		return fmt.Errorf("update users folders access: %w", err)
	}

	return nil
}
