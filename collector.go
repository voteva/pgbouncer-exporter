package main

import (
	"database/sql"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

type MetricDesc struct {
	Type prometheus.ValueType
	Desc prometheus.Desc
}

type Collector struct {
	db *sql.DB
	rw sync.Mutex

	// internal state
	up           prometheus.Gauge
	totalScrapes prometheus.Counter

	// metrics
	namespace             string
	metricDescMapInternal map[string]*MetricDesc
	metricDescMapList     map[string]*MetricDesc
}

func NewCollector(db *sql.DB, namespace string) *Collector {
	return &Collector{
		db:                db,
		namespace:         namespace,
		up:                buildGauge(InternalMetricUp),
		totalScrapes:      buildCounter(InternalMetricScrapeTotal),
		metricDescMapList: buildMetricDescMapList(),
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
	ch <- c.totalScrapes
}

func (c *Collector) scrape(ch chan<- prometheus.Metric) {
	c.rw.Lock()
	defer c.rw.Unlock()

	c.up.Set(1)
	c.totalScrapes.Inc()

	if err := c.scrapeLists(ch); err != nil {
		c.up.Set(0)
	}
}

func (c *Collector) scrapeLists(ch chan<- prometheus.Metric) error {
	rows, err := c.db.Query(`SHOW LISTS;`)
	if err != nil {
		return err
	}
	defer rows.Close()

	p, _ := rows.Columns()
	nColumn := len(p)

	columnData := make([]interface{}, nColumn)
	scanArgs := make([]interface{}, nColumn)
	for i := 0; i < nColumn; i++ {
		scanArgs[i] = &columnData[i]
	}

	listResult := make(map[string]float64)
	for rows.Next() {
		if err = rows.Scan(scanArgs...); err != nil {
			return err
		}
		listResult[Cast2string(columnData[0])] = Cast2Float64(columnData[1])
	}

	for k, v := range listResult {
		metricDesc := c.metricDescMapList[k]
		if metricDesc != nil {
			ch <- prometheus.MustNewConstMetric(&metricDesc.Desc, metricDesc.Type, v)
		}
	}
	return nil
}

func buildMetricDescMapList() map[string]*MetricDesc {
	prefix := "lists"
	m := make(map[string]*MetricDesc)

	for _, v := range MetricPropsList {
		m[v.Name] = &MetricDesc{
			Type: v.Type,
			Desc: *prometheus.NewDesc(fmt.Sprintf("%s_%s_%s", namespace, prefix, v.Name), v.Help, nil, nil),
		}
	}
	return m
}

func buildGauge(props MetricProps) prometheus.Gauge {
	return prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      props.Name,
		Help:      props.Help,
	})
}

func buildCounter(props MetricProps) prometheus.Counter {
	return prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      props.Name,
		Help:      props.Help,
	})
}
