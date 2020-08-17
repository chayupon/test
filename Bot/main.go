package main
import (
	"fmt"
	"net/http"
	//"log"
	"database/sql"
	"github.com/gin-gonic/gin"
	
	_ "github.com/lib/pq"
	
)
var db *sql.DB
var err error
const (
	host     = "192.168.1.39"
	port     = 5432 
	user     = "postgres"
	password = "Tonkla2639"
	dbName   = "BotApp"
)


//Member Table
type Member   struct {
	Email     string `json:"email"`
	UserID    string `json:"user_id"`
	UserName  string `json:"username"`
	Address   string `json:"address"`
	
}
func main() {
	dbBot := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbName)
	db, err := sql.Open("postgres", dbBot)
	if err != nil {
		fmt.Println("Connect Fail")
	}
	defer db.Close()
	r := gin.Default()
	r.POST("/register/email",register)
	r.GET("/member",listMember)
	r.Run()
}
func register(c *gin.Context) {
	sqlStr := "INSERT INTO member VALUES($1,$2,$3,$4)"
    var member Member
	err := c.BindJSON(&member)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(member)
	_, err = db.Exec(sqlStr, member.Email, member.UserID, member.UserName,member.Address)
	if err != nil {
		fmt.Println(err)
	}
}
func listMember(c *gin.Context) {
	sqlStr := "select username FROM member"
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println("failed")
		fmt.Println(err)
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

