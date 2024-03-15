package prometheus

import (
	"math/rand"
	"time"
)

var animalType = []string{
	"猴子",
	"大象",
	"长颈鹿",
	"狮子",
	"老虎",
	"野狼",
}

var animalMetric = &Metric{
	Name:     "animal",
	Help:     "This is an animal metric.",
	Duration: 10,
	Lables: MetricLables{
		Names:  []string{"type"},
		Values: []string{},
	},
	Value: 0,
}

func LoadAnimalMetrics() {
	// 注册自定义指标并赋初值
	registerMetric(animalMetric.Name, animalMetric.Help, animalMetric.Lables.Names)
	for {
		//获取值
		rand.Seed(time.Now().UnixNano())
		randomType := rand.Intn(5)
		animalMetric.Lables.Values = []string{animalType[randomType]}
		animalMetric.Value = float64(rand.Intn(100))
		//赋值
		for name := range allMetrics {
			if name == animalMetric.Name {
				updateMetric(name, animalMetric.Lables.Values, animalMetric.Value)
			}
		}
		time.Sleep(time.Duration(animalMetric.Duration) * time.Second)
	}
}
