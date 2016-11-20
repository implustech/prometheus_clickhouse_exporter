package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/roistat/go-clickhouse"
	"github.com/spf13/viper"
	"implus.co/prometheus_clickhouse_exporter/cmd"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
)

type ClickhouseMetrics struct {
	Name      string
	Counter   prometheus.Counter
	LastValue float64
}

func NameToUnderscore(src string) string {
	re := regexp.MustCompile(`(\w)([A-Z])`)
	return strings.ToLower(re.ReplaceAllString(src, "${1}_${2}"))
}

func main() {

	if err := cmd.RootCmd.Execute(); err != nil {
		logger.Println(err)
		os.Exit(-1)
	}

	initLogger()

	clickhouseConnection := viper.GetString("clickhouse.connection")
	logger.Info("clickhouse.connection ", clickhouseConnection)
	logger.Info("production ", viper.GetBool("production"))
	logger.Info("listen ", viper.GetString("listen"))

	clickhouseConnectionErrorsTotal := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: viper.GetString("prometheus.namespace"),
		Subsystem: viper.GetString("prometheus.subsystem"),
		Name:      "clickhouse_connection_errors_total",
		Help:      "clickhouse connection error counter",
	})

	prometheus.MustRegister(clickhouseConnectionErrorsTotal)
	prometheusHandler := prometheus.Handler()

	clickhouseMetrics := make(map[string]*ClickhouseMetrics)
	mutex := sync.Mutex{}

	mux := http.NewServeMux()
	mux.HandleFunc("/metrics", func(w http.ResponseWriter, req *http.Request) {
		mutex.Lock()
		defer mutex.Unlock()

		conn := clickhouse.NewConn(clickhouseConnection, clickhouse.NewHttpTransport())
		query := clickhouse.NewQuery("select * from system.events")
		iter := query.Iter(conn)
		if iter.Error() != nil {
			clickhouseConnectionErrorsTotal.Inc()
		}

		var name string
		var value float64

		for iter.Scan(&name, &value) {
			underscoreName := NameToUnderscore(name)
			metrics, ok := clickhouseMetrics[underscoreName]
			if !ok {
				metrics = &ClickhouseMetrics{}
				metrics.Counter = prometheus.NewCounter(prometheus.CounterOpts{
					Namespace: viper.GetString("prometheus.namespace"),
					Subsystem: viper.GetString("prometheus.subsystem"),
					Name:      underscoreName,
					Help:      name,
				})
				prometheus.MustRegister(metrics.Counter)
				clickhouseMetrics[underscoreName] = metrics
			}
			metrics.Counter.Add(value - metrics.LastValue)
			metrics.LastValue = value
		}

		prometheusHandler.ServeHTTP(w, req)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Clickhouse Exporter</title></head>
             <body>
             <h1>Clickhouse Exporter</h1>
             <p><a href='/metrics'>Metrics</a></p>
             </body>
             </html>`))
	})

	logger.Fatalln(http.ListenAndServe(viper.GetString("listen"), mux))

}
