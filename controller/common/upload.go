package common

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"errors"
	"strings"
	"mime"
	"github.com/zkc123/model"
	"os"
	"net/http"
	"github.com/zkc123/dao"
)
//upload file 
func Upload(c *gin.Context)(map[string]interface{},error ){
	file,err:=c.FormFile("upFile")
	if err!=nil{
		return nil,errors.New("参数错误")
	}
	var filename=file.Filename
	var index=strings.LastIndex(filename,".")
	if index <0{
		return nil,errors.New("无效的文件名")
	}
	ext:=filename[index:]
	if len(ext)==1{
		return nil,errors.New("无效的扩展名")
	}
	var mimeType=mime.TypeByExtension(ext)
	if mimeType==""{
		return nil,errors.New("无效的扩展类型")
	} 
	imgUploadedInfo:=model.GenerateImgUploadedInfo(ext)
	fmt.Println(imgUploadedInfo.UploadDir)
	if err:=os.MkdirAll(imgUploadedInfo.UploadDir,0777);err!=nil{
		fmt.Println(err.Error())
		return nil,errors.New("mkdir error")
	}
	if err:=c.SaveUploadedFile(file,imgUploadedInfo.UploadFilePath);err!=nil{
		fmt.Println(err.Error())
		return nil,errors.New("save img error!")
	}
	image :=model.Image{
		Title:imgUploadedInfo.FileName,
		OrignalTitle:filename,
		URL:imgUploadedInfo.ImgURL,
		Width:0,
		Heigth:0,
		Mime:mimeType,
	}
	if err:=dao.DB.Create(&image).Error;err!=nil{
		fmt.Println(err.Error())
		return nil,errors.New("create image error")
	}
	return map[string]interface{}{
		"id":image.ID,
		"url":imgUploadedInfo.ImgURL,
		"title":imgUploadedInfo.FileName,
		"original":filename,
		"type":mimeType,
	},nil
}
//UploadHandler 文件上传
func UploadHandler(c *gin.Context){
	data,err:=Upload(c)
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