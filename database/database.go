package database

import (
	"gorm.io/gorm"
)

// Database instance
type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

