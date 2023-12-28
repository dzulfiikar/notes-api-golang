package main

import (
	"notes-api-golang/framework/gocron"
	server "notes-api-golang/framework/http"
	"notes-api-golang/framework/mongo"
	sql "notes-api-golang/framework/sql"
)

func main() {
	sql.ConnectMysql()
	mongo.CreateMongoConnection()
	gocron.StartGoCronScheduler()
	server.StartServer()
}
