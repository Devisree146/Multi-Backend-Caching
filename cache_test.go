package cache_test

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"

	"github.com/your-username/mini-project/cache" // Update with your actual module path
)

func TestInMemoryCacheOperations(t *testing.T) {
	cache := cache.NewInMemoryCache(3)

	// Test Set and Get operations
	err := cache.Set(context.Background(), "key1", 10, 5*time.Minute)
	assert.NoError(t, err)

	val, err := cache.Get(context.Background(), "key1")
	assert.NoError(t, err)
	assert.Equal(t, 10, val.(int))

	// Test Delete operation
	err = cache.Delete(context.Background(), "key1")
	assert.NoError(t, err)

	_, err = cache.Get(context.Background(), "key1")
	assert.Error(t, err)
}

func TestRedisCacheOperations(t *testing.T) {
	// Initialize Redis client
	rdb := NewTestRedisClient()

	redisCache := cache.NewRedisCache(rdb)

	// Test Set and Get operations
	err := redisCache.Set(context.Background(), "key2", 20, 5*time.Minute)
	assert.NoError(t, err)

	val, err := redisCache.Get(context.Background(), "key2")
	assert.NoError(t, err)
	assert.Equal(t, 20, val.(int))

	// Test Delete operation
	err = redisCache.Delete(context.Background(), "key2")
	assert.NoError(t, err)

	_, err = redisCache.Get(context.Background(), "key2")
	assert.Error(t, err)

	// Close Redis client connection
	err = rdb.Close()
	assert.NoError(t, err)
}

// NewTestRedisClient creates a new Redis client for testing purposes.
func NewTestRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return rdb
}
