package repositories

import (
	"errors"

	"github.com/osuTitanic/titanic-go/internal/schemas"
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

func (r *BeatmapRepository) UpdateBySetId(updates *schemas.Beatmap, columns ...string) (int64, error) {
	if len(columns) == 0 {
		return 0, errors.New("at least one column must be specified")
	}
	result := r.db.Model(&schemas.Beatmap{}).Where("set_id = ?", updates.SetId).Select(columns).Updates(updates)
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

func (r *BeatmapRepository) GetCount() (int, error) {
	var count int64
	err := r.db.Model(&schemas.Beatmap{}).Count(&count).Error
	return int(count), err
}

func (r *BeatmapRepository) GetCountGroupedByStatus(mode int) (map[int]int, error) {
	var results []struct {
		Status int
		Count  int
	}

	err := r.db.Model(&schemas.Beatmap{}).
		Select("status, count(*) as count").
		Where("mode = ?", mode).
		Group("status").
		Scan(&results).Error

	counts := make(map[int]int)
	for _, res := range results {
		counts[res.Status] = res.Count
	}

	return counts, err
}
