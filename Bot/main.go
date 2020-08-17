package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)
var db *sql.DB
var err error
const (
	dbHost     = "192.168.1.39"
	dbUser     = "postgres"
	dbPassword = "tonkla727426"
	dbPort     = 5432
	dbName     = "BotApp"
)

// Member table
type Member struct {
	Email    string `json:"Email"`
	UserID    string `json:"user_id"`
	UserName     string `json:"username"`
	Address string `json:"address"`
	
}
func main() {
    dbsale := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err = sql.Open("postgres", dbsale)
	if err != nil {
		fmt.Println("Connect Fail")
		panic(err)
	}
	fmt.Println("Database Connected.")
	defer db.Close()

	r := gin.Default()

	r.POST("/register/email", register)
	r.GET("/member", listMember)
	r.Run()
}
func register(c *gin.Context) {
	sqlStr := "INSERT INTO member VALUES($1, $2, $3, $4)"

	var member Member
	err := c.BindJSON(&member)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(member)
	_, err = db.Exec(sqlStr, member.Email,member.UserID,member.UserName,member.Address)
	if err != nil {
		fmt.Println(err)
	}
}
func listMember(c *gin.Context) {
	sqlStr := "SELECT username  FROM member"
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println("FAIL")
		}
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"username":    username,
		})
	}
	if !rows.NextResultSet() {
		fmt.Println(rows.Err())
	}
}
