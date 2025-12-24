package store

import "fmt"

var ErrUniqueConstraint = fmt.Errorf("unique constraint violation")
