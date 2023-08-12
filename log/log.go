package log

import (
	"io"
	"log/slog"
	"os"
	"runtime"
	"time"

	"github.com/Sup3r-Us3r/barber-server/config"
)

func NewLogger(configs *config.Conf) {
	var logOutput io.Writer

	if configs.LogOutput == "stdout" {
		logOutput = os.Stdout
	} else {
		logOutput = logFile()
	}

	jsonHandler := slog.NewJSONHandler(logOutput, nil).WithAttrs(
		[]slog.Attr{
			slog.String("api-version", "v1"),
			slog.String("golang-version", runtime.Version()),
			slog.String("environment", configs.Environment),
		},
	)
	logger := slog.New(jsonHandler)
	slog.SetDefault(logger)
}

func logFile() io.Writer {
	file, err := os.OpenFile("api.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		Error(err.Error())
		panic(err)
	}

	return file
}

func Info(message string, args ...any) {
	slog.Info(message, args...)
}

func Debug(message string, args ...any) {
	slog.Debug(message, args...)
}

func Warn(message string, args ...any) {
	slog.Warn(message, args...)
}

func Error(message string, args ...any) {
	slog.Error(message, args...)
}

func Any(key string, value any) slog.Attr {
	return slog.Any(key, value)
}

func Group(key string, args ...any) slog.Attr {
	return slog.Group(key, args...)
}

func String(key string, value string) slog.Attr {
	return slog.String(key, value)
}

func Int64(key string, value int64) slog.Attr {
	return slog.Int64(key, value)
}

func Int(key string, value int) slog.Attr {
	return slog.Int(key, value)
}

func Uint64(key string, value uint64) slog.Attr {
	return slog.Uint64(key, value)
}

func Float64(key string, value float64) slog.Attr {
	return slog.Float64(key, value)
}

func Bool(key string, value bool) slog.Attr {
	return slog.Bool(key, value)
}

func Time(key string, value time.Time) slog.Attr {
	return slog.Time(key, value)
}

func Duration(key string, value time.Duration) slog.Attr {
	return slog.Duration(key, value)
}

type HandlerOptionalAttributes struct {
	QueryParams any
	Payload     any
}

type HandlerOptions struct {
	Method             string
	Path               string
	StatusCode         int
	IpAddress          string
	BytesWritten       int
	Duration           time.Duration
	OptionalAttributes HandlerOptionalAttributes
}

type handlerLevelFunc = func(message string, args ...any)

func Handler(handlerOptions HandlerOptions) {
	levels := map[string]handlerLevelFunc{
		"info":  Info,
		"warn":  Warn,
		"error": Error,
	}

	var level handlerLevelFunc

	if handlerOptions.StatusCode >= 200 && handlerOptions.StatusCode <= 299 {
		level = levels["info"]
	} else if handlerOptions.StatusCode >= 300 && handlerOptions.StatusCode <= 499 {
		level = levels["warn"]
	} else {
		level = levels["error"]
	}

	level(
		"Request",
		slog.String("method", handlerOptions.Method),
		slog.String("path", handlerOptions.Path),
		slog.String("ip-address", handlerOptions.IpAddress),
		slog.Int("status-code", handlerOptions.StatusCode),
		slog.Int("bytes-written", handlerOptions.BytesWritten),
		slog.String("duration", handlerOptions.Duration.String()),
		slog.Any("query", handlerOptions.OptionalAttributes.QueryParams),
		slog.Any("payload", handlerOptions.OptionalAttributes.Payload),
	)
}
