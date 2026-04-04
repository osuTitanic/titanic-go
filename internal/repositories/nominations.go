package repositories

import (
	"github.com/osuTitanic/titanic-go/internal/schemas"
	"gorm.io/gorm"
)

type NominationRepository struct {
	db *gorm.DB
}

func NewNominationRepository(db *gorm.DB) *NominationRepository {
	return &NominationRepository{db: db}
}

func (r *NominationRepository) Create(nomination *schemas.BeatmapNomination) error {
	return r.db.Create(nomination).Error
}

func (r *NominationRepository) Delete(nomination *schemas.BeatmapNomination) error {
	return r.db.Delete(nomination).Error
}

func (r *NominationRepository) DeleteAll(setId int) error {
	return r.db.Where("set_id = ?", setId).Delete(&schemas.BeatmapNomination{}).Error
}

func (r *NominationRepository) FetchByBeatmapset(setId int, preload ...string) ([]*schemas.BeatmapNomination, error) {
	var nominations []*schemas.BeatmapNomination
	err := Preloaded(r.db, preload).Where("set_id = ?", setId).Find(&nominations).Error
	return nominations, err
}
