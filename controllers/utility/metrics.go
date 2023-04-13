package utility

import (
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

var EdgeUp = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "current_active_edge",
		Help: "Active Edge",
	},
)

var EdgeDown = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "current_inactive_edge",
		Help: "InActive Edge",
	},
)

func init() {
	// Register custom metrics with the global prometheus registry
	metrics.Registry.MustRegister(EdgeUp, EdgeDown)
}
