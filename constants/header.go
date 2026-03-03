package constants

import "net/textproto"

var (
	XServiceName  = textproto.CanonicalMIMEHeaderKey("x - service - name")
	XApiKey       = textproto.CanonicalMIMEHeaderKey("x - api - key")
	XRequestId    = textproto.CanonicalMIMEHeaderKey("x - request - id")
	XrequestAt    = textproto.CanonicalMIMEHeaderKey("x - request - at")
	Authorization = textproto.CanonicalMIMEHeaderKey("authorization")
)
