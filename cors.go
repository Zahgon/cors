/*
Package cors is net/http handler to handle CORS related requests
as defined by http://www.w3.org/TR/cors/

You can configure it by passing an option struct to cors.New:

	c := cors.New(cors.Options{
	    AllowedOrigins:   []string{"foo.com"},
	    AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete},
	    AllowCredentials: true,
	})

Then insert the handler in the chain:

	handler = c.Handler(handler)

See Options documentation for more options.

The resulting handler is a standard net/http handler.
*/
package cors

import (
	"net/http"

	"github.com/rs/cors/internal"
)

var (
	headerVaryOrigin = []string{"Origin"}
	headerOriginAll  = []string{"*"}
	headerTrue       = []string{"true"}
)

// Options is a configuration container to setup the CORS middleware.
type Options struct {
	// AllowedOrigins is a list of origins a cross-domain request can be executed from.
	// If the special "*" value is present in the list, all origins will be allowed.
	// An origin may contain a wildcard (*) to replace 0 or more characters
	// (i.e.: http://*.domain.com). Usage of wildcards implies a small performance penalty.
	// Only one wildcard can be used per origin.
	// Default value is ["*"]
	AllowedOrigins []string
	// AllowOriginFunc is a custom function to validate the origin. It take the
	// origin as argument and returns true if allowed or false otherwise. If
	// this option is set, the content of `AllowedOrigins` is ignored.
	AllowOriginFunc func(origin string) bool
	// AllowOriginRequestFunc is a custom function to validate the origin. It
	// takes the HTTP Request object and the origin as argument and returns true
	// if allowed or false otherwise. If headers are used take the decision,
	// consider using AllowOriginVaryRequestFunc instead. If this option is set,
	// the contents of `AllowedOrigins`, `AllowOriginFunc` are ignored.
	//
	// Deprecated: use `AllowOriginVaryRequestFunc` instead.
	AllowOriginRequestFunc func(r *http.Request, origin string) bool
	// AllowOriginVaryRequestFunc is a custom function to validate the origin.
	// It takes the HTTP Request object and the origin as argument and returns
	// true if allowed or false otherwise with a list of headers used to take
	// that decision if any so they can be added to the Vary header. If this
	// option is set, the contents of `AllowedOrigins`, `AllowOriginFunc` and
	// `AllowOriginRequestFunc` are ignored.
	AllowOriginVaryRequestFunc func(r *http.Request, origin string) (bool, []string)
	// AllowedMethods is a list of methods the client is allowed to use with
	// cross-domain requests. Default value is simple methods (HEAD, GET and POST).
	AllowedMethods []string
	// AllowedHeaders is list of non simple headers the client is allowed to use with
	// cross-domain requests.
	// If the special "*" value is present in the list, all headers will be allowed.
	// Default value is [].
	AllowedHeaders []string
	// ExposedHeaders indicates which headers are safe to expose to the API of a CORS
	// API specification
	ExposedHeaders []string
	// MaxAge indicates how long (in seconds) the results of a preflight request
	// can be cached. Default value is 0, which stands for no
	// Access-Control-Max-Age header to be sent back, resulting in browsers
	// using their default value (5s by spec). If you need to force a 0 max-age,
	// set `MaxAge` to a negative value (ie: -1).
	MaxAge int
	// AllowCredentials indicates whether the request can include user credentials like
	// cookies, HTTP authentication or client side SSL certificates.
	AllowCredentials bool
	// AllowPrivateNetwork indicates whether to accept cross-origin requests over a
	// private network.
	AllowPrivateNetwork bool
	// OptionsPassthrough instructs preflight to let other potential next handlers to
	// process the OPTIONS method. Turn this on if your application handles OPTIONS.
	OptionsPassthrough bool
	// Provides a status code to use for successful OPTIONS requests.
	// Default value is http.StatusNoContent (204).
	OptionsSuccessStatus int
	// Debugging flag adds additional output to debug server side CORS issues
	Debug bool
	// Adds a custom logger, implies Debug is true
	Logger Logger
}

// Logger generic interface for logger
type Logger interface {
	Printf(string, ...any)
}

