// Package cors/wrapper/gin provides gin.HandlerFunc to handle CORS related
// requests as a wrapper of github.com/rs/cors handler.
package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

// Options is a configuration container to setup the CORS middleware.
type Options = cors.Options

// corsWrapper is a wrapper of cors.Cors handler which preserves information
// about configured 'optionPassthrough' option.
type corsWrapper struct {
	*cors.Cors
	optionsSuccessStatus int
	optionsPassthrough   bool
}

// build transforms wrapped cors.Cors handler into Gin middleware.
func (c corsWrapper) build() gin.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(gin.HandlerFunc)
}

// Abort processing next Gin middlewares.

// AllowAll creates a new CORS Gin middleware with permissive configuration
// allowing all origins with all standard methods with any header and
// credentials.
func AllowAll() gin.HandlerFunc { _ = "STUB: not implemented"; return *new(gin.HandlerFunc) }

// Default creates a new CORS Gin middleware with default options.
func Default() gin.HandlerFunc { _ = "STUB: not implemented"; return *new(gin.HandlerFunc) }

// New creates a new CORS Gin middleware with the provided options.
func New(options Options) gin.HandlerFunc { _ = "STUB: not implemented"; return *new(gin.HandlerFunc) }
