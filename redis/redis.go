package redis

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_redis "github.com/redis/go-redis/v9"
)

var rdb *_redis.Client

func InitRedis() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	redisAddr := os.Getenv("REDIS_ADDRESS")
	redisPass := os.Getenv("REDIS_PASSWORD")
	redisDbStr := os.Getenv("REDIS_DB_TOKEN")
	
	redisDB, err := strconv.Atoi(redisDbStr)
    if err != nil {
        log.Fatalf("Error converting REDIS_DB_TOKEN string to integer: %v", err)
    }

	rdb = _redis.NewClient(&_redis.Options{
		Addr:     redisAddr,
		Password: redisPass,
		DB:       redisDB,
	})

	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Error when connecting to redis: %s", err)
	}

	log.Println("Connected to Redis:", pong)
}

func StoreToken(token string) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	
	expireTokenTimeStr := os.Getenv("EXPIRE_TOKEN_TIME")
    expDuration, err := time.ParseDuration(expireTokenTimeStr)
	if err != nil {
        log.Fatalf("Error converting EXPIRE_TOKEN_TIME string to integer: %v", err)
    }
 
	return rdb.Set(context.Background(), token, "", expDuration).Err()
}

func GetToken(token string) (string, error) {
	return rdb.Get(context.Background(), token).Result()
}

func DeleteToken(token string) error {
	return rdb.Del(context.Background(), token).Err()
}