package mysql

import (
	"database/sql"
	"farhadiis/order-process/utils"
	"github.com/go-sql-driver/mysql"
	"log"
)

func Connect() *sql.DB {
	cfg := mysql.Config{
		Addr:                 utils.GetEnv("MYSQL_URI"),
		User:                 utils.GetEnv("MYSQL_USER"),
		Passwd:               utils.GetEnv("MYSQL_PASSWORD"),
		DBName:               utils.GetEnv("MYSQL_DATABASE"),
		Net:                  "tcp",
		AllowNativePasswords: true,
	}

	client, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("MySql initialize successful")
	return client
}

func Disconnect(client *sql.DB) {
	if err := client.Close(); err != nil {
		panic(err)
	}
}
