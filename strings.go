package fasthttp

var (
	defaultServerName  = []byte("fasthttp")
	defaultUserAgent   = []byte("fasthttp")
	defaultContentType = []byte("text/plain; charset=utf-8")
)

var (
	strSlash            = []byte("/")
	strSlashSlash       = []byte("//")
	strSlashDotDot      = []byte("/..")
	strSlashDotSlash    = []byte("/./")
	strSlashDotDotSlash = []byte("/../")
	strCRLF             = []byte("\r\n")
	strHTTP             = []byte("http")
	strHTTPS            = []byte("https")
	strHTTP11           = []byte("HTTP/1.1")
	strColon            = []byte(":")
	strColonSlashSlash  = []byte("://")
	strColonSpace       = []byte(": ")
	strGMT              = []byte("GMT")
	strAt               = []byte("@")

	strResponseContinue = []byte("HTTP/1.1 100 Continue\r\n\r\n")

	strGet     = []byte(MethodGet)
	strHead    = []byte(MethodHead)
	strPost    = []byte(MethodPost)
	strPut     = []byte(MethodPut)
	strDelete  = []byte(MethodDelete)
	strConnect = []byte(MethodConnect)
	strOptions = []byte(MethodOptions)
	strTrace   = []byte(MethodTrace)
	strPatch   = []byte(MethodPatch)

	strExpect           = []byte(HeaderExpect)
	strConnection       = []byte(HeaderConnection)
	strContentLength    = []byte(HeaderContentLength)
	strContentType      = []byte(HeaderContentType)
	strDate             = []byte(HeaderDate)
	strHost             = []byte(HeaderHost)
	strReferer          = []byte(HeaderReferer)
	strServer           = []byte(HeaderServer)
	strTransferEncoding = []byte(HeaderTransferEncoding)
	strContentEncoding  = []byte(HeaderContentEncoding)
	strAcceptEncoding   = []byte(HeaderAcceptEncoding)
	strUserAgent        = []byte(HeaderUserAgent)
	strCookie           = []byte(HeaderCookie)
	strSetCookie        = []byte(HeaderSetCookie)
	strLocation         = []byte(HeaderLocation)
	strIfModifiedSince  = []byte(HeaderIfModifiedSince)
	strLastModified     = []byte(HeaderLastModified)
	strAcceptRanges     = []byte(HeaderAcceptRanges)
	strRange            = []byte(HeaderRange)
	strContentRange     = []byte(HeaderContentRange)
	strAuthorization    = []byte(HeaderAuthorization)

	strCookieExpires        = []byte("expires")
	strCookieDomain         = []byte("domain")
	strCookiePath           = []byte("path")
	strCookieHTTPOnly       = []byte("HttpOnly")
	strCookieSecure         = []byte("secure")
	strCookieMaxAge         = []byte("max-age")
	strCookieSameSite       = []byte("SameSite")
	strCookieSameSiteLax    = []byte("Lax")
	strCookieSameSiteStrict = []byte("Strict")
	strCookieSameSiteNone   = []byte("None")

	strClose               = []byte("close")
	strGzip                = []byte("gzip")
	strDeflate             = []byte("deflate")
	strKeepAlive           = []byte("keep-alive")
	strUpgrade             = []byte("Upgrade")
	strChunked             = []byte("chunked")
	strIdentity            = []byte("identity")
	str100Continue         = []byte("100-continue")
	strPostArgsContentType = []byte("application/x-www-form-urlencoded")
	strMultipartFormData   = []byte("multipart/form-data")
	strBoundary            = []byte("boundary")
	strBytes               = []byte("bytes")
	strTextSlash           = []byte("text/")
	strApplicationSlash    = []byte("application/")
	strBasicSpace          = []byte("Basic ")
)
