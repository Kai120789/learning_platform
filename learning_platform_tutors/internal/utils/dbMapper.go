package utils

import (
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func DBInt8ToOptional(val pgtype.Int8) *int64 {
	if val.Valid {
		return &val.Int64
	} else {
		return nil
	}
}

func DBStringToOptional(val pgtype.Text) *string {
	if val.Valid {
		return &val.String
	} else {
		return nil
	}
}

func DBTimeToOptional(val pgtype.Timestamptz) *timestamppb.Timestamp {
	if val.Valid {
		return timestamppb.New(val.Time)
	} else {
		return nil
	}
}
