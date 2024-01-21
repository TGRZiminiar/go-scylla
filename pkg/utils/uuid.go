package utils

import (
	"github.com/gocql/gocql"
)

func GenrateUUId() gocql.UUID {
	return gocql.TimeUUID()
}
