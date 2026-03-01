package repositories

import (
	"github.com/osuTitanic/common-go/schemas"
	"gorm.io/gorm"
)

type InfringementRepository struct {
	db *gorm.DB
}

func NewInfringementRepository(db *gorm.DB) *InfringementRepository {
	return &InfringementRepository{db: db}
}

func (r *InfringementRepository) Create(infringement *schemas.Infringement) error {
	return r.db.Create(infringement).Error
}

func (r *InfringementRepository) Update(id int, updates map[string]interface{}) (int64, error) {
	result := r.db.Model(&schemas.Infringement{}).Where("id = ?", id).Updates(updates)
	return result.RowsAffected, result.Error
}

func (r *InfringementRepository) Delete(infringement *schemas.Infringement) error {
	return r.db.Delete(infringement).Error
}

func (r *InfringementRepository) ById(id int, preload ...string) (*schemas.Infringement, error) {
	var infringement schemas.Infringement
	err := Preloaded(r.db, preload).Where("id = ?", id).First(&infringement).Error
	if err != nil {
		return nil, err
	}
	return &infringement, nil
}

func (r *InfringementRepository) ManyByUserId(userId int, preload ...string) ([]*schemas.Infringement, error) {
	var infringements []*schemas.Infringement
	err := Preloaded(r.db, preload).Where("user_id = ?", userId).Order("time DESC").Find(&infringements).Error
	return infringements, err
}
