package seckill2

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

//Boost
type Boost struct {
	Name  string   `gorm:"primaryKey"`
	Num  int    `gorm:"column:num"`
	Price string `gorm:"column:price"`
	Time  time.Time `gorm:"column:time"`
	Maxnums  int `gorm:"column:max_nums"`
}
func (Boost) TableName() string {
	return "boost"
}
//room ...
type Purinfo struct {
	UserName           string  `gorm:"column:username"`
	GoodsName   string    `gorm:"column:goodsname"`
	Nums       int    `gorm:"column:nums"` //1男 0女
	Success       int    `gorm:"column:success"` //0 新生不可见
	Time time.Time `gorm:"column:time"`
	Reason string `gorm:"column:reason"`
}
func (Purinfo) TableName() string {
	return "purinfo"
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

func sayMore() {
	db := connect()
	tx := db.Begin()
	boost := Boost{}
	errBoost2 := db.Table("boost").Where(" name = ?", name).Where("max_nums>=?",maxNums).Where("num>=?",maxNums).First(&boost).Error
	if errors.Is(errBoost2,gorm.ErrRecordNotFound){
		fmt.Println("余量小于欲购件数")
		purinfo := Purinfo{
			UserName: username,
			GoodsName:  name,
			Nums: maxNums,
			Success: 0,
			Time: time.Now(),
			Reason: "余量小于欲购件数",
		}
		//插入数据
		db.Create(&purinfo)
		return
	}
	//对数据库进行修改。
	boost.Num =boost.Num-maxNums
	db.Save(&boost)
	//在purinfo表中添加用户购买信息
	purinfo := Purinfo{
		UserName: username,
		GoodsName:  name,
		Nums: maxNums,
		Success: 1,
		Time: time.Now(),
	}

	//插入数据
	db.Create(&purinfo)
	// 否则，提交事务
	tx.Commit()
	//
	//
	fmt.Println("------操作数据库成功---------")
}
var username string
var name string
var maxNums int
func Deal(u string ,n string ,m int ) {
	username=u
	name=n
	maxNums=m
	sayMore()
}
