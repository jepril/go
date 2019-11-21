package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

func main() {
	fmt.Println("This is wbserver base!")
	
	//注册
	fmt.Println("请注册")
	var name string
	var	password	int

	fmt.Println("name:")
	fmt.Scan(&name)
	fmt.Println("password:")
	fmt.Scan(&password)
	
	//想把这些name，password存到文件里，等待调用
	

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
    if userName == "wangbiao" && passWord == "123456" {
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
