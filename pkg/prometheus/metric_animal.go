package prometheus

import (
	"math/rand"
	"time"
)

const (
	animalMetric         = "animal"
	animalMetricHelp     = "This is an animal metric."
	animalMetricDuration = 1
)

func LoadAnimalMetrics() {
	// 注册自定义指标并赋初值
	registerMetric(animalMetric, animalMetricHelp)
	for {
		for name := range allMetrics {
			if name == animalMetric {
				updateMetric(name, "lable", rand.Float64()*100)
			}
		}
		time.Sleep(animalMetricDuration * time.Second)
	}
}
