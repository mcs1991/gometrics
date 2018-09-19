package main
// import "syscall"
// import "os"
import (
	"fmt"
	"os"
	"strconv"
	"time"
	"runtime"
	"os/signal"
)

var black string
var red string
var green string
var yellow string
var blue string
var purple string
var dgreen string
var white string
var mysql_switch string


func main() {
	//设置可用的最大CPU核数
	runtime.GOMAXPROCS(runtime.NumCPU())
	info := GetValue()
	// ss := basic{}
	second := Basic{}
	//捕获ctrl c信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		select {
		case s := <-c:
			fmt.Printf("\n\033[1;4;31m%s:Manual abort！\033[0m\n", s)
			ExecCommand("killall tcprstat")
			//不写就退不出来了
			os.Exit(1)
			// panic("退出")
		}
	}()
	//nocolor
	if info["nocolor"] == true {
		black = ""
		red = ""
		green = ""
		yellow = ""
		blue = ""
		purple = ""
		dgreen = ""
		white = ""

	} else {
		black = "black"
		red = "red"
		green = "green"
		yellow = "yellow"
		blue = "blue"
		purple = "purple"
		dgreen = "dgreen"
		white = "white"
	}
	//tcprstat
	if info["rt"] == true {
		go func() {
			var rt_cmd string
			rt_cmd = "tcprstat --no-header -p " + info["port"].(string) + " -t " + info["interval"].(string) + " -n 0 -l `/sbin/ifconfig | grep 'addr:[^ ]\\+' -o | cut -f 2 -d : | xargs echo | sed -e 's/ /,/g'` 1>/tmp/gometrics_tcprstat.log"
			fmt.Printf("rt:%s\n",rt_cmd)
			ExecCommand(rt_cmd)
		}()
	}
	//懒人模式
	if info["mysql"] == true {
		info["time"] = true
		info["com"] = true
		info["hit"] = true
		info["threads"] = true
		info["bytes"] = true
	}

	if info["innodb"] == true {
		info["time"] = true
		info["innodb_pages"] = true
		info["innodb_data"] = true
		info["innodb_log"] = true
		info["innodb_status"] = true
	}

	if info["sys"] == true {
		info["time"] = true
		info["load"] = true
		info["cpu"] = true
		info["swap"] = true
	}

	if info["lazy"] == true {
		info["time"] = true
		info["load"] = true
		info["cpu"] = true
		info["swap"] = true
		info["com"] = true
		info["hit"] = true
	}
	//mysql开关
	if info["username"].(string) != "" && info["password"].(string) != "" {
		myresult := CheckMysql(info)
		if myresult == "1" {
			mysql_switch = "on"
		} else{
			mysql_switch = "invalid"
		}
	}else {
		mysql_switch = "off"
	}
	//检查传参是否有效
	CheckFlag(info,mysql_switch)
	first := CreateCommand(info, 0, mysql_switch)

	interval := info["interval"].(string)
	interv, err := strconv.Atoi(interval)
	CheckErr(err)
	if info["count"] == 0 {
		i := 0
		//死循环
		for {
			second = CreateCommand(info, i, mysql_switch)
			ShowData(info, first, second, i, mysql_switch)
			first = second
			time.Sleep(time.Duration(interv) * time.Second)
			i++
		}
	} else {
		//有限循环
		for i := 0; i <= info["count"].(int); i++ {
			second = CreateCommand(info, i, mysql_switch)
			//second.Cpu_core, _ = strconv.ParseFloat(cpu_string, 64)
			ShowData(info, first, second, i, mysql_switch)
			first = second
			time.Sleep(time.Duration(interv) * time.Second)
		}
		ExecCommand("killall tcprstat")
		os.Exit(0)
	}
}