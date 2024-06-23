package metrics

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/opsdata-io/opsdata/pkg/config"
	"github.com/opsdata-io/opsdata/pkg/health"
	"github.com/opsdata-io/opsdata/pkg/logging"
	"github.com/opsdata-io/opsdata/pkg/version"
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
	mux.HandleFunc("/healthz", health.HealthzHandler())
	mux.HandleFunc("/readyz", health.ReadyzHandler())
	mux.HandleFunc("/version", version.GetVersion) // Use version.GetVersion for version endpoint

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
