package pkg

type Route struct {
	Name       string
	TravelTime int
}

type RouteIterator interface {
	HasNext() bool
	Next() *Route
}

type RouteCollection struct {
	Routes []*Route
	index  int
}

func (rc *RouteCollection) HasNext() bool {
	return rc.index < len(rc.Routes)
}

func (rc *RouteCollection) Next() *Route {
	route := rc.Routes[rc.index]
	rc.index++
	return route
}
