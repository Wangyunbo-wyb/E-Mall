package orm

import "github.com/zeromicro/go-zero/core/metric"

// define metric
const gormNamespace = "gorm_client"

// define two metrics for monitoring GORM client requests, helping developers monitor and analyze the performance and error conditions of GORM clients
var (
	// metricClientReqDur is used to monitor the duration of requests(ms)
	metricClientReqDur = metric.NewHistogramVec(&metric.HistogramVecOpts{
		Namespace: gormNamespace,
		Subsystem: "requests",
		Name:      "duration_ms",
		Help:      "gorm client requests duration(ms).",
		Labels:    []string{"table", "method"},
		Buckets:   []float64{5, 10, 25, 50, 100, 250, 500, 1000},
	})

	// metricClientReqErrTotal is used to monitor the number of errors in requests
	metricClientReqErrTotal = metric.NewCounterVec(&metric.CounterVecOpts{
		Namespace: gormNamespace,
		Subsystem: "requests",
		Name:      "error_total",
		Help:      "gorm client requests error count.",
		Labels:    []string{"table", "method", "is_error"},
	})
)
