package store

import "fmt"

var ErrUniqueConstraint = fmt.Errorf("unique constraint violation")
var ErrNotFound = fmt.Errorf("data not found")
