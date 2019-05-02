package route

import "net/url"

type RawURLMap map[string]string

type RouteMap map[string]*url.URL

func (m RouteMap) Route(name string) *url.URL {
	return m[name]
}

func (m RouteMap) Map(rmap RawURLMap) {
	for name, raw := range rmap {
		m[name] = parseURL(raw)
	}
}

func parseURL(rawURL string) *url.URL {
	url, _ := url.Parse(rawURL)
	return url
}
