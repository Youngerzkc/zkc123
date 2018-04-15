package dao

import (
	"github.com/zkc123/model"
	"fmt"
)
type File model.File

func init(){
	if DB.HasTable("files")!=true{
		DB.CreateTable(&File{})
	}else{
		fmt.Println("table is exist")
	}
	DB.AutoMigrate(&File{})
}