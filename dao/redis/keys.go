package redis

import "time"

const (
	KeyPrefix        = "fisco-cert-app:"
	KeyWatermark     = "watermark" // 帖子及发帖时间
	presignedURLTime = time.Hour * 24 * 7
)

func getRedisPrefixKey(key string) string {
	return KeyPrefix + key
}
