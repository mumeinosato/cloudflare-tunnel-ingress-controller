package controller

// AnnotationProxySSLVerify is the annotation key for proxy-ssl-verify, available values: "on" or "off", default "off".
const AnnotationProxySSLVerify = "cloudflared-ingress.mumeinosato.github.io/proxy-ssl-verify"
const AnnotationProxySSLVerifyOn = "on"
const AnnotationProxySSLVerifyOff = "off"

// AnnotationBackendProtocol is the annotation key for proxy-backend-protocol, default "http".
const AnnotationBackendProtocol = "cloudflared-ingress.mumeinosato.github.io/backend-protocol"

// AnnotationHTTPHostHeader is to set the HTTP Host header for the local webserver.
const AnnotationHTTPHostHeader = "cloudflared-ingress.mumeinosato.github.io/http-host-header"

// AnnotationOriginServerName is the hostname on the origin server certificate.
const AnnotationOriginServerName = "cloudflared-ingress.mumeinosato.github.io/origin-server-name"

// CustomAnnotations
const AnnotationDisableChunkedEncoding = "cloudflared-ingress.mumeinosato.github.io/disable-chunked-encoding"
