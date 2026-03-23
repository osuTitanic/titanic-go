package repositories

import (
	"github.com/osuTitanic/common-go/internal/schemas"
	"gorm.io/gorm"
)

type BeatmapRepository struct {
	db *gorm.DB
}

func NewBeatmapRepository(db *gorm.DB) *BeatmapRepository {
	return &BeatmapRepository{db: db}
}

func (r *BeatmapRepository) Create(beatmap *schemas.Beatmap) error {
	return r.db.Create(beatmap).Error
}

func (r *BeatmapRepository) Delete(beatmap *schemas.Beatmap) error {
	return r.db.Delete(beatmap).Error
}

func (r *BeatmapRepository) Update(updates *schemas.Beatmap, columns ...string) (int64, error) {
	if len(columns) == 0 {
		result := r.db.Save(updates)
		return result.RowsAffected, result.Error
	}
	result := r.db.Model(updates).Select(columns).Updates(updates)
	return result.RowsAffected, result.Error
}

func (r *BeatmapRepository) ById(id int, preload ...string) (*schemas.Beatmap, error) {
	var beatmap schemas.Beatmap
	err := Preloaded(r.db, preload).Where("id = ?", id).First(&beatmap).Error
	if err != nil {
		return nil, err
	}
	return &beatmap, nil
}

func (r *BeatmapRepository) ManyById(ids []int, preload ...string) ([]*schemas.Beatmap, error) {
	if len(ids) == 0 {
		return []*schemas.Beatmap{}, nil
	}

	var beatmaps []*schemas.Beatmap
	err := Preloaded(r.db, preload).Where("id IN ?", ids).Find(&beatmaps).Error
	return beatmaps, err
}
