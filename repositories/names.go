package repositories

import (
	"github.com/osuTitanic/common-go/schemas"
	"gorm.io/gorm"
)

type NameRepository struct {
	db *gorm.DB
}

func NewNameRepository(db *gorm.DB) *NameRepository {
	return &NameRepository{db: db}
}

func (r *NameRepository) Create(name *schemas.Name) error {
	return r.db.Create(name).Error
}

func (r *NameRepository) Update(id int, updates map[string]interface{}) (int64, error) {
	result := r.db.Model(&schemas.Name{}).Where("id = ?", id).Updates(updates)
	return result.RowsAffected, result.Error
}

func (r *NameRepository) Delete(name *schemas.Name) error {
	return r.db.Delete(name).Error
}

func (r *NameRepository) ById(id int, preload ...string) (*schemas.Name, error) {
	var name schemas.Name
	err := Preloaded(r.db, preload).Where("id = ?", id).First(&name).Error
	if err != nil {
		return nil, err
	}
	return &name, nil
}

func (r *NameRepository) ByName(value string, preload ...string) (*schemas.Name, error) {
	var name schemas.Name
	err := Preloaded(r.db, preload).Where("name = ?", value).First(&name).Error
	if err != nil {
		return nil, err
	}
	return &name, nil
}

func (r *NameRepository) ManyByUserId(userId int, preload ...string) ([]*schemas.Name, error) {
	var names []*schemas.Name
	err := Preloaded(r.db, preload).Where("user_id = ?", userId).Find(&names).Error
	return names, err
}
