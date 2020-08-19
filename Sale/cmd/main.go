package main

import (
	"database/sql"
	"fmt"
	"log"
    _ "github.com/lib/pq"
	"github.com/karnsl/exercise/gopark/service/v1/handler"
)
func main() {
//connect database
	const (
		dbHost     = "192.168.1.39"
		dbName     = "BotApp"
		dbUser     = "postgres"
		dbPassword = "tonkla727426"
		dbPort     = 5432
	)
    dbSale := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbSale)
	if err != nil {
		log.Fatal("Connect Fail")
	}
	log.Println("Connect")
	defer db.Close()
    handler.Router(db)
}
