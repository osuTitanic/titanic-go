package repositories

import (
	"github.com/osuTitanic/common-go/schemas"
	"gorm.io/gorm"
)

type BeatmapFavouriteRepository struct {
	db *gorm.DB
}

func NewBeatmapFavouriteRepository(db *gorm.DB) *BeatmapFavouriteRepository {
	return &BeatmapFavouriteRepository{db: db}
}

func (r *BeatmapFavouriteRepository) Create(favourite *schemas.BeatmapFavourite) error {
	return r.db.Create(favourite).Error
}

func (r *BeatmapFavouriteRepository) Update(userId int, setId int, updates map[string]interface{}) error {
	return r.db.Model(&schemas.BeatmapFavourite{}).Where("user_id = ? AND set_id = ?", userId, setId).Updates(updates).Error
}

func (r *BeatmapFavouriteRepository) Delete(favourite *schemas.BeatmapFavourite) error {
	return r.db.Delete(favourite).Error
}

func (r *BeatmapFavouriteRepository) ByUserAndSet(userId int, setId int) (*schemas.BeatmapFavourite, error) {
	var favourite schemas.BeatmapFavourite
	err := r.db.Where("user_id = ? AND set_id = ?", userId, setId).First(&favourite).Error
	if err != nil {
		return nil, err
	}
	return &favourite, nil
}

func (r *BeatmapFavouriteRepository) ManyByUserId(userId int) ([]*schemas.BeatmapFavourite, error) {
	var favourites []*schemas.BeatmapFavourite
	err := r.db.Where("user_id = ?", userId).Find(&favourites).Error
	return favourites, err
}

func (r *BeatmapFavouriteRepository) CountByUserId(userId int) (int, error) {
	var count int64
	err := r.db.Model(&schemas.BeatmapFavourite{}).Where("user_id = ?", userId).Count(&count).Error
	return int(count), err
}

func (r *BeatmapFavouriteRepository) CountBySetId(setId int) (int, error) {
	var count int64
	err := r.db.Model(&schemas.BeatmapFavourite{}).Where("set_id = ?", setId).Count(&count).Error
	return int(count), err
}
