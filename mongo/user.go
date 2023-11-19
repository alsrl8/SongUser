package mongo

import (
	"SongUser/auth"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type UserInfo struct {
	Name string `json:"name"`
	Id   string `json:"id"`
	Pw   string `json:"pw"`
}

func NewUserRepository(dbName, collName string) (UserRepository, error) {
	client, err := getMongoClient()
	if err != nil {
		log.Printf("Error getting mongo client: %+v", err)
		return nil, err
	}

	return &DBUserRepository{
		Client:     client,
		Collection: client.Database(dbName).Collection(collName),
	}, nil
}

type UserRepository interface {
	InsertOne(id, pw, name string) error
	FindUser(id string) (*UserInfo, error)
	Close() error
	clearCollection() error
}

type DBUserRepository struct {
	*mongo.Client
	*mongo.Collection
}

func (db *DBUserRepository) InsertOne(id, pw, name string) error {
	userInfo := &UserInfo{
		Name: name,
		Id:   id,
		Pw:   pw,
	}
	_, err := db.Collection.InsertOne(context.TODO(), userInfo)
	return err
}

func (db *DBUserRepository) FindUser(id string) (*UserInfo, error) {
	var result UserInfo
	filter := bson.D{{"id", id}}
	if err := db.Collection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		log.Printf("User not found for ID %s", id)
		return nil, errors.New("user not found")
	}
	return &result, nil
}

func (db *DBUserRepository) Close() error {
	err := db.Client.Disconnect(context.TODO())
	if err != nil {
		fmt.Printf("Error disconecting mongo client: %+v", err)
		return err
	}
	return nil
}

func (db *DBUserRepository) clearCollection() error {
	many, err := db.Collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	log.Printf("Clear mongo collection: %+v", many)
	return nil
}

func Login(id string, pw string, repo UserRepository) error {
	result, err := repo.FindUser(id)
	if err != nil {
		log.Printf(err.Error())
		return err
	}

	if matched := auth.CheckPwHash(pw, result.Pw); !matched {
		log.Printf("Password mismatch for user ID %s", id)
		return errors.New("password mismatch")
	}

	log.Printf("User authenticated: %s\n", id)
	return nil
}

func Register(id, pw, name string, repo UserRepository) error {
	_, err := repo.FindUser(id)
	if err == nil {
		return errors.New("user already exists")
	}

	hashPw, err := auth.HashPw(pw)
	if err != nil {
		return err
	}

	err = repo.InsertOne(id, hashPw, name)
	if err != nil {
		return err
	}

	return nil
}
