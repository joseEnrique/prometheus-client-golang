package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var responseMetric = prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Name:    "request_duration_milliseconds",
		Help:    "Request latency distribution",
		Buckets: prometheus.ExponentialBuckets(10.0, 1.13, 40),
	})

var (
	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current temperature of the CPU.",
	})
	hdFailures = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "hd_errors_total",
			Help: "Number of hard-disk errors.",
		},
		[]string{"device"},
	)
)

func main() {
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(hdFailures)
	prometheus.MustRegister(responseMetric)

	cpuTemp.Set(65.3)
	hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()

	http.Handle("/metrics", prometheus.Handler())
	http.ListenAndServe(":8080", nil)
	// Any other setup, then an http.ListenAndServe here
}
