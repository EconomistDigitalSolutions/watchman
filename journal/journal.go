package journal

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	stdlog "log"

	kitlog "github.com/go-kit/kit/log"
)

var (
	logger  kitlog.Logger
	Service string
)

func init() {
	logger = kitlog.NewJSONLogger(os.Stdout)
	stdlog.SetOutput(kitlog.NewStdlibAdapter(logger))
}

// SetLogger allows user to output log data to a writer.
func SetLogger(w io.Writer) {
	logger = kitlog.NewJSONLogger(w)
	stdlog.SetOutput(kitlog.NewStdlibAdapter(logger))
}

// SetLogFile allows the user to output log data to a file.
func SetLogFile(file string) {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		LogError(fmt.Sprintf("error opening log file: %v", err))
	}
	SetLogger(f)
}

// LogRequest logs details of an HTTP request.
func LogRequest(r *http.Request) {
	logger.Log("channel", "request", "service", Service, "method", r.Method, "url", r.URL.String(), "headers", r.Header, "ts", time.Now())
}

// LogChannel logs data to a log channel.
func LogChannel(channel string, message ...interface{}) {
	logger.Log("channel", channel, "service", Service, "message", message, "ts", time.Now())
}

// LogError logs error data to the error channel
func LogError(message string) {
	LogChannel("error", message)
}

// LogInfo logs informational data.
func LogInfo(message string) {
	LogChannel("information", message)
}

// LogWorker logs information around queue worker operations.
func LogWorker(message ...interface{}) {
	LogChannel("worker", message)
}
