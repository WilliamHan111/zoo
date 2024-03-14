package prometheus

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	allMetrics = make(map[string]*prometheus.GaugeVec)
	mutex      sync.Mutex
)

func registerMetric(name, help string) {
	mutex.Lock()
	defer mutex.Unlock()

	if _, ok := allMetrics[name]; !ok {
		gaugeVec := prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: name,
			Help: help,
		}, []string{"label"})
		allMetrics[name] = gaugeVec
		prometheus.MustRegister(gaugeVec)
	}
}

func unregisterMetric(name string) {
	mutex.Lock()
	defer mutex.Unlock()

	if gaugeVec, ok := allMetrics[name]; ok {
		prometheus.Unregister(gaugeVec)
		delete(allMetrics, name)
	}
}

func updateMetric(name, label string, value float64) {
	mutex.Lock()
	defer mutex.Unlock()

	if gaugeVec, ok := allMetrics[name]; ok {
		gaugeVec.WithLabelValues(label).Set(value)
	}
}
