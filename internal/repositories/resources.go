package repositories

import (
	"github.com/osuTitanic/titanic-go/internal/schemas"
	"gorm.io/gorm"
)

type ResourceMirrorRepository struct {
	db *gorm.DB
}

func NewResourceMirrorRepository(db *gorm.DB) *ResourceMirrorRepository {
	return &ResourceMirrorRepository{db: db}
}

func (r *ResourceMirrorRepository) Create(mirror *schemas.BeatmapMirror) error {
	return r.db.Create(mirror).Error
}

func (r *ResourceMirrorRepository) Delete(mirror *schemas.BeatmapMirror) error {
	return r.db.Delete(mirror).Error
}

func (r *ResourceMirrorRepository) Update(updates *schemas.BeatmapMirror, columns ...string) (int64, error) {
	return CommonUpdate(r.db, updates, columns...)
}
