package infrastructure

import (
	"context"

	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
	tredis "github.com/octoposprime/op-be-shared/tool/redis"
	"github.com/redis/go-redis/v9"
)

type RedisAdapter struct {
	redisClient *tredis.RedisClient
	Log         func(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error)
}

func NewRedisAdapter(redisClient *tredis.RedisClient) RedisAdapter {
	adapter := RedisAdapter{
		redisClient: redisClient,
		Log:         Log,
	}

	return adapter
}

// SetLogger sets logging call-back function
func (a *RedisAdapter) SetLogger(LoggerFunc func(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error)) {
	a.Log = LoggerFunc
}

// ConsumeChannel subscribes the given channel to Redis and returns *redis.PubSub
func (a *RedisAdapter) ConsumeChannel(ctx context.Context, channelName string) *redis.PubSub {
	return a.redisClient.Subscribe(ctx, channelName)
}
