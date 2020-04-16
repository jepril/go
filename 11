package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/net/publicsuffix"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"time"
)

//结构体
type accountReqeustParams struct {
	lt         string
	execution  string
	_eventId   string
	submit     string
	JSESSIONID string
}

type User struct {
	ID             int    `gorm:"id" json:"id"`
	User_id        string `gorm:"user_id" json:"user_id"`
	User_name      string `gorm:"user_name" json:"user_name"`
	Password       string `gorm:"password" json:"password"`
	Signture       string `gorm:"signture" json:"signture"`
	Image_url      string `gorm:"image_url" json:"image_url"`
	Background_url string `gorm:"background_url" json:"background_url"`
	Fans_num       int    `gorm:"fans_num" json:"fans_num"`
	Following_num  int    `gorm:"following_num" json:"following_num"`
	Upload_time    string `gorm:"upload_time" json:"upload_time"`
}

type Following_fans struct {
	Following_id string `json:"following_id"`
	Fans_id      string `json:"fans_id"`
}

type File struct{
	File_id		int			`gorm:"file_id" json:"file_id"`
	File_url	string		`gorm:"file_url" json:"file_url"`
	File_name	string		`gorm:"file_name" json:"file_name"`
	Format		string		`gorm:"format" json:"format"`
	Content		string		`gorm:"content" json:"content"`
	Subject		string		`gorm:"subject" json:"subject"`
	College		string		`gorm:"college" json:"college"`
	Type		string		`gorm:"type" json:"type"`
	Grade		int			`gorm:"grade" json:"grade"`
	Like_num	int			`gorn:"like_num" json:"like_num"`
	Collect_num	int			`gorm:"collect_num" json:"collect_num"`
	Download_num int		`gorm:"download_num" json:"download"`
}

//确认模拟登陆是否成功
func ConfirmUser(sid string, pwd string) bool {
	params, err := makeAccountPreflightRequest()
	if err != nil {
		log.Println(err)
		return false
	}

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Println(err)
		return false
	}
	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
		Jar:     jar,
	}
	//fmt.Println(params)
	result := makeAccountRequest(sid, pwd, params, &client)

	return result
}

// 预处理，打开 account.ccnu.edu.cn 获取模拟登陆需要的表单字段
func makeAccountPreflightRequest() (*accountReqeustParams, error) {
	var JSESSIONID string
	var lt string
	var execution string
	var _eventId string

	params := &accountReqeustParams{}

	// 初始化 http client
	client := http.Client{
		//Timeout: TIMEOUT,
	}

	// 初始化 http request
	request, err := http.NewRequest("GET", "https://account.ccnu.edu.cn/cas/login", nil)
	if err != nil {
		log.Println(err)
		return params, err
	}
	//request.Header.Add("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")
	// 发起请求
	resp, err := client.Do(request)
	if err != nil {

		log.Println(err)
		return params, err
	}

	// 读取 Body
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		log.Println(err)
		return params, err
	}

	// 获取 Cookie 中的 JSESSIONID
	for _, cookie := range resp.Cookies() {
		//fmt.Println(cookie.Value)
		if cookie.Name == "JSESSIONID" {
			JSESSIONID = cookie.Value
		}
	}

	if JSESSIONID == "" {
		log.Println("Can not get JSESSIONID")
		return params, errors.New("Can not get JSESSIONID")
	}

	// 正则匹配 HTML 返回的表单字段
	ltReg := regexp.MustCompile("name=\"lt\".+value=\"(.+)\"")
	executionReg := regexp.MustCompile("name=\"execution\".+value=\"(.+)\"")
	_eventIdReg := regexp.MustCompile("name=\"_eventId\".+value=\"(.+)\"")

	bodyStr := string(body)

	ltArr := ltReg.FindStringSubmatch(bodyStr)
	if len(ltArr) != 2 {
		log.Println("Can not get form paramater: lt")
		return params, errors.New("Can not get form paramater: lt")
	}
	lt = ltArr[1]

	execArr := executionReg.FindStringSubmatch(bodyStr)
	if len(execArr) != 2 {
		log.Println("Can not get form paramater: execution")
		return params, errors.New("Can not get form paramater: execution")
	}
	execution = execArr[1]

	_eventIdArr := _eventIdReg.FindStringSubmatch(bodyStr)
	if len(_eventIdArr) != 2 {
		log.Println("Can not get form paramater: _eventId")
		return params, errors.New("Can not get form paramater: _eventId")
	}
	_eventId = _eventIdArr[1]

	//log.Println("Get params successfully", lt, execution, _eventId)

	params.lt = lt
	params.execution = execution
	params._eventId = _eventId
	params.submit = "LOGIN"
	params.JSESSIONID = JSESSIONID

	return params, nil
}

