package model

import(
	"os"
	"unicode/utf8"
	"strings"
	"github.com/zkc123/utils"
	"github.com/satori/go.uuid"
)
//图片类型
type Image struct{
	ID uint `gorm:"primary_key" json:"id"`
	Title string `json:"title"`
	OrignalTitle string `json:"orignalTitle"`
	URL string `json:"url"`
	Width uint 	`json:"width"`
	Heigth uint `json:"heigth"`
	Mime string `json:"mime"`
}
//ImageUploadInfo 图片上传后的相关信息（目录、文件名，UUIDName，请求URL）
type ImageUploadInfo struct{
	UploadDir string
	UploadFilePath string
	FileName string
	UUIDName string
	ImgURL string
}
//GenerateImgUploadedInfo 创建一个ImageUploadedInfo 
func GenerateImgUploadedInfo(ext string) ImageUploadInfo{
	sep:=string(os.PathSeparator)
	uploadImgDir:="image"//暂时写死，后续在修改
	length:=utf8.RuneCountInString(uploadImgDir)
	lastChar :=uploadImgDir[length-1:]
	ymStr :=utils.GetToDayYM(sep)
	var uploadDir string
	if lastChar !=sep{
		uploadDir=uploadImgDir+sep+ymStr
	}else{
		uploadDir=uploadImgDir+ymStr
	}
	uuidName:=uuid.NewV4().String()
	filename:=uuidName+ext
	uploadFilePath:=uploadDir+sep+filename
	imgURL:=strings.Join([]string{
		uploadImgDir,ymStr,filename,
	},"/")

	return ImageUploadInfo{
		ImgURL:imgURL,
		UUIDName:uuidName,
		FileName:filename,
		UploadDir:uploadDir,
		UploadFilePath:uploadFilePath,
	} 

}
