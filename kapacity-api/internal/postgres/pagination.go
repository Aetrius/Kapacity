package postgres

import "gorm.io/gorm"

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {

	return func(db *gorm.DB) *gorm.DB {
		if page == 1 {
			return db.Limit(1) // Limit to just one row
		} else {
			offset := (page - 1) * pageSize
			return db.Offset(offset).Limit(pageSize)
		}
	}
}
