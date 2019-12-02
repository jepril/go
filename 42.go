package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"time"
)

 	var db, _ = gorm.Open("mysql", "root:258789hxr@/usersinfo?charset=utf8&parseTime=True&loc=Local")
	//var db *gorm.DB

type User struct {
	Id       int    `gorm:"AUTO_INCREMENT";json:"-"`
	Username     string `gorm:"type:varchar(100)"`
	Password string `grom:"type:varchar(100)"`
}

var userId int = 1
var user User

type Status struct {
	State  bool
	Detail string
}

var status = Status{false, ""}

type Newuser struct {
	Username    string
	Oldpassword string
	Newpassword string
}

func Existed(user User) bool {
	var l User
	res := db.Where("username= ? ",user.Username).Find(&l)
	//判断用户数据是否被查找到
	if res.RecordNotFound() {
			return  false
	}
	return true
}

func regist(userInfo []byte) {
	var user User
	json.Unmarshal(userInfo,&user)

	if Existed(user) {
		status = Status{false,"用户已存在"}
		return 
	}
	if user.Username == "" {
		status=Status{false,"用户名为空"}
		return
	}
	if user.Password == "" {
		status=Status{false,"密码为空"}
		return
	}
	db.Create(&user)//将用户信息储存到数据表中
	status = Status{true,"注册成功"}
}

func Modify(user Newuser) bool {
	var l User
	err := http.ListenAndServe("127.0.0.1:8081", nil)
	if err != nil {
		fmt.Println("ListenAndServe error:", err.Error())
	}
	res := db.Where("username = ? AND password = ?", user.Username, user.Oldpassword).Find(&l)
	if res.RecordNotFound() {
		//将Newuser结构体中的信息转换到User中
		l.Username = user.Username
		l.Password = user.Oldpassword
		return false
	}
	//更新
	db.Model(&l).Where("username = ?", user.Username).Update("password", user.Newpassword)
	return true
}

//修改密码
func Update_Password(userInfo []byte) {
	var user Newuser
	json.Unmarshal(userInfo, &user)
	if !Modify(user) {
		status = Status{false, "用户名或旧密码错误"}
		return
	}
	status = Status{true, "修改成功"}
}

func main() {
	//db, err := gorm.Open("mysql", "root:258789hxr@/usersinfo?charset=utf8&parseTime=True&loc=Local")
	//if err != nil {
	//	fmt.Println(err)
	//}
	defer db.Close()

	http.HandleFunc("/update", update_password)
	http.HandleFunc("/register", Regist)
	http.HandleFunc("/login", LoginTask)
	//服务器要监听的主机地址和端口号
	err := http.ListenAndServe("127.0.0.1:8081", nil)
	if err != nil {
		fmt.Println("ListenAndServe error:", err.Error())
	}

}

type BaseJsonBean struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewBaseJsonBean() *BaseJsonBean {
	return &BaseJsonBean{}
}

func update_password(res http.ResponseWriter, req *http.Request) {
	fmt.Println("update is running")

	time.Sleep(time.Second * 2)
	//获取客户端通过GET/POST方式传递的参数
	req.ParseForm()
	_, found1 := req.Form["username"]
	_, found2 := req.Form["password"]
	if !(found1 && found2) {
		fmt.Fprint(res, "请输入用户名和密码")
		return
	}
	result := NewBaseJsonBean()

	s, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	Update_Password(s)
	
	//向客户端发送json数据
	bytes, _ := json.Marshal(result)
	fmt.Fprint(res, string(bytes))
}

func LoginTask(w http.ResponseWriter, req *http.Request) {
	fmt.Println("loginTask is running")

	time.Sleep(time.Second * 2)
	//获取客户端通过GET/POST方式传递的参数
	req.ParseForm()
	_, found1 := req.Form["username"]
	_, found2 := req.Form["password"]
	if !(found1 && found2) {
		fmt.Fprint(w, "请输入用户名和密码")
		return
	}
	result := NewBaseJsonBean()
	//userName := param_userName[0]
	//passWord := param_passWord[0]

	var l User
	res := db.Where("username = ? AND password = ?", user.Username, user.Password).Find(&l)

	if res.RecordNotFound() {
		result.Code = 102
		result.Message = "用户名或密码不正确"
	} else {
		result.Code = 100
		result.Message = "登录成功"
	}
	//向客户端发送json数据
	bytes, _ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))
}

func Regist(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Regist is running")

	time.Sleep(time.Second * 2)
	
	s, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	
	regist(s)
	res.Write(Feedback(status))

}

func Feedback(a Status)[]byte {
	s,err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	return s
}
