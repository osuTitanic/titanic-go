package repositories

import (
	"github.com/osuTitanic/titanic-go/internal/schemas"
	"gorm.io/gorm"
)

type ActivityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) *ActivityRepository {
	return &ActivityRepository{db: db}
}

func (r *ActivityRepository) Create(activity *schemas.Activity) error {
	return r.db.Create(activity).Error
}

func (r *ActivityRepository) Delete(activity *schemas.Activity) error {
	return r.db.Delete(activity).Error
}

func (r *ActivityRepository) Update(updates *schemas.Activity, columns ...string) (int64, error) {
	return CommonUpdate(r.db, updates, columns...)
}
