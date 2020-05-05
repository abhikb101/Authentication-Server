package mongo_wrap_test

import (
	"log"
	"testing"
	root "user_auth/pkg"
	"user_auth/pkg/mock"
	"user_auth/pkg/mongo_wrap"
)

const (
	mongoUrl           = "localhost:27017"
	dbName             = "test_db"
	userCollectionName = "user"
	testUsername       = "integration_testser"
	testEmail          = "integration_test_email"
	testPassword       = "integration_test_pword"
)

func Test_UserService(t *testing.T) {
	t.Run("CreateUser", createUser_should_insert_user_into_mongo)
	t.Run("GetByUsername", GetByUsername_should_get_user_from_mongo)
	t.Run("GethUsers", GetUsers_should_get_all_users_from_mongo)
}

func createUser_should_insert_user_into_mongo(t *testing.T) {
	//Arrange
	client, err := mongo_wrap.NewClient(mongoUrl)
	if err != nil {
		log.Fatalf("Unable to connect to mongo: %s", err)
	}
	defer func() {
		client.Close()
	}()
	mockhash := mock.Hash{}
	userService := mongo_wrap.NewUserService(client, dbName, userCollectionName, &mockhash)

	user := root.User{
		Email:    testEmail,
		Username: testUsername,
		Password: testPassword}

	//Act
	err = userService.CreateUser(&user)

	//Assert
	if err != nil {
		t.Error("Unable to create user:", err)
	}
}
func GetByUsername_should_get_user_from_mongo(t *testing.T) {
	//Arrange
	client, err := mongo_wrap.NewClient(mongoUrl)
	if err != nil {
		log.Fatalf("Unable to connect to mongo: %s", err)
	}
	defer func() {
		client.Close()
	}()
	mockhash := mock.Hash{}
	userService := mongo_wrap.NewUserService(client, dbName, userCollectionName, &mockhash)
	user := root.User{
		Email:    testEmail,
		Username: testUsername,
		Password: testPassword}
	//Act
	var results *root.User
	results, err = userService.GetByUsername(testUsername)
	if err != nil {
		t.Error(err)
	}
	if results.Username != user.Username {
		t.Error("Incorrect Username. Expected , Got: ", testUsername, results.Username)
	}
}
func GetUsers_should_get_all_users_from_mongo(t *testing.T) {
	//Arrange
	client, err := mongo_wrap.NewClient(mongoUrl)
	if err != nil {
		log.Fatalf("Unable to connect to mongo: %s", err)
	}
	defer func() {
		client.Close()
	}()
	mockhash := mock.Hash{}
	userService := mongo_wrap.NewUserService(client, dbName, userCollectionName, &mockhash)
	//Act
	var result []root.User
	result, err = userService.GetUsers()
	if err != nil {
		t.Error(err)
	}
	t.Log(result)

}
