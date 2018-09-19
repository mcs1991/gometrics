package tool

import (
	"strconv"
	"fmt"
	"time"
	"strings"
)

//time to int 转换函数
func PerSecond_Int(before int, after int, time string) (string, bool) {
	var result interface{}
	var ok bool
	var rs string
	seconds, err := strconv.Atoi(time)
	if err != nil {
		fmt.Println(err.Error())
	}
	result = (after - before) / seconds
	// fmt.Println(result.(int))
	// tmp = fmt.Sprintf("%s", reflect.TypeOf(result))
	switch result.(type) {
	case int:
		// fmt.Println("int")
		rs = strconv.Itoa(result.(int))
	case int64:
		fmt.Println("int64")
		rs = strconv.FormatInt(result.(int64), 64)
	case float64:
		fmt.Println("float64")
		rs = strconv.FormatFloat(result.(float64), 'E', 5, 64)
	default:
		panic("not fount number type in perSecond")
	}
	if result.(int) > 0 {
		ok = true
	} else {
		ok = false
	}

	return rs, ok
}

//float to string 转换函数
func FloatToString(x float64, f int) string {
	rs := strconv.FormatFloat(x, 'f', f, 64)
	return rs
}

//time to float64 转换函数
func PerSecond_Float(before float64, after float64, time string) (string, bool) {
	var result interface{}
	var ok bool
	var rs string
	//转换时间为float64
	seconds, err := strconv.ParseFloat(time, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	result = (after - before) / seconds
	// tmp = fmt.Sprintf("%s", reflect.TypeOf(result))
	switch result.(type) {
	case int:
		// fmt.Println("int")
		rs = strconv.Itoa(result.(int))
	case int64:
		fmt.Println("int64")
		rs = strconv.FormatInt(result.(int64), 64)
		// case float32:
		// fmt.Println("float32")
		// rs = strconv.FormatFloat(result.(float32), 'f', 4, 32)
	case float64:
		fmt.Println("float64")
		rs = strconv.FormatFloat(result.(float64), 'f', 4, 64)
	default:
		panic("not fount number type in perSecond_Float")
	}

	if result.(int) > 0 {
		ok = true
	} else {
		ok = false
	}

	return rs, ok
}

//获取当前时间
func GetNowTime() string {
	f := fmt.Sprintf("%s", time.Now().Format("2006-01-02 15:04:05"))
	timeformatdate, _ := time.Parse("2006-01-02 15:04:05", f)
	convtime := fmt.Sprintf("%s", timeformatdate.Format("15:04:05"))
	return convtime
}

func String2float(stringtext string) (ftext []float64){
	var f []float64
	s := strings.Split(stringtext,",")
	if len(s) == 0{
		fmt.Printf("s is nil")
	}
	for i := range s {
		pf,_ := strconv.ParseFloat(s[i],64)
		f = append(f,pf)
	}
	return f
}