**gometrics统计指标:**

**loadavg:/proc/loadavg**
    
    1m 5m 15m

**cpu:/proc/stat**

    usr sys idle iowait
    %iowait = (cpu idle time)/(all cpu time)

**swap:/proc/vmstat**

    pswpin   虚拟内存中,从块设备swap区中读入的页数(swap pages)
	pswpout	 虚拟内存中,从块设备swap区中读入的页数(swap pages)

**net:ifconfig**

    recv   
	send

**tcprstat**

**disk:/proc/diskstats**

    r/s:每秒钟完成的读请求数量 
    w/s:每秒钟完成的写请求数量
    rkB/s:每秒钟读取的数量kb
    wkB/s:每秒钟写入的数量kb 
    queue:平均请求队列的长度
    await:平均每次请求的等待时间，单位毫秒  
    svctm:平均每次请求的服务时间，即磁盘读或写操作执行的时间
    util:设备的利用率

**global status**

    Aborted_connects	尝试连接MySQL服务器失败的次数
    Aborted_clients	由于客户端在没有正确关闭连接的情况下死亡而中止的连接数(已成功连接且断开)
    Com_insert	执行insert操作次数
    Com_update	执行update操作次数
    Com_delete	执行delete操作次数
    Com_select	执行select操作次数
    tps:(Com_commit+Com_rollback)/interval	tps
    qps:Questions/interval	qps
    	
    Innodb_rows_inserted	innodb表执行insert操作影响的行数
    Innodb_rows_updated	innodb表执行update操作影响的行数
    Innodb_rows_deleted	innodb表执行delete操作影响的行数
    Innodb_rows_read	innodb表执行select操作影响的行数
    	
    Innodb_buffer_pool_pages_data	缓冲池中包含数据的页数（包含脏页和干净页）
    Innodb_buffer_pool_pages_free	缓冲池中的空页数
    Innodb_buffer_pool_pages_dirty	缓冲池中的脏页数
    Innodb_buffer_pool_pages_flushed	缓冲池中已经刷盘的页数
    	
    Innodb_data_reads	从物理文件读取的次数
    Innodb_data_writes	向物理文件中写入的次数
    Innodb_data_read	从物理文件读取的总字节数
    Innodb_data_written	向物理文件写入的总字节数
    	
    Innodb_os_log_fsyncs	向日志文件完成fsync()写次数
    Innodb_os_log_written	向日志文件写入的总字节数
    	
    Threads_connected	当前已经建立的连接数，该值=show processlist
    Threads_created	为处理超出thread_cache_size的连接创建线程数
    Threads_running	非sleep状态的线程数
    Threads_cached	线程缓存中的线程数（目前都是0）
    thread_cache_hit:=1-(Threads_created/Connections)	线程缓存率，如果thread_cached=0且threads_created不断增大，考虑改大thread_cache_size
    	
    Bytes_received	从所有客户端接收的字节数
    Bytes_sent	向所有客户端发送的字节数
    	
    Rpl_semi_sync_master_tx_avg_wait_time	master节点等待每个事务的平均时间（毫秒）
    Rpl_semi_sync_master_no_tx	master未成功确认的事务数，未成功接收到slave ack的事务数
    Rpl_semi_sync_master_yes_tx	master成功确认的事务数
    Rpl_semi_sync_master_no_times	从半同步降级为异步的次数
    	
    Rpl_semi_sync_master_status	主节点上的半同步复制状态
    Rpl_semi_sync_slave_status	从节点上的半同步复制状态
    	
    Binlog_cache_disk_use	超过binlog_cache_size而使用临时文件来存储的事务数
    	
