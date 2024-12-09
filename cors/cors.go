package cors

import (
	"net/http"
)

type MethodType string

type CORS struct {
	allowedOrigins []string
}

func NewCORS(allowedOrigins []string) *CORS {
	return &CORS{allowedOrigins: allowedOrigins}
}

func (c *CORS) EnableCORS(next http.Handler, methods []string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			for _, allowedOrigin := range c.allowedOrigins {
				if origin == allowedOrigin {
					w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
				}
			}
			// Set CORS headers
			w.Header().Set("Access-Control-Allow-Methods", methodsToString(methods))
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			// Handle OPTIONS request
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		},
	)
}

func methodsToString(methods []string) string {
	res := ""
	for _, method := range methods {
		res += method
	}
	return res
}
