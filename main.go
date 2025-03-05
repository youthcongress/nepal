package main

import (
	"github.com/youthcongress/nepal/database"
	"github.com/youthcongress/nepal/server"
	"github.com/youthcongress/nepal/storage"
)

func main(){
	database.Connection()
	storage.Connection()
	
	server.StartServer()
}