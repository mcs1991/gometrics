package main

import (
	"flag"
	"fmt"
	"os"
)

type Basic struct {
	//basic info
	hostname string
	ip string
	db string
	var_binlog_format string
	var_max_binlog_cache_size int
	var_max_binlog_size int
	var_max_connect_errors string
	var_max_connections string
	var_max_user_connections string
	var_open_files_limit string
	var_sync_binlog string
	var_table_definition_cache string
	var_table_open_cache string
	var_thread_cache_size string
	var_innodb_adaptive_flushing string
	var_innodb_adaptive_hash_index string
	var_innodb_buffer_pool_size int
	var_innodb_file_per_table string
	var_innodb_flush_log_at_trx_commit string
	var_innodb_flush_method string
	var_innodb_io_capacity string
	var_innodb_lock_wait_timeout string
	var_innodb_log_buffer_size int
	var_innodb_log_file_size int
	var_innodb_log_files_in_group string
	var_innodb_max_dirty_pages_pct string
	var_innodb_open_files string
	var_innodb_read_io_threads string
	var_innodb_thread_concurrency string
	var_innodb_write_io_threads string
	//loadavg
	load_1 float64
	load_5 float64
	load_15 float64
	//cpu
	Cpu_core float64
	cpu_usr int
	cpu_nice int
	cpu_sys int
	cpu_idl int
	cpu_iow int
	cpu_irq int
	cpu_softirq int
	cpu_steal int
	cpu_guest int
	cpu_guest_nice int
	//swap
	swap_in int
	swap_out int
	//net
	net_recv int
	net_send int
	//disk
	io_1 int
	io_2 int
	io_3 int
	io_4 int
	io_5 int
	io_6 int
	io_7 int
	io_8 int
	io_9 int
	io_10 int
	io_11 int
	//tcprstat rt
	rt_count int
	rt_avg int
	rt_a5 int
	rt_a9 int
	//mysql -e "show global status" 不用\G
	//mysql -com
	Question int
	Com_select int
	Com_insert int
	Com_update int
	Com_delete int
	Com_commit int
	Com_rollback int
	//mysql -hit
	//while true;do s1=`mysql -e 'show global status'|grep -w -E 'Innodb_buffer_pool_read_requests|Innodb_buffer_pool_reads'|xargs echo|awk '{print $2}'`;s2=`mysql -e 'show global status'|grep -w -E 'Innodb_buffer_pool_read_requests|Innodb_buffer_pool_reads'|xargs echo|awk '{print $4}'`;sleep 1;ss1=`mysql -e 'show global status'|grep -w -E 'Innodb_buffer_pool_read_requests|Innodb_buffer_pool_reads'|xargs echo|awk '{print $2}'`;ss2=`mysql -e 'show global status'|grep -w -E 'Innodb_buffer_pool_read_requests|Innodb_buffer_pool_reads'|xargs echo|awk '{print $4}'`;rs1=$(($ss1-$s1+1));rs2=$(($ss2-$s2));rs3=$((1000000*($rs1-$rs2)/$rs1));echo $rs1,$rs2,$rs3;done
	// (Innodb_buffer_pool_read_requests - Innodb_buffer_pool_reads) / Innodb_buffer_pool_read_requests * 100%,每秒的计算
	Innodb_buffer_pool_read_requests int
	Innodb_buffer_pool_reads int
	//mysql -innodb_rows
	Innodb_rows_inserted int
	Innodb_rows_updated int
	Innodb_rows_deleted int
	Innodb_rows_read int
	//mysql -innodb_pages
	Innodb_buffer_pool_pages_data int
	Innodb_buffer_pool_pages_free int
	Innodb_buffer_pool_pages_dirty int
	Innodb_buffer_pool_pages_flushed int
	//mysql --innodb_data
	Innodb_data_reads int
	Innodb_data_writes int
	Innodb_data_read int
	Innodb_data_written int
	//mysql --innodb_log
	Innodb_os_log_fsyncs int
	Innodb_os_log_written int
	//mysql --threads
	Threads_running int
	Threads_connected int
	Threads_created int
	Threads_cached int
	//mysql --bytes
	Bytes_received int
	Bytes_sent int
	//mysql --innodb_status show engine innodb status
	//log unflushed = Log sequence number - Log flushed up to
	//uncheckpointed bytes = Log sequence number - Last checkpoint at
	//mysql -e "show engine innodb status\G"|grep -n -E -A4 -B1 "^TRANSACTIONS|LOG|ROW OPERATIONS"
	//mysql -e "show engine innodb status\G"|grep -E "Last checkpoint|read view|queries inside|queue"
	Log_sequence int
	Log_flushed int
	History_list int
	Last_checkpoint int
	Read_view int
	Query_inside int
	Query_queue int
	//addition
	//show status
	Max_used_connections int
	Aborted_connects string
	Aborted_clients string
	Select_full_join string
	Binlog_cache_disk_use string
	Binlog_cache_use string
	Opened_tables string
	//Thread_cache_hits = (1 - Threads_created / connections ) * 100%
	Connections int
	Qcache_hits int
	Handler_read_first int
	Handler_read_key int
	Handler_read_next int
	Handler_read_prev int
	Handler_read_rnd int
	Handler_read_rnd_next int
	Handler_rollback int
	Created_tmp_tables int
	Created_tmp_disk_tables int
	Slow_queries string
	Key_read_requests int
	Key_reads int
	Key_write_requests int
	Key_writes int
	Select_scan string
	//半同步
	Rpl_semi_sync_master_net_avg_wait_time int
	Rpl_semi_sync_master_no_times int
	Rpl_semi_sync_master_no_tx int
	Rpl_semi_sync_master_status string
	Rpl_semi_sync_master_tx_avg_wait_time int
	Rpl_semi_sync_master_wait_sessions int
	Rpl_semi_sync_master_yes_tx int
	Rpl_semi_sync_slave_status string
	rpl_semi_sync_master_timeout string
	//Slave状态监控
	Master_Host string
	Master_User string
	Master_Port string
	Slave_IO_Running string
	Slave_SQL_Running string
	Master_Server_Id string
	Seconds_Behind_Master int
	Read_Master_Log_Pos int
	Exec_Master_Log_Pos int
}

