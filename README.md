**gometrics**

**介绍：**

    监控，收集mysql,os性能数据，并且绘制折线图，直方图。

**安装:**

	export GOPATH=...
    export GOBIN=...
    go install .../gometrics/src/gometrics

**使用说明:**
    
	./gometrics -h
    Usage of gometrics:
    MySQL Dsn:
    -u     MySQL user (default root)
    -p 	   MySQL password (default 123456)
    -P     MySQL server port (default 3306)
    -H     MySQL server host (default 127.0.0.1)
    
    Os metrics:
    -c     Print cpu info
    -d     Print disk info
    -l     Print load info
    -s     Print swap info
    -n     Print net info
    
    MySQL metrics:
    -com   			 Print MySQL Status(Com_select,Com_insert,Com_update,Com_delete)
    -hit   			 Print Innodb Hit
    -T    			 Print Threads Status(Threads_running,Threads_connected,Threads_created,Threads_cached)
    -rt				 Print tcprstat info
    -B     			 Print Bytes received from/send to MySQL(Bytes_received,Bytes_sent)
    -semi  			 Semisynchronous monitoring info
    -slave 			 Print Slave info
    -innodb_rows     Print Innodb Rows Status(Innodb_rows_ inserted/updated/deleted/read)
    -innodb_pages    Print Innodb Buffer Pool Pages Status(Innodb_buffer_pool_pages_data/free/dirty/flushed)
    -innodb_data     Print Innodb Data Status(Innodb_data_reads/writes/read/written)
    -innodb_log      Print Innodb Log Status(Innodb_os_log_fsyncs/written)
    -innodb_status   Print Innodb Status from Command: 'Show Engine Innodb Status'

     Gometrics options:
    -i     			 Time interval(Default 1)
	-C     			 Gometrics run time
    -t     			 Print current time
    -nocolor         Do not display color
    -L     			 Print to Logfile
    -logfile_by_day  One day a logfile,the suffix of logfile is 'yyyy-mm-dd'
    
     Lazy command:
    -mysql           Print MySQLInfo (include -t,-com,-hit,-T,-B)
    -innodb          Print InnodbInfo(include -t,-innodb_pages,-innodb_data,-innodb_log,-innodb_status)
    -sys   			 Print SysInfo (include -t,-l,-c,-s)
    -lazy  			 Print Info (include -t,-l,-c,-s,-com,-hit)


**性能消耗：**

    指定展示全部的mysql和os指标约占用资源：
    CPU:0.2%
    MEM:0.2%