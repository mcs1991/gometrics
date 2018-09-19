package main

import (
	"fmt"
	"strconv"
	"time"
	"strings"
	"plot"
	"tool"
)

var queryps string
var tranps string
var usrpercent string
var syspercent string
var idlepercent string
var iowaitpercent string
var iobusy []float64

//根据传入参数展示结果
func ShowData(flag_info map[string]interface{}, first Basic, second Basic, count int, mysql_switch string) {
	var title_summit string
	var title_detail string
	var data_detail string
	var pic string
	interval, _ := strconv.Atoi(flag_info["interval"].(string))
	if count == 0 {
		var tmp_used string
		tmp_max_connections, _ := strconv.Atoi(second.var_max_connections)
		if second.Max_used_connections > (tmp_max_connections * 7 / 10) {
			tmp_used = ShowFont(strconv.Itoa(second.Max_used_connections), red, "", "", "y")
		} else {
			tmp_used = ShowFont(strconv.Itoa(second.Max_used_connections), "", "", "", "")
		}

		var tmptable string
		tmp_table_x := float64(second.Created_tmp_disk_tables) / float64(second.Created_tmp_tables) * 100
		if tmp_table_x < 10.0 {
			tmptable = ShowFont(tool.FloatToString(tmp_table_x, 2), green, "", "", "")
		} else {
			tmptable = ShowFont(tool.FloatToString(tmp_table_x, 2), red, "", "", "y")
		}
		systime := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
		pic += ShowFont(".==========================================================================================================.\n", green, "", "", "")
		pic += ShowFont("|",green,"","","") + "Program:" + ShowFont("GoMetrics",red,"","","") + "\n"
		pic += ShowFont("|", green, "", "", "") + "Author:" + ShowFont("MeiJia",red,"","","") + "\n"
		pic += ShowFont("|",green,"","","") + "datetime:" + ShowFont(systime,red,"","","")  + "\n"
		pic += ShowFont("'=========================================================================================================='\n\n", green, "", "", "")
		pic += ShowFont("HOST: ", red, "", "", "") + ShowFont(strings.Replace(second.hostname, "\n", "", -1), yellow, "", "", "")  + "\n"
		pic += ShowFont("IP: ", red, "", "", "") + ShowFont(strings.Replace(second.ip, "\n", "", -1), yellow, "", "", "") + "\n"
		if mysql_switch == "on" {
			pic += ShowFont("DB Info: ", red, "", "", "") + ShowFont(second.db, yellow, "", "", "") + "\n"
			pic += ShowFont(" binlog_format", purple, "", "", "") + "[" + second.var_binlog_format + "]" + ShowFont(" max_binlog_cache_size", purple, "", "", "") + "[" + ChangeUntils(second.var_max_binlog_cache_size) + "]" + ShowFont(" max_binlog_size", purple, "", "", "") + "[" + ChangeUntils(second.var_max_binlog_size) + "]" + ShowFont(" sync_binlog", purple, "", "", "") + "[" + second.var_sync_binlog + "]" + "\n"
			pic += ShowFont(" max_connect_errors", purple, "", "", "") + "[" + second.var_max_connect_errors + "]" + ShowFont(" max_connections", purple, "", "", "") + "[" + second.var_max_connections + "]" + ShowFont(" max_user_connections", purple, "", "", "") + "[" + second.var_max_user_connections + "]" + ShowFont(" max_used_connections", purple, "", "", "") + "[" + tmp_used + "]" + "\n"
			pic += ShowFont(" open_files_limit", purple, "", "", "") + "[" + second.var_open_files_limit + "]" + ShowFont(" table_definition_cache", purple, "", "", "") + "[" + second.var_table_definition_cache + "]" + ShowFont(" Aborted_connects", purple, "", "", "") + "[" + second.Aborted_connects + "]" + ShowFont(" Aborted_clients", purple, "", "", "") + "[" + second.Aborted_clients + "]" + "\n"
			pic += ShowFont(" Binlog_cache_disk_use", purple, "", "", "") + "[" + second.Binlog_cache_disk_use + "]" + ShowFont(" Select_scan", purple, "", "", "") + "[" + second.Select_scan + "]" + ShowFont(" Select_full_join", purple, "", "", "") + "[" + second.Select_full_join + "]" + ShowFont(" Slow_queries", purple, "", "", "") + "[" + second.Slow_queries + "]\n"
			if second.Rpl_semi_sync_master_status != "" {
				pic += ShowFont(" Rpl_semi_sync_master_status", purple, "", "", "") + "[" + second.Rpl_semi_sync_master_status + "]" + ShowFont(" Rpl_semi_sync_slave_status", purple, "", "", "") + "[" + second.Rpl_semi_sync_slave_status + "]" + ShowFont(" rpl_semi_sync_master_timeout", purple, "", "", "") + "[" + second.rpl_semi_sync_master_timeout + "]\n"
			}
			if second.Master_Host != "" {
				pic += ShowFont(" Master_Host", purple, "", "", "") + "[" + second.Master_Host + "]" + ShowFont(" Master_User", purple, "", "", "") + "[" + second.Master_User + "]" + ShowFont(" Master_Port", purple, "", "", "") + "[" + second.Master_Port + "]" + ShowFont(" Master_Server_Id", purple, "", "", "") + "[" + second.Master_Server_Id + "]\n"
				io := ""
				sql := ""
				if second.Slave_IO_Running != "Yes" {
					io = ShowFont("No", red, "", "", "y")
				} else {
					io = ShowFont("Yes", green, "", "", "")
				}
				if second.Slave_SQL_Running != "Yes" {
					sql = ShowFont("No", red, "", "", "y")
				} else {
					sql = ShowFont("Yes", green, "", "", "")
				}
				pic += ShowFont(" Slave_IO_Running", purple, "", "", "") + "[" + io + "]" + ShowFont(" Slave_SQL_Running", purple, "", "", "") + "[" + sql + "]\n"
			}
			pic += ShowFont(" table_open_cache", purple, "", "", "") + "[" + second.var_table_open_cache + "]" + ShowFont(" thread_cache_size", purple, "", "", "") + "[" + second.var_thread_cache_size + "]" + ShowFont(" Opened_tables", purple, "", "", "") + "[" + second.Opened_tables + "]" + ShowFont(" Created_tmp_disk_tables_ratio", purple, "", "", "") + "[" + tmptable + "]\n\n"

			pic += ShowFont(" innodb_adaptive_flushing", purple, "", "", "") + "[" + second.var_innodb_adaptive_flushing + "]" + ShowFont(" innodb_adaptive_hash_index", purple, "", "", "") + "[" + second.var_innodb_adaptive_hash_index + "]" + ShowFont(" innodb_buffer_pool_size", purple, "", "", "") + "[" + ChangeUntils(second.var_innodb_buffer_pool_size) + "]" + "\n"
			pic += ShowFont(" innodb_file_per_table", purple, "", "", "") + "[" + second.var_innodb_file_per_table + "]" + ShowFont(" innodb_flush_log_at_trx_commit", purple, "", "", "") + "[" + second.var_innodb_flush_log_at_trx_commit + "]" + ShowFont(" innodb_flush_method", purple, "", "", "") + "[" + second.var_innodb_flush_method + "]" + "\n"
			pic += ShowFont(" innodb_io_capacity", purple, "", "", "") + "[" + second.var_innodb_io_capacity + "]" + ShowFont(" innodb_lock_wait_timeout", purple, "", "", "") + "[" + second.var_innodb_lock_wait_timeout + "]" + ShowFont(" innodb_log_buffer_size", purple, "", "", "") + "[" + ChangeUntils(second.var_innodb_log_buffer_size) + "]" + "\n"
			pic += ShowFont(" innodb_log_file_size", purple, "", "", "") + "[" + ChangeUntils(second.var_innodb_log_file_size) + "]" + ShowFont(" innodb_log_files_in_group", purple, "", "", "") + "[" + second.var_innodb_log_files_in_group + "]" + ShowFont(" innodb_max_dirty_pages_pct", purple, "", "", "") + "[" + second.var_innodb_max_dirty_pages_pct + "]\n"
			pic += ShowFont(" innodb_read_io_threads", purple, "", "", "") + "[" + second.var_innodb_read_io_threads + "]" + ShowFont(" innodb_thread_concurrency", purple, "", "", "") + "[" + second.var_innodb_thread_concurrency + "]" + "\n"
			pic += ShowFont(" innodb_write_io_threads", purple, "", "", "") + "[" + second.var_innodb_write_io_threads + "]" + "\n"
		}
	}

	//必打时间time信息
	if flag_info["time"] == true {
		title_summit = ShowFont("--------|", dgreen, "", "", "")
		title_detail = ShowFont("  time  |", dgreen, "", "y", "")
		data_detail = ShowFont(tool.GetNowTime(), yellow, "", "", "") + ShowFont("|", dgreen, "", "", "")
	}

	//loadavg 信息
	if flag_info["load"] == true {
		title_summit += ShowFont("---load----avg---|", dgreen, "", "", "")
		title_detail += ShowFont("  1m    5m   15m |", dgreen, "", "y", "")
		// fmt.Println(strings.Repeat(" ", 5-len(tool.FloatToString(first.load_1, 2)))+tool.FloatToString(first.load_1, 2), tool.FloatToString(first.load_1, 2), len(tool.FloatToString(first.load_1, 2)))
		//load 1 min
		if second.load_1 > second.Cpu_core {
			if second.load_1 >= 10.0 {
				data_detail += ShowFont(strings.Repeat(" ", 5-len(tool.FloatToString(second.load_1, 2)))+tool.FloatToString(second.load_1, 2), red, "", "", "y")
			} else {
				data_detail += ShowFont(strings.Repeat(" ", 5-len(tool.FloatToString(second.load_1, 2)))+tool.FloatToString(second.load_1, 2), red, "", "", "y")
			}
		} else {
			data_detail += ShowFont(strings.Repeat(" ", 5-len(tool.FloatToString(second.load_1, 2)))+tool.FloatToString(second.load_1, 2), "", "", "", "")
		}

		if second.load_5 > second.Cpu_core {
			if second.load_1 >= 10.0 {
				data_detail += ShowFont(strings.Repeat(" ", 6-len(tool.FloatToString(second.load_5, 2)))+tool.FloatToString(second.load_5, 2), red, "", "", "y")
			} else {
				data_detail += ShowFont(strings.Repeat(" ", 6-len(tool.FloatToString(second.load_5, 2)))+tool.FloatToString(second.load_5, 2), red, "", "", "y")
			}
		} else {
			data_detail += ShowFont(strings.Repeat(" ", 6-len(tool.FloatToString(second.load_5, 2)))+tool.FloatToString(second.load_5, 2), "", "", "", "")
		}

		if second.load_15 > second.Cpu_core {
			if second.load_1 >= 10.0 {
				data_detail += ShowFont(strings.Repeat(" ", 6-len(tool.FloatToString(second.load_15, 2)))+tool.FloatToString(second.load_15, 2), red, "", "", "y") + ShowFont("|", dgreen, "", "", "")
			} else {
				data_detail += ShowFont(strings.Repeat(" ", 6-len(tool.FloatToString(second.load_15, 2)))+tool.FloatToString(second.load_15, 2), red, "", "", "y") + ShowFont("|", dgreen, "", "", "")
			}
		} else {
			data_detail += ShowFont(strings.Repeat(" ", 6-len(tool.FloatToString(second.load_15, 2)))+tool.FloatToString(second.load_15, 2), "", "", "", "") + ShowFont("|", dgreen, "", "", "")
		}

	}

	//cpu-usage
	if flag_info["cpu"] == true {
		title_summit += ShowFont("---cpu-usage%--|", dgreen, "", "", "")
		title_detail += ShowFont("usr sys idl iow|", dgreen, "", "y", "")

		cpu_total1 := first.cpu_usr + first.cpu_nice + first.cpu_sys + first.cpu_idl + first.cpu_iow + first.cpu_irq + first.cpu_softirq
		cpu_total2 := second.cpu_usr + second.cpu_nice + second.cpu_sys + second.cpu_idl + second.cpu_iow + second.cpu_irq + second.cpu_softirq

		usr := (second.cpu_usr - first.cpu_usr) * 100 / (cpu_total2 - cpu_total1)
		sys := (second.cpu_sys - first.cpu_sys) * 100 / (cpu_total2 - cpu_total1)
		idl := (second.cpu_idl - first.cpu_idl) * 100 / (cpu_total2 - cpu_total1)
		iow := (second.cpu_iow - first.cpu_iow) * 100 / (cpu_total2 - cpu_total1)
		//usr
		if usr > 10 {
			data_detail += ShowFont(strings.Repeat(" ", 3-len(strconv.Itoa(usr)))+strconv.Itoa(usr)+" ", red, "", "", "y")
		} else {
			data_detail += ShowFont(strings.Repeat(" ", 3-len(strconv.Itoa(usr)))+strconv.Itoa(usr)+" ", green, "", "", "")
		}

		if sys > 10 {
			data_detail += ShowFont(strings.Repeat(" ", 3-len(strconv.Itoa(sys)))+strconv.Itoa(sys)+" ", red, "", "", "y")
		} else {
			data_detail += ShowFont(strings.Repeat(" ", 3-len(strconv.Itoa(sys)))+strconv.Itoa(sys)+" ", "", "", "", "")
		}

		if 1 != 1 {
			data_detail += ShowFont(strings.Repeat(" ", 3-len(strconv.Itoa(idl)))+strconv.Itoa(idl)+" ", red, "", "", "y")
		} else {
			data_detail += ShowFont(strings.Repeat(" ", 3-len(strconv.Itoa(idl)))+strconv.Itoa(idl)+" ", "", "", "", "")
		}

		if iow > 10 {
			data_detail += ShowFont(strings.Repeat(" ", 3-len(strconv.Itoa(iow)))+strconv.Itoa(iow), red, "", "", "y")
		} else {
			data_detail += ShowFont(strings.Repeat(" ", 3-len(strconv.Itoa(iow)))+strconv.Itoa(iow), green, "", "", "")
		}
		data_detail += ShowFont("|", dgreen, "", "", "")
		if flag_info["count"] != 0 {
			if count == flag_info["count"] {
				usrpercent += strconv.Itoa(usr)
				syspercent += strconv.Itoa(sys)
				idlepercent += strconv.Itoa(idl)
				iowaitpercent += strconv.Itoa(iow)
				plot.PlotCpu("%CPU", "times", "%Percent", tool.String2float(usrpercent), tool.String2float(syspercent), tool.String2float(idlepercent), tool.String2float(iowaitpercent), count)
			} else {
				usrpercent += strconv.Itoa(usr) + ","
				syspercent += strconv.Itoa(sys) + ","
				idlepercent += strconv.Itoa(idl) + ","
				iowaitpercent += strconv.Itoa(iow) + ","
			}
		}
	}

	//swap
	if flag_info["swap"] == true {
		title_summit += ShowFont("---swap---|", dgreen, "", "", "")
		title_detail += ShowFont("  si   so |", dgreen, "", "y", "")
		if flag_info["interval"] == "1" && count == 0 {
			data_detail += "00" + ShowFont("|", dgreen, "", "", "y")
		} else if flag_info["interval"] == "1" && count > 0 {
			si := second.swap_in - first.swap_in
			so := second.swap_out - first.swap_out
			// fmt.Println(second.swap_in, first.swap_in, si, second.swap_out, first.swap_out, so)
			si_string := strconv.Itoa(si)
			so_string := strconv.Itoa(so)

			in := strings.Repeat(" ", 5-len(si_string)) + si_string
			out := strings.Repeat(" ", 5-len(so_string)) + so_string
			if si > 0 {
				data_detail += ShowFont(in, red, "", "", "y")
			} else {
				data_detail += ShowFont(in, "", "", "", "")
			}

			if so > 0 {
				data_detail += ShowFont(out, red, "", "", "y")
			} else {
				data_detail += ShowFont(out, "", "", "", "")
			}

			data_detail += ShowFont("|", dgreen, "", "", "")
		}
	}

	//net
	//swap
	if flag_info["net"] != "none" {
		title_summit += ShowFont("----net(B)---- ", dgreen, "", "", "")
		title_detail += ShowFont("  recv  send  |", dgreen, "", "y", "")
		if flag_info["interval"] == "1" && count == 0 {
			data_detail += " 0 0" + ShowFont("|", dgreen, "", "", "y")
		} else if flag_info["interval"] == "1" && count > 0 {
			net_in := float64((second.net_recv-first.net_recv) / interval)
			net_out := float64((second.net_send-first.net_send) / interval)

			if net_in/1024/1024 >= 1.0 {
				data_detail += ShowFont(strings.Repeat(" ", 6-len(tool.FloatToString(net_in/1024/1024, 1)))+tool.FloatToString(net_in/1024/1024, 1)+"m", red, "", "", "y")
			} else if net_in/1024 < 1.0 {
				data_detail += ShowFont(strings.Repeat(" ", 7-len(strconv.Itoa(int(net_in))))+strconv.Itoa(int(net_in)), "", "", "", "")
			} else if net_in/1024/1024 < 1.0 && net_in/1024 >= 1.0 {
				data_detail += ShowFont(strings.Repeat(" ", 6-len(strconv.Itoa(int(net_in)/1024)))+strconv.Itoa(int(net_in)/1024)+"k", "", "", "", "")
			}

			if net_out/1024/1024 >= 1.0 {
				data_detail += ShowFont(strings.Repeat(" ", 6-len(tool.FloatToString(float64(net_out)/1024/1024, 1)))+tool.FloatToString(float64(net_out)/1024/1024, 1)+"m", red, "", "", "y")
			} else if net_out/1024 < 1.0 {
				data_detail += ShowFont(strings.Repeat(" ", 7-len(strconv.Itoa(int(net_out))))+strconv.Itoa(int(net_out)), "", "", "", "")
			} else if net_out/1024/1024 < 1.0 && net_out/1024 >= 1.0 {
				data_detail += ShowFont(strings.Repeat(" ", 6-len(strconv.Itoa(int(net_out)/1024)))+strconv.Itoa(int(net_out)/1024)+"k", "", "", "", "")
			}

			data_detail += ShowFont("|", dgreen, "", "", "")
		}
	}

	//disk
	//r/s:每秒读次数
	//w/s:每秒写次数
	//rkB/s:每秒读kb
	//wkb/s:每秒写kb
	//queue:disk请求队列长度
	//svctm:平均每次IO请求的处理时间，即磁盘读或写操作执行的时间，包括寻道，旋转时延，和数据传输等时间
	//await:平均每次IO请求的等待时间，单位毫秒
	//%util:设备的利用率
	if flag_info["disk"] != "none" {
		title_summit += ShowFont("--------------------------io-usage-------------------------|", dgreen, "", "", "")
		title_detail += ShowFont("  r/s   w/s   rkB/s   wkB/s   queue    await   svctm  "+"%"+"util|", dgreen, "", "y", "")
		if count == 0 {
			data_detail += ShowFont("0.0 0.0 0.0 0.0 0.0 0.0 0.0 0.0|", "", "", "", "")
		} else {
			// fmt.Printf("rs_disk is float64(%d-%d)/0.999\n", second.io_1, first.io_1)
			rs_disk := float64((second.io_1-first.io_1) / interval)
			// fmt.Printf("ws_disk is float64(%d-%d)/0.999\n", second.io_5, first.io_5)
			ws_disk := float64((second.io_5-first.io_5) / interval)

			// fmt.Printf("rkbs_disk is float64(%d-%d)/1.999\n", second.io_3, first.io_3)
			rkbs_disk := float64((second.io_3-first.io_3) / 2 / interval)
			// fmt.Printf("wkbs_disk is float64(%d-%d)/1.999\n", second.io_7, first.io_7)
			wkbs_disk := float64((second.io_7-first.io_7) / 2 / interval)

			queue_disk := strconv.Itoa(second.io_9)

			var await_disk float64
			var svctm_disk float64
			if (rs_disk + ws_disk) == 0.0 {
				await_disk = float64(second.io_4+second.io_8-first.io_4-first.io_8) / (rs_disk + ws_disk + 1)
				svctm_disk = float64(second.io_10-first.io_10) / (rs_disk + ws_disk + 1)
			} else {
				await_disk = float64(second.io_4+second.io_8-first.io_4-first.io_8) / (rs_disk + ws_disk)
				svctm_disk = float64(second.io_10-first.io_10) / (rs_disk + ws_disk)
			}

			util_disk := float64((second.io_10-first.io_10) *100 / 1000 / interval)
			//fmt.Printf("util:%f\n",util_disk)
			//usr
			// fmt.Println(rs_disk, ws_disk, rkbs_disk, wkbs_disk, queue_disk, await_disk, svctm_disk, util_disk)
			// fmt.Println(strings.Repeat(" ", 6-len(tool.FloatToString(rs_disk, 1))) + tool.FloatToString(rs_disk, 1))

			data_detail += ShowFont(strings.Repeat(" ", 5-len(tool.FloatToString(rs_disk, 1)))+tool.FloatToString(rs_disk, 1), "", "", "", "")

			if 1 != 1 {
				data_detail += ShowFont(strings.Repeat(" ", 7-len(tool.FloatToString(ws_disk, 1)))+tool.FloatToString(ws_disk, 1), red, "", "", "y")
			} else {
				data_detail += ShowFont(strings.Repeat(" ", 7-len(tool.FloatToString(ws_disk, 1)))+tool.FloatToString(ws_disk, 1), "", "", "", "")
			}

			if rkbs_disk > 1024.0 {
				data_detail += ShowFont(strings.Repeat(" ", 7-len(tool.FloatToString(rkbs_disk, 1)))+tool.FloatToString(rkbs_disk, 1), red, "", "", "y")
			} else {
				data_detail += ShowFont(strings.Repeat(" ", 7-len(tool.FloatToString(rkbs_disk, 1)))+tool.FloatToString(rkbs_disk, 1), "", "", "", "")
			}

			if wkbs_disk > 1024.0 {
				data_detail += ShowFont(strings.Repeat(" ", 9-len(tool.FloatToString(wkbs_disk, 1)))+tool.FloatToString(wkbs_disk, 1), red, "", "", "y")
			} else {
				data_detail += ShowFont(strings.Repeat(" ", 9-len(tool.FloatToString(wkbs_disk, 1)))+tool.FloatToString(wkbs_disk, 1), "", "", "", "")
			}

			if second.io_9 > 10 {
				data_detail += ShowFont(strings.Repeat(" ", 5-len(queue_disk))+queue_disk+".0 ", red, "", "", "y")
			} else {
				data_detail += ShowFont(strings.Repeat(" ", 5-len(queue_disk))+queue_disk+".0 ", "", "", "", "")
			}

			if await_disk > 5.0 {
				data_detail += ShowFont(strings.Repeat(" ", 8-len(tool.FloatToString(await_disk, 1)))+tool.FloatToString(await_disk, 1), red, "", "", "y")
			} else {
				data_detail += ShowFont(strings.Repeat(" ", 8-len(tool.FloatToString(await_disk, 1)))+tool.FloatToString(await_disk, 1), green, "", "", "")
			}

			if svctm_disk > 5.0 {
				data_detail += ShowFont(strings.Repeat(" ", 8-len(tool.FloatToString(svctm_disk, 1)))+tool.FloatToString(svctm_disk, 1), red, "", "", "y")
			} else {
				data_detail += ShowFont(strings.Repeat(" ", 8-len(tool.FloatToString(svctm_disk, 1)))+tool.FloatToString(svctm_disk, 1), "", "", "", "")
			}

			if util_disk > 80.0 && util_disk < 100.0 {
				data_detail += ShowFont(strings.Repeat(" ", 7-len(tool.FloatToString(util_disk, 1)))+tool.FloatToString(util_disk, 1), red, "", "", "")
				iobusy = append(iobusy,util_disk)
			}
			if util_disk >= 100.0 {
				data_detail += ShowFont(strings.Repeat(" ",2)+"100.0", red, "", "", "y")
				iobusy = append(iobusy,100.0)
			}
			if util_disk <= 80.0 {
				data_detail += ShowFont(strings.Repeat(" ", 7-len(tool.FloatToString(util_disk, 1)))+tool.FloatToString(util_disk, 1), green, "", "", "")
				iobusy = append(iobusy,util_disk)
			}

			data_detail += ShowFont("|", dgreen, "", "", "")

			if count == flag_info["count"] {
				plot.PlotSingle("%IOBusy", "times", "%Percent", iobusy, count)
			}
		}

	}

	//Com_*
	if flag_info["com"] == true && mysql_switch == "on" {
		title_summit += ShowFont("-------QPS--------------TPS------|", green, blue, "", "")
		title_detail += ShowFont("  ins  upd   del  sel   qps   tps|", green, "", "y", "")
		if count == 0 {
			data_detail += ShowFont("0 0 0  0 0", "", "", "", "") + ShowFont("|", green, "", "", "")
		} else {
			insert_diff := (second.Com_insert - first.Com_insert) / interval
			update_diff := (second.Com_update - first.Com_update) / interval
			delete_diff := (second.Com_delete - first.Com_delete) / interval
			select_diff := (second.Com_select - first.Com_select) / interval
			commit_diff := (second.Com_commit - first.Com_commit) / interval
			rollback_diff := (second.Com_rollback - first.Com_rollback) / interval
			tps := rollback_diff + commit_diff
			qps := (second.Question - first.Question) / interval
			data_detail += ShowFont(strings.Repeat(" ", 4-len(strconv.Itoa(insert_diff)))+strconv.Itoa(insert_diff), "", "", "", "")
			data_detail += ShowFont(strings.Repeat(" ", 5-len(strconv.Itoa(update_diff)))+strconv.Itoa(update_diff), "", "", "", "")
			data_detail += ShowFont(strings.Repeat(" ", 6-len(strconv.Itoa(delete_diff)))+strconv.Itoa(delete_diff), "", "", "", "")
			data_detail += ShowFont(strings.Repeat(" ", 6-len(strconv.Itoa(select_diff)))+strconv.Itoa(select_diff), "", "", "", "")
			data_detail += ShowFont(strings.Repeat(" ", 6-len(strconv.Itoa(qps)))+strconv.Itoa(qps), yellow, "", "", "")
			data_detail += ShowFont(strings.Repeat(" ", 6-len(strconv.Itoa(tps)))+strconv.Itoa(tps), yellow, "", "", "")
			data_detail += ShowFont("|", green, "", "", "")
			if flag_info["count"] != 0 {
				if count == flag_info["count"] {
					queryps += strconv.Itoa(qps)
					plot.PlotSingle("QPS", "times", "Querys", tool.String2float(queryps), count)
					tranps += strconv.Itoa(tps)
					plot.PlotSingle("TPS", "times", "Transactions", tool.String2float(tranps), count)
				} else {
					queryps += strconv.Itoa(qps) + ","
					tranps += strconv.Itoa(tps) + ","
				}
			}
		}
	}

	//hit   	
	if flag_info["hit"] == true && mysql_switch == "on" {
		title_summit += ShowFont("----indexhit----innodbhit----(%)|", green, blue, "", "")
		title_detail += ShowFont(" idxhitcur  idxhitall  bufferhit|", green, "", "y", "")
		if count == 0 {
			data_detail += ShowFont("100.00 100.00 100.00|", "", "", "", "") + ShowFont("|", green, "", "", "")
		} else {
			read_request := (second.Innodb_buffer_pool_read_requests - first.Innodb_buffer_pool_read_requests) / interval
			read := (second.Innodb_buffer_pool_reads - first.Innodb_buffer_pool_reads) / interval
			//innodb hit
			hrr := (second.Handler_read_rnd - first.Handler_read_rnd) / interval
			hrrn := (second.Handler_read_rnd_next - first.Handler_read_rnd_next) / interval
			hrf := (second.Handler_read_first - first.Handler_read_first) / interval
			hrk := (second.Handler_read_key - first.Handler_read_key) / interval
			hrn := (second.Handler_read_next - first.Handler_read_next) / interval
			hrp := (second.Handler_read_prev - first.Handler_read_prev) / interval
			index_total_hit := (100 - (100*(float64(second.Handler_read_rnd+second.Handler_read_rnd_next) + 0.0001) / (0.0001 + float64(second.Handler_read_first+second.Handler_read_key+second.Handler_read_next+second.Handler_read_prev+second.Handler_read_rnd+second.Handler_read_rnd_next))))
			index_current_hit := 100.00
			if hrr+hrrn != 0 {
				index_current_hit = (100 - (100 * (float64(hrr+hrrn) + 0.0001) / (0.0001 + float64(hrf+hrk+hrn+hrp+hrr+hrrn))))
			}
			//remove qcache
			/*
			query_hits_s := (second.Qcache_hits - first.Qcache_hits) / interval
			com_select_s := (second.Com_select - first.Com_select) / interval

			query_hit := (float64(query_hits_s) + 0.0001) / (float64(query_hits_s+com_select_s) + 0.0001) * 100
			*/
			//innodb buffer pool hit
			innodb_hit := ((float64(read_request-read) + 0.0001) / (float64(read_request) + 0.0001)) * 100

			data_detail += Hit(9, index_current_hit)
			data_detail += Hit(11, index_total_hit)
			// lor = read_request
			//data_detail += ShowFont(strings.Repeat(" ", 8-len(strconv.Itoa(read_request)))+strconv.Itoa(read_request), "", "", "", "")
			data_detail += Hit(12, innodb_hit)
			data_detail += ShowFont("|", green, "", "", "")
		}
	}

	//innodb_rows
	if flag_info["innodb_rows"] == true && mysql_switch == "on" {
		title_summit += ShowFont("---innodb rows status---|", green, blue, "", "")
		title_detail += ShowFont("  ins   upd   del   read|", green, "", "y", "")
		if count == 0 {
			data_detail += ShowFont("0 0 0 0", "", "", "", "") + ShowFont("|", green, "", "", "")
		} else {
			innodb_rows_inserted_diff := (second.Innodb_rows_inserted - first.Innodb_rows_inserted) / interval
			innodb_rows_updated_diff := (second.Innodb_rows_updated - first.Innodb_rows_updated) / interval
			innodb_rows_deleted_diff := (second.Innodb_rows_deleted - first.Innodb_rows_deleted) / interval
			innodb_rows_read_diff := (second.Innodb_rows_read - first.Innodb_rows_read) / interval

			data_detail += ShowFont(strings.Repeat(" ", 5-len(strconv.Itoa(innodb_rows_inserted_diff)))+strconv.Itoa(innodb_rows_inserted_diff), "", "", "", "")
			data_detail += ShowFont(strings.Repeat(" ", 6-len(strconv.Itoa(innodb_rows_updated_diff)))+strconv.Itoa(innodb_rows_updated_diff), "", "", "", "")
			data_detail += ShowFont(strings.Repeat(" ", 6-len(strconv.Itoa(innodb_rows_deleted_diff)))+strconv.Itoa(innodb_rows_deleted_diff), "", "", "", "")
			data_detail += ShowFont(strings.Repeat(" ", 7-len(strconv.Itoa(innodb_rows_read_diff)))+strconv.Itoa(innodb_rows_read_diff), "", "", "", "")

			data_detail += ShowFont("|", green, "", "", "")
		}
	}

	//innodb_pages
	if flag_info["innodb_pages"] == true && mysql_switch == "on" {
		title_summit += ShowFont("---innodb bp pages status--|", green, blue, "", "")
		title_detail += ShowFont("  data   free  dirty  flush|", green, "", "y", "")
		if count == 0 {
			data_detail += ShowFont(" 0 0 0 0", "", "", "", "") + ShowFont("|", green, "", "", "")
		} else {
			flush := (second.Innodb_buffer_pool_pages_flushed - first.Innodb_buffer_pool_pages_flushed) / interval

			data_detail += ShowFont(strings.Repeat(" ", 7-len(strconv.Itoa(second.Innodb_buffer_pool_pages_data)))+strconv.Itoa(second.Innodb_buffer_pool_pages_data), "", "", "", "")
			data_detail += ShowFont(strings.Repeat(" ", 7-len(strconv.Itoa(second.Innodb_buffer_pool_pages_free)))+strconv.Itoa(second.Innodb_buffer_pool_pages_free), "", "", "", "")
			data_detail += ShowFont(strings.Repeat(" ", 7-len(strconv.Itoa(second.Innodb_buffer_pool_pages_dirty)))+strconv.Itoa(second.Innodb_buffer_pool_pages_dirty), yellow, "", "", "")
			data_detail += ShowFont(strings.Repeat(" ", 6-len(strconv.Itoa(flush)))+strconv.Itoa(flush), yellow, "", "", "")

			data_detail += ShowFont("|", green, "", "", "")
		}
	}

	//innodb_data
	if flag_info["innodb_data"] == true && mysql_switch == "on" {
		title_summit += ShowFont("-----innodb data status------|", green, blue, "", "")
		title_detail += ShowFont(" reads  writes  read  written|", green, "", "y", "")
		if count == 0 {
			data_detail += ShowFont(" 0 0 0 0", "", "", "", "") + ShowFont("|", green, "", "", "")
		} else {
			innodb_data_reads_diff := (second.Innodb_data_reads - first.Innodb_data_reads) / interval
			innodb_data_writes_diff := (second.Innodb_data_writes - first.Innodb_data_writes) / interval
			innodb_data_read_diff := (second.Innodb_data_read - first.Innodb_data_read) / interval
			innodb_data_written_diff := (second.Innodb_data_written - first.Innodb_data_written) / interval

			data_detail += ShowFont(strings.Repeat(" ", 6-len(strconv.Itoa(innodb_data_reads_diff)))+strconv.Itoa(innodb_data_reads_diff), "", "", "", "")
			data_detail += ShowFont(strings.Repeat(" ", 7-len(strconv.Itoa(innodb_data_writes_diff)))+strconv.Itoa(innodb_data_writes_diff), "", "", "", "")

			if innodb_data_read_diff/1024/1024 > 9 {
				data_detail += ShowFont(strings.Repeat(" ", 6-len(tool.FloatToString(float64(innodb_data_read_diff)/1024/1024, 1)))+tool.FloatToString(float64(innodb_data_read_diff)/1024/1024, 1)+"m", red, "", "", "y")
			} else if innodb_data_read_diff/1024/1024 <= 9 && innodb_data_read_diff/1024/1024 >= 1 {
				data_detail += ShowFont(strings.Repeat(" ", 6-len(tool.FloatToString(float64(innodb_data_read_diff)/1024/1024, 1)))+tool.FloatToString(float64(innodb_data_read_diff)/1024/1024, 1)+"m", "", "", "", "")
			} else if innodb_data_read_diff/1024 >= 1 && innodb_data_read_diff/1024/1024 < 1 {
				data_detail += ShowFont(strings.Repeat(" ", 6-len(strconv.Itoa(innodb_data_read_diff/1024)))+strconv.Itoa(innodb_data_read_diff/1024)+"k", "", "", "", "")
			} else if innodb_data_read_diff/1024 < 1 {
				data_detail += ShowFont(strings.Repeat(" ", 7-len(strconv.Itoa(innodb_data_read_diff)))+strconv.Itoa(innodb_data_read_diff), "", "", "", "")
			}

			if innodb_data_written_diff/1024/1024 > 9 {
				data_detail += ShowFont(strings.Repeat(" ", 8-len(tool.FloatToString(float64(innodb_data_written_diff)/1024/1024, 1)))+tool.FloatToString(float64(innodb_data_written_diff)/1024/1024, 1)+"m", red, "", "", "y")
			} else if innodb_data_written_diff/1024/1024 <= 9 && innodb_data_written_diff/1024/1024 >= 1 {
				data_detail += ShowFont(strings.Repeat(" ", 8-len(tool.FloatToString(float64(innodb_data_written_diff)/1024/1024, 1)))+tool.FloatToString(float64(innodb_data_written_diff)/1024/1024, 1)+"m", "", "", "", "")
			} else if innodb_data_written_diff/1024 >= 1 && innodb_data_written_diff/1024/1024 < 1 {
				data_detail += ShowFont(strings.Repeat(" ", 8-len(strconv.Itoa(innodb_data_written_diff/1024)))+strconv.Itoa(innodb_data_written_diff/1024)+"k", "", "", "", "")
			} else if innodb_data_written_diff/1024 < 1 {
				data_detail += ShowFont(strings.Repeat(" ", 9-len(strconv.Itoa(innodb_data_written_diff)))+strconv.Itoa(innodb_data_written_diff), "", "", "", "")
			}

			data_detail += ShowFont("|", green, "", "", "")
		}
	}

	//innodb_log
	if flag_info["innodb_log"] == true && mysql_switch == "on" {
		title_summit += ShowFont("--innodb log--|", green, blue, "", "")
		title_detail += ShowFont("fsyncs written|", green, "", "y", "")
		if count == 0 {
			data_detail += ShowFont(" 0 0", "", "", "", "") + ShowFont("|", green, "", "", "")
		} else {

			innodb_os_log_fsyncs_diff := (second.Innodb_os_log_fsyncs - first.Innodb_os_log_fsyncs) / interval
			innodb_os_log_written_diff := (second.Innodb_os_log_written - first.Innodb_os_log_written) / interval

			data_detail += ShowFont(strings.Repeat(" ", 4-len(strconv.Itoa(innodb_os_log_fsyncs_diff)))+strconv.Itoa(innodb_os_log_fsyncs_diff), "", "", "", "")

			if innodb_os_log_written_diff/1024/1024 >= 1 {
				data_detail += ShowFont(strings.Repeat(" ", 8-len(tool.FloatToString(float64(innodb_os_log_written_diff)/1024/1024, 1)))+tool.FloatToString(float64(innodb_os_log_written_diff)/1024/1024, 1)+"m", red, "", "", "y")
			} else if innodb_os_log_written_diff/1024/1024 < 1 && innodb_os_log_written_diff/1024 >= 1 && innodb_os_log_written_diff/1024 < 100 {
				data_detail += ShowFont(strings.Repeat(" ", 8-len(strconv.Itoa(int(float64(innodb_os_log_written_diff)/1024/1024+0.5))))+strconv.Itoa(int(float64(innodb_os_log_written_diff)/1024))+"k", yellow, "", "", "")
			}else if innodb_os_log_written_diff/1024 > 100 {
				data_detail += ShowFont(strings.Repeat(" ", 7-len(strconv.Itoa(int(float64(innodb_os_log_written_diff)/1024/1024+0.5))))+strconv.Itoa(int(float64(innodb_os_log_written_diff)/1024))+"k", yellow, "", "", "")
			} else if innodb_os_log_written_diff/1024 < 1 {
				data_detail += ShowFont(strings.Repeat(" ", 7-len(strconv.Itoa(innodb_os_log_written_diff)))+strconv.Itoa(innodb_os_log_written_diff)+"B", "", "", "", "")
			}

			data_detail += ShowFont("|", green, "", "", "")
		}
	}

	//innodb_status
	//list:History list length
	//uflush:unflushed_log=log_bytes_written-log_bytes_flushed
	//uckpt:checkpoint_age=log_bytes_written-last_checkpoint
	//
	if flag_info["innodb_status"] == true && mysql_switch == "on" {
		title_summit += ShowFont("---his---log(byte)---read ---query---|", green, blue, "", "")
		title_detail += ShowFont(" list  uflush  uckpt  view inside que|", green, "", "y", "")
		if count == 0 {
			data_detail += ShowFont("0 0 0 0 0 0", "", "", "", "") + ShowFont("|", green, "", "", "")
		} else {

			//mysql --innodb_status show engine innodb status
			//log unflushed = Log sequence number - Log flushed up to
			//uncheckpointed bytes = Log sequence number - Last checkpoint at
			//mysql -e "show engine innodb status\G"|grep -n -E -A4 -B1 "^TRANSACTIONS|LOG|ROW OPERATIONS"
			//mysql -e "show engine innodb status\G"|grep -E "Last checkpoint|read view|queries inside|queue"
			// Log_sequenceint
			// Log_flushed int
			//History_listint
			// Last_checkpoint int
			// Read_view int
			// Query_insideint
			// Query_queue int

			unflushed_log := second.Log_sequence - second.Log_flushed
			checkpoint_age := second.Log_sequence - second.Last_checkpoint
			//History_list
			data_detail += ShowFont(strings.Repeat(" ", 5-len(strconv.Itoa(second.History_list)))+strconv.Itoa(second.History_list), "", "", "", "")
			//unflushed_log
			if unflushed_log/1024/1024 >= 1 {
				data_detail += ShowFont(strings.Repeat(" ", 6-len(tool.FloatToString(float64(unflushed_log)/1024/1024+0.5, 1)))+tool.FloatToString(float64(unflushed_log)/1024/1024+0.5, 1)+"m", yellow, "", "", "")
			} else if unflushed_log/1024/1024 < 1 && unflushed_log/1024 >= 1 {
				data_detail += ShowFont(strings.Repeat(" ", 6-len(strconv.Itoa(int(float64(unflushed_log)/1024+0.5))))+strconv.Itoa(int(float64(unflushed_log)/1024+0.5))+"k", yellow, "", "", "")
			} else if unflushed_log/1024 < 1 {
				data_detail += ShowFont(strings.Repeat(" ", 7-len(strconv.Itoa(unflushed_log)))+strconv.Itoa(unflushed_log), yellow, "", "", "")
			}

			//checkpoint_age
			if checkpoint_age/1024/1024 >= 1 {
				data_detail += ShowFont(strings.Repeat(" ", 6-len(tool.FloatToString(float64(checkpoint_age)/1024/1024+0.5, 1)))+tool.FloatToString(float64(checkpoint_age)/1024/1024+0.5, 1)+"m", yellow, "", "", "")
			} else if checkpoint_age/1024/1024 < 1 && checkpoint_age/1024 >= 1 {
				data_detail += ShowFont(strings.Repeat(" ", 6-len(strconv.Itoa(int(float64(checkpoint_age)/1024+0.5))))+strconv.Itoa(int(float64(checkpoint_age)/1024+0.5))+"k", yellow, "", "", "")
			} else if checkpoint_age/1024 < 1 {
				data_detail += ShowFont(strings.Repeat(" ", 7-len(strconv.Itoa(checkpoint_age)))+strconv.Itoa(checkpoint_age), yellow, "", "", "")
			}

			//Read_views
			data_detail += ShowFont(strings.Repeat(" ", 6-len(strconv.Itoa(second.Read_view)))+strconv.Itoa(second.Read_view), "", "", "", "")
			//inside
			data_detail += ShowFont(strings.Repeat(" ", 6-len(strconv.Itoa(second.Query_inside)))+strconv.Itoa(second.Query_inside), "", "", "", "")
			//queue
			data_detail += ShowFont(strings.Repeat(" ", 6-len(strconv.Itoa(second.Query_queue)))+strconv.Itoa(second.Query_queue), "", "", "", "")

			data_detail += ShowFont("|", green, "", "", "")
		}
	}

	//threads ------threads------
	if flag_info["threads"] == true && mysql_switch == "on" {
		title_summit += ShowFont("----------threads---------|", green, blue, "", "")
		title_detail += ShowFont("  run  con  cre  cac  "+"%"+"hit|", green, "", "y", "")
		if count == 0 {
			data_detail += ShowFont(" 0 0 0 0 0", "", "", "", "") + ShowFont("|", green, "", "", "")
		} else {
			connections_dirr := (second.Connections - first.Connections)

			threads_created_diff := (second.Threads_created - first.Threads_created)

			thread_cache_hit := (1 - float64(threads_created_diff)/float64(connections_dirr)) * 100
			data_detail += ShowFont(strings.Repeat(" ", 4-len(strconv.Itoa(second.Threads_running)))+strconv.Itoa(second.Threads_running), "", "", "", "")

			data_detail += ShowFont(strings.Repeat(" ", 5-len(strconv.Itoa(second.Threads_connected)))+strconv.Itoa(second.Threads_connected), "", "", "", "")

			data_detail += ShowFont(strings.Repeat(" ", 5-len(strconv.Itoa(threads_created_diff)))+strconv.Itoa(threads_created_diff), "", "", "", "")

			data_detail += ShowFont(strings.Repeat(" ", 5-len(strconv.Itoa(second.Threads_cached)))+strconv.Itoa(second.Threads_cached), "", "", "", "")
			if thread_cache_hit > 99.0 {
				data_detail += ShowFont(strings.Repeat(" ", 7-len(tool.FloatToString(thread_cache_hit, 2)))+tool.FloatToString(thread_cache_hit, 2), green, "", "", "")
			} else if thread_cache_hit <= 99.0 && thread_cache_hit > 90.0 {
				data_detail += ShowFont(strings.Repeat(" ", 7-len(tool.FloatToString(thread_cache_hit, 2)))+tool.FloatToString(thread_cache_hit, 2), yellow, "", "", "")
			} else {
				data_detail += ShowFont(strings.Repeat(" ", 7-len(tool.FloatToString(thread_cache_hit, 2)))+tool.FloatToString(thread_cache_hit, 2), red, "", "", "")
			}

			data_detail += ShowFont("|", green, "", "", "")
		}
	}

	//bytes
	if flag_info["bytes"] == true && mysql_switch == "on" {
		title_summit += ShowFont("---bytes---|", green, blue, "", "")
		title_detail += ShowFont(" recv  send|", green, "", "y", "")
		if count == 0 {
			data_detail += ShowFont(" 0 0", "", "", "", "") + ShowFont("|", green, "", "", "")
		} else {

			bytes_received_diff := (second.Bytes_received - first.Bytes_received) / interval
			bytes_sent_diff := (second.Bytes_sent - first.Bytes_sent) / interval

			if bytes_received_diff/1024/1024 >= 1 {
				data_detail += ShowFont(strings.Repeat(" ", 4-len(tool.FloatToString(float64(bytes_received_diff)/1024/1024+0.5, 1)))+tool.FloatToString(float64(bytes_received_diff)/1024/1024+0.5, 1)+"m", red, "", "", "y")
			} else if bytes_received_diff/1024/1024 < 1 && bytes_received_diff/1024 >= 1 {
				data_detail += ShowFont(strings.Repeat(" ", 4-len(strconv.Itoa(int(float64(bytes_received_diff)/1024+0.5))))+strconv.Itoa(int(float64(bytes_received_diff)/1024+0.5))+"k", "", "", "", "")
			} else if bytes_received_diff/1024 < 1 {
				data_detail += ShowFont(strings.Repeat(" ", 5-len(strconv.Itoa(bytes_received_diff)))+strconv.Itoa(bytes_received_diff), "", "", "", "")
			}

			if bytes_sent_diff/1024/1024 >= 1 {
				data_detail += ShowFont(strings.Repeat(" ", 5-len(tool.FloatToString(float64(bytes_sent_diff)/1024/1024+0.5, 1)))+tool.FloatToString(float64(bytes_sent_diff)/1024/1024+0.5, 1)+"m", red, "", "", "y")
			} else if bytes_sent_diff/1024/1024 < 1 && bytes_sent_diff/1024 >= 1 {
				data_detail += ShowFont(strings.Repeat(" ", 5-len(strconv.Itoa(int(float64(bytes_sent_diff)/1024+0.5))))+strconv.Itoa(int(float64(bytes_sent_diff)/1024+0.5))+"k", "", "", "", "")
			} else if bytes_sent_diff/1024 < 1 {
				data_detail += ShowFont(strings.Repeat(" ", 6-len(strconv.Itoa(bytes_sent_diff)))+strconv.Itoa(bytes_sent_diff), "", "", "", "")
			}

			data_detail += ShowFont("|", green, "", "", "")
		}
	}

	//semi
	//tx:master等待事务的平均时间（微秒）
	//notx:slave未成功确认的提交数
	//yestx:slave成功确认的提交数
	//notime:master关闭半同步复制的次数
	if flag_info["semi"] == true && mysql_switch == "on" {
		title_summit += ShowFont("--avg_wait---tx_times--semi|", green, blue, "", "")
		title_detail += ShowFont("   tx   notx  yestx  notime|", green, "", "y", "")
		if count == 0 {
			data_detail += ShowFont(" 100ms 1000 1000 1000", "", "", "", "") + ShowFont("|", green, "", "", "")
		} else {
			// fmt.Printf("1 %d 2 %d 3 %d 4 %d 5 %d", second.Rpl_semi_sync_master_net_avg_wait_time, second.Rpl_semi_sync_master_tx_avg_wait_time, second.Rpl_semi_sync_master_no_tx, second.Rpl_semi_sync_master_yes_tx, second.Rpl_semi_sync_master_no_times)
			/*该参数已废弃
			if second.Rpl_semi_sync_master_net_avg_wait_time < 1000 {
				data_detail += ShowFont(strings.Repeat(" ", 3-len(strconv.Itoa(second.Rpl_semi_sync_master_net_avg_wait_time)))+strconv.Itoa(second.Rpl_semi_sync_master_net_avg_wait_time)+"us", "", "", "", "")
			} else if second.Rpl_semi_sync_master_net_avg_wait_time >= 1000 && second.Rpl_semi_sync_master_net_avg_wait_time/1000/1000 <= 1 {
				data_detail += ShowFont(strings.Repeat(" ", 3-len(strconv.Itoa(second.Rpl_semi_sync_master_net_avg_wait_time/1000)))+strconv.Itoa(second.Rpl_semi_sync_master_net_avg_wait_time/1000)+"ms", "", "", "", "")
			} else if second.Rpl_semi_sync_master_net_avg_wait_time/1000/1000 > 1 {
				data_detail += ShowFont(strings.Repeat(" ", 4-len(strconv.Itoa(second.Rpl_semi_sync_master_net_avg_wait_time/1000/1000)))+strconv.Itoa(second.Rpl_semi_sync_master_net_avg_wait_time/1000/1000)+"s", red, "", "", "y")
			}
			*/
			if second.Rpl_semi_sync_master_tx_avg_wait_time < 1000 {
				data_detail += ShowFont(strings.Repeat(" ", 4-len(strconv.Itoa(second.Rpl_semi_sync_master_tx_avg_wait_time)))+strconv.Itoa(second.Rpl_semi_sync_master_tx_avg_wait_time)+"us", "", "", "", "")
			} else if second.Rpl_semi_sync_master_tx_avg_wait_time > 1000 && second.Rpl_semi_sync_master_tx_avg_wait_time/1000/1000 <= 1 {
				data_detail += ShowFont(strings.Repeat(" ", 4-len(strconv.Itoa(second.Rpl_semi_sync_master_tx_avg_wait_time/1000)))+strconv.Itoa(second.Rpl_semi_sync_master_tx_avg_wait_time/1000)+"ms", "", "", "", "")
			} else if second.Rpl_semi_sync_master_tx_avg_wait_time/1000/1000 > 1 {
				data_detail += ShowFont(strings.Repeat(" ", 5-len(strconv.Itoa(second.Rpl_semi_sync_master_tx_avg_wait_time/1000/1000)))+strconv.Itoa(second.Rpl_semi_sync_master_tx_avg_wait_time/1000/1000)+"s", red, "", "", "y")
			}

			if second.Rpl_semi_sync_master_no_tx > 1 {
				data_detail += ShowFont(strings.Repeat(" ", 5-len(strconv.Itoa(second.Rpl_semi_sync_master_no_tx)))+strconv.Itoa(second.Rpl_semi_sync_master_no_tx), red, "", "", "y")
			} else {
				data_detail += ShowFont(strings.Repeat(" ", 5-len(strconv.Itoa(second.Rpl_semi_sync_master_no_tx)))+strconv.Itoa(second.Rpl_semi_sync_master_no_tx), "", "", "", "y")
			}

			data_detail += ShowFont(strings.Repeat(" ", 6-len(ChangeUntils(second.Rpl_semi_sync_master_yes_tx)))+ChangeUntils(second.Rpl_semi_sync_master_yes_tx), "", "", "", "y")

			if second.Rpl_semi_sync_master_no_times > 1 {
				data_detail += ShowFont(strings.Repeat(" ", 10-len(strconv.Itoa(second.Rpl_semi_sync_master_no_times)))+strconv.Itoa(second.Rpl_semi_sync_master_no_times), red, "", "", "y")
			} else {
				data_detail += ShowFont(strings.Repeat(" ", 10-len(strconv.Itoa(second.Rpl_semi_sync_master_no_times)))+strconv.Itoa(second.Rpl_semi_sync_master_no_times), "", "", "", "y")
			}
			data_detail += ShowFont("|", green, "", "", "")
		}
	}

	//threads ------threads------
	if flag_info["slave"] == true && mysql_switch == "on" {
		title_summit += ShowFont("-------------SlaveStatus-----------|", green, blue, "", "")
		title_detail += ShowFont("  ReadMLP   ExecMLP   chkRE   SecBM|", green, "", "y", "")
		if count == 0 {
			data_detail += ShowFont(" 1066312331 1066312331 6312331 6312331", "", "", "", "") + ShowFont("|", green, "", "", "")
		} else {

			checkNum := second.Read_Master_Log_Pos - second.Exec_Master_Log_Pos

			data_detail += ShowFont(strings.Repeat(" ", 6-len(strconv.Itoa(second.Read_Master_Log_Pos)))+strconv.Itoa(second.Read_Master_Log_Pos), "", "", "", "")

			data_detail += ShowFont(strings.Repeat(" ", 11-len(strconv.Itoa(second.Exec_Master_Log_Pos)))+strconv.Itoa(second.Exec_Master_Log_Pos), "", "", "", "")

			data_detail += ShowFont(strings.Repeat(" ", 10-len(strconv.Itoa(checkNum)))+strconv.Itoa(checkNum), "", "", "", "")

			if second.Seconds_Behind_Master > 300 {
				data_detail += ShowFont(strings.Repeat(" ", 8-len(strconv.Itoa(second.Seconds_Behind_Master)))+strconv.Itoa(second.Seconds_Behind_Master), red, "", "", "")
			} else {
				data_detail += ShowFont(strings.Repeat(" ", 8-len(strconv.Itoa(second.Seconds_Behind_Master)))+strconv.Itoa(second.Seconds_Behind_Master), green, "", "", "")
			}

			data_detail += ShowFont("|", green, "", "", "")
		}
	}

	//tcprstat
	//count:完成的 request 个数
	//avg:平均 response 时间(微秒)
	//95-avg:95% response时间(微秒)
	//99-avg:99% response时间(微秒)
	if flag_info["rt"] == true {
		// 1,000 皮秒 = 1纳秒
		// 1,000,000 皮秒 = 1微秒
		// 1,000,000,000 皮秒 = 1毫秒
		// 1,000,000,000,000 皮秒 = 1秒
		var rt_count, rt_avg, rt_95avg, rt_99avg string
		title_summit += ShowFont("--------tcprstat(us)--------|", green, blue, "", "") + " "
		title_detail += ShowFont("  count  avg  95-avg  99-avg|", green, "", "y", "")
		if count == 0 {
			data_detail = ShowFont(" 0 0 0 0", "", "", "", "") + ShowFont("|", green, "", "", "")
		} else {
			if second.rt_count > 1000 {
				rt_count = ShowFont(strings.Repeat(" ", 7-len(strconv.Itoa(second.rt_count)))+strconv.Itoa(second.rt_count), red, "", "", "")
			} else {
				rt_count = ShowFont(strings.Repeat(" ", 7-len(strconv.Itoa(second.rt_count)))+strconv.Itoa(second.rt_count), "", "", "", "")
			}

			if second.rt_avg/1000 > 60 {
				rt_avg = ShowFont(strings.Repeat(" ", 7-len(strconv.Itoa(second.rt_avg)))+strconv.Itoa(second.rt_avg), red, "", "", "")
			} else {
				rt_avg = ShowFont(strings.Repeat(" ", 7-len(strconv.Itoa(second.rt_avg)))+strconv.Itoa(second.rt_avg), green, "", "", "")
			}

			if second.rt_a5/1000 > 100 && second.rt_a5 != 0 {
				rt_95avg = ShowFont(strings.Repeat(" ", 7-len(strconv.Itoa(second.rt_a5)))+strconv.Itoa(second.rt_a5), red, "", "", "")
			} else {
				rt_95avg = ShowFont(strings.Repeat(" ", 7-len(strconv.Itoa(second.rt_a5)))+strconv.Itoa(second.rt_a5), green, "", "", "")
			}

			if second.rt_a9/1000 == 100 && second.rt_a9 != 0 {
				rt_99avg = ShowFont(strings.Repeat(" ", 7-len(strconv.Itoa(second.rt_a9)))+strconv.Itoa(second.rt_a9), red, "", "", "")
			} else {
				rt_99avg = ShowFont(strings.Repeat(" ", 7-len(strconv.Itoa(second.rt_a9)))+strconv.Itoa(second.rt_a9), green, "", "", "")
			}

			data_detail += rt_count + rt_avg + rt_95avg + rt_99avg + ShowFont("|", green, "", "", "")
		}

	}
	//展示标题
	if count == 0 {
		fmt.Println(pic)
		fmt.Println(title_summit)
		fmt.Println(title_detail)
		Add_log(flag_info, pic)
		Add_log(flag_info, title_summit)
		Add_log(flag_info, title_detail)
	}
	//每隔20行打印一次标题
	if count != 0 && count%20 == 0 {
		fmt.Println(title_summit)
		fmt.Println(title_detail)
		Add_log(flag_info, title_summit)
		Add_log(flag_info, title_detail)
	}
	if count != 0 && count%20 != 0 {
		fmt.Println(data_detail)
		Add_log(flag_info, data_detail)

	}

}


