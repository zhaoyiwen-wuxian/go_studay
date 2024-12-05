package service

import (
	"encoding/json"
	"go_dome/cache"
	"go_dome/dao"
	"go_dome/models"
	"strconv"
)

func CreateUser(user *models.User) error {
	return dao.CreateUser(user)
}

func GetUserById(id int) (*models.User, error) {
	key := "user:" + strconv.Itoa(id)
	cacheUser, err := cache.RedisClient.Get(cache.Ctx, key).Result()
	if err == nil {
		var user models.User
		json.Unmarshal([]byte(cacheUser), &user)
		return &user, nil
	}

	user, err := dao.GetUserById(id)

	if err != nil {
		return nil, err
	}
	data, _ := json.Marshal(user)
	cache.RedisClient.Set(cache.Ctx, key, data, 0)
	return user, nil
}
