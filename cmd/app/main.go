package main

import (
	"fmt"
	"user_auth/pkg/crypto"
	"user_auth/pkg/mongo_wrap"
	"user_auth/pkg/server"
)

func main() {

	client, err := mongo_wrap.NewClient("localhost:27017")
	if err != nil {
		fmt.Println("Cannot Connect to mongo server")
		return
	}
	hash := crypto.Hash{}
	userservice := mongo_wrap.NewUserService(client, "UserDatabse", "Users", &hash)
	//token := token.TokenService{}
	serv := server.NewServer(userservice)
	serv.Start()
}