func Hit(num int, in float64) string {
	var result string
	if in > 99.0 {
		result = ShowFont(strings.Repeat(" ", num-len(tool.FloatToString(in, 2)))+tool.FloatToString(in, 2), green, "", "", "")
	} else if in > 90.0 && in <= 99.0 {
		result = ShowFont(strings.Repeat(" ", num-len(tool.FloatToString(in, 2)))+tool.FloatToString(in, 2), yellow, "", "", "")
	} else if in < 0.01 {
		result = ShowFont(strings.Repeat(" ", num-len("0.00"))+"0.00", red, "", "", "")
	} else {
		result = ShowFont(strings.Repeat(" ", num-len(tool.FloatToString(in, 2)))+tool.FloatToString(in, 2), red, "", "", "y")
	}
	return result
}


func ChangeUntils(in int) string {
	var result string
	if in/1024 < 1 {
		tmp := strconv.Itoa(in)
		result = tmp
	} else if in/1024 >= 1 && in/1024/1024 < 1 {
		tmp := strconv.Itoa(in / 1024)
		result = tmp + "k"
	} else if in/1024/1024 >= 1 && in/1024/1024/1024 < 1 {
		tmp := strconv.Itoa(in / 1024 / 1024)
		result = tmp + "m"
	} else if in/1024/1024/1024 >= 1 && in/1024/1024/1024/1024 < 1 {
		tmp := strconv.Itoa(in / 1024 / 1024 / 1024)
		result = tmp + "g"
	} else if in/1024/1024/1024/1024 >= 1 {
		tmp := strconv.Itoa(in / 1024 / 1024 / 1024 / 1024)
		result = tmp + "pg"
	}
	return result
}

