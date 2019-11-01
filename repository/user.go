package repository

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"

	"github.com/airlanggatirta/pawon-warga/model"
)

const (
	userStatusVerified = "VERIFIED"
)

type UserRepository struct {
	RepositoryOption
}

func (ur *UserRepository) GetOneUserByEmailWithoutFilter(email string) (user model.User, err error) {
	conn := ur.Cache.Get()
	defer conn.Close()

	cacheData, err := redis.Bytes(conn.Do("GET", fmt.Sprintf("%s:%s", ur.AppConfig.Cache.Items.GetOneUserByEmailWithoutFilter.Prefix, email)))
	if err != nil {
		err = ur.DB.Preload("Privilege").
			Preload("UserStatus").
			Preload("Location").
			Where("email = ?", email).
			Find(&user).Error

		if err != nil {
			return model.User{}, err
		}

		data, err := json.Marshal(user)
		if err != nil {
			return model.User{}, err
		}

		_, err = conn.Do("SETEX", fmt.Sprintf("%s:%s", ur.AppConfig.Cache.Items.GetOneUserByEmailWithoutFilter.Prefix, email), ur.AppConfig.Cache.Items.GetOneUserByEmailWithoutFilter.TTL.Seconds(), data)
		if err != nil {
			return model.User{}, err
		}

		return user, nil
	}

	err = json.Unmarshal(cacheData, &user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) GetOneUserByWhatsappWithoutFilter(waNumber string) (user model.User, err error) {
	conn := ur.Cache.Get()
	defer conn.Close()

	cacheData, err := redis.Bytes(conn.Do("GET", fmt.Sprintf("%s:%s", ur.AppConfig.Cache.Items.GetOneUserByWhatsappWithoutFilter.Prefix, waNumber)))
	if err != nil {
		err = ur.DB.Preload("Privilege").
			Preload("UserStatus").
			Preload("Location").
			Where("wa_phone_number = ?", waNumber).
			Find(&user).Error

		if err != nil {
			return model.User{}, err
		}

		data, err := json.Marshal(user)
		if err != nil {
			return model.User{}, err
		}

		_, err = conn.Do("SETEX", fmt.Sprintf("%s:%s", ur.AppConfig.Cache.Items.GetOneUserByWhatsappWithoutFilter.Prefix, waNumber), ur.AppConfig.Cache.Items.GetOneUserByEmailWithoutFilter.TTL.Seconds(), data)
		if err != nil {
			return model.User{}, err
		}

		return user, nil
	}

	err = json.Unmarshal(cacheData, &user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
