package dao

import (

	"errors"
	"encoding/json"
	"crypto/md5"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/zkc123/model"
	
)
//redis 相关常量，为防止从redis中取数据时key混乱了，在此定义个key的名字
const(
	//LoginUser 用户信息
	LoginUser ="loginUser"

)

type User model.User

func (user *User)CheckPassWord(password string) bool   {
	var newUser User
	if password==""{
		return false
	}
	if err:=DB.Where("name=?",user.Name).First(&newUser).Error;err!=nil{
		return false
	}
	    return EncryptPassword(password,newUser.Salt())==newUser.Pass
}
func (user *User) Salt() string {
	var userSalt string
	userSalt="helloWorld"
	return userSalt
}
//用户信息保存到mysql数据库
func SaveUserMysql(user *model.User ) error{
	var err error
	var newUser User
	if err=DB.Where("name=?",user.Name).Find(&newUser).Error;err!=nil{
		err=DB.Save(user).Error
		return err
	} else {
		return errors.New("user name exist!")
	}
}
//mysql中删除用户信息
func DeleteUserMysql(name string)error{
	var err error
	var user User
	if err=DB.Where("name=?",name).Find(&user).Error;err!=nil{
		return err
	} 
	err=DB.Delete(&user).Error
	return err
}
//查找用户信息
func SelectUserMysql(name string)(*model.User,error){
	var err error
	var user model.User
	fmt.Println("查找用户信息")
	if err=DB.Where("name=?",name).Find(&user).Error;err!=nil{
		return nil,err
	} 
	return &user,nil
}
//修改用户密码
func UpdateUserMysql(user *model.User)error{
	var err error
	var newUser model.User
	fmt.Println("修改密码")
	if err=DB.Where("name=?",user.Name).Find(&newUser).Error;err!=nil{
		return err
	}
	newUser.Pass=user.Pass
	return 	SaveUserMysql(&newUser)
}
//给密码加密
func EncryptPassword(password ,salt string) (hash string)  {
	password=fmt.Sprintf("%x",md5.Sum([]byte(password)))
	hash=salt+password
	hash=fmt.Sprintf("%x",md5.Sum([]byte(hash)))
	fmt.Println("hash ",hash)
	return hash
}
//获取密码
func GetUserPasswd(name string) string {
	var passwd string
	var newUser User
	if err:=DB.Where("name=?",name).Find(&newUser).Error;err==nil{
		passwd=newUser.Pass
	}
	fmt.Println("user passwd",passwd)
	return passwd
}

//UserToRedis 将用户信息存到redis
func  UserToRedis(user *model.User) error {
	userBytes,err:=json.Marshal(user)
	if err!=nil{
		fmt.Println(err)
		return errors.New("error")
	}
	loginUserKey:=fmt.Sprintf("%s%d",LoginUser,user.ID)

	RedisCoon:=RedisPool.Get()
	defer RedisCoon.Close()
	
	if _,redisErr:=RedisCoon.Do("SET",loginUserKey,userBytes);redisErr!=nil{
		fmt.Println("redis set failed:",redisErr.Error())
		return errors.New("error")
	}
	return nil
}
//从redis中取出用户信息
func UserFromRedis(userID int)(model.User,error){
	var user model.User
	loginUser:=fmt.Sprintf("%s%d",LoginUser,userID)
	redisConn:=RedisPool.Get()
	defer redisConn.Close()
	userBytes,err:=redis.Bytes(redisConn.Do("GET",loginUser))
	if err!=nil{
		fmt.Println("未从redis 获取到数据")
		return user ,errors.New("未登录")
	}
	bytesErr:=json.Unmarshal(userBytes,&user)
	if bytesErr!=nil{
		fmt.Println(bytesErr)
		return user,errors.New("未登录")
	}
	fmt.Println("get redis user  ",user)
	return user,nil
}
func init(){
	if DB.HasTable("users")!=true{
		DB.CreateTable(&User{})
	}else{
		fmt.Println("table is exist")
	}
	DB.AutoMigrate(&User{})
}