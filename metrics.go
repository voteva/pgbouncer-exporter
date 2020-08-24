package main

import "github.com/prometheus/client_golang/prometheus"

type MetricDescriptor struct {
	Prefix      string
	Labels      []string
	MetricProps []MetricProps
}

type MetricProps struct {
	Type prometheus.ValueType
	Name string
	Help string
}

var InternalMetricUp = MetricProps{
	Type: prometheus.GaugeValue, Name: "up", Help: "Whether pgbouncer is alive",
}

var InternalMetricErrors = MetricProps{
	Type: prometheus.GaugeValue, Name: "errors", Help: "Errors per scrape",
}

var InternalMetricScrapeLastTime = MetricProps{
	Type: prometheus.GaugeValue, Name: "scrape_last_time", Help: "Last timestamp of scrape in unix epoch",
}

var InternalMetricScrapeTotal = MetricProps{
	Type: prometheus.CounterValue, Name: "scrape_total", Help: "Total number of times pgbouncer has been scraped for metrics",
}

var MetricDescriptorLists = MetricDescriptor{
	Prefix: "lists",
	Labels: []string{},
	MetricProps: []MetricProps{
		{Type: prometheus.GaugeValue, Name: "databases", Help: "Count of databases"},
		{Type: prometheus.GaugeValue, Name: "users", Help: "Count of users"},
		{Type: prometheus.GaugeValue, Name: "pools", Help: "Count of pools"},
		{Type: prometheus.GaugeValue, Name: "free_clients", Help: "Count of free clients"},
		{Type: prometheus.GaugeValue, Name: "used_clients", Help: "Count of used clients"},
		{Type: prometheus.GaugeValue, Name: "login_clients", Help: "Count of clients in login state"},
		{Type: prometheus.GaugeValue, Name: "free_servers", Help: "Count of free servers"},
		{Type: prometheus.GaugeValue, Name: "used_servers", Help: "Count of used servers"},
		{Type: prometheus.GaugeValue, Name: "dns_names", Help: "Count of DNS names in the cache"},
		{Type: prometheus.GaugeValue, Name: "dns_zones", Help: "Count of DNS zones in the cache"},
		{Type: prometheus.GaugeValue, Name: "dns_queries", Help: "Count of in-flight DNS queries"},
		{Type: prometheus.GaugeValue, Name: "dns_pending", Help: "Count of DNS pending queries"},
	},
}

var MetricDescriptorStats = MetricDescriptor{
	Prefix: "stats",
	Labels: []string{"database"},
	MetricProps: []MetricProps{
		{Type: prometheus.CounterValue, Name: "total_xact_count", Help: "Total number of SQL transactions pooled"},
		{Type: prometheus.CounterValue, Name: "total_query_count", Help: "Total number of SQL queries pooled"},
		{Type: prometheus.CounterValue, Name: "total_received", Help: "Total volume in bytes of network traffic received by pgbouncer, shown as bytes"},
		{Type: prometheus.CounterValue, Name: "total_sent", Help: "Total volume in bytes of network traffic sent by pgbouncer, shown as bytes"},
		{Type: prometheus.CounterValue, Name: "total_xact_time", Help: "Total number of microseconds spent by pgbouncer when connected to PostgreSQL in a transaction, either idle in transaction or executing queries"},
		{Type: prometheus.CounterValue, Name: "total_query_time", Help: "Total number of microseconds spent by pgbouncer when actively connected to PostgreSQL, executing queries"},
		{Type: prometheus.CounterValue, Name: "total_wait_time", Help: "Time spent by clients waiting for a server in microseconds"},
		{Type: prometheus.GaugeValue, Name: "avg_xact_count", Help: "Average transactions per second in last stat period"},
		{Type: prometheus.GaugeValue, Name: "avg_query_count", Help: "Average queries per second in last stat period"},
		{Type: prometheus.GaugeValue, Name: "avg_recv", Help: "Average received (from clients) bytes per second"},
		{Type: prometheus.GaugeValue, Name: "avg_sent", Help: "Average sent (to clients) bytes per second"},
		{Type: prometheus.GaugeValue, Name: "avg_xact_time", Help: "Average transaction duration in microseconds"},
		{Type: prometheus.GaugeValue, Name: "avg_query_time", Help: "Average query duration in microseconds"},
		{Type: prometheus.GaugeValue, Name: "avg_wait_time", Help: "Time spent by clients waiting for a server in microseconds (average per second)"},
	},
}

