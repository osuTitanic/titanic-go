package rankings

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RankingsService struct {
	ctx    context.Context
	client *redis.Client
}

func NewRankingsService(client *redis.Client) *RankingsService {
	return &RankingsService{
		ctx:    context.Background(),
		client: client,
	}
}
