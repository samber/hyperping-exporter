package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"runtime"

	"github.com/alecthomas/kingpin/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/samber/hyperping_exporter/exporter"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"

	hyperpingToken = kingpin.Flag("hyperping.token", "Hyperping token").Envar("HYPERPING_TOKEN").Required().String()
	namespace      = kingpin.Flag("namespace", "Namespace for metrics").Envar("HYPERPING_EXPORTER_NAMESPACE").Default("hyperping").String()
	listenAddress  = kingpin.Flag("web.listen-address", "Address to listen on for web interface and telemetry.").Envar("HYPERPING_EXPORTER_WEB_LISTEN_ADDRESS").Default(":9312").String()
	metricPath     = kingpin.Flag("web.telemetry-path", "Path under which to expose metrics.").Envar("HYPERPING_EXPORTER_WEB_TELEMETRY_PATH").Default("/metrics").String()
	logFormat      = kingpin.Flag("log.format", "Log format, valid options are txt and json").Envar("HYPERPING_EXPORTER_LOG_FORMAT").Default("txt").String()
	// logLevel       = kingpin.Flag("log.level", "Log level").Envar("HYPERPING_EXPORTER_LOG_FORMAT").Default("debug").String()
)

func main() {
	kingpin.Version(version)
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	switch *logFormat {
	case "json":
		slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, nil)))
	default:
		slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, nil)))
	}

	slog.Info("Hyperping Metrics Exporter",
		slog.Any("build.time", date),
		slog.Any("build.release", version),
		slog.Any("build.commit", commit),
		slog.Any("go.version", runtime.Version()),
		slog.Any("go.os", runtime.GOOS),
		slog.Any("go.arch", runtime.GOARCH))

	exporter := exporter.NewExporter(
		*hyperpingToken,
		*namespace,
	)

	slog.Info(fmt.Sprintf("Providing metrics at %s%s", *listenAddress, *metricPath))

	registerer := prometheus.NewRegistry()
	registerer.MustRegister(exporter)

	http.Handle(*metricPath, promhttp.HandlerFor(registerer, promhttp.HandlerOpts{}))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//nolint:errcheck
		w.Write([]byte(`<html>
			<head><title>HyperPing Exporter</title></head>
			<body>
			<h1>HyperPing Exporter</h1>
			<p><a href="` + *metricPath + `">Metrics</a></p>
			</body>
			</html>`))
	})

	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
