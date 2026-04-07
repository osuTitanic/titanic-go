package repositories

import (
	"github.com/osuTitanic/titanic-go/internal/schemas"
	"gorm.io/gorm"
)

type BanchoActivityRepository struct {
	db *gorm.DB
}

func NewBanchoActivityRepository(db *gorm.DB) *BanchoActivityRepository {
	return &BanchoActivityRepository{db: db}
}

func (r *BanchoActivityRepository) Create(activity *schemas.BanchoActivity) error {
	return r.db.Create(activity).Error
}

func (r *BanchoActivityRepository) Delete(activity *schemas.BanchoActivity) error {
	return r.db.Delete(activity).Error
}

func (r *BanchoActivityRepository) Update(updates *schemas.BanchoActivity, columns ...string) (int64, error) {
	return CommonUpdate(r.db, updates, columns...)
}
