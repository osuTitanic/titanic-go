package repositories

import "gorm.io/gorm"

func Preloaded(db *gorm.DB, preload []string) *gorm.DB {
	for _, p := range preload {
		db = db.Preload(p)
	}
	return db
}

func CommonUpdate[T any](db *gorm.DB, updates *T, columns ...string) (int64, error) {
	if len(columns) == 0 {
		result := db.Save(updates)
		return result.RowsAffected, result.Error
	}
	result := db.Model(updates).Select(columns).Updates(updates)
	return result.RowsAffected, result.Error
}
