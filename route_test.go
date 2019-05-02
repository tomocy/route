package route

import (
	"path"
	"testing"
)

var (
	routeMap = make(RouteMap)
)

func TestRoute(t *testing.T) {
	routeMap.Map(
		RawURLMap{
			"user.index": localhost("/users"),
			"user.new":   localhost("/users/new"),
		},
	)

	tests := []struct {
		name     string
		expected string
	}{
		{
			name:     "user.index",
			expected: "http://localhost/users",
		},
		{
			name:     "user.new",
			expected: "http://localhost/users/new",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := routeMap.Route(test.name)
			errorfIfNotEqual(t, actual.String(), test.expected)
		})
	}
}

func errorfIfNotEqual(t *testing.T, actual, expected interface{}) {
	if actual != expected {
		t.Errorf("Got %v, Expected %v\n", actual, expected)
	}
}

func localhost(p string) string {
	return "http://localhost" + path.Join("/", p)
}
