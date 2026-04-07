package repositories

import (
	"github.com/osuTitanic/titanic-go/internal/schemas"
	"gorm.io/gorm"
)

type ScreenshotRepository struct {
	db *gorm.DB
}

func NewScreenshotRepository(db *gorm.DB) *ScreenshotRepository {
	return &ScreenshotRepository{db: db}
}

func (r *ScreenshotRepository) Create(screenshot *schemas.Screenshot) error {
	return r.db.Create(screenshot).Error
}

func (r *ScreenshotRepository) Delete(screenshot *schemas.Screenshot) error {
	return r.db.Delete(screenshot).Error
}

func (r *ScreenshotRepository) Update(updates *schemas.Screenshot, columns ...string) (int64, error) {
	return CommonUpdate(r.db, updates, columns...)
}
