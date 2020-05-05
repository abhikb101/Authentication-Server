package mongo_wrap

import (
	"context"
	"errors"
	"fmt"
	"time"
	root "user_auth/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	collection *mongo.Collection
	hash       root.Hash
}

func NewUserService(client *Client, dbName string, collectionName string, hash root.Hash) *UserService {
	collection := client.GetCollection(dbName, collectionName)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, _ = collection.Indexes().CreateMany(ctx, userModelIndex())
	return &UserService{collection, hash}
}

func (p *UserService) CreateUser(u *root.User) error {
	user := newUserModel(u)
	hashed, err := p.hash.Generate(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashed
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err = p.collection.InsertOne(ctx, &user)
	return err
}

func (p *UserService) GetByUsername(username string) (*root.User, error) {
	model := UserModel{}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := p.collection.FindOne(ctx, bson.D{{"username", username}}).Decode(&model)
	model.Password = "-"
	return model.toRootUser(), err
}

func (p *UserService) Login(u *root.User) (*root.User, error) {
	model := UserModel{}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := errors.New("nil")
	if u.Email != "" {
		err = p.collection.FindOne(ctx, bson.D{{"email", u.Email}}).Decode(&model)
	} else {
		err = p.collection.FindOne(ctx, bson.D{{"username", u.Username}}).Decode(&model)

	}
	if err != nil {
		return model.toRootUser(), err
	}
	err = p.hash.Compare(model.Password, u.Password)
	if err != nil {
		return model.toRootUser(), errors.New("Username/Email or Passoword is incorrect")
	}
	model.Password = "-"
	return model.toRootUser(), err
}

func (p *UserService) GetUsers() ([]root.User, error) {
	var users []root.User
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := p.collection.Find(ctx, bson.D{{}})
	fmt.Println(cursor, err)
	defer cursor.Close(ctx)
	if err != nil {
		return users, nil
	}
	for cursor.Next(ctx) {
		model := UserModel{}
		err := cursor.Decode(&model)
		if err != nil {
			return users, err
		}
		model.Password = "-"
		users = append(users, *model.toRootUser())
	}
	return users, nil
}
