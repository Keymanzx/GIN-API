package redis

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

func connectRedis() *redis.Client {
	// สร้าง Redis client
	redisHost := os.Getenv("REDIS_HOST")
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisHost, // ที่อยู่ของ Redis server
		Password: "",        // รหัสผ่าน (ถ้ามี)
		DB:       0,         // เลือก database (0 คือ default)
	})

	// ทดสอบการเชื่อมต่อ Redis server
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return nil
	}
	fmt.Println("Connected to Redis:", pong)

	return rdb
}

func SetKey(key, value string) {
	rdb := connectRedis()
	// เก็บค่าใน Redis
	err := rdb.Set(context.Background(), key, value, 0).Err()
	if err != nil {
		fmt.Println("Error setting value in Redis:", err)
		panic(err)
	}
}

func GetValue(key string) string {
	rdb := connectRedis()
	// ดึงค่าจาก Redis
	val, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		fmt.Println("Error getting value from Redis:", err)
		return ""
	}

	return val
}

func ClearCache(keys ...string) {
	rdb := connectRedis()
	// ลบค่าใน Redis
	_, err := rdb.Del(context.Background(), keys...).Result()
	if err != nil {
		fmt.Println("Error clearing cache in Redis:", err)
		panic(err)
	}
}
