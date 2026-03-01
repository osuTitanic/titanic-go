package repositories

import "gorm.io/gorm"

func Preloaded(db *gorm.DB, preload []string) *gorm.DB {
	for _, p := range preload {
		db = db.Preload(p)
	}
	return db
}
