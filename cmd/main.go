package main

import (
	server "notes-api-golang/framework/http"
	"notes-api-golang/framework/mongo"
	sql "notes-api-golang/framework/sql"
)

func main() {
	sql.ConnectMysql()
	mongo.CreateMongoConnection()
	server.StartServer()
}
