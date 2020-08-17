
package main

import (
	"fmt"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	_ "github.com/lib/pq"
	
)
const (
	host     = "192.168.1.39"
	port     = 5432 
	user     = "postgres"
	password = "tonkla727426"
	dbname   = "BotApp"
)

//Member Table
type Member struct {

	Email string `json :"email"`
	UserID string `json :"userid"`
	UserName string `json :"username"`
	Address string  `json :"address"`
}

func main() {
	dbBot :=fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s,dbHost,dbPort,dbUser,dbPassword,dbName")
	db, err := sql.Open("postgres", dbBot)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer db.Close()
	r := gin.Default()
	r.POST("/Register/Email",register)
	r.GET("/member",listMember)
	r.RUN()
}
func register(c *gin.Context){
	sqlStr :="INCERT INTO member VALUES($1,$2,$3,$4)"
	var member Member
	err :=c.BindJSON()
}