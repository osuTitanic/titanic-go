package repositories

import (
	"github.com/osuTitanic/titanic-go/internal/constants"
	"github.com/osuTitanic/titanic-go/internal/schemas"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ScoreRepository struct {
	db *gorm.DB
}

func NewScoreRepository(db *gorm.DB) *ScoreRepository {
	return &ScoreRepository{db: db}
}

func (r *ScoreRepository) Create(score *schemas.Score) error {
	return r.db.Create(score).Error
}

func (r *ScoreRepository) Delete(score *schemas.Score) error {
	return r.db.Delete(score).Error
}

func (r *ScoreRepository) Update(updates *schemas.Score, columns ...string) (int64, error) {
	if len(columns) == 0 {
		result := r.db.Save(updates)
		return result.RowsAffected, result.Error
	}
	result := r.db.Model(updates).Select(columns).Updates(updates)
	return result.RowsAffected, result.Error
}

func (r *ScoreRepository) ById(id int64, preload ...string) (*schemas.Score, error) {
	var score schemas.Score
	err := Preloaded(r.db, preload).Where("id = ?", id).First(&score).Error
	if err != nil {
		return nil, err
	}
	return &score, nil
}

func (r *ScoreRepository) ManyById(ids []int64, preload ...string) ([]*schemas.Score, error) {
	if len(ids) == 0 {
		return []*schemas.Score{}, nil
	}

	var scores []*schemas.Score
	err := Preloaded(r.db, preload).Where("id IN ?", ids).Find(&scores).Error
	return scores, err
}

func (r *ScoreRepository) GetCount() (int64, error) {
	var count int64
	err := r.db.Model(&schemas.Score{}).Count(&count).Error
	return count, err
}

func (r *ScoreRepository) FetchScoreIndexById(scoreId int64, beatmapId int, mode constants.Mode) (int, error) {
	var rank int
	err := r.db.Raw(`
		SELECT ranked.rank
		FROM (
			SELECT
				id,
				RANK() OVER (ORDER BY total_score DESC) AS rank
			FROM scores
			WHERE beatmap_id = ?
				AND mode = ?
				AND hidden = FALSE
				AND status_score = 3
		) AS ranked
		WHERE ranked.id = ?
		LIMIT 1
	`, beatmapId, mode, scoreId).Scan(&rank).Error
	if err != nil {
		return 0, err
	}

	return rank, nil
}

func (r *ScoreRepository) FetchScoreIndexByTscore(totalScore int64, beatmapId int, mode constants.Mode) (int, error) {
	var closestScore schemas.Score
	err := r.db.Model(&schemas.Score{}).
		Where("total_score > ?", totalScore).
		Where("beatmap_id = ?", beatmapId).
		Where("mode = ?", mode).
		Where("status_score = 3").
		Where("hidden = FALSE").
		Order(clause.Expr{SQL: "ABS(total_score - ?)", Vars: []interface{}{totalScore}}).
		Order("id ASC").
		First(&closestScore).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 1, nil
		}
		return 0, err
	}

	rank, err := r.FetchScoreIndexById(closestScore.Id, beatmapId, mode)
	if err != nil {
		return 0, err
	}
	if rank == 0 {
		return 1, nil
	}
	return rank + 1, nil
}
