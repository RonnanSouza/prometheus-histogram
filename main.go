package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	histogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "example_histogram",
		Help:    "Just a histogram example from 0 - 10",
		Buckets: prometheus.LinearBuckets(1, 1, 10),
	})
)

func init() {
	prometheus.MustRegister(histogram)
}

func main() {

	go func() {
		for {
			fmt.Println("type a value from 1 to 10")
			var metric float64
			fmt.Print("-> ")
			fmt.Scanf("%f", &metric)
			histogram.Observe(metric)
		}
	}()

	// Expose the registered metrics via HTTP.
	http.Handle("/metrics", promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			// Opt into OpenMetrics to support exemplars.
			EnableOpenMetrics: true,
		},
	))
	log.Fatal(http.ListenAndServe(":9091", nil))
}