func makeAccountRequest(sid, password string, params *accountReqeustParams, client *http.Client) bool {
	v := url.Values{}
	v.Set("username", sid)
	v.Set("password", password)
	v.Set("lt", params.lt)
	v.Set("execution", params.execution)
	v.Set("_eventId", params._eventId)
	v.Set("submit", params.submit)
	//fmt.Println(strings.NewReader(v.Encode()))
	request, err := http.NewRequest("POST", "https://account.ccnu.edu.cn/cas/login;jsessionid="+params.JSESSIONID, strings.NewReader(v.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")

	resp, err := client.Do(request)
	if err != nil {
		log.Print(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return false
	}

	// check
	reg := regexp.MustCompile("class=\"success\"")
	matched := reg.MatchString(string(body))
	if !matched {
		//log.Println("Wrong sid or pwd")
		return false
	}

	//log.Println("Login successfully")
	return true
}

//注册用户
func CreateUser(user_id string, user_name string, password string) {
	DB.Self.Model(&User{}).Create(&User{User_id: user_id, User_name: user_name, Password: password})
}

func CheckUserByUser_id(user_id string) bool {
	var l User
	res := DB.Self.Model(&User{}).Table("user").Where(User{User_id: user_id}).First(&l)
	if res.RecordNotFound() {
		return false
	}
	return true
}

func CreateToken(user_id string) string {
	keyInfo := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()"
	dataByte, _ := json.Marshal(user_id)
	var dataStr = string(dataByte)

	//使用Claim保存json

	data := jwt.StandardClaims{Subject: dataStr}
	tokenInfo := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	//生成token字符串
	token, _ := tokenInfo.SignedString([]byte(keyInfo))
	return token
}

func Viewing(user_id string) User {
	var l User
	fmt.Println(user_id)
	DB.Self.Model(&User{}).Table("user").Where(User{User_id: user_id}).First(&l)

	return l
}

func Token_info(Token string) (string, bool) {
	keyInfo := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()"
	//将token字符串转换为token对象（结构体更确切点吧，go很孤单，没有对象。。。）
	tokenInfo, _ := jwt.Parse(Token, func(token *jwt.Token) (i interface{}, e error) {
		return keyInfo, nil
	})

	//校验错误（基本）
	err := tokenInfo.Claims.Valid()
	if err != nil {
		println(err.Error())
		return err.Error(), false
	}

	finToken := tokenInfo.Claims.(jwt.MapClaims)
	//校验下token是否过期
	succ := finToken.VerifyExpiresAt(time.Now().Unix()+100000, true)
	var a string
	if succ {
		return a, false
	} else {
		//fmt.Println(finToken["sub"])
		return finToken["sub"].(string)[1:11], true

	}

	//fmt.Println("succ",succ)
	//获取token中保存的用户信息
	//fmt.Println(finToken["sub"])
	//return finToken["sub"]
}

func Background_modify(user_id string, background_url string) error {
	var tmpUser []User
	//tmpUser.Background_url = background_url
	//tmpUser.ID=2
	if err := DB.Self.Model(&tmpUser).Table("user").Where("user_id =?", user_id).Update("background_url", background_url).Error; err != nil {
		return err
	}
	fmt.Println(user_id)
	fmt.Println(background_url)
	return nil
}

func Image_modify(user_id string, image_url string) {
	//DB.Self.First(&User{})
	DB.Self.Model(&User{}).Table("user").Where(User{User_id: user_id}).Update(User{Image_url: image_url})
}

func Signture_modify(user_id string, signture string) {
	DB.Self.Model(&User{}).Table("user").Where(User{User_id: user_id}).Update(User{Signture: signture})
}

func CreateFollowing(fans_id string, following_id string) {
	DB.Self.Model(&Following_fans{}).Create(&Following_fans{Fans_id: fans_id, Following_id: following_id})
}

func GetDownFileid(uid string) (file_id []int,err error)  {
	if err = DB.Self.Table("file_downloader").Where("downloader_id", uid).Pluck("file_id",&file_id).Error ; err != nil {
		err = nil
		return
	}
	return
}

func List(fid []int)(file []File,err error){
	var File []File
	for _ ,data := range fid {
		File = nil
		if err := DB.Self.Table("file").Where("file_id",data).Joins("JOIN file_uploader ON file.file_id = file_uploader.file_id ").Select("file.file_name,file.content,file.subject,file.college,file_uploader.upload_time,file.file_url,file.collect_num,file.down_num,file.format").Scan(&File).Error ; err != nil {
			err = nil
		}
		for _ ,data1 := range File {
			file = append(file,data1)
		}
	}
	return
}

func GetUpFileid(uid string) (file_id []int,err error)  {
	if err = DB.Self.Table("file_uploader").Where("uploader_id", uid).Pluck("file_id",&file_id).Error ; err != nil {
		err = nil
		return
	}
	return
}


func GetCollectionFileid(uid string) (file_id []int,err error)  {
	if err = DB.Self.Table("file_collecter").Where("collecter_id", uid).Pluck("file_id",&file_id).Error ; err != nil {
		err = nil
		return
	}
	return
}

func CheckFollowingByFans_id(following_id string,fans_id string)bool  {
	var l Following_fans
	res := DB.Self.Model(&Following_fans{}).Table("following_fans").Where(Following_fans{Following_id: following_id,Fans_id :fans_id}).First(&l)
	if res.RecordNotFound() {
		return false
	}
	return true
}

func DeleteFollowing(fans_id string,following_id string)error  {
	if err :=DB.Self.Table("following_fans").Where("fans_id",fans_id,"following_id",following_id).Delete(Following_fans{}).Error ; err != nil {
		return err
	}
	return nil
}

func GetFansid(uid string) (fans_id []string,err error)  {
	if err = DB.Self.Table("following_fans").Where("following_id", uid).Pluck("fans_id",&fans_id).Error ; err != nil {
		err = nil

	}
	return
}

func FansList(fid []string)(user []User,err error){
	var User []User
	for _ ,data := range fid {
		User = nil
		if err := DB.Self.Table("user").Where("user_id",data).Select("user.id,user_name,image_url,signture").Scan(&User).Error ; err != nil {
			err = nil
		}
		for _ ,data1 := range User {
			user = append(user,data1)
		}
	}
	return
}

func FansNum(uid string)(num int ,err error)  {
	if err := DB.Self.Table("following_fans").Where("following_id",uid).Count(&num).Error; err != nil {
		err = nil

	}
	return
}

func GetFollowingid(uid string) (following_id []string,err error)  {
	if err = DB.Self.Table("following_fans").Where("fans_id", uid).Pluck("following_id",&following_id).Error ; err != nil {
		err = nil

	}
	return
}

type Following struct{
	ID             int    `gorm:"id" json:"id"`
	User_id        string `gorm:"user_id" json:"user_id"`
	User_name      string `gorm:"user_name" json:"user_name"`
	Signture       string `gorm:"signture" json:"signture"`
	Image_url      string `gorm:"image_url" json:"image_url"`
	Following_up   []File `gorm:"following_up" json:"following_up"`
}



func FollowingList(fid []string)(following []Following,err error){
	var Following []Following
	for _ ,data := range fid {
		Following = nil
		m ,_ := GetUpFileid(data)
		l ,_ := List(m)
		if err := DB.Self.Table("user").Where("user_id",data).Select("user.id,user_name,image_url,signture").Scan(&Following).Error ; err != nil {
			err = nil
		}
		for _ ,data1 := range Following {
			data1.Following_up= l
			following = append(following,data1)
			// following.Following_up= l
		}
	}
	return
}