package prometheus

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Run() {
	rand.Seed(time.Now().UnixNano())
	// 注册自定义指标并赋初值
	registerMetric("example_custom_metric", "This is an example custom metric.")
	go func() {
		for {
			for name := range allMetrics {
				updateMetric(name, "lable", rand.Float64()*100)
			}
			time.Sleep(1 * time.Second)
		}
	}()
	http.HandleFunc("/config", handleConfig)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":18080", nil)
}
func handleConfig(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	action := r.Form.Get("action")
	label := r.Form.Get("label")
	value, _ := strconv.ParseFloat(r.Form.Get("value"), 64)

	if name == "" {
		http.Error(w, "Missing 'name' parameter", http.StatusBadRequest)
		return
	}

	switch action {
	case "register":
		help := r.Form.Get("help")
		if help == "" {
			http.Error(w, "Missing 'help' parameter", http.StatusBadRequest)
			return
		}
		registerMetric(name, help)
		fmt.Fprintf(w, "Metric '%s' registered.\n", name)
	case "unregister":
		unregisterMetric(name)
		fmt.Fprintf(w, "Metric '%s' unregistered.\n", name)
	case "update":
		if label == "" {
			http.Error(w, "Missing 'label' parameter", http.StatusBadRequest)
			return
		}
		updateMetric(name, label, value)
		fmt.Fprintf(w, "Metric '%s' updated with label '%s' and value %f.\n", name, label, value)
	default:
		http.Error(w, "Invalid 'action' parameter", http.StatusBadRequest)
		return
	}
}