// Cors http handler
type Cors struct {
	// Debug logger
	Log Logger
	// Normalized list of plain allowed origins
	allowedOrigins []string
	// List of allowed origins containing wildcards
	allowedWOrigins []wildcard
	// Optional origin validator function
	allowOriginFunc func(r *http.Request, origin string) (bool, []string)
	// Normalized list of allowed headers
	// Note: the Fetch standard guarantees that CORS-unsafe request-header names
	// (i.e. the values listed in the Access-Control-Request-Headers header)
	// are unique and sorted;
	// see https://fetch.spec.whatwg.org/#cors-unsafe-request-header-names.
	allowedHeaders internal.SortedSet
	// Normalized list of allowed methods
	allowedMethods []string
	// Pre-computed normalized list of exposed headers
	exposedHeaders []string
	// Pre-computed maxAge header value
	maxAge []string
	// Set to true when allowed origins contains a "*"
	allowedOriginsAll bool
	// Set to true when allowed headers contains a "*"
	allowedHeadersAll bool
	// Status code to use for successful OPTIONS requests
	optionsSuccessStatus int
	allowCredentials     bool
	allowPrivateNetwork  bool
	optionPassthrough    bool
	preflightVary        []string
}

// New creates a new Cors handler with the provided options.
func New(options Options) *Cors { _ = "STUB: not implemented"; return nil }

// Allowed origins

// Default is all origins

// Note: for origins matching, the spec requires a case-sensitive matching.
// As it may error prone, we chose to ignore the spec here.

// If "*" is present in the list, turn the whole list into a match all

// Split the origin in two: start and end string without the *

// Allowed Headers
// Note: the Fetch standard guarantees that CORS-unsafe request-header names
// (i.e. the values listed in the Access-Control-Request-Headers header)
// are lowercase; see https://fetch.spec.whatwg.org/#cors-unsafe-request-header-names.

// Use sensible defaults

// Allowed Methods

// Default is spec's "simple" methods

// Options Success Status Code

// Pre-compute exposed headers header value

// Pre-compute prefight Vary header to save allocations

// Precompute max-age

// Default creates a new Cors handler with default options.
func Default() *Cors { _ = "STUB: not implemented"; return nil }

// AllowAll create a new Cors handler with permissive configuration allowing all
// origins with all standard methods with any header and credentials.
func AllowAll() *Cors { _ = "STUB: not implemented"; return nil }

// Handler apply the CORS specification on the request, and add relevant CORS headers
// as necessary.
func (c *Cors) Handler(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// Preflight requests are standalone and should stop the chain as some other
// middleware may not handle OPTIONS requests correctly. One typical example
// is authentication middleware ; OPTIONS requests won't carry authentication
// headers (see #1)

// HandlerFunc provides Martini compatible handler
func (c *Cors) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Negroni compatible interface
func (c *Cors) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_ = "STUB: not implemented"
	return
}

// Preflight requests are standalone and should stop the chain as some other
// middleware may not handle OPTIONS requests correctly. One typical example
// is authentication middleware ; OPTIONS requests won't carry authentication
// headers (see #1)

// handlePreflight handles pre-flight CORS requests
func (c *Cors) handlePreflight(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Always set Vary headers
// see https://github.com/rs/cors/issues/10,
//     https://github.com/rs/cors/commit/dbdca4d95feaa7511a46e6f1efb3b3aa505bc43f#commitcomment-12352001

// Note: the Fetch standard guarantees that at most one
// Access-Control-Request-Headers header is present in the preflight request;
// see step 5.2 in https://fetch.spec.whatwg.org/#cors-preflight-fetch-0.
// However, some gateways split that header into multiple headers of the same name;
// see https://github.com/rs/cors/issues/184.

// Spec says: Since the list of methods can be unbounded, simply returning the method indicated
// by Access-Control-Request-Method (if supported) can be enough

// Spec says: Since the list of headers can be unbounded, simply returning supported headers
// from Access-Control-Request-Headers can be enough

// handleActualRequest handles simple cross-origin requests, actual request or redirects
func (c *Cors) handleActualRequest(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Always set Vary, see https://github.com/rs/cors/issues/10

// Note that spec does define a way to specifically disallow a simple method like GET or
// POST. Access-Control-Allow-Methods is only used for pre-flight requests and the
// spec doesn't instruct to check the allowed methods for simple cross-origin requests.
// We think it's a nice feature to be able to have control on those methods though.

// convenience method. checks if a logger is set.
func (c *Cors) logf(format string, a ...any) { _ = "STUB: not implemented"; return }

// check the Origin of a request. No origin at all is also allowed.
func (c *Cors) OriginAllowed(r *http.Request) bool { _ = "STUB: not implemented"; return false }

// isOriginAllowed checks if a given origin is allowed to perform cross-domain requests
// on the endpoint
func (c *Cors) isOriginAllowed(r *http.Request, origin string) (allowed bool, varyHeaders []string) {
	_ = "STUB: not implemented"
	return false, nil
}

// isMethodAllowed checks if a given method can be used as part of a cross-domain request
// on the endpoint
func (c *Cors) isMethodAllowed(method string) bool { _ = "STUB: not implemented"; return false }

// If no method allowed, always return false, even for preflight request

// Always allow preflight requests
