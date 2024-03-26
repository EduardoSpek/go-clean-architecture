package middlewares

import (
	"net/http"
)

// CorsMiddleware é um middleware para permitir CORS
func CorsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Headers", "*")
        w.Header().Set("Access-Control-Allow-Methods", "*")
        next.ServeHTTP(w, r)
    })
}