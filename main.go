package main

import (
	"github/mohanapranes/book_trust_go/config"
	"github/mohanapranes/book_trust_go/config/database"
	"github/mohanapranes/book_trust_go/config/server"
)

func main() {
	config := config.GetConfig()
	database := database.NewPostgresDatabase(config)
	server.Start(config, database)
}
