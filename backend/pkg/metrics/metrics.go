package metrics

import (
	"net/http"
	"strconv"
	"time"

	"github.com/opsdata-io/opsdata/backend/pkg/config"
	"github.com/opsdata-io/opsdata/backend/pkg/logging"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var logger = logging.SetupLogging()

// Prometheus metrics
var (
	numberOfCustomers = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "number_of_customers",
			Help: "Number of customers in the database",
		},
		[]string{"status"},
	)
)

// Set up Prometheus metrics
func init() {
	logger.Debug("Initializing Prometheus metrics")
	prometheus.MustRegister(numberOfCustomers)
}

// StartMetricsServer starts the metrics server
func StartMetricsServer() {
	logger.Debug("Starting metrics server setup")
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	serverPortStr := strconv.Itoa(config.CFG.MetricsPort)
	logger.Printf("Metrics server starting on port %s\n", serverPortStr)

	srv := &http.Server{
		Addr:         ":" + serverPortStr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.Fatalf("Metrics server failed to start: %v", err)
	}
}