//文字字体设置函数，可设置字体大小，亮度，颜色
func ShowFont(text string, status string, background string, underline string, highshow string) string {
	out_one := "\033["
	out_two := ""
	out_three := ""
	out_four := ""
	//可动态配置字体颜色 背景色 高亮
	// 显示：0(默认)、1(粗体/高亮)、22(非粗体)、4(单条下划线)、24(无下划线)、5(闪烁)、25(无闪烁)、7(反显、翻转前景色和背景色)、27(无反显)
	// 颜色：0(黑)、1(红)、2(绿)、 3(黄)、4(蓝)、5(洋红)、6(青)、7(白)
	// 前景色为30+颜色值，如31表示前景色为红色；背景色为40+颜色值，如41表示背景色为红色。
	if underline == "y" && highshow == "y" {
		out_four = ";1;4m" //高亮
	} else if underline != "y" && highshow == "y" {
		out_four = ";1m"
	} else if underline == "y" && highshow != "y" {
		out_four = ";4m"
	} else {
		out_four = ";22m"
	}

	switch status {
	case "black":
		out_two = "30"
	case "red":
		out_two = "31"
	case "green":
		out_two = "32"
	case "yellow":
		out_two = "33"
	case "blue":
		out_two = "34"
	case "purple":
		out_two = "35"
	case "dgreen":
		out_two = "36"
	case "white":
		out_two = "37"
	default:
		out_two = ""
	}

	switch background {
	case "black":
		out_three = "40;"
	case "red":
		out_three = "41;"
	case "green":
		out_three = "42;"
	case "yellow":
		out_three = "43;"
	case "blue":
		out_three = "44;"
	case "purple":
		out_three = "45;"
	case "dgreen":
		out_three = "46;"
	case "white":
		out_three = "47;"
	default:
		out_three = ""
	}
	return out_one + out_three + out_two + out_four + text + "\033[0m"
}