package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/gocraft/health"
	//"os"
	"time"
	"fmt"
)

const (
	defaultStreamKey = "gin-health"
)

/**
	go get github.com/gocraft/health/cmd/healthd
	HEALTHD_MONITORED_HOSTPORTS=localhost:5020 HEALTHD_SERVER_HOSTPORT=:5032 ./healthd

	go get github.com/gocraft/health/cmd/healthtop
	./healthtop jobs
 */

func Stats() gin.HandlerFunc {
	var stream = health.NewStream()
	stream.EventKv("starting_app", health.Kvs{"app": "store"})
	//stream.AddSink(&health.WriterSink{os.Stdout})

	sink := health.NewJsonPollingSink(time.Minute, 5*time.Minute)
	stream.AddSink(sink)
	sink.StartServer("0.0.0.0:5020")

	return func(c *gin.Context) {
		c.Set(defaultStreamKey, stream)
		c.Next()
	}
}

/**
// Timings:
startTime := time.Now()
// Do something...
job.Timing("fetch_user", time.Since(startTime).Nanoseconds()) // NOTE: Nanoseconds!

// Timings also support keys/values:
job.TimingKv("fetch_user", time.Since(startTime).Nanoseconds(),
	health.Kvs{"user_email": userEmail})
 */
func Instrument(metricPath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		stream := c.MustGet(defaultStreamKey).(*health.Stream)
		job := stream.NewJob(metricPath)
		startTime := time.Now()
		c.Next()
		job.Timing(metricPath, time.Since(startTime).Nanoseconds())

		if c.IsAborted() {
			for _,err := range c.Errors {
				stream.EventErr(fmt.Sprintf("Panic at %v", c.Request.RequestURI), err)
			}

			job.Complete(health.Error)
		} else {
			job.Complete(health.Success)
		}
	}
}
