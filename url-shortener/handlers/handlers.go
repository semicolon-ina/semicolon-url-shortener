package handlers

import (
	"github.com/semicolon-ina/semicolon-url-shortener/repo/domain/url"
)

type HTTPHandler struct {
	uSvc url.URLInterface
}

func NewHTTPHandler(uSvc url.URLInterface) *HTTPHandler {
	return &HTTPHandler{uSvc: uSvc}
}

type ShortenRequest struct {
	URL string `json:"url"`
}
