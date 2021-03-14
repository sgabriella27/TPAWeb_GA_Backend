package database

import (
	"gorm.io/gorm"
)

func Paginate(r int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := r

		pageSize := 10

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