// type mysql struct {
// }

type Flags struct {
	interval string //时间间隔 默认1秒
	count int //运行时间 默认无限
	time bool //打印当前时间
	nocolor bool //不显示颜色
	load bool //打印Load info
	cpu bool //打印Cpu info
	swap bool //打印swap info
	disk string //打印Disk info
	net string // 打印net info
	slave bool // 打印slave info
	username string //mysql用户名
	password string //mysql密码
	host string //mysql连接主机
	port string // mysql连接断开
	socket string //mysql socket连接文件
	com bool //Print MySQL Status(Com_select,Com_insert,Com_update,Com_delete).
	hit bool //Print Innodb Hit%.
	innodb_rows bool //Print Innodb Rows Status(Innodb_rows_inserted/updated/deleted/read).
	innodb_pages bool //Print Innodb Buffer Pool Pages Status(Innodb_buffer_pool_pages_data/free/dirty/flushed)
	innodb_data bool //Print Innodb Data Status(Innodb_data_reads/writes/read/written)
	innodb_log bool //Print Innodb Log Status(Innodb_os_log_fsyncs/written)
	innodb_status bool //Print Innodb Status from Command: 'Show Engine Innodb Status'
	//(history list/ log unflushed/uncheckpointed bytes/ read views/ queries inside/queued)
	threads bool //Print Threads Status(Threads_running,Threads_connected,Threads_created,Threads_cached).
	rt bool //Print MySQL DB RT(us).
	bytes bool //Print Bytes received from/send to MySQL(Bytes_received,Bytes_sent).

	mysql bool //Print MySQLInfo (include -t,-com,-hit,-T,-B).
	innodb bool //Print InnodbInfo(include -t,-innodb_pages,-innodb_data,-innodb_log,-innodb_status)
	sys bool //Print SysInfo (include -t,-l,-c,-s).
	lazy bool //Print Info (include -t,-l,-c,-s,-com,-hit).

	logfile string //Print to Logfile.
	logfile_by_day bool //One day a logfile,the suffix of logfile is 'yyyy-mm-dd';
	semi bool //半同步设置
	other []string
	//and is valid with -L.
}

