package handler

import (
	"github.com/redis/go-redis/v9"
	"context"
	"crypto/md5"
	"encoding/hex"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
 }
 var ctx = context.Background()
 var rdb *redis.Client
 func init() {
	 rdb = redis.NewClient(&redis.Options{
		 Addr: "redis:6379", Password: "", DB: 0,
	 })
 }