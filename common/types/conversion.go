package types

import (
	"encoding/json"

	"github.com/jackc/pgx/v5/pgtype"
)

func ConvertTimestampToUnix(ts pgtype.Timestamp) int64 {
    return ts.Time.Unix()
}

func TypeConverter[R any](data any) (*R, error) {
    var result R
    b, err := json.Marshal(&data)
    if err != nil {
      return nil, err
    }
    err = json.Unmarshal(b, &result)
    if err != nil {
      return nil, err
    }
    return &result, err
}