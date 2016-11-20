package aero

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"unsafe"

	"github.com/OneOfOne/xxhash"
	cache "github.com/patrickmn/go-cache"
	"github.com/valyala/fasthttp"
)

// This should be close to the MTU size of a TCP packet.
// Regarding performance it makes no sense to compress smaller files.
// Bandwidth can be saved however the savings are minimal for small files
// and the overhead of compressing can lead up to a 75% reduction
// in server speed under high load. Therefore in this case
// we're trying to optimize for performance, not bandwidth.
const gzipThreshold = 1450

var (
	serverHeader               = []byte("Server")
	server                     = []byte("Aero")
	cacheControlHeader         = []byte("Cache-Control")
	cacheControlAlwaysValidate = []byte("no-cache")
	contentTypeOptionsHeader   = []byte("X-Content-Type-Options")
	contentTypeOptions         = []byte("nosniff")
	xssProtectionHeader        = []byte("X-XSS-Protection")
	xssProtection              = []byte("1; mode=block")
	etagHeader                 = []byte("ETag")
	contentTypeHeader          = []byte("Content-Type")
	contentTypeHTML            = []byte("text/html; charset=utf-8")
	contentTypeJSON            = []byte("application/json")
	contentTypePlainText       = []byte("text/plain; charset=utf-8")
	contentEncodingHeader      = []byte("Content-Encoding")
	contentEncodingGzip        = []byte("gzip")
	responseTimeHeader         = []byte("X-Response-Time")
	ifNoneMatchHeader          = []byte("If-None-Match")
)

var ifNoneMatchHeaderBytes []byte

func init() {
	ifNoneMatchHeaderBytes = []byte(ifNoneMatchHeader)
}

// Context ...
type Context struct {
	// Keep this as the first parameter for quick pointer acquisition.
	requestCtx *fasthttp.RequestCtx

	// A pointer to the application this request occured on.
	App *Application

	// Start time
	start time.Time
}

// Handle ...
type Handle func(*Context) string

// JSON encodes the object to a JSON strings and responds.
func (ctx *Context) JSON(value interface{}) string {
	bytes, _ := json.Marshal(value)

	ctx.requestCtx.Response.Header.SetBytesKV(contentTypeHeader, contentTypeJSON)
	return string(bytes)
}

// HTML sends a HTML string.
func (ctx *Context) HTML(html string) string {
	ctx.requestCtx.Response.Header.SetBytesKV(contentTypeHeader, contentTypeHTML)
	return html
}

// Text sends a plain text string.
func (ctx *Context) Text(text string) string {
	ctx.requestCtx.Response.Header.SetBytesKV(contentTypeHeader, contentTypePlainText)
	return text
}

// Error should be used for sending HTML error messages to the user.
func (ctx *Context) Error(statusCode int, html string) string {
	ctx.SetStatusCode(statusCode)
	ctx.requestCtx.Response.Header.SetBytesKV(contentTypeHeader, contentTypeHTML)
	return html
}

// SetStatusCode sets the status code of the request.
func (ctx *Context) SetStatusCode(status int) {
	ctx.requestCtx.SetStatusCode(status)
}

// SetHeader sets header to value.
func (ctx *Context) SetHeader(header string, value string) {
	ctx.requestCtx.Response.Header.Set(header, value)
}

// Get retrieves an URL parameter.
func (ctx *Context) Get(param string) string {
	return fmt.Sprint(ctx.requestCtx.UserValue(param))
}

// GetInt retrieves an URL parameter as an integer.
func (ctx *Context) GetInt(param string) (int, error) {
	return strconv.Atoi(ctx.Get(param))
}

// Respond responds either with raw code or gzipped if the
// code length is greater than the gzip threshold.
func (ctx *Context) Respond(code string) {
	ctx.RespondBytes(StringToBytesUnsafe(code))
}

// RespondBytes responds either with raw code or gzipped if the
// code length is greater than the gzip threshold. Requires a byte slice.
func (ctx *Context) RespondBytes(b []byte) {
	http := ctx.requestCtx

	// Headers
	http.Response.Header.SetBytesKV(cacheControlHeader, cacheControlAlwaysValidate)
	http.Response.Header.SetBytesKV(serverHeader, server)
	http.Response.Header.SetBytesKV(contentTypeOptionsHeader, contentTypeOptions)
	http.Response.Header.SetBytesKV(xssProtectionHeader, xssProtection)
	// http.Response.Header.Set(responseTimeHeader, strconv.FormatInt(time.Since(ctx.start).Nanoseconds()/1000, 10)+" us")

	if ctx.App.Security.Certificate != nil {
		http.Response.Header.Set("Content-Security-Policy", "default-src https:; script-src 'self'; style-src 'sha256-"+ctx.App.cssHash+"'; connect-src https: wss:")
	}

	// Body
	if ctx.App.Config.GZip && len(b) >= gzipThreshold {
		http.Response.Header.SetBytesKV(contentEncodingHeader, contentEncodingGzip)

		// ETag generation
		h := xxhash.NewS64(0)
		h.Write(b)
		etag := strconv.FormatUint(h.Sum64(), 16)

		// If client cache is up to date, send 304 with no response body.
		clientETag := http.Request.Header.PeekBytes(ifNoneMatchHeader)

		if etag == *(*string)(unsafe.Pointer(&clientETag)) {
			http.SetStatusCode(304)
			return
		}

		// Set ETag
		http.Response.Header.SetBytesK(etagHeader, etag)

		if ctx.App.Config.GZipCache {
			cachedResponse, found := ctx.App.gzipCache.Get(etag)

			if found {
				http.Write(cachedResponse.([]byte))
				return
			}
		}

		fasthttp.WriteGzipLevel(http.Response.BodyWriter(), b, 1)

		if ctx.App.Config.GZipCache {
			body := http.Response.Body()
			gzipped := make([]byte, len(body))
			copy(gzipped, body)
			ctx.App.gzipCache.Set(etag, gzipped, cache.DefaultExpiration)
		}
	} else {
		http.Write(b)
	}
}
