// Package status provides a comprehensive set of HTTP-like status codes
// and utility functions for managing and retrieving their descriptions.
// It defines a custom `StatusCode` type and a mapping of codes to
// human-readable descriptions, facilitating consistent status handling
// across applications.
package status

// StatusCode is a custom integer type representing an HTTP-like status code.
type StatusCode int

// HTTP 1xx (Informational)
const (
	Continue           StatusCode = 100 // 100 Continue: The server has received the request headers and the client should proceed to send the request body.
	SwitchingProtocols StatusCode = 101 // 101 Switching Protocols: The requester has asked the server to switch protocols.
	Processing         StatusCode = 102 // 102 Processing: A WebDAV extension, indicates that the server has accepted the complete request but has not yet completed it.
	EarlyHints         StatusCode = 103 // 103 Early Hints: Used to return some response headers before final HTTP message.
)

// HTTP 2xx (Success)
const (
	OK                          StatusCode = 200 // 200 OK: The request has succeeded.
	Created                     StatusCode = 201 // 201 Created: The request has been fulfilled and resulted in a new resource being created.
	Accepted                    StatusCode = 202 // 202 Accepted: The request has been accepted for processing, but the processing has not been completed.
	NonAuthoritativeInformation StatusCode = 203 // 203 Non-Authoritative Information: The server is a transforming proxy that received a 200 OK from its origin, but is returning a modified version of the origin's response.
	NoContent                   StatusCode = 204 // 204 No Content: The server successfully processed the request and is not returning any content.
	ResetContent                StatusCode = 205 // 205 Reset Content: The server successfully processed the request, but is not returning any content. The client should reset the document view.
	PartialContent              StatusCode = 206 // 206 Partial Content: The server is delivering only part of the resource due to a range header sent by the client.
	MultiStatusCode             StatusCode = 207 // 207 Multi-Status: The message body that follows is an XML message and may contain a number of separate response codes, depending on how many sub-requests were made.
	AlreadyReported             StatusCode = 208 // 208 Already Reported: The members of a DAV binding have already been enumerated in a preceding part of the (multistatus) response, and are not being enumerated again.
	IMUsed                      StatusCode = 226 // 226 IM Used: The server has fulfilled a GET request for the resource, and the response is a representation of the result of one or more instance-manipulations applied to the current instance.
)

// HTTP 3xx (Redirection)
const (
	MultipleChoices   StatusCode = 300 // 300 Multiple Choices: Indicates multiple options for the resource from which the client may choose.
	MovedPermanently  StatusCode = 301 // 301 Moved Permanently: The requested resource has been assigned a new permanent URI.
	Found             StatusCode = 302 // 302 Found: The requested resource resides temporarily under a different URI.
	SeeOther          StatusCode = 303 // 303 See Other: The response to the request can be found under another URI using the GET method.
	NotModified       StatusCode = 304 // 304 Not Modified: The resource has not been modified since the version specified by the request headers.
	UseProxy          StatusCode = 305 // 305 Use Proxy: The requested resource MUST be accessed through the proxy given by the Location field.
	TemporaryRedirect StatusCode = 307 // 307 Temporary Redirect: The request should be repeated with another URI; however, future requests should still use the original URI.
	PermanentRedirect StatusCode = 308 // 308 Permanent Redirect: The request and all future requests should be repeated using another URI.
)

