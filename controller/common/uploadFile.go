package common

import(
	"github.com/gin-gonic/gin"
	"errors"
	"net/http"
	"github.com/zkc123/model"
	"strings"
	"os"
	"fmt"
	"github.com/zkc123/dao"
)
func UploadFile(c *gin.Context)(map[string]interface{},error){
	file,err:=c.FormFile("file")
	if err!=nil{
		return nil,errors.New("参数错误")
	}
	var filename =file.Filename
	var index = strings.LastIndex(filename,".")
	if index<0{
		return nil,errors.New("无效的文件名")
	}
	ext:=filename[index:]
	if len(ext)==1{
		return nil,errors.New("无效的扩展名")
	}
	fileUploadInfo:=model.GenerateFileUploadedInfo(ext)
	if err:=os.MkdirAll(fileUploadInfo.UploadDir,0777);err!=nil{
		fmt.Println(err.Error())	
		return nil,errors.New("mkdir error")
	}
	// desfileIndex:=strings.LastIndex(fileUploadInfo.UploadFilePath,"/")
	// desFIleName:=fileUploadInfo.UploadFilePath[:desfileIndex+1]+file.Filename
	// fmt.Println("FIle INFO is",fileUploadInfo.UploadFilePath)
	// fmt.Println("desFileName is",desFIleName)
	// 
	if err:=c.SaveUploadedFile(file,fileUploadInfo.UploadFilePath);err!=nil{
		return nil,errors.New("save file error!")
	}
	fileInfo:=model.File{
		FileName:filename,
		Url:fileUploadInfo.FileURL,
		FileSize:file.Size,
	}
	if err:=dao.DB.Create(&fileInfo).Error;err!=nil{
		return nil,errors.New("create file error")
	}
	return map[string]interface{}{
		"id":fileInfo.Id,
		"url":fileInfo.Url,
		"filename":fileInfo.FileName,
		"filesize":fileInfo.FileSize,
	},nil
}
func UploadFileHandler(c *gin.Context){
	data,err:=UploadFile(c)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"errNo":402,
			"msg":err.Error(),
			"data":gin.H{},
		})
		return 
	}
	c.JSON(http.StatusOK,gin.H{
		"msg":"SUCCESS",
		"data":data,
	})
}