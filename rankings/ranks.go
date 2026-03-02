package rankings

import (
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func (service *RankingsService) RankByKey(key string, userId int) (int, error) {
	if service == nil || service.client == nil {
		return 0, fmt.Errorf("rankings: redis client is not initialized")
	}

	rank, err := service.client.ZRevRank(service.ctx, key, strconv.Itoa(userId)).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, nil
		}
		return 0, err
	}

	return int(rank + 1), nil
}

func (service *RankingsService) GlobalRank(userId int, mode int) (int, error) {
	return service.RankByKey(service.RankingKey(mode, "performance", nil), userId)
}

func (service *RankingsService) PPv1Rank(userId int, mode int) (int, error) {
	return service.RankByKey(service.RankingKey(mode, "ppv1", nil), userId)
}

func (service *RankingsService) CountryRank(userId int, mode int, country string) (int, error) {
	return service.RankByKey(service.RankingKey(mode, "performance", &country), userId)
}

func (service *RankingsService) ScoreRank(userId int, mode int) (int, error) {
	return service.RankByKey(service.RankingKey(mode, "rscore", nil), userId)
}

func (service *RankingsService) ClearsRank(userId int, mode int) (int, error) {
	return service.RankByKey(service.RankingKey(mode, "clears", nil), userId)
}

func (service *RankingsService) TotalScoreRank(userId int, mode int) (int, error) {
	return service.RankByKey(service.RankingKey(mode, "tscore", nil), userId)
}

func (service *RankingsService) LeaderScoresRank(userId int, mode int) (int, error) {
	return service.RankByKey(service.RankingKey(mode, "leader", nil), userId)
}

func (service *RankingsService) ScoreRankCountry(userId int, mode int, country string) (int, error) {
	return service.RankByKey(service.RankingKey(mode, "rscore", &country), userId)
}

func (service *RankingsService) ClearsRankCountry(userId int, mode int, country string) (int, error) {
	return service.RankByKey(service.RankingKey(mode, "clears", &country), userId)
}

func (service *RankingsService) PPv1CountryRank(userId int, mode int, country string) (int, error) {
	return service.RankByKey(service.RankingKey(mode, "ppv1", &country), userId)
}

func (service *RankingsService) TotalScoreRankCountry(userId int, mode int, country string) (int, error) {
	return service.RankByKey(service.RankingKey(mode, "tscore", &country), userId)
}

func (service *RankingsService) Rank(userId int, mode int, rankType string, country *string) (int, error) {
	return service.RankByKey(service.RankingKey(mode, rankType, country), userId)
}
