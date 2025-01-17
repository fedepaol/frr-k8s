// SPDX-License-Identifier:Apache-2.0

package main

import (
	"flag"
	"fmt"
	stdlog "log"
	"net/http"
	"os"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/metallb/frrk8s/frr-tools/metrics/collector"
	"github.com/metallb/frrk8s/frr-tools/metrics/liveness"
	"github.com/metallb/frrk8s/frr-tools/metrics/vtysh"
	"github.com/metallb/frrk8s/internal/logging"
	"github.com/metallb/frrk8s/internal/version"
)

var (
	metricsPort        = flag.Uint("metrics-port", 7573, "Port to listen on for web interface.")
	metricsBindAddress = flag.String("metrics-bind-address", "127.0.0.1", "The address the metric endpoint binds to")
	metricsPath        = flag.String("metrics-path", "/metrics", "Path under which to expose metrics.")
)

func metricsHandler(logger log.Logger) http.Handler {
	BGPCollector := collector.NewBGP(logger)
	BFDCollector := collector.NewBFD(logger)

	registry := prometheus.NewRegistry()
	registry.MustRegister(BGPCollector)
	registry.MustRegister(BFDCollector)

	gatherers := prometheus.Gatherers{
		prometheus.DefaultGatherer,
		registry,
	}

	handlerOpts := promhttp.HandlerOpts{
		ErrorLog:      stdlog.New(log.NewStdlibAdapter(level.Error(logger)), "", 0),
		ErrorHandling: promhttp.ContinueOnError,
		Registry:      registry,
	}

	return promhttp.HandlerFor(gatherers, handlerOpts)
}

func main() {
	flag.Parse()

	logger, err := logging.Init("error")
	if err != nil {
		fmt.Printf("failed to initialize logging: %s\n", err)
		os.Exit(1)
	}

	level.Info(logger).Log("version", version.Version(), "commit", version.CommitHash(), "branch", version.Branch(), "goversion", version.GoString(), "msg", "FRR metrics exporter starting "+version.String())

	mux := http.NewServeMux()
	mux.Handle(*metricsPath, metricsHandler(logger))
	mux.Handle("/livez", liveness.Handler(vtysh.Run, logger))
	level.Info(logger).Log("msg", "Starting exporter", "metricsPath", metricsPath, "port", metricsPort)

	srv := &http.Server{
		Addr:        fmt.Sprintf("%s:%d", *metricsBindAddress, *metricsPort),
		ReadTimeout: 3 * time.Second,
		Handler:     mux,
	}

	if err := srv.ListenAndServe(); err != nil {
		level.Error(logger).Log("error", err)
		os.Exit(1)
	}
}
