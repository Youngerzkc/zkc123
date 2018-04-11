package dao

import (

	"strconv"
	"time"
	"fmt"
	"os"
	"io/ioutil"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_"github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

//DB数据库连接
var DB *gorm.DB

//RedisPool Redis连接池
var RedisPool *redis.Pool


type PraseConf struct{
	Redis []Redisf	
	Mysql []Mysqlf
}
type Redisf struct{
	Port int
	Host string
	MaxActive int
    MaxIdle int
}
type Mysqlf struct{
	Port int
	Host string
	User string
	Password string
	Database string
	Charset string
	MaxIdleConns int
	MaxOpenConns int
}



func GetPraseConf()  *PraseConf  {
	tmp:=&PraseConf{}
	pwdPath,_:=os.Getwd()
	// fmt.Println("pwdpath is ",pwdPath)
	yamlPath:=pwdPath+"/conf/conf.yml"
	// fmt.Println("yamlPath is",yamlPath)
	yamlFile,err:=ioutil.ReadFile(yamlPath)
	// fmt.Println("yamlPath",yamlFile)
	if err!=nil{
		fmt.Printf("yamlFile get err %#v",err)
	}
	err=yaml.Unmarshal(yamlFile,tmp)
	// fmt.Println("yaml ",*tmp)
	if err!=nil{
		fmt.Printf("Unmarshal:%v",err)
	}
	// fmt.Println("yaml ",)
	return tmp
}
func initDB(){
	user:=GetPraseConf().Mysql[0].User
	host:=GetPraseConf().Mysql[0].Host
	port:=GetPraseConf().Mysql[0].Port
	passwd:=GetPraseConf().Mysql[0].Password
	database:=GetPraseConf().Mysql[0].Database
	charset:=GetPraseConf().Mysql[0].Charset
	url:="@tcp("+host+":"+strconv.Itoa(port)+")/"+database+"?"+"charset="+charset+"&parseTime=True&loc=Local"
	db,err:=gorm.Open("mysql",user+":"+passwd+url)
	if err!=nil{
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	db.DB().SetMaxIdleConns(GetPraseConf().Mysql[0].MaxIdleConns)
	db.DB().SetMaxOpenConns(GetPraseConf().Mysql[0].MaxOpenConns)
	DB=db
}

func  initRedis()  {
	RedisPool=&redis.Pool{
		MaxActive:	GetPraseConf().Redis[0].MaxActive,
		MaxIdle:	GetPraseConf().Redis[0].MaxIdle,
		IdleTimeout:	20*time.Second,
		Dial: func()(redis.Conn,error){
			c,err:=redis.Dial("tcp",GetPraseConf().Redis[0].Host+":"+strconv.Itoa(GetPraseConf().Redis[0].Port))
			if err!=nil{
				return nil,err
			}
			return c,nil
		},
	}

}
func  init()  {
	initRedis()
	initDB()
	
}