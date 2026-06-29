package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"learning-platform/lessons/internal/dto"
	"learning-platform/lessons/internal/models"
)

type LessonMediaStorage struct {
	conn *pgxpool.Pool
}

func NewLessonMediaStorage(conn *pgxpool.Pool) *LessonMediaStorage {
	return &LessonMediaStorage{
		conn: conn,
	}
}

func (lm *LessonMediaStorage) GetAllLessonMedias(lessonID int64) ([]models.LessonMedia, error) {
	var resLessonMedias []models.LessonMedia
	query := `
		SELECT id, lesson_id, s3_link, s3_preview, type
		FROM lesson_medias
		WHERE lesson_id = $1
	`

	rows, err := lm.conn.Query(context.Background(), query, lessonID)
	if err != nil {
		return nil, fmt.Errorf("get all medias for lesson %d from db: %w", lessonID, err)
	}

	for rows.Next() {
		var oneLesson models.LessonMedia

		err := rows.Scan(
			&oneLesson.ID,
			&oneLesson.LessonID,
			&oneLesson.S3Link,
			&oneLesson.S3Preview,
			&oneLesson.Type,
		)
		if err != nil {
			return nil, fmt.Errorf("scan one lesson media from db %w: ", err)
		}

		resLessonMedias = append(resLessonMedias, oneLesson)
	}

	return resLessonMedias, nil
}

func (lm *LessonMediaStorage) SetLessonMedias(lessonID int64, lessonMedias []dto.MediaItem) error {
	mediasCount := len(lessonMedias)
	s3Links := make([]string, 0, mediasCount)
	s3Previews := make([]string, 0, mediasCount)
	types := make([]string, 0, mediasCount)
	lessonIDs := make([]int64, 0, mediasCount)

	for _, oneMedia := range lessonMedias {
		s3Links = append(s3Links, oneMedia.S3Link)
		s3Previews = append(s3Previews, oneMedia.S3Preview)
		types = append(types, string(oneMedia.Type))
		lessonIDs = append(lessonIDs, lessonID)
	}

	fmt.Println(lessonID)

	query := `
		INSERT INTO lesson_medias (
			s3_link,
			s3_preview,
			type,
			lesson_id
		)
		SELECT *
		FROM unnest(
			$1::text[],
			$2::text[],
			$3::type_enum[],
			$4::bigint[]
		)
	`

	_, err := lm.conn.Exec(
		context.Background(),
		query,
		s3Links,
		s3Previews,
		types,
		lessonIDs,
	)
	if err != nil {
		return fmt.Errorf("add medias to lesson %d: %w", lessonID, err)
	}

	return nil
}

func (lm *LessonMediaStorage) DeleteLessonMedias(lessonID int64, mediaIDs []int64) error {
	query := `
		DELETE FROM lesson_medias
		WHERE lesson_id = $1
		AND id = ANY($2::bigint[])
	`

	_, err := lm.conn.Exec(
		context.Background(),
		query,
		lessonID,
		mediaIDs,
	)
	if err != nil {
		return fmt.Errorf("delete lesson medias : %w", err)
	}

	return nil
}
