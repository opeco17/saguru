package metrics

import (
	"runtime"
	"strconv"
	"time"

	"github.com/labstack/echo-contrib/prometheus"
	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type Metrics struct {
	functionCallDuration *prometheus.Metric
	cacheAccess          *prometheus.Metric
}

func NewMetrics() *Metrics {
	return &Metrics{
		functionCallDuration: &prometheus.Metric{
			Name:        "function_call_duration",
			Description: "Duration for function call",
			Type:        "summary_vec",
			Args:        []string{"function_name"},
		},
		cacheAccess: &prometheus.Metric{
			Name:        "cache_access",
			Description: "Cache access count",
			Type:        "counter_vec",
			Args:        []string{"key", "hit"},
		},
	}
}

func (m *Metrics) MetricList() []*prometheus.Metric {
	return []*prometheus.Metric{
		m.functionCallDuration,
		m.cacheAccess,
	}
}

func (m *Metrics) ObservefunctionCallDuration(since time.Time) {
	pc, _, _, ok := runtime.Caller(1)
	if ok != true {
		logrus.Warn("Failed to measure function call duration.")
		return
	}
	labels := prom.Labels{"function_name": runtime.FuncForPC(pc).Name()}
	m.functionCallDuration.MetricCollector.(*prom.SummaryVec).With(labels).Observe(time.Since(since).Seconds())
}

func (m *Metrics) CountCacheAccess(key string, hit bool) {
	labels := prom.Labels{"key": key, "hit": strconv.FormatBool(hit)}
	m.cacheAccess.MetricCollector.(*prom.CounterVec).With(labels).Inc()
}

var M = NewMetrics()
