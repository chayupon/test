package register

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	
)
//Query input to db
type Query struct{
	db *sql.DB
}
// Member table
type member struct {
	Email     string `json:"email"`
	UserID    string `json:"user_id"`
	UserName  string `json:"username"`
	Address   string `json:"address"`
	
}
//register
func (q Query) register(c *gin.Context) {
	sqlStr := "INSERT INTO member VALUES($1, $2, $3, $4)"
    var m member
	err := c.BindJSON(&m)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return 
	}
	//fmt.Printf("%+v",member)
	_, err = q.db.Exec(sqlStr, m.Email,m.UserID,m.UserName,m.Address)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status" :"OK",
	})
}
//Get (Select database)
func (q Query) listMember(c *gin.Context) {
	sqlStr := `SELECT "UserName" FROM member`
	rows, err := q.db.Query(sqlStr)
	if err != nil {
		fmt.Println("FAIL")
		c.String(http.StatusBadRequest, err.Error())
		return
		}
    defer rows.Close()
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"username": username,
		})
	
    }
}

