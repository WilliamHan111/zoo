package prometheus

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	allMetrics = make(map[string]*prometheus.GaugeVec)
	mutex      sync.Mutex
)

type Metric struct {
	Name     string
	Help     string
	Duration int
	Lables   MetricLables
	Value    float64
}

type MetricLables struct {
	Names  []string
	Values []string
}

// 注册指标
func registerMetric(name, help string, lables []string) {
	mutex.Lock()
	defer mutex.Unlock()

	if _, ok := allMetrics[name]; !ok {
		gaugeVec := prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: name,
			Help: help,
		}, lables)
		allMetrics[name] = gaugeVec
		prometheus.MustRegister(gaugeVec)
	}
}

// 取消注册指标
func unregisterMetric(name string) {
	mutex.Lock()
	defer mutex.Unlock()

	if gaugeVec, ok := allMetrics[name]; ok {
		prometheus.Unregister(gaugeVec)
		delete(allMetrics, name)
	}
}

// 指标更新
func updateMetric(name string, labels []string, value float64) {
	mutex.Lock()
	defer mutex.Unlock()

	if gaugeVec, ok := allMetrics[name]; ok {
		gaugeVec.WithLabelValues(labels...).Set(value)
	}
}
