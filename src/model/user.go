package model

import (
	"context"
	"fmt"

	"github.com/igeargeek/igg-go-acl/config"

	"github.com/igeargeek/igg-golang-api-response/response"
	"github.com/qiniu/qmgo/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var collectionName string = "users"

type userModel interface {
	GetUserPaginate(
		limit int64,
		page int64,
		filter interface{},
		sort string,
	) (response.Pagination, error)
	GetUserById(ID string) (User, error)
	GetUserOne(filter interface{}) (User, error)
	CreateUser(data *User) error
	UpdateUserByID(data bson.M, ID string) error
}

type User struct {
	field.DefaultField `bson:",inline"`
	Username           string `json:"Username" bson:"username"`
	Password           string `json:"-" bson:"password"`
	Fullname           string `json:"Fullname" bson:"fullname"`
	Roles              string `json:"Roles" bson:"roles"`
}

func NewUser() userModel {
	return &User{}
}

func (u *User) BeforeInsert() error {
	u.Password = generatePassword(u.Password)
	return nil
}

func generatePassword(str string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 14)
	if err != nil {
		fmt.Println("Bcrypt password error", err)
		return ""
	}
	return string(bytes)
}

func (u *User) GetUserPaginate(
	limit int64,
	page int64,
	filter interface{},
	sort string,
) (response.Pagination, error) {
	var results response.Pagination
	data := []User{}
	results, err := getPaginate(&data, collectionName, limit, page, filter, sort)
	if err != nil {
		return results, err
	}
	return results, err
}

func (u *User) GetUserById(ID string) (User, error) {
	result := User{}
	collection := config.GetDBClient().Collection(collectionName)
	docID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		fmt.Println(err)
		return result, err
	}

	if err := collection.Find(context.TODO(), bson.M{"_id": docID}).One(&result); err != nil {
		fmt.Println(err)
		return result, err
	}
	return result, nil
}

func (u *User) GetUserOne(filter interface{}) (User, error) {
	var result User
	collection := config.GetDBClient().Collection(collectionName)
	if err := collection.Find(context.TODO(), filter).One(&result); err != nil {
		return result, err
	}
	return result, nil
}

func (u *User) CreateUser(data *User) error {
	collection := config.GetDBClient().Collection(collectionName)
	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) UpdateUserByID(data bson.M, ID string) error {
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	collection := config.GetDBClient().Collection(collectionName)
	if err != nil {
		return err
	}
	if data["password"] != nil {
		data["password"] = generatePassword(data["password"].(string))
	}
	update := bson.D{
		{"$set", data},
	}
	if err := collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update); err != nil {
		return err
	}
	return nil
}
