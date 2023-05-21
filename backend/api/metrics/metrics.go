package metrics

import (
	"runtime"
	"time"

	"github.com/labstack/echo-contrib/prometheus"
	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type Metrics struct {
	functionCallDuration *prometheus.Metric
}

func NewMetrics() *Metrics {
	return &Metrics{
		functionCallDuration: &prometheus.Metric{
			Name:        "function_call_duration",
			Description: "Duration for function call",
			Type:        "summary_vec",
			Args:        []string{"function_name"},
		},
	}
}

func (m *Metrics) MetricList() []*prometheus.Metric {
	return []*prometheus.Metric{
		m.functionCallDuration,
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

var M = NewMetrics()
