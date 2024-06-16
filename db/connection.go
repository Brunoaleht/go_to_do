package db

import (
	"database/sql"
	"fmt"
	"log"
	"modules/configs"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error){ 
	config := configs.GetDBConfig()

	screenConnection := fmt.Sprintf("host=%s port=%s dbname=%s sslmode=disable", config.Host, config.Port, config.DBName)

	conn, err := sql.Open("postgres", screenConnection)
	if err != nil {
		log.Fatalf("Critical error opening connection on database: %v", err)
	}

	err = conn.Ping()

	return conn, err
}