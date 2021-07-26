package repositories

import (
	"gorm.io/gorm"
)

type baseRepository struct {
	dbCtx *gorm.DB
}