func (e *Flags) init() {
	interval := flag.String("i", "1", "Time interval")
	count := flag.Int("C", 0, "Gometrics run time")
	time := flag.Bool("t", false, "Print current time")
	nocolor := flag.Bool("nocolor", false, "Do not display color")
	load := flag.Bool("l", false, "Print load info")
	cpu := flag.Bool("c", false, "Print cpu info")
	swap := flag.Bool("s", false, "Print swap info")
	disk := flag.String("d", "none", "Print disk info")
	net := flag.String("n", "none", "Print net info")
	slave := flag.Bool("slave", false, "Print Slave info")
	username := flag.String("u", "", "MySQL user")
	password := flag.String("p", "", "MySQL password")
	host := flag.String("H", "127.0.0.1", "MySQL host")
	port := flag.String("P", "3306", "MySQL port")
	socket := flag.String("S", "/mysql/data/mysql.sock", "MySQL socket")
	com := flag.Bool("com", false, "Print MySQL Status(Com_select,Com_insert,Com_update,Com_delete).")
	hit := flag.Bool("hit", false, "Print Innodb Hit%.")
	innodb_rows := flag.Bool("innodb_rows", false, "Print Innodb Rows Status(Innodb_rows_inserted/updated/deleted/read).")
	innodb_pages := flag.Bool("innodb_pages", false, "Print Innodb Buffer Pool Pages Status(Innodb_buffer_pool_pages_data/free/dirty/flushed)")
	innodb_data := flag.Bool("innodb_data", false, "Print Innodb Data Status(Innodb_data_reads/writes/read/written)")
	innodb_log := flag.Bool("innodb_log", false, "Print Innodb Log Status(Innodb_os_log_fsyncs/written)")
	innodb_status := flag.Bool("innodb_status", false, "Print Innodb Status from Command: 'Show Engine Innodb Status'")
	threads := flag.Bool("T", false, "Print Threads Status(Threads_running,Threads_connected,Threads_created,Threads_cached).")
	rt := flag.Bool("rt", false, "Print tcprstat info.")
	bytes := flag.Bool("B", false, "Print Bytes received from/send to MySQL(Bytes_received,Bytes_sent).")
	mysql := flag.Bool("mysql", false, "Print MySQLInfo (include -t,-com,-hit,-T,-B).")
	innodb := flag.Bool("innodb", false, "Print InnodbInfo(include -t,-innodb_pages,-innodb_data,-innodb_log,-innodb_status)")
	sys := flag.Bool("sys", false, "Print SysInfo (include -t,-l,-c,-s).")
	lazy := flag.Bool("lazy", false, "Print Info (include -t,-l,-c,-s,-com,-hit).")
	semi := flag.Bool("semi", false, "Semisynchronous monitoring")
	logfile := flag.String("L", "none", "Print to Logfile.")
	logfile_by_day := flag.Bool("logfile_by_day", false, "One day a logfile,the suffix of logfile is 'yyyy-mm-dd';")

	flag.Parse()

	e.interval = *interval
	e.count = *count
	e.time = *time
	e.nocolor = *nocolor
	e.load = *load
	e.cpu = *cpu
	e.swap = *swap
	e.disk = *disk
	e.net = *net
	e.slave = *slave
	e.username = *username
	e.password = *password
	e.host = *host
	e.port = *port
	e.socket = *socket
	e.com = *com
	e.hit = *hit
	e.innodb_rows = *innodb_rows
	e.innodb_pages = *innodb_pages
	e.innodb_data = *innodb_data
	e.innodb_log = *innodb_log
	e.innodb_status = *innodb_status
	e.threads = *threads
	e.rt = *rt
	e.bytes = *bytes
	e.mysql = *mysql
	e.innodb = *innodb
	e.sys = *sys
	e.lazy = *lazy
	e.logfile = *logfile
	e.logfile_by_day = *logfile_by_day
	e.semi = *semi
	e.other = flag.Args()

	//NFlag返回解析时进行了设置的flag的数量
	if flag.NFlag() == 0 {
		fmt.Println("请输入【-h】查看帮助!")

		os.Exit(1)
	}
	flag.Usage = func() {
		Helpinfo()
	}
}

