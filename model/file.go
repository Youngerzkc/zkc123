package model
import (
	"strings"
	"github.com/satori/go.uuid"
	"github.com/zkc123/utils"
	"unicode/utf8"
	"os"

)
//文件基本信息
type File struct{
	Id uint  `josn:"id" gorm:"pk"`
	FileName string `json:"filename"`
	FileSize int64  `json:"filesize"`
	Url string		`json:"url"`
}
//上传文件之后的信息
type FileUploadInfo struct{
	UploadDir string
	UploadFilePath string
	FileName string
	UUIDName string
	FileURL string
}
func GenerateFileUploadedInfo (ext string) FileUploadInfo{
	sep:=string(os.PathSeparator)
	uploaddir:="image/file"
	lenth:=utf8.RuneCountInString(uploaddir)
	lastChar:=uploaddir[lenth-1:]
	ymStr:=utils.GetToDayYM(sep)
	var uploadDir string
	if lastChar!=sep{
		uploadDir=uploaddir+sep+ymStr
	}else{
		uploadDir=uploaddir+ymStr
	}
	uuidName:=uuid.NewV4().String()
	filename:=uuidName+ext
	uploadPath:=uploadDir+sep+filename
	filenameUrl:=strings.Join([]string{
		uploadDir,ymStr,filename,
	},"/")
	return FileUploadInfo{
		UploadDir:uploadDir,
		UploadFilePath:uploadPath,
		FileName: filename,
		UUIDName: uuidName,
		FileURL:filenameUrl,
	}
}