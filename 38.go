package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
    "os"
    "io"
)

//var Users = make ([]User,0)
var userId  int=1

type    User    struct{
    Id  int `json:"-"`
    Name    string
    Password    int
}

func WriteFile(path string)  {
    f,err := os.Create(path)
    if err  !=nil{
        fmt.Println("err = ",err)
        return
    }

    defer f.Close()

    var user User
    fmt.Println("name")
    fmt.Scan(&user.Name)
    fmt.Println("password")
    fmt.Scan(&user.Password)
//    var user User
//json.Unmarshal(infor,&user)

//user.Id = userId
//userId+=1
//Users=append(Users,user)
}

func Readfile(path string,a int) (buf []byte ) {
f , err := os.Open(path)
if err != nil {
fmt.Println("err",err)
return
}

defer	f.Close()
//buf = make([]User, 1024*2)

n , err1 := f.Read(buf)
if err != nil && err != io.EOF {
    fmt.Println("err1 = ",err1)
    return 
}
    if a==3 {
        fmt.Println("buf = ",string(buf[:n]))
    }
return
}

func main() {
    
    var a int
    var buf []byte
	//注册
    fmt.Println("1注册")
    fmt.Println("2登陆")
    fmt.Println("3查看信息")
    fmt.Scan(&a)
	
	//想把这些name，password存到文件里，等待调用
    path:="1.txt"
    WriteFile(path)
    Readfile(path,a)
    buf1:=User(buf)

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
    param_userName, found1 := req.Form["userName"]
    param_passWord, found2 := req.Form["passWord"]
    if !(found1 && found2) {
        fmt.Fprint(w, "请输入用户名和密码")
        return
    }
    result := NewBaseJsonBean()
    userName := param_userName[0]
    passWord := param_passWord[0]
    s := "userName:" + userName + ",password:" + passWord
    fmt.Println(s)
    if userName ==  buf.Name&& passWord == buf.Password  {
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
