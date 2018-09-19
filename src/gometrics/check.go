package main

import (
	"fmt"
	"os"
	"strings"
)


func CheckErr(errinfo error) {
	if errinfo != nil {
		fmt.Println(errinfo.Error())
		//os.Exit(-1)
	}
}

//检查输入的mysql参数是否有效（能否正常连接mysql）
func CheckMysql(info map[string]interface{})(myresult string){
	var cmd string
	var mysqlresult []string
	if info["host"].(string) != "" && info["host"].(string) != "localhost" {
		cmd = "mysql -u" + info["username"].(string) + " -p" + info["password"].(string) + " -h" + info["host"].(string) + " -P" + info["port"].(string) + " -e 'select 1'"
	} else {
		if info["host"].(string) != "" {
			cmd = "mysql -u" + info["username"].(string) + " -p" + info["password"].(string) + " -P" + info["port"].(string) + " -S " + info["socket"].(string) + " -e 'select 1'"
		}else {
			cmd = "mysql -u" + info["username"].(string) + " -p" + info["password"].(string) + " -h" + info["host"].(string) + " -P" + info["port"].(string) + " -S" + info["socket"].(string) + " -e 'select 1'"
		}
	}
	result := ExecCommand(cmd)
	resultstring := strings.Replace(result,"\n",",",1)
	mysqlresult = strings.Split(resultstring,",")
	return mysqlresult[0]
}

//检查传入的参数是否有效:
func CheckFlag(info map[string]interface{},mysql_switch string){
	if mysql_switch == "off" && (info["com"] == true || info["hit"] == true || info["threads"] == true|| info["bytes"] == true || info["innodb_pages"] == true || info["innodb_data"] == true || info["innodb_log"] == true || info["innodb_status"] == true || info["mysql"] == true || info["innodb"] == true){
		fmt.Println(ShowFont("[error]:You can't use the mysql parameter unless you specify the mysql configuration!","red","","",""))
		os.Exit(-1)
	}
	if mysql_switch == "invalid" {
		fmt.Println(ShowFont("[error]:your mysql dsn is incorrect!","red","","",""))
		fmt.Println(ShowFont("[error]:will not show mysql metrics!","red","","",""))
	}
	if info["interval"].(string) != "1" && (info["swap"] == true || info["net"] == true) {
		fmt.Println(ShowFont("[error]:can not collect swap or net metrics when the parameter interval <> 1","red","","",""))
		os.Exit(-1)
	}
}