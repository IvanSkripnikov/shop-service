package httphandler

import (
	"net/http"
	"regexp"

	"loyalty_system/controllers"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

var routes = []route{
	newRoute(http.MethodGet, "/health", controllers.HealthCheck),
	newRoute(http.MethodGet, "/v1/users/list", controllers.GetUsersListV1),
}
