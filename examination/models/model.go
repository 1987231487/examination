package models
/*
1.这里是数据库的表对应的模板
2.初始化数据库
*/
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"examination/pkg/setting"
	"log"
)
var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`   //设置一些tags
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init(){
	sec,err:=setting.Cfg.GetSection("database")//获取 数据库分组配置文件
	if err!=nil{
		log.Fatal(err)
	}
	dbType:=sec.Key("TYPE").String()     //读取ini里的数据库配置文件
	dbName:=sec.Key("NAME").String()
	user:=sec.Key("USER").String()
	password:= sec.Key("PASSWORD").String()
	host:= sec.Key("HOST").String()



	args:=fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True",
		user,
		password,
		host,
		dbName,
	)
	db,err=gorm.Open(dbType,args)   //链接数据库
	if err!=nil{
		log.Println("open db fail=",err)
	}

	db.SingularTable(true)   //禁用默认表名的复数形式
	db.LogMode(true)      // 启用Logger，显示详细日志
	db.DB().SetMaxIdleConns(10) //设置最大的空闲连接数为10
	db.DB().SetMaxOpenConns(100) //设置最大的并发打开连接数为100
}

func CloseDB(){
	defer db.Close()
}