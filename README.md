**gometrics**

**介绍：**

    监控，收集mysql,os性能数据，并且绘制折线图，直方图。

**安装:**

	export GOPATH=...
    export GOBIN=...
    go install .../gometrics/src/gometrics

**使用说明:**
    
	./gometrics -h
    Usage of ./gometrics:
      -B	Print Bytes received from/send to MySQL(Bytes_received,Bytes_sent).
      -C int
    	运行时间 默认无限
      -H string
    	Mysql连接主机，默认127.0.0.1 (default "127.0.0.1")
      -L string
    	Print to Logfile. (default "none")
      -P string
    	Mysql连接端口,默认3306 (default "3306")
      -S string
    	mysql socket连接文件地址 (default "/mysql/data/mysql.sock")
      -T	Print Threads Status(Threads_running,Threads_connected,Threads_created,Threads_cached).
      -c	打印Cpu info
      -com
    	Print MySQL Status(Com_select,Com_insert,Com_update,Com_delete).
      -d string
    	打印Disk info (default "none")
      -hit
    	Print Innodb Hit%.
      -i string
    	时间间隔 默认1秒 (default "1")
      -innodb
    	Print InnodbInfo(include -t,-innodb_pages,-innodb_data,-innodb_log,-innodb_status)
      -innodb_data
    	Print Innodb Data Status(Innodb_data_reads/writes/read/written)
      -innodb_log
    	Print Innodb Log Status(Innodb_os_log_fsyncs/written)
      -innodb_pages
    	Print Innodb Buffer Pool Pages Status(Innodb_buffer_pool_pages_data/free/dirty/flushed)
      -innodb_rows
    	Print Innodb Rows Status(Innodb_rows_inserted/updated/deleted/read).
      -innodb_status
    	Print Innodb Status from Command: 'Show Engine Innodb Status'
      -l	打印Load info
      -lazy
    	Print Info (include -t,-l,-c,-s,-com,-hit).
      -logfile_by_day
    	One day a logfile,the suffix of logfile is 'yyyy-mm-dd';
      -mysql
    	Print MySQLInfo (include -t,-com,-hit,-T,-B).
      -n string
    	打印net info (default "none")
      -nocolor
    	不显示颜色
      -p string
    	mysql密码
      -rt
    	Print MySQL DB RT(us).
      -s	打印swap info
      -semi
    	半同步监控
      -slave
    	打印Slave info
      -sys
    	Print SysInfo (include -t,-l,-c,-s).
      -t	打印当前时间
      -u string
    	mysql用户名

**性能消耗：**

    指定展示全部的mysql和os指标约占用资源：
    CPU:0.2%
    MEM:0.2%