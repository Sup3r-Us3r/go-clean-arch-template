package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/Sup3r-Us3r/barber-server/log"
)

type responseLogger struct {
	http.ResponseWriter
	statusCode   int
	bytesWritten int
}

func (rl *responseLogger) WriteHeader(statusCode int) {
	rl.statusCode = statusCode
	rl.ResponseWriter.WriteHeader(statusCode)
}

func (rl *responseLogger) Write(data []byte) (int, error) {
	n, err := rl.ResponseWriter.Write(data)
	// Track the number of bytes written
	rl.bytesWritten += n

	return n, err
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startRequest := time.Now()

		rl := &responseLogger{ResponseWriter: w}

		var payload any = nil

		body, err := io.ReadAll(r.Body)
		if err == nil {
			bodyReader := bytes.NewReader(body)

			// Restore request body
			r.Body = io.NopCloser(bodyReader)

			json.Unmarshal(body, &payload)
		}

		defer func() {
			log.Handler(
				log.HandlerOptions{
					Method:       r.Method,
					Path:         r.URL.String(),
					StatusCode:   rl.statusCode,
					IpAddress:    r.RemoteAddr,
					BytesWritten: rl.bytesWritten,
					Duration:     time.Since(startRequest),
					OptionalAttributes: log.HandlerOptionalAttributes{
						QueryParams: r.URL.Query().Encode(),
						Payload:     payload,
					},
				},
			)
		}()

		next.ServeHTTP(rl, r)
	})
}
