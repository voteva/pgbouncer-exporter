package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"sync"
	"time"
)

const MaxErrors = 5

type MetricGroup struct {
	Labels  []string
	Metrics map[string]*MetricDesc
}

type MetricDesc struct {
	Type prometheus.ValueType
	Desc prometheus.Desc
}

type Collector struct {
	db SQLDB
	rw sync.Mutex

	// internal state
	up             prometheus.Gauge
	errors         prometheus.Gauge
	scrapeLastTime prometheus.Gauge
	totalScrapes   prometheus.Counter

	// metrics
	namespace            string
	metricGroupLists     *MetricGroup
	metricGroupStats     *MetricGroup
	metricGroupPools     *MetricGroup
	metricGroupDatabases *MetricGroup
	metricGroupConfig    *MetricGroup
}

func NewCollector(db SQLDB, namespace string) *Collector {
	return &Collector{
		db:                   db,
		namespace:            namespace,
		up:                   prometheus.NewGauge(buildGaugeOpts(InternalMetricUp)),
		errors:               prometheus.NewGauge(buildGaugeOpts(InternalMetricErrors)),
		scrapeLastTime:       prometheus.NewGauge(buildGaugeOpts(InternalMetricScrapeLastTime)),
		totalScrapes:         prometheus.NewCounter(buildCounterOpts(InternalMetricScrapeTotal)),
		metricGroupLists:     buildMetricGroup(MetricDescriptorLists),
		metricGroupStats:     buildMetricGroup(MetricDescriptorStats),
		metricGroupPools:     buildMetricGroup(MetricDescriptorPools),
		metricGroupDatabases: buildMetricGroup(MetricDescriptorDatabases),
		metricGroupConfig:    buildMetricGroup(MetricDescriptorConfig),
	}
}

func (c *Collector) Close() {
	c.rw.Lock()
	defer c.rw.Unlock()
	c.db.Close()
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	metricCh := make(chan prometheus.Metric)
	doneCh := make(chan struct{})

	go func() {
		for m := range metricCh {
			ch <- m.Desc()
		}
		close(doneCh)
	}()

	c.Collect(metricCh)
	close(metricCh)
	<-doneCh
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	c.scrape(ch)
	ch <- c.up
	ch <- c.errors
	ch <- c.scrapeLastTime
	ch <- c.totalScrapes
}

func (c *Collector) scrape(ch chan<- prometheus.Metric) {
	c.rw.Lock()
	defer c.rw.Unlock()

	errors := 0
	c.up.Set(1)
	c.scrapeLastTime.Set(Cast2Float64(time.Now()))
	c.totalScrapes.Inc()

	metrics, err := c.extractMetrics("SHOW LISTS;", c.metricGroupLists, extractKeyValue)
	if err = c.handleExtractedMetrics(ch, metrics, err); err != nil {
		log.Error("Failed to extract metrics LISTS: ", err)
		errors++
	}

	metrics, err = c.extractMetrics("SHOW STATS;", c.metricGroupStats, extractWithLabels)
	if err = c.handleExtractedMetrics(ch, metrics, err); err != nil {
		log.Error("Failed to extract metrics STATS: ", err)
		errors++
	}

	metrics, err = c.extractMetrics("SHOW POOLS;", c.metricGroupPools, extractWithLabels)
	if err = c.handleExtractedMetrics(ch, metrics, err); err != nil {
		log.Error("Failed to extract metrics POOLS: ", err)
		errors++
	}

	metrics, err = c.extractMetrics("SHOW DATABASES;", c.metricGroupDatabases, extractWithLabels)
	if err = c.handleExtractedMetrics(ch, metrics, err); err != nil {
		log.Error("Failed to extract metrics DATABASES: ", err)
		errors++
	}

	metrics, err = c.extractMetrics("SHOW CONFIG;", c.metricGroupConfig, extractKeyValue)
	if err = c.handleExtractedMetrics(ch, metrics, err); err != nil {
		log.Error("Failed to extract metrics CONFIG: ", err)
		errors++
	}

	c.errors.Set(float64(errors))
	if errors >= MaxErrors {
		c.up.Set(0)
	}
}

func (c *Collector) handleExtractedMetrics(ch chan<- prometheus.Metric, metrics []prometheus.Metric, err error) error {
	if err != nil {
		return err
	}
	for _, m := range metrics {
		ch <- m
	}
	return nil
}

func (c *Collector) extractMetrics(query string, metricGroup *MetricGroup, extractFunc ExtractFunc) ([]prometheus.Metric, error) {
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	nColumn := len(columns)

	columnData := make([]interface{}, nColumn)
	scanArgs := make([]interface{}, nColumn)
	for i := 0; i < nColumn; i++ {
		scanArgs[i] = &columnData[i]
	}

	var resultMetrics []prometheus.Metric
	for rows.Next() {
		if err = rows.Scan(scanArgs...); err != nil {
			return nil, err
		}

		metrics := extractFunc(metricGroup, columns, columnData)
		resultMetrics = append(resultMetrics, metrics...)
	}

	return resultMetrics, nil
}

type ExtractFunc func(metricGroup *MetricGroup, columns []string, columnData []interface{}) []prometheus.Metric

func extractKeyValue(metricGroup *MetricGroup, columns []string, columnData []interface{}) []prometheus.Metric {
	var result []prometheus.Metric
	metricValue := Cast2Float64(columnData[1])
	metricDesc := metricGroup.Metrics[Cast2string(columnData[0])]
	if metricDesc != nil {
		result = append(result, prometheus.MustNewConstMetric(&metricDesc.Desc, metricDesc.Type, metricValue))
	}
	return result
}

func extractWithLabels(metricGroup *MetricGroup, columns []string, columnData []interface{}) []prometheus.Metric {
	var result []prometheus.Metric
	var labelValues []string

	// collect labels
	for i, colName := range columns {
		if Contains(metricGroup.Labels, colName) {
			labelValues = append(labelValues, Cast2string(columnData[i]))
			continue
		}
	}
	// collect metrics
	for i, colName := range columns {
		if Contains(metricGroup.Labels, colName) {
			continue
		}
		metricDesc := metricGroup.Metrics[colName]
		if metricDesc != nil {
			metricValue := Cast2Float64(columnData[i])
			result = append(result, prometheus.MustNewConstMetric(&metricDesc.Desc, metricDesc.Type, metricValue, labelValues...))
		}
	}
	return result
}

func buildMetricGroup(descriptor MetricDescriptor) *MetricGroup {
	m := make(map[string]*MetricDesc)

	for _, v := range descriptor.MetricProps {
		m[v.Name] = &MetricDesc{
			Type: v.Type,
			Desc: *prometheus.NewDesc(fmt.Sprintf("%s_%s_%s", namespace, descriptor.Prefix, v.Name), v.Help, descriptor.Labels, nil),
		}
	}
	return &MetricGroup{
		Labels:  descriptor.Labels,
		Metrics: m,
	}
}

func buildGaugeOpts(props MetricProps) prometheus.GaugeOpts {
	return prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      props.Name,
		Help:      props.Help,
	}
}

func buildCounterOpts(props MetricProps) prometheus.CounterOpts {
	return prometheus.CounterOpts{
		Namespace: namespace,
		Name:      props.Name,
		Help:      props.Help,
	}
}
