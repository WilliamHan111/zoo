package prometheus

import (
	"fmt"
	"math/rand"
	"time"
)

// 车辆类型：小汽车、卡车、巴士、摩托车、三轮车
var carType = []string{
	"小汽车",
	"卡车",
	"巴士",
	"摩托车",
	"三轮车",
}

// 停车场区域：A B C D E
var carLocation = []string{
	"A区",
	"B区",
	"C区",
	"D区",
	"E区",
}

var carMetric = &Metric{
	Name:     "car",
	Help:     "This is an car metric.",
	Duration: 10,
	Lables: MetricLables{
		Names:  []string{"plat", "type", "location"},
		Values: []string{},
	},
	Value: 0,
}

func LoadCarMetrics() {
	// 注册自定义指标并赋初值
	registerMetric(carMetric.Name, carMetric.Help, carMetric.Lables.Names)
	registerMetric("car_fee", "car fee", []string{"plat"})
	for {
		//获取值
		carMetric.Lables.Values = []string{newPlateNumber(), newType(), newLocation()}
		carMetric.Value = newValue()
		//赋值
		for name := range allMetrics {
			if name == carMetric.Name {
				updateMetric(name, carMetric.Lables.Values, carMetric.Value)
				updateMetric("car_fee", []string{newPlateNumber()}, newFeeValue())
			}
		}
		time.Sleep(time.Duration(carMetric.Duration) * time.Second)
	}

}
func newFeeValue() float64 {
	rand.Seed(time.Now().UnixNano())
	return float64(rand.Intn(100))
}
func newValue() float64 {
	return float64(1)
}

func newType() string {
	rand.Seed(time.Now().UnixNano())
	return carType[rand.Intn(5)]
}

var provinces = []string{"京", "沪", "津", "渝", "冀", "豫", "云", "辽", "黑", "湘", "皖", "鲁", "新", "苏", "浙", "赣", "鄂", "桂", "甘", "晋", "蒙", "陕", "吉", "闽", "贵", "粤", "青", "藏", "川", "宁", "琼"}

// 生成随机车牌号
func newPlateNumber() string {
	rand.Seed(time.Now().UnixNano())

	// 随机选择省份简称
	province := provinces[rand.Intn(len(provinces))]

	// 生成一个随机字母
	letter := string('A' + rand.Intn(26))

	// 生成五个随机数字
	numbers := make([]byte, 5)
	for i := range numbers {
		numbers[i] = byte('0' + rand.Intn(10))
	}

	// 组合车牌号
	return fmt.Sprintf("%s%s%s", province, letter, string(numbers))
}

func newLocation() string {
	rand.Seed(time.Now().UnixNano())
	return carLocation[rand.Intn(5)]
}
