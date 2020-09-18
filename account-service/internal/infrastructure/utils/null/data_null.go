package user_service

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

func TimeToNull(t time.Time) pq.NullTime {
	r := pq.NullTime{}
	r.Time = t

	if !t.IsZero() {
		r.Valid = true
	}

	return r
}

func StringToNull(s string) sql.NullString {
	r := sql.NullString{}
	r.String = s

	if s != "" {
		r.Valid = true
	}

	return r
}

func Int64ToNull(i int64) sql.NullInt64 {
	r := sql.NullInt64{}
	r.Int64 = i

	if i > 0 {
		r.Valid = true
	}

	return r
}

func Float64ToNull(f float64) sql.NullFloat64 {
	r := sql.NullFloat64{}
	r.Float64 = f

	if f > 0 {
		r.Valid = true
	}
	
	return r
}

func BollToNull(b *bool) sql.NullBool {
	r := sql.NullBool{}

	if b != nil {
		r.Bool = *b
		r.Valid = true
	}
	
	return r
}