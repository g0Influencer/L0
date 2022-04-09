package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var db *sql.DB

func Close() {
	db.Close()
}
func Connect(){
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	var err error
	db,err = sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		log.Println(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println(err)
	}
}

func GetDB() *sql.DB{
	return db
}