package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
  "time"
  "github.com/jinzhu/gorm"
  "strings"
  "io/ioutil"
  _ "github.com/jinzhu/gorm/dialects/mysql"

)

var db,_ = gorm.Open("mysql","root:huanglingyun0130@/usersinfo?charset=utf8&parseTime=True&loc=Local")

type User struct {
	Id       int `json:"-"`
	Name     string
	Password string
}

var userId int = 1
var user User

type Status struct {
	State   bool
	Detail  string 
}

var status =Status{false,""}

type Newuser struct {
	Username    string
	Oldpassword string
	Newpassword string
}

func regist(userInfo []byte)  {
	
	var a User
	json.Unmarshal(userInfo, &user) //将json转换成结构体

	res , _ =db.where("Name= ?",user.Name).Find(&a)
	if  res.RecordNotFound(){              //判断是否已经注册过
	user.Id = userId		
	userId += 1
	
	if user.Username == "" {
		status=Status{false,"用户名为空"}
		return
	}
	if user.Password == "" {
		status=Status{false,"密码为空"}
		return
	}

	status = Status{true, "注册成功"}

	db.Create(&user)
  
} else{
	status = Status{false, "用户名已存在"} //将状态回馈信息写入
}

func Modify (user Newuser) bool {
	var l User
	res := db.Where("username = ? AND password = ?",user.Username,user.Oldpassword).Find(&l)
	if res.RecordNotFound() {
		//将Newuser结构体中的信息转换到User中
		l.Name = user.Username
		l.Password = user.Oldpassword
		return false
	}
	//更新
	db.Model(&l).Where("username = ?",user.Name).Update("password",user.Newpassword)
	return true
}

//修改密码
func Update_Password(userInfo []byte) {
	var user Newuser
	json.Unmarshal(userInfo,&user)
	if !Modify(user) {
		status = Status{false,"用户名或旧密码错误"}
		return
	}
	status = Status{true,"修改成功"}
}

func main() {

    //db , = gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
  	defer db.Close()
	
	http.HandleFunc("/update",update_password)
	http.HandleFunc("/register",Regist)
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
	param_userName, found1 := req.Form["username"]
	param_passWord, found2 := req.Form["password"]
	if !(found1 && found2) {
		fmt.Fprint(w, "请输入用户名和密码")
		return
	}

	s,err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err) 
		}
	Update_Password(s)

	//向客户端发送json数据
	bytes, _ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))
}

func LoginTask(w http.ResponseWriter, req *http.Request) {
	fmt.Println("loginTask is running")

	time.Sleep(time.Second * 2)
	//获取客户端通过GET/POST方式传递的参数
	req.ParseForm()
	param_userName, found1 := req.Form["username"]
	param_passWord, found2 := req.Form["password"]
	if !(found1 && found2) {
		fmt.Fprint(w, "请输入用户名和密码")
		return
	}
	result := NewBaseJsonBean()
	userName := param_userName[0]
	passWord := param_passWord[0]
	
	var l User
	res := db.Where("username = ? AND password = ?",user.Username,user.Password).Find(&l)
	
	if  res.RecordNotFound(){
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

func Regist(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Regist is running")

	time.Sleep(time.Second * 2)
	//获取客户端通过GET/POST方式传递的参数
	req.ParseForm()
	param_userName, found1 := req.Form["username"]
	param_passWord, found2 := req.Form["password"]
	if !(found1 && found2) {
		fmt.Fprint(w, "请输入用户名和密码")
		return
	}

	s,err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err) 
		}
	Register(s)

	//向客户端发送json数据
	bytes, _ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))
}