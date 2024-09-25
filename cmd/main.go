package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/phildehovre/go-gym/cmd/api"
	"github.com/phildehovre/go-gym/config"
	"github.com/phildehovre/go-gym/db"
)

func main() {

	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	initStorage(db)

	server := api.NewAPIServer(":5000", db)
	if err := server.Run(); err != nil {
		log.Fatalf("error running server %v", err)
	}

	if err != nil {
		log.Fatalf("error running server %v", err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database sucessfully connected")
}
