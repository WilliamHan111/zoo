package prometheus

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 车辆类型：小汽车、卡车、巴士、摩托车、三轮车
var userType = []string{
	"员工",
	"游客",
}

// 动物园区域
var userLocation = []string{
	"猴子区",
	"大象区",
	"休息区",
	"大门",
	"停车场",
	"老虎区",
}

var userMetric = &Metric{
	Name:     "user",
	Help:     "This is an user metric.",
	Duration: 10,
	Lables: MetricLables{
		Names:  []string{"name", "type", "location"},
		Values: []string{},
	},
	Value: 0,
}

func LoadUserMetrics() {
	// 注册自定义指标并赋初值
	registerMetric(userMetric.Name, userMetric.Help, userMetric.Lables.Names)
	registerMetric("fee", "fee", []string{"id"})
	for {
		//获取值
		userMetric.Lables.Values = []string{randomString(5), newUserType(), newUserLocation()}
		userMetric.Value = float64(rand.Intn(500))
		//赋值
		for name := range allMetrics {
			if name == userMetric.Name {
				updateMetric(name, userMetric.Lables.Values, userMetric.Value)
				updateMetric("fee", []string{randomString(5)}, newFeeValue()+10)
			}
		}
		time.Sleep(time.Duration(userMetric.Duration) * time.Second)
	}

}

// 生成指定长度的随机字符串
func randomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func newUserType() string {
	rand.Seed(time.Now().UnixNano())
	return userType[rand.Intn(2)]
}

func newUserLocation() string {
	rand.Seed(time.Now().UnixNano())
	return userLocation[rand.Intn(6)]
}
