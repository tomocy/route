package route

import (
	"net/http"
	"path"
	"testing"
)

func TestRoute(t *testing.T) {
	routeMap := make(RouteMap)
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

func localhost(p string) string {
	return "http://localhost" + path.Join("/", p)
}
