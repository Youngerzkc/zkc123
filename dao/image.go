package dao

import (
	"github.com/zkc123/model"
	"fmt"
)
type Image model.Image

func init(){
	if DB.HasTable("images")!=true{
		DB.CreateTable(&Image{})
	}else{
		fmt.Println("table is exist")
	}
	DB.AutoMigrate(&Image{})
}