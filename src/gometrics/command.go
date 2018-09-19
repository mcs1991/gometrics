package main

import (
	"strings"
	"fmt"
	"os"
	"strconv"
	"os/exec"
)

//通用执行shell命令调用函数
func ExecCommand(commands string) string {
	out, err := exec.Command("bash", "-c", commands).Output()
	CheckErr(err)
	return string(out)
}

//性能指标收集函数
func CreateCommand(info map[string]interface{}, count int,mysql_switch string) Basic {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("\033[1;4;31m数据获取异常,请检查输入参数:%s\033[0m\n", err)
			os.Exit(1)
		}
	}()
	ss := Basic{}
	//info
	cpu_core_cmd := "grep processor /proc/cpuinfo | wc -l"
	cpu_string := ExecCommand(cpu_core_cmd)
	cpu_string = strings.Replace(cpu_string, "\n", "", -1)
	ss.Cpu_core, _ = strconv.ParseFloat(cpu_string, 64)
	if count == 0 {
		host_cmd := "hostname"
		host_string := ExecCommand(host_cmd)
		// host_string = strings.Replace(host_string, "\n", "", -1)
		// host_result := strings.Split(rt_string, ",")
		ss.hostname = host_string

		ip_cmd := "ip a|grep global|grep -v lo:|head -1|awk '{print $2}'|cut -d '/' -f1"
		ip_string := ExecCommand(ip_cmd)
		// ip_string = strings.Replace(ip_string, "\n", "", -1)
		ss.ip = ip_string
		if mysql_switch == "on" {
			//if info["username"].(string) != "" &&  info["password"].(string) != ""
			db_cmd := "mysql -u" + info["username"].(string) + " -p" + info["password"].(string) + " -e 'show databases;'|xargs echo|sed 's/ /|/g'"
			db_string := ExecCommand(db_cmd)
			db_string = strings.Replace(db_string, "\n", "", -1)
			ss.db = db_string

			variables_cmd := "mysql -u" + info["username"].(string) + " -p" + info["password"].(string) + " --host=" + info["host"].(string) + " --socket=" + info["socket"].(string) + " --port=" + info["port"].(string) + " -e 'show global variables'|grep -E -w 'binlog_format|max_binlog_cache_size|max_binlog_size|max_connect_errors|max_connections|max_user_connections|open_files_limit|sync_binlog|table_definition_cache|table_open_cache|thread_cache_size|innodb_adaptive_flushing|innodb_adaptive_hash_index|innodb_buffer_pool_size|innodb_file_per_table|innodb_flush_log_at_trx_commit|innodb_io_capacity|innodb_lock_wait_timeout|innodb_log_buffer_size|innodb_log_file_size|innodb_log_files_in_group|innodb_max_dirty_pages_pct|innodb_open_files|innodb_read_io_threads|innodb_thread_concurrency|innodb_write_io_threads'|awk '{print $2}'|xargs echo|sed 's/ /,/g'"
			variables_string := ExecCommand(variables_cmd)
			variables_string = strings.Replace(variables_string, "\n", "", -1)
			variables_result := strings.Split(variables_string, ",")

			ss.var_binlog_format = variables_result[0]
			ss.var_innodb_adaptive_flushing = variables_result[1]
			ss.var_innodb_adaptive_hash_index = variables_result[2]
			ss.var_innodb_buffer_pool_size, _ = strconv.Atoi(variables_result[3])
			ss.var_innodb_file_per_table = variables_result[4]
			ss.var_innodb_flush_log_at_trx_commit = variables_result[5]
			ss.var_innodb_io_capacity = variables_result[6]
			ss.var_innodb_lock_wait_timeout = variables_result[7]
			ss.var_innodb_log_buffer_size, _ = strconv.Atoi(variables_result[8])
			ss.var_innodb_log_file_size, _ = strconv.Atoi(variables_result[9])
			ss.var_innodb_log_files_in_group = variables_result[10]
			ss.var_innodb_max_dirty_pages_pct = variables_result[11]
			ss.var_innodb_open_files = variables_result[12]
			ss.var_innodb_read_io_threads = variables_result[13]
			ss.var_innodb_thread_concurrency = variables_result[14]
			ss.var_innodb_write_io_threads = variables_result[15]
			ss.var_max_binlog_cache_size, _ = strconv.Atoi(variables_result[16])
			ss.var_max_binlog_size, _ = strconv.Atoi(variables_result[17])
			ss.var_max_connect_errors = variables_result[18]
			ss.var_max_connections = variables_result[19]
			ss.var_max_user_connections = variables_result[20]
			ss.var_open_files_limit = variables_result[21]
			ss.var_sync_binlog = variables_result[22]
			ss.var_table_definition_cache = variables_result[23]
			ss.var_table_open_cache = variables_result[24]
			ss.var_thread_cache_size = variables_result[25]

			innodb_flush_method_cmd := "mysql -u" + info["username"].(string) + " -p" + info["password"].(string) + " --host=" + info["host"].(string) + " --socket=" + info["socket"].(string) + " --port=" + info["port"].(string) + " -e 'show variables'|grep -E -w 'innodb_flush_method'|awk '{print $2}'"
			innodb_flush_method_string := ExecCommand(innodb_flush_method_cmd)
			innodb_flush_method_string = strings.Replace(innodb_flush_method_string, "\n", "", -1)

			ss.var_innodb_flush_method = innodb_flush_method_string

			//semi
			semi_tmp := "mysql -u" + info["username"].(string) + " -p" + info["password"].(string) + " --host=" + info["host"].(string) + " --socket=" + info["socket"].(string) + " --port=" + info["port"].(string) + " -e 'show variables'|grep -E -w 'rpl_semi_sync_master_timeout'|awk '{print $2}'"
			semi_string := ExecCommand(semi_tmp)
			semi_string = strings.Replace(semi_string, "\n", "", -1)

			ss.rpl_semi_sync_master_timeout = semi_string

			//mysql global status
			innodb_cmd := "mysql -u" + info["username"].(string) + " -p" + info["password"].(string) + " --host=" + info["host"].(string) + " --socket=" + info["socket"].(string) + " --port=" + info["port"].(string) + " -e 'show global status'|grep -w -E 'Max_used_connections|Aborted_connects|Aborted_clients|Select_full_join|Binlog_cache_disk_use|Binlog_cache_use|Opened_tables|Connections|Qcache_hits|Handler_read_first|Handler_read_key|Handler_read_next|Handler_read_prev|Handler_read_rnd|Handler_read_rnd_next|Handler_rollback|Created_tmp_tables|Created_tmp_disk_tables|Slow_queries|Key_read_requests|Key_reads|Key_write_requests|Key_writes|Select_scan|Rpl_semi_sync_master_status|Rpl_semi_sync_slave_status'|awk '{print $2}'|xargs echo|sed 's/[[:space:]]/,/g'"

			innodb_string := ExecCommand(innodb_cmd)
			innodb_string = strings.Replace(innodb_string, "\n", "", -1)
			innodb_result := strings.Split(innodb_string, ",")

			lens := len(innodb_result)

			ss.Aborted_clients = innodb_result[0]
			ss.Aborted_connects = innodb_result[1]
			ss.Binlog_cache_disk_use = innodb_result[2]
			ss.Binlog_cache_use = innodb_result[3]
			ss.Connections, _ = strconv.Atoi(innodb_result[4])
			ss.Created_tmp_disk_tables, _ = strconv.Atoi(innodb_result[5])
			ss.Created_tmp_tables, _ = strconv.Atoi(innodb_result[6])
			ss.Handler_read_first, _ = strconv.Atoi(innodb_result[7])
			ss.Handler_read_key, _ = strconv.Atoi(innodb_result[8])
			ss.Handler_read_next, _ = strconv.Atoi(innodb_result[9])
			ss.Handler_read_prev, _ = strconv.Atoi(innodb_result[10])
			ss.Handler_read_rnd, _ = strconv.Atoi(innodb_result[11])
			ss.Handler_read_rnd_next, _ = strconv.Atoi(innodb_result[12])
			ss.Handler_rollback, _ = strconv.Atoi(innodb_result[13])
			ss.Key_read_requests, _ = strconv.Atoi(innodb_result[14])
			ss.Key_reads, _ = strconv.Atoi(innodb_result[15])
			ss.Key_write_requests, _ = strconv.Atoi(innodb_result[16])
			ss.Key_writes, _ = strconv.Atoi(innodb_result[17])
			ss.Max_used_connections, _ = strconv.Atoi(innodb_result[18])
			ss.Opened_tables = innodb_result[19]
			ss.Qcache_hits, _ = strconv.Atoi(innodb_result[20])
			if lens == 26 {
				ss.Rpl_semi_sync_master_status = innodb_result[21]
				ss.Rpl_semi_sync_slave_status = innodb_result[22]
				ss.Select_full_join = innodb_result[23]
				ss.Select_scan = innodb_result[24]
				ss.Slow_queries = innodb_result[25]
			} else {
				ss.Select_full_join = innodb_result[21]
				ss.Select_scan = innodb_result[22]
				ss.Slow_queries = innodb_result[23]
			}

			slave_cmd := "mysql -u" + info["username"].(string) + " -p" + info["password"].(string) + " --host=" + info["host"].(string) + " --socket=" + info["socket"].(string) + " --port=" + info["port"].(string) + " -e 'show slave status\\G'|grep -E -w 'Master_Host|Master_User|Master_Port|Slave_IO_Running|Slave_SQL_Running|Seconds_Behind_Master|Master_Server_Id|Read_Master_Log_Pos|Exec_Master_Log_Pos'|awk '{print $2}'|xargs echo|sed 's/[[:space:]]/,/g'"
			slave_string := ExecCommand(slave_cmd)
			slave_string = strings.Replace(slave_string, "\n", "", -1)
			slave_result := strings.Split(slave_string, ",")
			if slave_result[0] == "" {
				ss.Master_Host = ""
			} else {
				ss.Master_Host = slave_result[0]
				ss.Master_User = slave_result[1]
				ss.Master_Port = slave_result[2]
				ss.Slave_IO_Running = slave_result[4]
				ss.Slave_SQL_Running = slave_result[5]
				ss.Master_Server_Id = slave_result[8]
			}
			// fmt.Println(semi_cmd)
			// fmt.Println(semi_result)
			// //mysql engine innodb status
		}
	}
	//rt
	if info["rt"] == true {
		if count == 0 {
			ss.rt_count = 0
			ss.rt_avg = 0
			ss.rt_a5 = 0
			ss.rt_a9 = 0
		} else {
			var rt_cmd string
			rt_cmd = "tail -1 /tmp/gometrics_tcprstat.log |awk '{print $2,$5,$9,$12}'|sed 's/[[:space:]]/,/g'"

			rt_string := ExecCommand(rt_cmd)
			// fmt.Println(rt_string)
			rt_string = strings.Replace(rt_string, "\n", "", -1)
			rt_result := strings.Split(rt_string, ",")

			ss.rt_count, _ = strconv.Atoi(rt_result[0])
			ss.rt_avg, _ = strconv.Atoi(rt_result[1])
			ss.rt_a5, _ = strconv.Atoi(rt_result[2])
			ss.rt_a9, _ = strconv.Atoi(rt_result[3])
		}
	}

	// swap bool //打印swap info
	if info["swap"] == true || info["load"] == true || info["cpu"] == true {
		basic_cmd := "cat /proc/loadavg /proc/stat /proc/vmstat |sed 's/\\// china /g'|grep -w -E 'china|cpu|pswpin|pswpout'|xargs echo|awk '{print $1,$2,$3,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$20,$22}'|sed 's/[[:space:]]/,/g'"
		basic_string := ExecCommand(basic_cmd)
		basic_string = strings.Replace(basic_string, "\n", "", -1)
		basic_result := strings.Split(basic_string, ",")

		ss.load_1, _ = strconv.ParseFloat(basic_result[0], 64)
		ss.load_5, _ = strconv.ParseFloat(basic_result[1], 64)
		ss.load_15, _ = strconv.ParseFloat(basic_result[2], 64)
		ss.cpu_usr, _ = strconv.Atoi(basic_result[3])
		ss.cpu_nice, _ = strconv.Atoi(basic_result[4])
		ss.cpu_sys, _ = strconv.Atoi(basic_result[5])
		ss.cpu_idl, _ = strconv.Atoi(basic_result[6])
		ss.cpu_iow, _ = strconv.Atoi(basic_result[7])
		ss.cpu_irq, _ = strconv.Atoi(basic_result[8])
		ss.cpu_softirq, _ = strconv.Atoi(basic_result[9])
		ss.cpu_steal, _ = strconv.Atoi(basic_result[10])
		ss.cpu_guest, _ = strconv.Atoi(basic_result[11])
		ss.cpu_guest_nice, _ = strconv.Atoi(basic_result[12])
		ss.swap_in, _ = strconv.Atoi(basic_result[13])
		ss.swap_out, _ = strconv.Atoi(basic_result[14])
	}

	//disk
	if info["disk"] != "none" {
		disk_cmd := "cat /proc/diskstats |grep -w -E '" + info["disk"].(string) + "'|awk '{print $4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14}'|sed 's/[[:space:]]/,/g'"
		// fmt.Println(disk_cmd)
		disk_string := ExecCommand(disk_cmd)
		disk_string = strings.Replace(disk_string, "\n", "", -1)
		// fmt.Println(disk_string)
		disk_result := strings.Split(disk_string, ",")
		// fmt.Println(disk_result)
		ss.io_1, _ = strconv.Atoi(disk_result[0])
		ss.io_2, _ = strconv.Atoi(disk_result[1])
		ss.io_3, _ = strconv.Atoi(disk_result[2])
		ss.io_4, _ = strconv.Atoi(disk_result[3])
		ss.io_5, _ = strconv.Atoi(disk_result[4])
		ss.io_6, _ = strconv.Atoi(disk_result[5])
		ss.io_7, _ = strconv.Atoi(disk_result[6])
		ss.io_8, _ = strconv.Atoi(disk_result[7])
		ss.io_9, _ = strconv.Atoi(disk_result[8])
		ss.io_10, _ = strconv.Atoi(disk_result[9])
		ss.io_11, _ = strconv.Atoi(disk_result[10])
	}

	//net
	if info["net"] != "none" {
		net_cmd := "ifconfig " + info["net"].(string) + "|grep bytes|sed 's/:/ /g'|awk '{print $3,$8}'|sed 's/ /,/g'"
		// net_cmd := "cat /proc/net/dev |grep -E -w '" + info["net"].(string) + "'|awk '{print $2*8,$10*8}'|sed 's/[[:space:]]/,/g'"
		// fmt.Println(net_cmd)
		net_string := ExecCommand(net_cmd)
		net_string = strings.Replace(net_string, "\n", "", -1)
		net_result := strings.Split(net_string, ",")

		ss.net_recv, _ = strconv.Atoi(net_result[0])
		ss.net_send, _ = strconv.Atoi(net_result[1])
	}
	if info["innodb_status"] == true {
		engine_cmd := "mysql -u" + info["username"].(string) + " -p" + info["password"].(string) + " --host=" + info["host"].(string) + " --socket=" + info["socket"].(string) + " --port=" + info["port"].(string) + " -e 'show engine innodb status\\G'|grep -w -E 'History list|Log sequence|Log flushed|queries inside|queue|read views|Last checkpoint'|xargs echo|awk '{print $4,$8,$13,$17,$18,$22,$26}'|sed 's/[[:space:]]/,/g'"
		engine_string := ExecCommand(engine_cmd)
		engine_string = strings.Replace(engine_string, "\n", "", -1)
		engine_result := strings.Split(engine_string, ",")

		ss.History_list, _ = strconv.Atoi(engine_result[0])
		ss.Log_sequence, _ = strconv.Atoi(engine_result[1])
		ss.Log_flushed, _ = strconv.Atoi(engine_result[2])
		ss.Last_checkpoint, _ = strconv.Atoi(engine_result[3])
		ss.Query_inside, _ = strconv.Atoi(engine_result[4])
		ss.Query_queue, _ = strconv.Atoi(engine_result[5])
		ss.Read_view, _ = strconv.Atoi(engine_result[6])
	}
	if mysql_switch == "on" {
		// //mysql global status
		if info["com"] == true || info["hit"] == true || info["innodb_rows"] == true || info["innodb_pages"] == true || info["innodb_data"] == true || info["innodb_log"] == true || info["threads"] == true || info["bytes"] == true {
			global_cmd := "mysql -u" + info["username"].(string) + " -p" + info["password"].(string) + " --host=" + info["host"].(string) + " --socket=" + info["socket"].(string) + " --port=" + info["port"].(string) + " -e 'show global status'|grep -w -E 'Questions|Com_select|Com_insert|Com_update|Com_delete|Com_commit|Com_rollback|Innodb_buffer_pool_read_requests|Innodb_buffer_pool_reads|Innodb_rows_inserted|Innodb_rows_updated|Innodb_rows_deleted|Innodb_rows_read|Innodb_buffer_pool_pages_data|Innodb_buffer_pool_pages_free|Innodb_buffer_pool_pages_dirty|Innodb_buffer_pool_pages_flushed|Innodb_data_reads|Innodb_data_writes|Innodb_data_read|Innodb_data_written|Innodb_os_log_fsyncs|Innodb_os_log_written|Threads_running|Threads_connected|Threads_created|Threads_cached|Bytes_received|Bytes_sent|Max_used_connections|Aborted_connects|Aborted_clients|Select_full_join|Binlog_cache_disk_use|Binlog_cache_use|Opened_tables|Connections|Qcache_hits|Handler_read_first|Handler_read_key|Handler_read_next|Handler_read_prev|Handler_read_rnd|Handler_read_rnd_next|Handler_rollback|Created_tmp_tables|Created_tmp_disk_tables|Slow_queries|Key_read_requests|Key_reads|Key_write_requests|Key_writes'|awk '{print $2}'|xargs echo|sed 's/[[:space:]]/,/g'"

			global_string := ExecCommand(global_cmd)
			global_string = strings.Replace(global_string, "\n", "", -1)
			global_result := strings.Split(global_string, ",")

			// ss.Aborted_clients = global_result[0]
			// ss.Aborted_connects = global_result[1]
			// ss.Binlog_cache_disk_use = global_result[2]
			// ss.Binlog_cache_use, _ = strconv.Atoi(global_result[3])
			ss.Bytes_received, _ = strconv.Atoi(global_result[4])
			ss.Bytes_sent, _ = strconv.Atoi(global_result[5])
			ss.Com_commit, _ = strconv.Atoi(global_result[6])
			ss.Com_delete, _ = strconv.Atoi(global_result[7])
			ss.Com_insert, _ = strconv.Atoi(global_result[8])
			ss.Com_rollback, _ = strconv.Atoi(global_result[9])
			ss.Com_select, _ = strconv.Atoi(global_result[10])
			ss.Com_update, _ = strconv.Atoi(global_result[11])
			ss.Connections, _ = strconv.Atoi(global_result[12])
			ss.Created_tmp_disk_tables, _ = strconv.Atoi(global_result[13])
			ss.Created_tmp_tables, _ = strconv.Atoi(global_result[14])
			ss.Handler_read_first, _ = strconv.Atoi(global_result[15])
			ss.Handler_read_key, _ = strconv.Atoi(global_result[16])
			ss.Handler_read_next, _ = strconv.Atoi(global_result[17])
			ss.Handler_read_prev, _ = strconv.Atoi(global_result[18])
			ss.Handler_read_rnd, _ = strconv.Atoi(global_result[19])
			ss.Handler_read_rnd_next, _ = strconv.Atoi(global_result[20])
			ss.Handler_rollback, _ = strconv.Atoi(global_result[21])
			ss.Innodb_buffer_pool_pages_data, _ = strconv.Atoi(global_result[22])
			ss.Innodb_buffer_pool_pages_dirty, _ = strconv.Atoi(global_result[23])
			ss.Innodb_buffer_pool_pages_flushed, _ = strconv.Atoi(global_result[24])
			ss.Innodb_buffer_pool_pages_free, _ = strconv.Atoi(global_result[25])
			ss.Innodb_buffer_pool_read_requests, _ = strconv.Atoi(global_result[26])
			ss.Innodb_buffer_pool_reads, _ = strconv.Atoi(global_result[27])
			ss.Innodb_data_read, _ = strconv.Atoi(global_result[28])
			ss.Innodb_data_reads, _ = strconv.Atoi(global_result[29])
			ss.Innodb_data_writes, _ = strconv.Atoi(global_result[30])
			ss.Innodb_data_written, _ = strconv.Atoi(global_result[31])
			ss.Innodb_os_log_fsyncs, _ = strconv.Atoi(global_result[32])
			ss.Innodb_os_log_written, _ = strconv.Atoi(global_result[33])
			ss.Innodb_rows_deleted, _ = strconv.Atoi(global_result[34])
			ss.Innodb_rows_inserted, _ = strconv.Atoi(global_result[35])
			ss.Innodb_rows_read, _ = strconv.Atoi(global_result[36])
			ss.Innodb_rows_updated, _ = strconv.Atoi(global_result[37])
			ss.Key_read_requests, _ = strconv.Atoi(global_result[38])
			ss.Key_reads, _ = strconv.Atoi(global_result[39])
			ss.Key_write_requests, _ = strconv.Atoi(global_result[40])
			ss.Key_writes, _ = strconv.Atoi(global_result[41])
			ss.Max_used_connections, _ = strconv.Atoi(global_result[42])
			// ss.Opened_tables, _ = strconv.Atoi(global_result[43])
			//ss.Qcache_hits, _ = strconv.Atoi(global_result[44])
			// ss.Select_full_join, _ = strconv.Atoi(global_result[45])
			ss.Question, _ = strconv.Atoi(global_result[45])
			ss.Slow_queries = global_result[47]
			ss.Threads_cached, _ = strconv.Atoi(global_result[48])
			ss.Threads_connected, _ = strconv.Atoi(global_result[49])
			ss.Threads_created, _ = strconv.Atoi(global_result[50])
			ss.Threads_running, _ = strconv.Atoi(global_result[51])

		}

		// //mysql engine innodb status
		if info["semi"] == true {
			semi_cmd := "mysql -u" + info["username"].(string) + " -p" + info["password"].(string) + " --host=" + info["host"].(string) + " --socket=" + info["socket"].(string) + " --port=" + info["port"].(string) + " -e 'show status'|grep -E Rpl_semi|awk '{print $2}'|xargs echo|sed 's/[[:space:]]/,/g'"
			semi_string := ExecCommand(semi_cmd)
			semi_string = strings.Replace(semi_string, "\n", "", -1)
			semi_result := strings.Split(semi_string, ",")
			if semi_result[0] == "" {
				fmt.Println(ShowFont("semi半同步未开启", red, "", "", "y"))
				os.Exit(1)
			}
			// fmt.Println(semi_cmd)
			// fmt.Println(semi_result)
			ss.Rpl_semi_sync_master_net_avg_wait_time, _ = strconv.Atoi(semi_result[1])
			ss.Rpl_semi_sync_master_no_times, _ = strconv.Atoi(semi_result[4])
			ss.Rpl_semi_sync_master_no_tx, _ = strconv.Atoi(semi_result[5])
			// ss.Rpl_semi_sync_master_status = semi_result[6]
			ss.Rpl_semi_sync_master_tx_avg_wait_time, _ = strconv.Atoi(semi_result[8])
			ss.Rpl_semi_sync_master_wait_sessions, _ = strconv.Atoi(semi_result[12])
			ss.Rpl_semi_sync_master_yes_tx, _ = strconv.Atoi(semi_result[13])
			// ss.Rpl_semi_sync_slave_status = semi_result[14]
		}

		// slave status
		if info["slave"] == true {
			slave_cmd := "mysql -u" + info["username"].(string) + " -p" + info["password"].(string) + " --host=" + info["host"].(string) + " --socket=" + info["socket"].(string) + " --port=" + info["port"].(string) + " -e 'show slave status\\G'|grep -E -w 'Master_Host|Master_User|Master_Port|Slave_IO_Running|Slave_SQL_Running|Seconds_Behind_Master|Master_Server_Id|Read_Master_Log_Pos|Exec_Master_Log_Pos'|awk '{print $2}'|xargs echo|sed 's/[[:space:]]/,/g'"
			slave_string := ExecCommand(slave_cmd)
			slave_string = strings.Replace(slave_string, "\n", "", -1)
			slave_result := strings.Split(slave_string, ",")
			if slave_result[0] == "" {
				fmt.Println(ShowFont("该主机Mysql不是Slave端", red, "", "", "y"))
				os.Exit(1)
			}
			// fmt.Println(semi_cmd)
			// fmt.Println(semi_result)
			// ss.Master_Host = slave_result[0]
			// ss.Master_User = slave_result[1]
			// ss.Master_Port = slave_result[2]
			ss.Read_Master_Log_Pos, _ = strconv.Atoi(slave_result[3])
			// ss.Slave_IO_Running = slave_result[4]
			// ss.Slave_SQL_Running = slave_result[5]
			ss.Exec_Master_Log_Pos, _ = strconv.Atoi(slave_result[6])
			ss.Seconds_Behind_Master, _ = strconv.Atoi(slave_result[7])
			// ss.Master_Server_Id = slave_result[8]
		}
	}
	return ss
}