// HTTP 4xx (Client Error)
const (
	BadRequest                  StatusCode = 400 // 400 Bad Request: The server cannot or will not process the request due to something that is perceived to be a client error.
	Unauthorized                StatusCode = 401 // 401 Unauthorized: Authentication is required and has failed or has not yet been provided.
	PaymentRequired             StatusCode = 402 // 402 Payment Required: Reserved for future use. The original intention was that this code might be used as part of some form of digital cash or micropayment scheme.
	Forbidden                   StatusCode = 403 // 403 Forbidden: The request was a valid request, but the server is refusing to respond to it.
	NotFound                    StatusCode = 404 // 404 Not Found: The requested resource could not be found but may be available again in the future.
	MethodNotAllowed            StatusCode = 405 // 405 Method Not Allowed: A request method is not supported for the requested resource.
	NotAcceptable               StatusCode = 406 // 406 Not Acceptable: The requested resource is capable of generating only content not acceptable according to the Accept headers sent in the request.
	ProxyAuthenticationRequired StatusCode = 407 // 407 Proxy Authentication Required: The client must first authenticate itself with the proxy.
	RequestTimeout              StatusCode = 408 // 408 Request Timeout: The server timed out waiting for the request.
	Conflict                    StatusCode = 409 // 409 Conflict: The request could not be completed due to a conflict with the current state of the resource.
	Gone                        StatusCode = 410 // 410 Gone: The requested resource is no longer available at the server and no forwarding address is known.
	LengthRequired              StatusCode = 411 // 411 Length Required: The server rejected the request because the Content-Length header field is not defined and the server requires it.
	PreconditionFailed          StatusCode = 412 // 412 Precondition Failed: The server does not meet one of the preconditions that the requester put on the request.
	ContentTooLarge             StatusCode = 413 // 413 Content Too Large: The request is larger than the server is willing or able to process.
	URITooLong                  StatusCode = 414 // 414 URI Too Long: The URI provided was too long for the server to process.
	UnsupportedMediaType        StatusCode = 415 // 415 Unsupported Media Type: The request entity has a media type which the server or resource does not support.
	RangeNotSatisfiable         StatusCode = 416 // 416 Range Not Satisfiable: The client has asked for a portion of the file, but the server cannot supply that portion.
	ExpectationFailed           StatusCode = 417 // 417 Expectation Failed: The server cannot meet the requirements of the Expect request-header field.
	IMATeapot                   StatusCode = 418 // 418 I'm a teapot: This code was defined in 1998 as one of the traditional IETF April Fools' jokes, in RFC 2324, Hyper Text Coffee Pot Control Protocol.
	MisdirectedRequest          StatusCode = 421 // 421 Misdirected Request: The request was directed at a server that is not able to produce a response.
	UnprocessableContent        StatusCode = 422 // 422 Unprocessable Content: The request was well-formed but was unable to be followed due to semantic errors.
	Locked                      StatusCode = 423 // 423 Locked: The resource that is being accessed is locked.
	FailedDependency            StatusCode = 424 // 424 Failed Dependency: The method could not be performed on the resource because the requested action depended on another action and that action failed.
	TooEarly                    StatusCode = 425 // 425 Too Early: Indicates that the server is unwilling to risk processing a request that might be replayed.
	UpgradeRequired             StatusCode = 426 // 426 Upgrade Required: The client should switch to a different protocol.
	PreconditionRequired        StatusCode = 428 // 428 Precondition Required: The origin server requires the request to be conditional.
	TooManyRequests             StatusCode = 429 // 429 Too Many Requests: The user has sent too many requests in a given amount of time.
	RequestHeaderFieldsTooLarge StatusCode = 431 // 431 Request Header Fields Too Large: The server is unwilling to process the request because its header fields are too large.
	UnavailableForLegalReasons  StatusCode = 451 // 451 Unavailable For Legal Reasons: A server operator has received a legal demand to deny access to a resource or to a set of resources that includes the requested resource.
)

// HTTP 5xx (Server Error)
const (
	InternalServerError           StatusCode = 500 // 500 Internal Server Error: A generic error message, given when an unexpected condition was encountered and no more specific message is suitable.
	NotImplemented                StatusCode = 501 // 501 Not Implemented: The server either does not recognize the request method, or it lacks the ability to fulfill the request.
	BadGateway                    StatusCode = 502 // 502 Bad Gateway: The server was acting as a gateway or proxy and received an invalid response from the upstream server.
	ServiceUnavailable            StatusCode = 503 // 503 Service Unavailable: The server is currently unable to handle the request due to a temporary overload or scheduled maintenance.
	GatewayTimeout                StatusCode = 504 // 504 Gateway Timeout: The server was acting as a gateway or proxy and did not receive a timely response from the upstream server.
	HTTPVersionNotSupported       StatusCode = 505 // 505 HTTP Version Not Supported: The server does not support the HTTP protocol version used in the request.
	VariantAlsoNegotiates         StatusCode = 506 // 506 Variant Also Negotiates: The server has an internal configuration error: the chosen variant resource is configured to engage in transparent content negotiation itself, and is therefore not a proper end point in the negotiation process.
	InsufficientStorage           StatusCode = 507 // 507 Insufficient Storage: The server is unable to store the representation needed to complete the request.
	LoopDetected                  StatusCode = 508 // 508 Loop Detected: The server detected an infinite loop while processing the request.
	NotExtended                   StatusCode = 510 // 510 Not Extended: The policy for accessing the resource has not been met in the request's Etag or If-Modified-Since headers.
	NetworkAuthenticationRequired StatusCode = 511 // 511 Network Authentication Required: The client needs to authenticate to gain network access.
)

