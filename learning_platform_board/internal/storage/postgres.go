package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	postgresModels "learning-platform/board/internal/models/postgres"
)

type PostgresStorage struct {
	conn *pgxpool.Pool
}

func NewPostgresStorage(conn *pgxpool.Pool) *PostgresStorage {
	return &PostgresStorage{
		conn: conn,
	}
}

func PostgresConnection(connectStr string) (*pgxpool.Pool, error) {
	dbConn, err := pgxpool.New(context.Background(), connectStr)
	if err != nil {
		return nil, fmt.Errorf("unable connect to postgres: %w", err)
	}
	return dbConn, nil
}

func (p *PostgresStorage) GetBoard(boardID int64) (*postgresModels.Board, error) {
	var resBoard postgresModels.Board
	query := `
		SELECT id, title, tutor_id, created_at, updated_at FROM boards
		WHERE id = $1
	`

	err := p.conn.QueryRow(
		context.Background(),
		query,
		boardID,
	).Scan(
		&resBoard.ID,
		&resBoard.Title,
		&resBoard.TutorID,
		&resBoard.CreatedAt,
		&resBoard.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("get board: %w", err)
	}

	return &resBoard, nil
}

func (p *PostgresStorage) CreateBoard(title string, tutorID, tutorBoardsLimit int64) (*postgresModels.Board, error) {
	ctx := context.Background()

	tx, err := p.conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)
	_, err = tx.Exec(ctx, `SELECT pg_advisory_xact_lock($1)`, tutorID)
	if err != nil {
		return nil, err
	}

	var boardsCount int64
	queryLimit := `
		SELECT COUNT(*)
		FROM boards
		WHERE tutor_id = $1
	`

	err = tx.QueryRow(
		context.Background(),
		queryLimit,
		tutorID,
	).Scan(&boardsCount)
	if err != nil {
		return nil, fmt.Errorf("check limit: %w", err)
	}

	if boardsCount >= tutorBoardsLimit {
		return nil, fmt.Errorf("board limit reached")
	}

	var resBoard postgresModels.Board
	query := `
		INSERT INTO boards (title, tutor_id)
		VALUES ($1, $2)
		RETURNING id, title, tutor_id, created_at, updated_at
	`

	err = tx.QueryRow(
		context.Background(),
		query,
		title,
		tutorID,
	).Scan(
		&resBoard.ID,
		&resBoard.Title,
		&resBoard.TutorID,
		&resBoard.CreatedAt,
		&resBoard.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("create board: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &resBoard, nil
}

func (p *PostgresStorage) GetTutorBoards(tutorID int64) ([]postgresModels.Board, error) {
	var resBoards []postgresModels.Board
	query := `
		SELECT id, title, tutor_id, created_at, updated_at FROM boards
		WHERE tutor_id = $1
		ORDER BY updated_at DESC
	`

	rows, err := p.conn.Query(
		context.Background(),
		query,
		tutorID,
	)
	if err != nil {
		return nil, fmt.Errorf("get tutor boards: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var oneBoard postgresModels.Board
		err := rows.Scan(
			&oneBoard.ID,
			&oneBoard.Title,
			&oneBoard.TutorID,
			&oneBoard.CreatedAt,
			&oneBoard.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan one board: %w", err)
		}

		resBoards = append(resBoards, oneBoard)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate boards: %w", err)
	}

	return resBoards, nil
}

func (p *PostgresStorage) DeleteBoard(boardID, tutorID int64) error {
	query := `
		DELETE FROM boards
		WHERE id = $1 AND tutor_id = $2
	`

	_, err := p.conn.Exec(
		context.Background(),
		query,
		boardID,
		tutorID,
	)
	if err != nil {
		return fmt.Errorf("delete board: %w", err)
	}

	return nil
}

func (p *PostgresStorage) UpdateBoard(title string, tutorID, boardID int64) (*postgresModels.Board, error) {
	var resBoard postgresModels.Board
	query := `
		UPDATE boards
		SET 
		    title = $3, 
		    updated_at = now()
		WHERE id = $1 AND tutor_id = $2
		RETURNING id, title, tutor_id, created_at, updated_at
	`

	err := p.conn.QueryRow(
		context.Background(),
		query,
		boardID,
		tutorID,
		title,
	).Scan(
		&resBoard.ID,
		&resBoard.Title,
		&resBoard.TutorID,
		&resBoard.CreatedAt,
		&resBoard.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("update board: %w", err)
	}

	return &resBoard, nil
}
