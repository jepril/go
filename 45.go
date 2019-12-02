package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"  // 导入驱动
	"github.com/gin-gonic/gin"
    "fmt"
)

type RegisterPayload struct {
    Username  string  `json:"username"`
    Password  string  `json:"password"`
}

func GetDatabase(username, password, host, port, dbname string) (*sql.DB, error) {
    address := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)
    db, err := sql.Open("mysql", address)
    if err != nil { // 打开失败
        return nil, err
    }
    err = db.Ping()
    if err != nil { // 连接失败
        return nil, err
    }
    return db, nil
}

func main()  {
	router := gin.Default()
	db, err := GetDatabase("root", "password", "127.0.0.1", "3306", "go_blog")
    if err != nil {
        panic("Link to database failed! Reason: " + err.Error())  // 结束程序并且打印错误原因
	}
	defer db.Close()

	userID := 0

	router.POST("/register", func(c *gin.Context) {
        var data RegisterPayload
		err = c.BindJSON(&data)
    	if err = db.QueryRow("SELECT id FROM users WHERE username = ?", data.Username).Scan(&userID); err != nil {
        	db.Query("INSERT INTO users (username, password) VALUES (?, ?)", data.Username, data.Password)
        	c.JSON(401, gin.H{
        	    "message": data.Username + data.Password,
        	})
        	return
    	} else {
        	c.JSON(401, gin.H{
            	"message": "User already existed.",
        	})
		}
		
	})
	router.Run()	
}