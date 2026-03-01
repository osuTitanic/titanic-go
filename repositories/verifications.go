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

func (r *VerificationRepository) Update(id int, updates map[string]interface{}) error {
	return r.db.Model(&schemas.Verification{}).Where("id = ?", id).Updates(updates).Error
}

func (r *VerificationRepository) Delete(verification *schemas.Verification) error {
	return r.db.Delete(verification).Error
}

func (r *VerificationRepository) ById(id int) (*schemas.Verification, error) {
	var verification schemas.Verification
	err := r.db.Where("id = ?", id).First(&verification).Error
	if err != nil {
		return nil, err
	}
	return &verification, nil
}

func (r *VerificationRepository) ByToken(token string) (*schemas.Verification, error) {
	var verification schemas.Verification
	err := r.db.Where("token = ?", token).First(&verification).Error
	if err != nil {
		return nil, err
	}
	return &verification, nil
}

func (r *VerificationRepository) ManyByUserId(userId int) ([]*schemas.Verification, error) {
	var verifications []*schemas.Verification
	err := r.db.Where("user_id = ?", userId).Find(&verifications).Error
	return verifications, err
}

func (r *VerificationRepository) DeleteByToken(token string) error {
	return r.db.Where("token = ?", token).Delete(&schemas.Verification{}).Error
}
