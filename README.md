# PgBouncer metrics exporter

Prometheus exporter for PgBouncer.
Exports metrics at `8989/metrics`

## Build
```
make build
```
Build a docker image:
```
docker build -t pgbouncer-exporter .
```

## Run
```
pgbouncer-exporter [-p <telemetry port>] [-d <data source>]
```

## Flags
* ``` -p ``` - Port to listen on for web interface and telemetry
* ``` -d ``` - PgBouncer connection url

## Metrics
#### Internal
```
pgbouncer_up{}
pgbouncer_errors{}
pgbouncer_scrape_last_time{}
pgbouncer_scrape_total{}
```
#### Lists
```
pgbouncer_lists_databases{}
pgbouncer_lists_users{}
pgbouncer_lists_pools{}
pgbouncer_lists_free_clients{}
pgbouncer_lists_used_clients{}
pgbouncer_lists_login_clients{}
pgbouncer_lists_free_servers{}
pgbouncer_lists_used_servers{}
pgbouncer_lists_dns_names{}
pgbouncer_lists_dns_zones{}
pgbouncer_lists_dns_queries{}
pgbouncer_lists_dns_pending{}
```
#### Stats
```
pgbouncer_stats_total_xact_count{database}
pgbouncer_stats_total_query_count{database}
pgbouncer_stats_total_received{database}
pgbouncer_stats_total_sent{database}
pgbouncer_stats_total_xact_time{database}
pgbouncer_stats_total_query_time{database}
pgbouncer_stats_total_wait_time{database}
pgbouncer_stats_avg_xact_count{database}
pgbouncer_stats_avg_query_count{database}
pgbouncer_stats_avg_recv{database}
pgbouncer_stats_avg_sent{database}
pgbouncer_stats_avg_xact_time{database}
pgbouncer_stats_avg_query_time{database}
pgbouncer_stats_avg_wait_time{database}
```
#### Pools
```
pgbouncer_pools_cl_active{database,user,pool_mode}
pgbouncer_pools_cl_waiting{database,user,pool_mode}
pgbouncer_pools_sv_active{database,user,pool_mode}
pgbouncer_pools_sv_idle{database,user,pool_mode}
pgbouncer_pools_sv_used{database,user,pool_mode}
pgbouncer_pools_sv_tested{database,user,pool_mode}
pgbouncer_pools_sv_login{database,user,pool_mode}
pgbouncer_pools_maxwait{database,user,pool_mode}
pgbouncer_pools_maxwait_us{database,user,pool_mode}
```
#### Databases
```
pgbouncer_databases_pool_size{name,host,port,database,force_user,pool_mode}
pgbouncer_databases_reserve_pool{name,host,port,database,force_user,pool_mode}
pgbouncer_databases_max_connections{name,host,port,database,force_user,pool_mode}
pgbouncer_databases_current_connections{name,host,port,database,force_user,pool_mode}
pgbouncer_databases_paused{name,host,port,database,force_user,pool_mode}
pgbouncer_databases_disabled{name,host,port,database,force_user,pool_mode}
```
#### Config
```
pgbouncer_config_listen_backlog{}
pgbouncer_config_disable_pqexec{}
pgbouncer_config_pkt_buf{}
pgbouncer_config_max_client_conn{}
pgbouncer_config_default_pool_size{}
pgbouncer_config_min_pool_size{}
pgbouncer_config_reserve_pool_size{}
pgbouncer_config_reserve_pool_timeout{}
pgbouncer_config_max_db_connections{}
pgbouncer_config_max_user_connections{}
pgbouncer_config_autodb_idle_timeout{}
pgbouncer_config_server_reset_query_always{}
pgbouncer_config_server_check_delay{}
pgbouncer_config_query_timeout{}
pgbouncer_config_query_wait_timeout{}
pgbouncer_config_client_idle_timeout{}
pgbouncer_config_client_login_timeout{}
pgbouncer_config_idle_transaction_timeout{}
pgbouncer_config_server_lifetime{}
pgbouncer_config_server_idle_timeout{}
pgbouncer_config_server_connect_timeout{}
pgbouncer_config_server_login_retry{}
pgbouncer_config_server_round_robin{}
pgbouncer_config_suspend_timeout{}
pgbouncer_config_dns_max_ttl{}
pgbouncer_config_dns_nxdomain_ttl{}
pgbouncer_config_max_packet_size{}
pgbouncer_config_sbuf_loopcnt{}
pgbouncer_config_tcp_defer_accept{}
pgbouncer_config_tcp_socket_buffer{}
pgbouncer_config_tcpkeepalive{}
pgbouncer_config_tcp_keepcnt{}
pgbouncer_config_tcp_keepidle{}
pgbouncer_config_tcp_keepintvl{}
pgbouncer_config_verbose{}
pgbouncer_config_stats_period{}
pgbouncer_config_log_connections{}
pgbouncer_config_log_disconnections{}
pgbouncer_config_log_pooler_errors{}
pgbouncer_config_application_name_add_host{}
```
