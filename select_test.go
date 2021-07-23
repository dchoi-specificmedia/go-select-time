package worker

import (
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

func BenchmarkSelect(b *testing.B) {
	obj := map[float64]float64{0.0: 0.1, 0.025: 0.01, 0.16: 0.01, 0.5: 0.01, 0.84: 0.01, 0.95: 0.01, 0.975: 0.01, 1.0: 0.1}

	waitSummary := prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "wait",
		Objectives: obj,
	})

	excessWaitSummary := prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "excess_wait",
		Objectives: obj,
	})

	skipSummary := prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "skip",
		Objectives: obj,
	})

	prometheus.Register(waitSummary)
	prometheus.Register(excessWaitSummary)
	prometheus.Register(skipSummary)

	waitMicrosStr := os.Getenv("WAIT_MICROS")
	waitMicros := 500
	if waitMicrosStr != "" {
		waitMicro, err := strconv.Atoi(waitMicrosStr)
		if err != nil {
			log.Fatal(err)
		}
		waitMicros = waitMicro
	}

	b.Run("test", func(b *testing.B) {
		dur := time.Duration(time.Duration(waitMicros) * time.Microsecond)

		for i := 0; i < b.N; i++ {
			start := time.Now()
			select {
			case <-time.After(dur):
			default:
				skipSummary.Observe(float64(time.Since(start).Microseconds()))
			}

			start = time.Now()
			select {
			case <-time.After(dur):
				waitSummary.Observe(float64(time.Since(start).Microseconds()))
			}

			excessWaitSummary.Observe(float64((time.Since(start) - dur).Microseconds()))
		}
	})

	prom := prometheus.DefaultGatherer
	prometheus.WriteToTextfile("./select.prom.txt", prom)
}
