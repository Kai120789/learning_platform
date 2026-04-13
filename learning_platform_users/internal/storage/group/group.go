package group

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type GroupStorage struct {
	GroupBaseStorage *GroupBaseStorage
	GroupUserStorage *GroupUserStorage
}

func NewGroupStorage(
	conn *pgxpool.Pool,
) *GroupStorage {
	return &GroupStorage{
		GroupUserStorage: NewGroupUserStorage(conn),
		GroupBaseStorage: NewGroupBaseStorage(conn),
	}
}
