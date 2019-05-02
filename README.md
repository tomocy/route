# Route

[![CircleCI](https://circleci.com/gh/tomocy/route.svg?style=svg)](https://circleci.com/gh/tomocy/route)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

route manager

## Installation
```
go get github.com/tomocy/route
```

## Useage
Map routes
```go
func MapRoutes() {
    Web.Map(webRaw)
    API.Map(apiRaw)
}

var (
    Web = make(route.RouteMap)
    API = make(route.RouteMap)
)

var (
    webRaw = route.RawMap{
        "user.index": "http://localhost/users",
        "user.new": "http://localhost/users/new",
        "user.create": "http://localhost/users",
    }
    apiRaw = route.RawMap{
        "user.index": "http://localhost/api/users/index",
    }
)
```
`(RouteMap).Route(string)` returns `*url.URL`
```go
http.HandlerFunc(route.Web.Route("user.new").Path, func(w http.ResponseWriter, r *http.Request) {
    // handle user.new
})
http.HandlerFunc(route.Web.Route("user.create").Path, {
    // handle user.create

    // after creating a user
    http.Redirect(w, r, route.Web.Route("user.index").String(), http.StatusSeeOther)
})
```

## Author
[tomocy](https://github.com/tomocy)