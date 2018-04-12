package utils
import(
		"testing"
		"fmt"
		"strconv"
		"time"
		"strings"
)
func TestStrToIntMonth(t *testing.T){
	if StrToIntMonth("January")==0{
		t.Log("pass")
		fmt.Println("passd")	
	}else{
		t.Log("filed")
	}

}
func TestGetToDayYMD(t *testing.T){
	str:=GetToDayYMD("--")
	now:=time.Now()
	year:=now.Year()
	month:=StrToIntMonth(now.Month().String())
	date:=now.Day()
	var monthStr string
	var dateStr string

	if month<9 {
		monthStr="0"+strconv.Itoa(month+1)
	}else{
		monthStr=strconv.Itoa(month+1)
	}
	if date<10 {
		dateStr="0"+strconv.Itoa(date)
	}else{
		dateStr=strconv.Itoa(date)
	}
	testStr:=strconv.Itoa(year)+"--"+monthStr+"--"+dateStr
	if testStr==str{
		fmt.Println("TestGetToDayYMD passed")
	}
}
func TestGetToDayYM(t *testing.T){
	str:=GetToDayYM("++")
	now:=time.Now()
	year:=now.Year()
	month:= StrToIntMonth( now.Month().String())
	var monthStr string
	if month<9{
		monthStr="0"+strconv.Itoa(month+1)
	}else{
		monthStr=strconv.Itoa(month+1)
	}
	testStr:=strconv.Itoa(year)+"++"+monthStr
	if testStr==str{
		fmt.Println("PASSED TestGetToDayYM")
	}
}
func TestGetYesterdayYMD(t *testing.T){
	str:=GetYesterdayYMD("HH")
	now:=time.Now()
	todaySec:=now.Unix() //秒
	yesterdaySec:=todaySec-24*60*60;//昨天的秒
	yesterdaytime:=time.Unix(yesterdaySec,0)
	yesterdayYMD:=yesterdaytime.Format("2006-01-02")
	tsetStr:=strings.Replace(yesterdayYMD,"-","HH",-1)
	if str==tsetStr{
		fmt.Println("PASSED TestGetYesterdayYMD")
	}
}
func TestGetTomorrowYMD(t *testing.T){
	str:=GetTomorrowYMD("AA")
	now :=time.Now()
	todaySec:=now.Unix()
	tomorrowSec:=todaySec+24*60*60
	tomorrowTime:=time.Unix(tomorrowSec,0)
	tomorrowYMD:=tomorrowTime.Format("2006-01-02")
	testStr:= strings.Replace(tomorrowYMD,"-","AA",-1)
	if str==testStr{
		fmt.Println("PASSED TestGetTomorrowYMD")
	}
}