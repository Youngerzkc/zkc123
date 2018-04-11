package logs

import (
	"fmt"
	"os"
	"bytes"
	"io"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// Instances a Logger middleware that will write the logs to gin.DefaultWriter
// By default gin.DefaultWriter = os.Stdout
func New() gin.HandlerFunc {
	return NewWithWriter(gin.DefaultWriter)
}

func checkFileExist(filename string ) bool  {
	var exist bool=true
	if _,err:=os.Stat(filename);os.IsNotExist(err){
		exist=false
	}
	return exist
}

func FileLogs(filename string) io.Writer  {

	var f *os.File
	var err error
	//file is exist?
	if checkFileExist(filename){
		f,err =os.OpenFile(filename,os.O_APPEND|os.O_RDWR,0666)
		if err!=nil{
			fmt.Println("open is failed")
		}
		fmt.Println("file is exist")
		return f
	}else{
		f,err =os.Create(filename)
		if err!=nil{
			fmt.Println("file create error")
		}
		fmt.Println("create file")
	}
	return f
}

// Instance a Logger middleware with the specified writter buffer.
// Example: os.Stdout, a file opened in write mode, a socket...
func NewWithWriter(out io.Writer) gin.HandlerFunc {
	pool := &sync.Pool{
		New: func() interface{} {
			buf := new(bytes.Buffer)
			return buf
		},
	}
	return func(c *gin.Context) {
		// Process request
		c.Next()

		//127.0.0.1 user-identifier frank [10/Oct/2000:13:55:36 -0700] "GET /apache_pb.gif HTTP/1.0" 200 2326
		w := pool.Get().(*bytes.Buffer)

		w.Reset()
		w.WriteString(c.ClientIP())
		w.WriteString(" - - ")
		w.WriteString(time.Now().Format("[02/Jan/2006:15:04:05 -0700] "))
		w.WriteString("\"")
		w.WriteString(c.Request.Method)
		w.WriteString(" ")
		w.WriteString(c.Request.URL.Path)
		w.WriteString(" ")
		w.WriteString(c.Request.Proto)
		w.WriteString("\" ")
		w.WriteString(strconv.Itoa(c.Writer.Status()))
		w.WriteString(" ")
		w.WriteString(strconv.Itoa(c.Writer.Size()))
		
		w.WriteString("\n")
		
		w.WriteTo(out)
		pool.Put(w)
	}
}
