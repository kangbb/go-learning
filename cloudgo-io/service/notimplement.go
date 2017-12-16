package service

import (
	"net/http"
)

func NotImplementedHandler() http.Handler { return http.HandlerFunc(NotImplemented) }

func NotImplemented(w http.ResponseWriter, r *http.Request) { http.Error(w, "501 Not Implemented", 501) }

