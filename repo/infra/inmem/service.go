package inmem

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/semicolon-ina/semicolon-url-shortener/repo/common/config"
)

// 1. DEFINISI INTERFACE (Contract)
// Ini yang bakal lo inject ke Repository.
// Mocking di unit test nanti tinggal bikin struct yang implement interface ini.
type RedisItf interface {
	Ping(ctx context.Context) error
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Exists(ctx context.Context, key string) (bool, error)
	Close() error
}

// 2. IMPLEMENTASI REAL (Wrapper)
// Struct ini ngebungkus *redis.Client asli biar sesuai sama Interface di atas.
type redisWrapper struct {
	client *redis.Client
}

// Init sekarang return Interface (RedisItf), BUKAN struct (*redis.Client)
func Init(cfg config.RedisConfig) RedisItf {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// Wrap struct asli ke dalam wrapper kita
	wrapper := &redisWrapper{client: rdb}

	// Test Connection (Ping) pake method wrapper
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := wrapper.Ping(ctx); err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	fmt.Println("Connected to Redis at", cfg.Host+":"+cfg.Port)
	return wrapper
}

// --- METHOD IMPLEMENTATION (Supaya redisWrapper comply sama RedisClient interface) ---

func (r *redisWrapper) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

func (r *redisWrapper) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	if fmt.Sprintf("%T", value) != "string" {
		saveToRedis, err := json.Marshal(value)
		if err != nil {
			return err
		}

		return r.client.Set(ctx, key, saveToRedis, ttl).Err()
	}
	return r.client.Set(ctx, key, value, ttl).Err()
}

func (r *redisWrapper) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *redisWrapper) Exists(ctx context.Context, key string) (bool, error) {
	count, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *redisWrapper) Close() error {
	return r.client.Close()
}
