package main

import (
	"fmt"
	"strings"
	"time"
)

func main()  {
	timeString := strings.Replace("2019-05-02T10:24:26.248898138","T"," ",1)

	startTime := strings.Split(timeString,".")
	fmt.Println(startTime)

	tm,err := time.ParseInLocation("2006-01-02 15:04:05",startTime[0],time.Local)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tm)
	addTime,_ := time.ParseDuration("8h")
	fmt.Println(tm.Add(addTime))

	//fmt.Println(tm.Format("2006-01-02 15:04:05"))
	//timeString2 := "2019-05-04 10:24:26"
	//tm2,err :=  time.ParseInLocation("2006-01-02 15:04:05",timeString2,time.Local)
	//data := tm2.Sub(tm)
	//fmt.Println(data)
	//
	//startsTimeStr := "2019-05-02T10:24:26.248898138"
	//startsTimeStr = strings.Replace(startsTimeStr,"T"," ",1)
	//startsTime,err := time.ParseInLocation("2006-01-02 15:04:05",strings.Split(startsTimeStr,".")[0],time.Local)
	//endsTimeStr := "2019-05-02T12:24:26.248898138"
	//endsTimeStr = strings.Replace(endsTimeStr,"T"," ",1)
	//endsTime,err := time.ParseInLocation("2006-01-02 15:04:05",strings.Split(endsTimeStr,".")[0],time.Local)
	//
	//fmt.Println(endsTime.Sub(startsTime))


}
