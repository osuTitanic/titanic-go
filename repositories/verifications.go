package repositories

import (
	"github.com/osuTitanic/common-go/schemas"
	"gorm.io/gorm"
)

type VerificationRepository struct {
	db *gorm.DB
}

func NewVerificationRepository(db *gorm.DB) *VerificationRepository {
	return &VerificationRepository{db: db}
}

func (r *VerificationRepository) Create(verification *schemas.Verification) error {
	return r.db.Create(verification).Error
}

func (r *VerificationRepository) Delete(verification *schemas.Verification) error {
	return r.db.Delete(verification).Error
}

func (r *VerificationRepository) Update(updates *schemas.Verification, columns ...string) (int64, error) {
	if len(columns) == 0 {
		result := r.db.Save(updates)
		return result.RowsAffected, result.Error
	}
	result := r.db.Model(updates).Select(columns).Updates(updates)
	return result.RowsAffected, result.Error
}

func (r *VerificationRepository) ById(id int, preload ...string) (*schemas.Verification, error) {
	var verification schemas.Verification
	err := Preloaded(r.db, preload).Where("id = ?", id).First(&verification).Error
	if err != nil {
		return nil, err
	}
	return &verification, nil
}

func (r *VerificationRepository) ByToken(token string, preload ...string) (*schemas.Verification, error) {
	var verification schemas.Verification
	err := Preloaded(r.db, preload).Where("token = ?", token).First(&verification).Error
	if err != nil {
		return nil, err
	}
	return &verification, nil
}

func (r *VerificationRepository) ManyByUserId(userId int, preload ...string) ([]*schemas.Verification, error) {
	var verifications []*schemas.Verification
	err := Preloaded(r.db, preload).Where("user_id = ?", userId).Find(&verifications).Error
	return verifications, err
}

func (r *VerificationRepository) DeleteByToken(token string) error {
	return r.db.Where("token = ?", token).Delete(&schemas.Verification{}).Error
}