var MetricDescriptorPools = MetricDescriptor{
	Prefix: "pools",
	Labels: []string{"database", "user", "pool_mode"},
	MetricProps: []MetricProps{
		{Type: prometheus.GaugeValue, Name: "cl_active", Help: "Client connections linked to server connection and able to process queries, shown as connection"},
		{Type: prometheus.GaugeValue, Name: "cl_waiting", Help: "Client connections waiting on a server connection, shown as connection"},
		{Type: prometheus.GaugeValue, Name: "sv_active", Help: "Server connections linked to a client connection, shown as connection"},
		{Type: prometheus.GaugeValue, Name: "sv_idle", Help: "Server connections idle and ready for a client query, shown as connection"},
		{Type: prometheus.GaugeValue, Name: "sv_used", Help: "Server connections idle more than server_check_delay, needing server_check_query, shown as connection"},
		{Type: prometheus.GaugeValue, Name: "sv_tested", Help: "Server connections currently running either server_reset_query or server_check_query, shown as connection"},
		{Type: prometheus.GaugeValue, Name: "sv_login", Help: "Server connections currently in the process of logging in, shown as connection"},
		{Type: prometheus.GaugeValue, Name: "maxwait", Help: "Age of oldest unserved client connection, shown as second"},
		{Type: prometheus.GaugeValue, Name: "maxwait_us", Help: ""},
	},
}

var MetricDescriptorDatabases = MetricDescriptor{
	Prefix: "databases",
	Labels: []string{"name", "host", "port", "database", "force_user", "pool_mode"},
	MetricProps: []MetricProps{
		{Type: prometheus.GaugeValue, Name: "pool_size", Help: "Maximum number of pool backend connections"},
		{Type: prometheus.GaugeValue, Name: "reserve_pool", Help: "Maximum amount that the pool size can be exceeded temporarily"},
		{Type: prometheus.GaugeValue, Name: "max_connections", Help: "Maximum number of client connections allowed"},
		{Type: prometheus.GaugeValue, Name: "current_connections", Help: "Current number of client connections"},
		{Type: prometheus.GaugeValue, Name: "paused", Help: "Boolean indicating whether a pgbouncer PAUSE is currently active for this database"},
		{Type: prometheus.GaugeValue, Name: "disabled", Help: "Boolean indicating whether a pgbouncer DISABLE is currently active for this database"},
	},
}