**global variables:**	
    	
    Master_Host	主节点host
    Master_User	主节点复制用户
    Master_Port	主节点端口
    Master_Server_Id	主节点server_id
    	
    binlog_format	binlog模式:row,statement,mix
    max_binlog_cache_size	binlog可以使用的最大缓存大小，18446744073709547520
    max_binlog_size	binlog单个日志文件大小
    sync_binlog	binlog刷盘机制(=0，commit提交后写入os buffer,不同步刷，=1，commit提交后写入os buffer,调用fsync()同步刷盘，=N，执行N个事务后写入os buffer,调用fsync()同步刷盘)
    max_connect_errors	多个连接请求未成功而中断的次数超过该参数值，则服务器会阻止该客户端继续连接。取消阻止使用flush host命令将错误计数清零
    max_connections	最大并发连接数
    max_user_connections	同一用户允许的最大连接数(0表示无限)
    max_used_connections	自服务器启动以来的最高连接数
    open_files_limit	mysqld可用的最大文件描述符数目
    table_definition_cache	可以存储在定义高速缓存中的表定义（来自.frm文件）的数量
    Select_scan	全表扫描次数
    Select_full_join	没有使用索引的join次数
    Slow_queries	超过long_query_time秒的查询数
    	
    rpl_semi_sync_master_timeout	半同步降级为异步超时时间
    	
    Slave_IO_Running	从库io线程运行状态
    Slave_SQL_Running	从库SQL线程运行状态
    	
    table_open_cache	table高速缓存的数量，缓存表的文件描述符，open_tables=table_open_cache且opened_tables不断增大，则需要增大该值
    thread_cache_size	线程缓存数
    Opened_tables	打开表的数量，该值过大，则表示table_open_cache设置过小
    Created_tmp_disk_tables	在磁盘上创建临时表的数量
    Created_tmp_tables	创建临时表的数量（在磁盘上和内存中）
    	
    innodb_adaptive_flushing	指定是否根据工作负载动态调整InnoDB缓冲池中刷新脏页的速率
    innodb_adaptive_hash_index	是否启用或禁用InnoDB自适应哈希索引
    innodb_buffer_pool_size	innodb缓冲池大小
    innodb_file_per_table	使用独立表空间
    innodb_flush_log_at_trx_commit	控制是否在事务提交时刷新redolog,0:每秒刷日志，写入os buffer,并且fsync()落盘。1：commit刷日志，写入os buffer,并且fsync()并且落盘，2：commit刷日志，不会立即落盘。
    innodb_flush_method	定义用于将数据刷新到InnoDB数据文件和日志文件的方法。O_DIRECT官方文档均使用fsync() flush data and log
    innodb_io_capacity	innodb每秒能够执行的IO数上限。控制缓冲池每秒最多能写多少dirty pages
    innodb_lock_wait_timeout	放弃innodb事务等待行锁定的时长。
    innodb_log_buffer_size	日志缓冲大小
    innodb_log_file_size	单个日志文件大小
    innodb_log_files_in_group	日志组中的日志成员数
    innodb_max_dirty_pages_pct	InnoDB尝试从缓冲池中刷新数据，以便脏页的百分比不超过此值
    innodb_open_files	指定MySQL可以一次保持打开的最大.ibd文件数，最小值为10.如果未启用innodb_file_per_table，则默认值为300，否则为300和table_open_cache中的较高值。
    innodb_read_io_threads	InnoDB中读取操作的I / O线程数
    innodb_thread_concurrency	InnoDB内同时保持操作系统线程的数量限制（InnoDB使用操作系统线程来处理用户事务），超过此限制的线程均会处于等待状态。
    innodb_write_io_threads	InnoDB中写操作的I / O线程数
    	
**engine innodb status:**	
    	
    History_list	undo中待回收的事务数
    unflushed_log:Log sequence number-Log flushed up to	尚未被刷盘的日志字节数
    checkpoint age:=Log sequence number-Last checkpoint at	尚未被刷盘的脏页字节数
    read views open inside InnoDB	打开的只读视图个数
    queries inside InnoDB	有多少个正在查询的操作个数
    queries in queue	在队列中等待的查询请求数
    	
**slave status:**	
    	
    Read_Master_Log_Pos	io线程已经获取到主库binlog的位置
    Exec_Master_Log_Pos	sql线程已经执行到主库Binlog的位置
    Seconds_Behind_Master	从主库执行事务到从库apply事务的延迟时间