func GetValue() map[string]interface{} {
	u := Flags{}
	u.init()
	info := map[string]interface{}{
		"interval": u.interval,
		"count": u.count,
		"time": u.time,
		"nocolor": u.nocolor,
		"load": u.load,
		"cpu": u.cpu,
		"swap": u.swap,
		"disk": u.disk,
		"net": u.net,
		"slave": u.slave,
		"username": u.username,
		"password": u.password,
		"host": u.host,
		"port": u.port,
		"socket": u.socket,
		"com": u.com,
		"hit": u.hit,
		"innodb_rows": u.innodb_rows,
		"innodb_pages": u.innodb_pages,
		"innodb_data": u.innodb_data,
		"innodb_log": u.innodb_log,
		"innodb_status": u.innodb_status,
		"threads": u.threads,
		"rt": u.rt,
		"bytes": u.bytes,
		"mysql": u.mysql,
		"innodb": u.innodb,
		"sys": u.sys,
		"semi": u.semi,
		"lazy": u.lazy,
		"logfile": u.logfile,
		"logfile_by_day": u.logfile_by_day,
		"other": u.other,
	}
	return info
}

func Helpinfo() {
	fmt.Printf("Usage of gometrics:\n")
	fmt.Printf("MySQL Dsn:\n")
	fmt.Printf("-u     MySQL user (default root)\n")
	fmt.Printf("-p     MySQL password (default 123456)\n")
	fmt.Printf("-P     MySQL server port (default 3306)\n")
	fmt.Printf("-H     MySQL server host (default 127.0.0.1)\n")
	fmt.Printf("\nOs metrics:\n")
	fmt.Printf("-c     Print cpu info\n")
	fmt.Printf("-d     Print disk info\n")
	fmt.Printf("-l     Print load info\n")
	fmt.Printf("-s     Print swap info\n")
	fmt.Printf("-n     Print net info\n")
	fmt.Printf("\nMySQL metrics:\n")
	fmt.Printf("-com   Print MySQL Status(Com_select,Com_insert,Com_update,Com_delete)\n")
	fmt.Printf("-hit   Print Innodb Hit%\n")
	fmt.Printf("-innodb_rows     Print Innodb Rows Status(Innodb_rows_inserted/updated/deleted/read)\n")
	fmt.Printf("-innodb_pages    Print Innodb Buffer Pool Pages Status(Innodb_buffer_pool_pages_data/free/dirty/flushed)\n")
	fmt.Printf("-innodb_data     Print Innodb Data Status(Innodb_data_reads/writes/read/written)\n")
	fmt.Printf("-innodb_log      Print Innodb Log Status(Innodb_os_log_fsyncs/written)\n")
	fmt.Printf("-innodb_status   Print Innodb Status from Command: 'Show Engine Innodb Status'\n")
	fmt.Printf("-T     Print Threads Status(Threads_running,Threads_connected,Threads_created,Threads_cached)\n")
	fmt.Printf("-rt    Print tcprstat info\n")
	fmt.Printf("-B     Print Bytes received from/send to MySQL(Bytes_received,Bytes_sent)\n")
	fmt.Printf("-semi  Semisynchronous monitoring info\n")
	fmt.Printf("-slave Print Slave info\n")
	fmt.Printf("\n Gometrics options:\n")
	fmt.Printf("-i     Time interval(Default 1)")
	fmt.Printf("-C     Gometrics run time\n")
	fmt.Printf("-t     Print current time\n")
	fmt.Printf("-nocolor         Do not display color\n")
	fmt.Printf("-L     Print to Logfile\n")
	fmt.Printf("-logfile_by_day  One day a logfile,the suffix of logfile is 'yyyy-mm-dd'\n")
	fmt.Printf("\n Lazy command:\n")
	fmt.Printf("-mysql           Print MySQLInfo (include -t,-com,-hit,-T,-B)\n")
	fmt.Printf("-innodb          Print InnodbInfo(include -t,-innodb_pages,-innodb_data,-innodb_log,-innodb_status)\n")
	fmt.Printf("-sys   Print SysInfo (include -t,-l,-c,-s)\n")
	fmt.Printf("-lazy  Print Info (include -t,-l,-c,-s,-com,-hit)\n")
	os.Exit(1)
}