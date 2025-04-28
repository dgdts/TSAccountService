package model

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type UserBasic struct {
	mgm.DefaultModel `bson:",inline"`
	Nickname         string `bson:"nickname"`
	Password         string `bson:"password"`
	Phone            string `bson:"phone"`
	Email            string `bson:"email"`
	Avatar           string `bson:"avatar"`

	GoogleID string `bson:"google_id"`
	WechatID string `bson:"wechat_id"`
	GithubID string `bson:"github_id"`
}

func CreateUser(user *UserBasic) error {
	return mgm.Coll(user).Create(user)
}

func FindUserByPhone(phone string) (*UserBasic, error) {
	var user UserBasic
	err := mgm.Coll(&user).First(bson.M{"phone": phone}, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserByEmail(email string) (*UserBasic, error) {
	var user UserBasic
	err := mgm.Coll(&user).First(bson.M{"email": email}, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserByID(id string) (*UserBasic, error) {
	var user UserBasic
	err := mgm.Coll(&user).FindByID(id, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(user *UserBasic) error {
	return mgm.Coll(user).Update(user)
}

func FindUserByPhoneAndPassword(phone string, password string) (*UserBasic, error) {
	var user UserBasic
	err := mgm.Coll(&user).First(bson.M{"phone": phone, "password": password}, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserByEmailAndPassword(email string, password string) (*UserBasic, error) {
	var user UserBasic
	err := mgm.Coll(&user).First(bson.M{"email": email, "password": password}, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
