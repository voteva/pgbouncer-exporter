package main

import "github.com/prometheus/client_golang/prometheus"

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

var MetricPropsList = []MetricProps{
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
}

var MetricPropsStatTotal = []MetricProps{

}
