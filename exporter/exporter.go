package exporter

import (
	"crypto/tls"
	"log/slog"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/samber/lo"
)

const HyperPingEndpoint = "https://api.hyperping.io/v1"

type Exporter struct {
	namespace string
	client    *http.Client
	token     string

	metricScrapeFailures prometheus.Counter
	metricMonitorActive  *prometheus.GaugeVec
	metricMonitorStatus  *prometheus.GaugeVec
}

var _ prometheus.Collector = (*Exporter)(nil)

func NewExporter(token string, namespace string) *Exporter {
	return &Exporter{
		namespace: namespace,
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{},
			},
			Timeout: 10 * time.Second,
		},
		token: token,

		metricScrapeFailures: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "scrape_failures_total",
			Help:      "Number of errors while scraping HyperPing.",
		}),
		metricMonitorActive: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "monitor_active",
				Help:      "Probe active (0==paused, 1==active)",
			},
			[]string{"project_id", "monitor_id", "name", "protocol"},
		),
		metricMonitorStatus: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "monitor_status",
				Help:      "Probe status (0==down, 1==up)",
			},
			[]string{"project_id", "monitor_id", "name", "protocol"},
		),
	}
}

// Describe describes all the metrics ever exported by the hyperping exporter. It
// implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.metricScrapeFailures.Desc()
	e.metricMonitorActive.Describe(ch)
	e.metricMonitorStatus.Describe(ch)
}

// implements prometheus.Collector.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	monitors, err := e.getMonitors()
	if err != nil {
		e.metricScrapeFailures.Inc()
		slog.Error(err.Error())
		return
	}

	for _, m := range monitors {
		e.metricMonitorActive.WithLabelValues(m.ProjectUUID, m.UUID, m.Name, m.Protocol).Set(lo.Ternary(!m.Paused, 1.0, 0.0))
		e.metricMonitorStatus.WithLabelValues(m.ProjectUUID, m.UUID, m.Name, m.Protocol).Set(lo.Ternary(m.Status == "up", 1.0, 0.0))
	}

	e.metricScrapeFailures.Collect(ch)
	e.metricMonitorActive.Collect(ch)
	e.metricMonitorStatus.Collect(ch)
}
