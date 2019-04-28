package route

import "net/url"

type RawSet struct {
	Method string
	URL    string
}

type RawMap map[string]*RawSet

type RouteSet struct {
	Method string
	URL    *url.URL
}

type RouteMap map[string]*RouteSet

func (m RouteMap) Route(name string) *RouteSet {
	return m[name]
}

func (m RouteMap) URL(name string) *url.URL {
	return m[name].URL
}

func (m RouteMap) Map(rawMap RawMap) {
	for name, rawSet := range rawMap {
		m[name] = &RouteSet{
			Method: rawSet.Method,
			URL:    parseURL(rawSet.URL),
		}
	}
}

func parseURL(rawURL string) *url.URL {
	url, _ := url.Parse(rawURL)
	return url
}
