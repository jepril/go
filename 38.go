package main

import (
	"encoding/json"
	"fmt"
	// "io"
	"net/http"
	"os"
    "time"
    "strings"
    "io/ioutil"
)

//var Users = make ([]User,0)

type User struct {
	Id       int `json:"-"`
	Name     string
	Password string
}

var userId int = 1
var buf User

func WriteFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer f.Close()

	var user User
    fmt.Println("name")
    fmt.Scan(&user.Name)
	fmt.Fprintf(f, "%s ", user.Name)
    fmt.Println("password")
    fmt.Scan(&user.Password)
	fmt.Fprintf(f, "%s", user.Password)
	//    var user User
	//json.Unmarshal(infor,&user)

	//user.Id = userId
	//userId+=1
	//Users=append(Users,user)
}

func Readfile(path string, a int) (buff User) {
	// f, err := os.Open(path)
	// if err != nil {
	// 	fmt.Println("err", err)
	// 	return
	// }

	// defer f.Close()

	//var buf1 []byte
    buf1, err := ioutil.ReadFile(path)
    if err != nil {
		fmt.Println("err = ", err)
		return
	}

	index := strings.Index(string(buf1), " ")
	if index == -1 {
		return
    }
    buff.Name =string(buf1[:index])
    buff.Password =string(buf1[index+1:])

	
	if a == 3 {
		fmt.Println("buff = ", buff)
	}
	return
}

func main() {

	var a int

	//注册
	fmt.Println("1注册")
	fmt.Println("2登陆")
	fmt.Println("3查看信息")
	fmt.Scan(&a)

	//想把这些name，password存到文件里，等待调用
    path := "1.txt"

	WriteFile(path)
	buf = Readfile(path, a)

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
	s := "userUame:" + userName + ",password:" + passWord
    fmt.Println(s)
    fmt.Println(buf)
    fmt.Println(buf.Password)
    fmt.Println(buf.Name) 
	if userName == buf.Name && passWord == buf.Password {
		result.Code = 100
		result.Message = "登录成功"
	} else {
		result.Code = 102
		result.Message = "用户名或密码不正确"
	}
	//向客户端发送json数据
	bytes, _ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))
}
