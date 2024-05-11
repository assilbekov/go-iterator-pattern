package main

import "iterator/pkg"

var routes = pkg.RouteCollection{
	Routes: []*pkg.Route{
		{Name: "Route 1", TravelTime: 10},
		{Name: "Route 2", TravelTime: 20},
		{Name: "Route 3", TravelTime: 30},
	},
}

func main() {
	for routes.HasNext() {
		route := routes.Next()
		println(route.Name, route.TravelTime)
	}
}
