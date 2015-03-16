package kangen

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type route struct {
	pattern *regexp.Regexp
	handler func(http.ResponseWriter, *http.Request)
}

type regexpHandler struct {
	routes []*route
}

func (h *regexpHandler) addRoute(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	r := &route{regexp.MustCompile(pattern), handler}
	h.routes = append(h.routes, r)
}

func (h *regexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.routes {
		if route.pattern.MatchString(r.URL.Path) {
			route.handler(w, r)
			return
		}
	}
}

func kangenHandler(w http.ResponseWriter, r *http.Request) {
	shorten := strings.TrimLeft(r.URL.Path, "/")
	conn := connectRedis()
	url := getURL(conn, shorten)

	if url == "" {
		http.NotFound(w, r)
	} else {
		http.Redirect(w, r, url, 301)
	}
}

func Server(port int) {
	writePid()

	handlers := new(regexpHandler)
	handlers.addRoute("/*", kangenHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), handlers)
}