var MetricDescriptorConfig = MetricDescriptor{
	Prefix: "config",
	Labels: []string{},
	MetricProps: []MetricProps{
		{Type: prometheus.CounterValue, Name: "listen_backlog", Help: "Maximum number of backlogged listen connections before further connection attempts are dropped"},
		{Type: prometheus.CounterValue, Name: "disable_pqexec", Help: "Boolean; 1 means pgbouncer enforce Simple Query Protocol; 0 means it allows multiple queries in a single packet"},
		{Type: prometheus.CounterValue, Name: "pkt_buf", Help: "Internal buffer size for packets. See docs"},
		{Type: prometheus.GaugeValue, Name: "max_client_conn", Help: "Maximum number of client connections allowed"},
		{Type: prometheus.GaugeValue, Name: "default_pool_size", Help: "The default for how many server connections to allow per user/database pair"},
		{Type: prometheus.GaugeValue, Name: "min_pool_size", Help: "Minimum number of backends a pool will always retain"},
		{Type: prometheus.GaugeValue, Name: "reserve_pool_size", Help: "How many additional connections to allow to a pool once it's crossed it's maximum"},
		{Type: prometheus.GaugeValue, Name: "reserve_pool_timeout", Help: "If a client has not been serviced in this many seconds, pgbouncer enables use of additional connections from reserve pool"},
		{Type: prometheus.GaugeValue, Name: "max_db_connections", Help: "Server level maximum connections enforced for a given db, irregardless of pool limits"},
		{Type: prometheus.GaugeValue, Name: "max_user_connections", Help: "Maximum number of connections a user can open irregardless of pool limits"},
		{Type: prometheus.GaugeValue, Name: "autodb_idle_timeout", Help: "Unused pools created via '*' are reclaimed after this interval"},
		{Type: prometheus.GaugeValue, Name: "server_reset_query_always", Help: "Boolean indicating whether or not server_reset_query is enforced for all pooling modes, or just session"},
		{Type: prometheus.GaugeValue, Name: "server_check_delay", Help: "How long to keep released connections available for immediate re-use, without running sanity-check queries on it. If 0 then the query is ran always"},
		{Type: prometheus.GaugeValue, Name: "query_timeout", Help: "Maximum time that a query can run for before being cancelled"},
		{Type: prometheus.GaugeValue, Name: "query_wait_timeout", Help: "Maximum time that a query can wait to be executed before being cancelled"},
		{Type: prometheus.GaugeValue, Name: "client_idle_timeout", Help: "Client connections idling longer than this many seconds are closed"},
		{Type: prometheus.GaugeValue, Name: "client_login_timeout", Help: "Maximum time in seconds for a client to either login, or be disconnected"},
		{Type: prometheus.GaugeValue, Name: "idle_transaction_timeout", Help: "If client has been in 'idle in transaction' state longer than this amount in seconds, it will be disconnected"},
		{Type: prometheus.GaugeValue, Name: "server_lifetime", Help: "The pooler will close an unused server connection that has been connected longer than this many seconds"},
		{Type: prometheus.GaugeValue, Name: "server_idle_timeout", Help: "If a server connection has been idle more than this many seconds it will be dropped"},
		{Type: prometheus.GaugeValue, Name: "server_connect_timeout", Help: "Maximum time allowed for connecting and logging into a backend server"},
		{Type: prometheus.GaugeValue, Name: "server_login_retry", Help: "If connecting to a backend failed, this is the wait interval in seconds before retrying"},
		{Type: prometheus.GaugeValue, Name: "server_round_robin", Help: "Boolean; if 1, pgbouncer uses backends in a round robin fashion.  If 0, it uses LIFO to minimize connectivity to backends"},
		{Type: prometheus.GaugeValue, Name: "suspend_timeout", Help: "Timeout for how long pgbouncer waits for buffer flushes before killing connections during pgbouncer admin SHUTDOWN and SUSPEND invocations"},
		{Type: prometheus.GaugeValue, Name: "dns_max_ttl", Help: "Irregardless of DNS TTL, this is the TTL that pgbouncer enforces for dns lookups it does for backends"},
		{Type: prometheus.GaugeValue, Name: "dns_nxdomain_ttl", Help: "Irregardless of DNS TTL, this is the period enforced for negative DNS answers"},
		{Type: prometheus.GaugeValue, Name: "dns_zone_check_period", Help: "Period to check if zone serial has changed"},
		{Type: prometheus.GaugeValue, Name: "max_packet_size", Help: "Maximum packet size for postgresql packets that pgbouncer will relay to backends"},
		{Type: prometheus.GaugeValue, Name: "sbuf_loopcnt", Help: "How many results to process for a given connection's packet results before switching to others to ensure fairness"},
		{Type: prometheus.GaugeValue, Name: "tcp_defer_accept", Help: "Configurable for TCP_DEFER_ACCEPT"},
		{Type: prometheus.GaugeValue, Name: "tcp_socket_buffer", Help: "Configurable for tcp socket buffering; 0 is kernel managed"},
		{Type: prometheus.GaugeValue, Name: "tcpkeepalive", Help: "Boolean; if 1, tcp keepalive is enabled w/ OS defaults.  If 0, disabled"},
		{Type: prometheus.GaugeValue, Name: "tcp_keepcnt", Help: "See TCP documentation for this field"},
		{Type: prometheus.GaugeValue, Name: "tcp_keepidle", Help: "See TCP documentation for this field"},
		{Type: prometheus.GaugeValue, Name: "tcp_keepintvl", Help: "See TCP documentation for this field"},
		{Type: prometheus.GaugeValue, Name: "verbose", Help: "If log verbosity is increased.  Only relevant as a metric if log volume begins exceeding log consumption"},
		{Type: prometheus.GaugeValue, Name: "stats_period", Help: "Periodicity in seconds of pgbouncer recalculating internal stats"},
		{Type: prometheus.GaugeValue, Name: "log_connections", Help: "Whether connections are logged or not"},
		{Type: prometheus.GaugeValue, Name: "log_disconnections", Help: "Whether connection disconnects are logged"},
		{Type: prometheus.GaugeValue, Name: "log_pooler_errors", Help: "Whether pooler errors are logged or not"},
		{Type: prometheus.GaugeValue, Name: "application_name_add_host", Help: "Whether pgbouncer add the client host address and port to the application name setting set on connection start or not"},
	},
}
