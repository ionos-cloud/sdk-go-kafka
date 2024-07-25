/*
 * Kafka as a Service API
 *
 * An managed Apache Kafka cluster is designed to be highly fault-tolerant and scalable, allowing large volumes of data to be ingested, stored, and processed in real-time. By distributing data across multiple brokers, Kafka achieves high throughput and low latency, making it suitable for applications requiring real-time data processing and analytics.
 *
 * API version: 1.4.0
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"log"
	"os"
	"strings"
)

type LogLevel uint

func (l *LogLevel) Get() LogLevel {
	if l != nil {
		return *l
	}
	return Off
}

// Satisfies returns true if this LogLevel is at least high enough for v
func (l *LogLevel) Satisfies(v LogLevel) bool {
	return l.Get() >= v
}

const (
	Off LogLevel = 0x100 * iota
	Debug
	// Trace We recommend you only set this field for debugging purposes.
	// Disable it in your production environments because it can log sensitive data.
	// It logs the full request and response without encryption, even for an HTTPS call.
	// Verbose request and response logging can also significantly impact your application's performance.
	Trace
)

var LogLevelMap = map[string]LogLevel{
	"off":   Off,
	"debug": Debug,
	"trace": Trace,
}

// getLogLevelFromEnv - gets LogLevel type from env variable IONOS_LOG_LEVEL
// returns Off if an invalid log level is encountered
func getLogLevelFromEnv() LogLevel {
	strLogLevel := "off"
	if os.Getenv(IonosLogLevelEnvVar) != "" {
		strLogLevel = os.Getenv(IonosLogLevelEnvVar)
	}

	logLevel, ok := LogLevelMap[strings.ToLower(strLogLevel)]
	if !ok {
		log.Printf("Cannot set logLevel for value: %s, setting loglevel to Off", strLogLevel)
	}
	return logLevel
}

type Logger interface {
	Printf(format string, args ...interface{})
}

func NewDefaultLogger() Logger {
	return &defaultLogger{
		logger: log.New(os.Stderr, "IONOSLOG ", log.LstdFlags),
	}
}

type defaultLogger struct {
	logger *log.Logger
}

func (l defaultLogger) Printf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}
