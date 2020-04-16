package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func main()  {
	keyInfo := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()"
 
	//将部分用户信息保存到map并转换为json
	info := map[string]interface{}{}
	info["userName"] = "Shisan"
	dataByte,_:= json.Marshal(info)
	var dataStr = string(dataByte)
 
	//使用Claim保存json
	//这里是个例子，并包含了一个故意签发一个已过期的token
	data := jwt.StandardClaims{Subject:dataStr,ExpiresAt:time.Now().Unix()-1000}
	tokenInfo := jwt.NewWithClaims(jwt.SigningMethodHS256,data)
	//生成token字符串
	tokenStr,_ := tokenInfo.SignedString([]byte(keyInfo))
	fmt.Println("myToken is: ",tokenStr)
 
	//将token字符串转换为token对象（结构体更确切点吧，go很孤单，没有对象。。。）
	tokenInfo , _ = jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, e error) {
		return keyInfo,nil
	})
 
	//校验错误（基本）
	err := tokenInfo.Claims.Valid()
	if err!=nil{
		println(err.Error())
	}
	
	finToken := tokenInfo.Claims.(jwt.MapClaims)
    //校验下token是否过期
	succ := finToken.VerifyExpiresAt(time.Now().Unix(),true)
	fmt.Println("succ",succ)
    //获取token中保存的用户信息
	fmt.Println(finToken["sub"])

}