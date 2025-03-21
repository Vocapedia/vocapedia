package utils

import (
	"encoding/json"
	"math/rand"
	"time"
)

func StructToMap(object any) (map[string]any, error) {
	var newMap map[string]any
	data, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &newMap)
	if err != nil {
		return nil, err
	}
	return newMap, err
}

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)

	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}
