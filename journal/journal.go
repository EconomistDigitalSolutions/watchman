package journal

import (
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

// Simplified set of logging functions based on this:
// http://dave.cheney.net/2015/11/05/lets-talk-about-logging

// If you want to do weird logging in your own services by all
// means create convenience functions that delegate to LogInfo.

// LogRequest logs details of an HTTP request.
func LogRequest(r *http.Request) {
	logger.Log("channel", "request", "service", Service, "method", r.Method, "url", r.URL.String(), "headers", r.Header, "ts", time.Now())
}

// LogInfo logs informational data.
func LogInfo(message ...interface{}) {
	logger.Log("level", "INFO", "service", Service, "message", message, "ts", time.Now())
}

// LogDebug logs debug data.
func LogDebug(message ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		logger.Log("level", "DEBUG", "service", Service, "message", message, "ts", time.Now())
	}
}

// These functions are open for discussion and are being used in
// at least one service right now.

// LogError logs error data to the error channel, but allows some extra info to be passed along as a top level concern.
func LogErrorWithInfo(message string, infoPairs ...interface{}) {
	keyVals := append([]interface{}{"channel", "error", "service", Service, "ts", time.Now(), "message", message}, infoPairs...)
	logger.Log(keyVals...)
}

// LogEvent logs event information to the event channel.
func LogEvent(eventName string) {
	logger.Log("channel", "event", "service", Service, "event", eventName, "ts", time.Now())
}

// LogEvent logs event information to the event channel, but allows some extra info to be passed along as a top level concern.
func LogEventWithInfo(eventName string, infoPairs ...interface{}) {
	keyVals := append([]interface{}{"channel", "event", "service", Service, "ts", time.Now(), "event", eventName}, infoPairs...)
	logger.Log(keyVals...)
}
