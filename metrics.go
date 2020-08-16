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

var InternalMetricScrapeLastTime = MetricProps{
	Type: prometheus.GaugeValue, Name: "scrape_last_time", Help: "", // TODO
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
		{Type: prometheus.CounterValue, Name: "total_xact_count", Help: ""},
		{Type: prometheus.CounterValue, Name: "total_query_count", Help: ""},
		{Type: prometheus.CounterValue, Name: "total_received", Help: ""},
		{Type: prometheus.CounterValue, Name: "total_sent", Help: ""},
		{Type: prometheus.CounterValue, Name: "total_xact_time", Help: ""},
		{Type: prometheus.CounterValue, Name: "total_query_time", Help: ""},
		{Type: prometheus.CounterValue, Name: "total_wait_time", Help: ""},
		{Type: prometheus.GaugeValue, Name: "avg_xact_count", Help: ""},
		{Type: prometheus.GaugeValue, Name: "avg_query_count", Help: ""},
		{Type: prometheus.GaugeValue, Name: "avg_recv", Help: ""},
		{Type: prometheus.GaugeValue, Name: "avg_sent", Help: ""},
		{Type: prometheus.GaugeValue, Name: "avg_xact_time", Help: ""},
		{Type: prometheus.GaugeValue, Name: "avg_query_time", Help: ""},
		{Type: prometheus.GaugeValue, Name: "avg_wait_time", Help: ""},
	},
}

var MetricDescriptorPools = MetricDescriptor{
	Prefix: "pools",
	Labels: []string{"database", "user"}, // TODO pool_mode?
	MetricProps: []MetricProps{
		{Type: prometheus.GaugeValue, Name: "cl_active", Help: ""},
		{Type: prometheus.GaugeValue, Name: "cl_waiting", Help: ""},
		{Type: prometheus.GaugeValue, Name: "sv_active", Help: ""},
		{Type: prometheus.GaugeValue, Name: "sv_idle", Help: ""},
		{Type: prometheus.GaugeValue, Name: "sv_used", Help: ""},
		{Type: prometheus.GaugeValue, Name: "sv_tested", Help: ""},
		{Type: prometheus.GaugeValue, Name: "sv_login", Help: ""},
		{Type: prometheus.GaugeValue, Name: "maxwait", Help: ""},
	},
}

var MetricDescriptorDatabases = MetricDescriptor{
	Prefix: "databases",
	Labels: []string{"name", "host", "port", "database"}, // TODO pool_mode?
	MetricProps: []MetricProps{
		{Type: prometheus.GaugeValue, Name: "force_user", Help: ""},
		{Type: prometheus.GaugeValue, Name: "pool_size", Help: ""},
		{Type: prometheus.GaugeValue, Name: "reserve_pool", Help: ""},
		{Type: prometheus.GaugeValue, Name: "max_connections", Help: ""},
		{Type: prometheus.GaugeValue, Name: "current_connections", Help: ""},
		{Type: prometheus.GaugeValue, Name: "paused", Help: ""},
		{Type: prometheus.GaugeValue, Name: "disabled", Help: ""},
	},
}

var MetricDescriptorConfig = MetricDescriptor{
	Prefix: "config",
	Labels: []string{},
	MetricProps: []MetricProps{
		// TODO
	},
}
