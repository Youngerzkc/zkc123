package utils

import (

	"strings"
	"strconv"
	"time"

)
//strToIntMonth 字符串月份转整数月份
func StrToIntMonth (month string) int{
	var data =map[string]int{
		"January"   : 0,
        "February"  : 1,
        "March"     : 2,
        "April"     : 3,
        "May"       : 4,
        "June"      : 5,
        "July"      : 6,
        "August"    : 7,
        "September" : 8,
        "October"   : 9,
        "November"  : 10,
        "December"  : 11,
	}
	return data[month]
}
//GetToDayYMD 得到以sep为分隔符的年月日字符串
func GetToDayYMD(sep string)string{
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
	return strconv.Itoa(year)+sep+monthStr+sep+dateStr
}
//GetTodayYM 得到以sep为分隔符的年，月字符串(今天所属于的月份)
func GetToDayYM(sep string) string{
	now:=time.Now()
	year:=now.Year()
	month:= StrToIntMonth( now.Month().String())
	var monthStr string
	if month<9{
		monthStr="0"+strconv.Itoa(month+1)
	}else{
		monthStr=strconv.Itoa(month+1)
	}
	return strconv.Itoa(year)+sep+monthStr
}
//GetYesterdayYMD 得到以sep为分割的年月日字符串（昨天）
func GetYesterdayYMD(sep string)string{
	
	now:=time.Now()
	todaySec:=now.Unix() //秒
	yesterdaySec:=todaySec-24*60*60;//昨天的秒
	yesterdaytime:=time.Unix(yesterdaySec,0)
	yesterdayYMD:=yesterdaytime.Format("2006-01-02")
	return strings.Replace(yesterdayYMD,"-",sep,-1)
}
//GetTomorrowYMD得到以sep为分隔符的年月日字符串（明天）
func GetTomorrowYMD(sep string)string{
	now :=time.Now()
	todaySec:=now.Unix()
	tomorrowSec:=todaySec+24*60*60
	tomorrowTime:=time.Unix(tomorrowSec,0)
	tomorrowYMD:=tomorrowTime.Format("2006-01-02")
	return strings.Replace(tomorrowYMD,"-",sep,-1)
}