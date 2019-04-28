package route

import (
	"net/http"
	"path"
	"testing"
)

var (
	routeMap = make(RouteMap)
)

func TestRoute(t *testing.T) {
	routeMap.Map(
		RawMap{
			"user.index": &RawSet{
				Method: http.MethodGet,
				URL:    localhost("/users"),
			},
		},
	)

	tests := []struct {
		name           string
		expectedMethod string
		expectedURL    string
	}{
		{
			name:           "user.index",
			expectedMethod: http.MethodGet,
			expectedURL:    "http://localhost/users",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := routeMap.Route(test.name)
			if actual.Method != test.expectedMethod {
				t.Errorf("Got %v, Expected %v\n", actual.Method, test.expectedMethod)
			}
			if actual.URL.String() != test.expectedURL {
				t.Errorf("Got %v, Expected %v\n", actual.URL.String(), test.expectedURL)
			}
		})
	}
}

func TestURL(t *testing.T) {
	routeMap.Map(
		RawMap{
			"user.create": &RawSet{
				Method: http.MethodPost,
				URL:    localhost("/users"),
			},
		},
	)

	tests := []struct {
		tname    string
		name     string
		expected string
	}{
		{
			tname:    "success",
			name:     "user.index",
			expected: "http://localhost/users",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := routeMap.URL(test.name)
			testURL(t, actual.String(), test.expected)
		})
	}
}

func testURL(t *testing.T, actual, expected string) {
	if actual != expected {
		errorf(t, actual, expected)
	}
}

func TestMethod(t *testing.T) {
	routeMap.Map(
		RawMap{
			"user.delete": &RawSet{
				Method: http.MethodDelete,
				URL:    localhost("/users/{id}"),
			},
		},
	)

	tests := []struct {
		tname    string
		name     string
		expected string
	}{
		{
			tname:    "success",
			name:     "user.delete",
			expected: http.MethodDelete,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := routeMap.Method(test.name)
			errorfIfNotEqual(t, actual, test.expected)
		})
	}
}

func errorfIfNotEqual(t *testing.T, actual, expected interface{}) {
	if actual != expected {
		errorf(t, actual, expected)
	}
}

func errorf(t *testing.T, actual, expected interface{}) {
	t.Errorf("Got %v, Expected %v\n", actual, expected)
}

func localhost(p string) string {
	return "http://localhost" + path.Join("/", p)
}