// statusDescriptions maps each StatusCode to its corresponding human-readable string description.
var statusDescriptions = map[StatusCode]string{
	100: "Continue",
	101: "Switching Protocols",
	102: "Processing",
	103: "Early Hints",
	200: "OK",
	201: "Created",
	202: "Accepted",
	203: "Non-Authoritative Information",
	204: "No Content",
	205: "Reset Content",
	206: "Partial Content",
	207: "Multi-Status",
	208: "Already Reported",
	226: "IM Used",
	300: "Multiple Choices",
	301: "Moved Permanently",
	302: "Found",
	303: "See Other",
	304: "Not Modified",
	305: "Use Proxy",
	307: "Temporary Redirect",
	308: "Permanent Redirect",
	400: "Bad Request",
	401: "Unauthorized",
	402: "Payment Required",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	406: "Not Acceptable",
	407: "Proxy Authentication Required",
	408: "Request Timeout",
	409: "Conflict",
	410: "Gone",
	411: "Length Required",
	412: "Precondition Failed",
	413: "Content Too Large",
	414: "URI Too Long",
	415: "Unsupported Media Type",
	416: "Range Not Satisfiable",
	417: "Expectation Failed",
	418: "I'm a teapot",
	421: "Misdirected Request",
	422: "Unprocessable Content",
	423: "Locked",
	424: "Failed Dependency",
	425: "Too Early",
	426: "Upgrade Required",
	428: "Precondition Required",
	429: "Too Many Requests",
	431: "Request Header Fields Too Large",
	451: "Unavailable For Legal Reasons",
	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Service Unavailable",
	504: "Gateway Timeout",
	505: "HTTP Version Not Supported",
	506: "Variant Also Negotiates",
	507: "Insufficient Storage",
	508: "Loop Detected",
	510: "Not Extended",
	511: "Network Authentication Required",
}

// GetValue returns the integer representation of the StatusCode.
func (c StatusCode) GetValue() int {
	return int(c)
}

// GetDescription returns the human-readable string description for the StatusCode.
// If the StatusCode is not recognized, it returns "Unknown Status Code".
func (c StatusCode) GetDescription() string {
	if desc, exists := statusDescriptions[c]; exists {
		return desc
	}
	return "Unknown Status Code"
}

// NewStatusCode creates a StatusCode from an integer value.
// It also returns a boolean indicating whether the created StatusCode
// is a known, defined HTTP status code.
//
// Parameters:
//
//	value: The integer value to convert to a StatusCode.
//
// Returns:
//
//	A StatusCode type and a boolean indicating if the status code is known.
func NewStatusCode(value int) (StatusCode, bool) {
	s := StatusCode(value)
	// Check if the integer value corresponds to a known status code description.
	if _, exists := statusDescriptions[s]; exists {
		return s, true
	}
	return 0, false // Return 0 and false if the status code is unknown.
}

// GetStatusTexts returns a copy of the internal map that maps each
// StatusCode to its string description. This allows external access to
// all defined status code descriptions without modifying the internal map.
func GetStatusTexts() map[StatusCode]string {
	// Create a new map to ensure the original map is not modified externally.
	copyMap := make(map[StatusCode]string, len(statusDescriptions))
	for key, value := range statusDescriptions {
		copyMap[key] = value
	}
	return copyMap
}
