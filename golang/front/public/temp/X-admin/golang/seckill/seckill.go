package main

import (
	"Send"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)
type Ret struct {
	Code int    `json:"code,int"`
	Data string `json:"data"`
}
type Boost struct {
	Name  string   `gorm:"primaryKey"`
	Num  int    `gorm:"column:num"`
	Price string `gorm:"column:price"`
	Time  time.Time `gorm:"column:time"`
	MaxNums  int `gorm:"column:max_nums"`
}
func (Boost) TableName() string {
	return "boost"
}
func connect() *gorm.DB {
	dsn := "admin:admin@tcp(182.92.172.68:3306)/admin?charset=utf8mb4&parseTime=True&loc=Local"
	//连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("------------连接数据库成功-----------")
	return db
}
func sayMore(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	name := r.FormValue("BoostName")
	maxNums,_ := strconv.ParseInt(r.FormValue("max_nums"),10,64)
	fmt.Println("username",username,"name:", name, "max_nums:",maxNums)
	b:=maxNums +'0'
	db := connect()
	boost := Boost{}
	errBoost1 := db.Table("boost").Where(" name = ?", name).First(&boost).Error
	errBoost2 := db.Table("boost").Where(" name = ?", name).Where("max_nums>=?",int(maxNums)).First(&boost).Error

	if errors.Is(errBoost1, gorm.ErrRecordNotFound) {
		fmt.Println("商品名称错误")

		return
	}
	if  boost.Time.After(time.Now()) {
		fmt.Println("未到抢购时间")
		return
	}
	if errors.Is(errBoost2,gorm.ErrRecordNotFound){
		fmt.Println("超过抢购上限")
		return
	}
	Send.S(username,name,string(b))
	r.ParseForm() //解析参数，默认是不会解析的
	printRequest(w, r)
}

func printRequest(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/zhongzhuan.html")
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
func tzfunc(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("html/status.html")
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/order", sayMore)
	http.HandleFunc("/tz", tzfunc)
	err := http.ListenAndServe(":8888", nil)    //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
