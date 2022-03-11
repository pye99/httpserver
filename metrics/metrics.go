package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "cloudnative"
	subsystem = "httpserver"
)

var RequestsCost = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "requests_cost_seconds",
		Help:      "request(ms) cost seconds",
	},
	[]string{"method", "path"},
)

func init() {
	prometheus.MustRegister(RequestsCost)
}
