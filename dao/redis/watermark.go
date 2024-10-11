package redis

// SavePresignedURL 中存储presignedURL
func SavePresignedURL(name, url string) (err error) {
	key := getRedisPrefixKey(name)
	_, err = client.Set(key, url, presignedURLTime).Result()
	return
}

// GetPresignedURL 获取presignedURL
func GetPresignedURL(name string) (url string, err error) {
	key := getRedisPrefixKey(name)
	url, err = client.Get(key).Result()
	return
}
