package repositories

import (
	"time"

	"github.com/osuTitanic/titanic-go/internal/constants"
	"github.com/osuTitanic/titanic-go/internal/rankings"
	"github.com/osuTitanic/titanic-go/internal/schemas"
	"gorm.io/gorm"
)

type HistoryRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) *HistoryRepository {
	return &HistoryRepository{db: db}
}

func (r *HistoryRepository) UpdatePlays(userId int, mode constants.Mode) error {
	now := time.Now()
	year, month, _ := now.Date()

	result := r.db.Model(&schemas.PlayHistory{}).
		Where("user_id = ? AND mode = ? AND year = ? AND month = ?", userId, mode, year, int(month)).
		Update("plays", gorm.Expr("plays + 1"))

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected != 0 {
		return nil
	}

	entry := &schemas.PlayHistory{
		UserId: userId,
		Mode:   mode,
		Year:   year,
		Month:  int(month),
		Plays:  1,
	}
	return r.db.Create(entry).Error
}

func (r *HistoryRepository) FetchPlaysHistory(userId int, mode constants.Mode, until time.Time) ([]*schemas.PlayHistory, error) {
	var history []*schemas.PlayHistory
	err := r.db.Where("user_id = ? AND mode = ? AND created_at >= ?", userId, mode, until).
		Order("created_at DESC").
		Find(&history).Error
	return history, err
}

func (r *HistoryRepository) FetchPlaysHistoryAll(userId int, mode constants.Mode) ([]*schemas.PlayHistory, error) {
	var history []*schemas.PlayHistory
	err := r.db.Where("user_id = ? AND mode = ?", userId, mode).
		Order("created_at DESC").
		Find(&history).Error
	return history, err
}

func (r *HistoryRepository) UpdateReplayViews(userId int, mode constants.Mode) error {
	now := time.Now()
	year, month, _ := now.Date()

	result := r.db.Model(&schemas.ReplayHistory{}).
		Where("user_id = ? AND mode = ? AND year = ? AND month = ?", userId, mode, year, int(month)).
		Update("replay_views", gorm.Expr("replay_views + 1"))

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected != 0 {
		return nil
	}

	entry := &schemas.ReplayHistory{
		UserId:      userId,
		Mode:        mode,
		Year:        year,
		Month:       int(month),
		ReplayViews: 1,
	}
	return r.db.Create(entry).Error
}

func (r *HistoryRepository) FetchReplayHistory(userId int, mode constants.Mode, until time.Time) ([]*schemas.ReplayHistory, error) {
	var history []*schemas.ReplayHistory
	err := r.db.Where("user_id = ? AND mode = ? AND created_at >= ?", userId, mode, until).
		Order("created_at DESC").
		Find(&history).Error
	return history, err
}

func (r *HistoryRepository) FetchReplayHistoryAll(userId int, mode constants.Mode) ([]*schemas.ReplayHistory, error) {
	var history []*schemas.ReplayHistory
	err := r.db.Where("user_id = ? AND mode = ?", userId, mode).
		Order("created_at DESC").
		Find(&history).Error
	return history, err
}

func (r *HistoryRepository) AddRank(history *schemas.RankHistory) error {
	return r.db.Create(history).Error
}

func (r *HistoryRepository) UpdateRank(stats *schemas.Stats, country string, rankingsService *rankings.RankingsService) error {
	globalRank, _ := rankingsService.GlobalRank(stats.UserId, stats.Mode)
	countryRank, _ := rankingsService.CountryRank(stats.UserId, stats.Mode, country)
	scoreRank, _ := rankingsService.ScoreRank(stats.UserId, stats.Mode)
	ppv1Rank, _ := rankingsService.PPv1Rank(stats.UserId, stats.Mode)

	if globalRank <= 0 || countryRank <= 0 || scoreRank <= 0 || ppv1Rank <= 0 {
		return nil
	}

	entry := &schemas.RankHistory{
		UserId:      stats.UserId,
		Mode:        stats.Mode,
		Rscore:      stats.Rscore,
		PP:          int(stats.PP),
		PPv1:        int(stats.PPv1),
		GlobalRank:  globalRank,
		CountryRank: countryRank,
		ScoreRank:   scoreRank,
		PPv1Rank:    ppv1Rank,
		Time:        time.Now(),
	}
	return r.db.Create(entry).Error
}

func (r *HistoryRepository) FetchRankHistory(userId int, mode constants.Mode, until time.Time) ([]*schemas.RankHistory, error) {
	var history []*schemas.RankHistory
	err := r.db.Where("user_id = ? AND mode = ? AND time > ?", userId, mode, until).
		Order("time DESC").
		Find(&history).Error
	return history, err
}

func (r *HistoryRepository) FetchPeakGlobalRank(userId int, mode constants.Mode) (int, error) {
	var peakRank int
	err := r.db.Model(&schemas.RankHistory{}).
		Select("COALESCE(MIN(global_rank), 0)").
		Where("user_id = ? AND mode = ? AND global_rank != 0", userId, mode).
		Scan(&peakRank).Error
	return peakRank, err
}
