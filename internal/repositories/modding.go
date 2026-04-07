package repositories

import (
	"github.com/osuTitanic/titanic-go/internal/schemas"
	"gorm.io/gorm"
)

type BeatmapModdingRepository struct {
	db *gorm.DB
}

func NewBeatmapModdingRepository(db *gorm.DB) *BeatmapModdingRepository {
	return &BeatmapModdingRepository{db: db}
}

func (r *BeatmapModdingRepository) Create(modding *schemas.BeatmapModding) error {
	return r.db.Create(modding).Error
}

func (r *BeatmapModdingRepository) Delete(modding *schemas.BeatmapModding) error {
	return r.db.Delete(modding).Error
}

func (r *BeatmapModdingRepository) Update(updates *schemas.BeatmapModding, columns ...string) (int64, error) {
	return CommonUpdate(r.db, updates, columns...)
}
