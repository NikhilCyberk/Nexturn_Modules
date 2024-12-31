package middleware

import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"
)

func JSONValidator() func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            if r.Method == http.MethodPost || r.Method == http.MethodPut {
                contentType := r.Header.Get("Content-Type")
                if contentType != "application/json" {
                    http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
                    return
                }

                // Read the body
                bodyBytes, err := io.ReadAll(r.Body)
                if err != nil {
                    http.Error(w, "Error reading request body", http.StatusBadRequest)
                    return
                }
                r.Body.Close()

                // Validate JSON format
                var js json.RawMessage
                if err := json.Unmarshal(bodyBytes, &js); err != nil {
                    http.Error(w, "Invalid JSON format", http.StatusBadRequest)
                    return
                }

                // Create new ReadCloser from the bytes for subsequent reads
                r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
            }
            next.ServeHTTP(w, r)
        })
    }
}