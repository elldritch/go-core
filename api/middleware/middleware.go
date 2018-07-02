// Package middleware is unimplemented
package middleware

import (
	"log"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	log.Println("WARNING: middleware.Logger is not implemented")
	return next
}
