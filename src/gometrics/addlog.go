package main

import (
	"time"
	"fmt"
	"os"
	"log"
)

func Add_log(flag_log map[string]interface{}, info string) {
	var file_name string
	if flag_log["logfile_by_day"].(bool) == true {
		t := time.Now()
		if flag_log["logfile"].(string) != "none" {
			file_name = flag_log["logfile"].(string) + "_" + fmt.Sprintf("%s", t.Format("2006-01-02")) + ".log"
		} else {
			file_name = "/tmp/gometrics" + "_" + fmt.Sprintf("%s", t.Format("2006-01-02")) + ".log"
		}
	} else {
		if flag_log["logfile"].(string) != "none" {
			file_name = flag_log["logfile"].(string)
		} else {
			file_name = "/tmp/gometrics.log"
		}
	}

	lf, err := os.OpenFile(file_name, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0606)
	CheckErr(err)

	defer lf.Close()

	l := log.New(lf, "", os.O_APPEND)

	l.Printf("%s\n", info)

}
