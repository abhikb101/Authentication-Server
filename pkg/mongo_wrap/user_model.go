package mongo_wrap

import (
	root "user_auth/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserModel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string
	Username string `json:"username"`
	Password string
}

func newUserModel(u *root.User) *UserModel {
	return &UserModel{
		Email:    u.Email,
		Username: u.Username,
		Password: u.Password,
	}
}
func userModelIndex() []mongo.IndexModel {
	return []mongo.IndexModel{
		mongo.IndexModel{
			Keys: bson.M{
				"username": 1,
			},
			// create UniqueIndex option
			Options: options.Index().SetUnique(true),
		},
		mongo.IndexModel{
			Keys: bson.M{
				"email": 1,
			},
			// create UniqueIndex option
			Options: options.Index().SetUnique(true),
		},
	}
}
func (u *UserModel) toRootUser() *root.User {
	return &root.User{
		ID:       u.ID.Hex(),
		Email:    u.Email,
		Username: u.Username,
		Password: u.Password,
	}
}
