package prometheus

import (
	"math/rand"
	"time"
)

var carType = []string{
	"汽车",
}

type CarMetric struct {
	Name     string
	Help     string
	Duration int
	Lables   CarMetricLables
	Value    float64
}

type CarMetricLables struct {
	Names  []string
	Values []string
}

var carMetric = &CarMetric{
	Name:     "car",
	Help:     "This is an car metric.",
	Duration: 10,
	Lables: CarMetricLables{
		Names:  []string{"license", "io", "type"},
		Values: []string{},
	},
	Value: 0,
}

func LoadCarMetrics() {
	// 注册自定义指标并赋初值
	registerMetric(carMetric.Name, carMetric.Help, carMetric.Lables.Names)
	for {
		//获取值
		rand.Seed(time.Now().UnixNano())
		randomType := rand.Intn(5)
		carMetric.Lables.Values = []string{animalType[randomType]}
		carMetric.Value = float64(rand.Intn(100))
		//赋值
		for name := range allMetrics {
			if name == carMetric.Name {
				updateMetric(name, carMetric.Lables.Values, carMetric.Value)
			}
		}
		time.Sleep(time.Duration(carMetric.Duration) * time.Second)
	}
}
