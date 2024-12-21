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
	// system
	newRoute(http.MethodGet, "/health", controllers.HealthCheck),
	// users
	newRoute(http.MethodGet, "/v1/users/list", controllers.GetUsersListV1),
	newRoute(http.MethodGet, "/v1/users/get", controllers.GetUserV1),
	newRoute(http.MethodPost, "/v1/users/add-loyalty", controllers.AddLoyaltyV1),
	newRoute(http.MethodDelete, "/v1/users/remove-loyalty", controllers.RemoveLoyaltyV1),
	newRoute(http.MethodPost, "/v1/users/create", controllers.CreateUserV1),
	newRoute(http.MethodPut, "/v1/users/update", controllers.UpdateUserV1),
	newRoute(http.MethodPost, "/v1/users/block", controllers.BlockUserV1),
	newRoute(http.MethodPost, "/v1/users/reset-password", controllers.ResetUserPasswordV1),
	newRoute(http.MethodGet, "/v1/users/statistics", controllers.GetStatisticsV1),
	// users categories
	newRoute(http.MethodGet, "/v1/user-category/list", controllers.GetUserCategoriesListV1),
	newRoute(http.MethodGet, "/v1/user-category/get", controllers.HealthCheck),
	newRoute(http.MethodPost, "/v1/user-category/create", controllers.HealthCheck),
	newRoute(http.MethodPut, "/v1/user-category/update", controllers.HealthCheck),
	newRoute(http.MethodDelete, "/v1/user-category/remove", controllers.HealthCheck),
	// items
	newRoute(http.MethodGet, "/v1/items/list", controllers.HealthCheck),
	newRoute(http.MethodGet, "/v1/items/get", controllers.HealthCheck),
	newRoute(http.MethodPost, "/v1/items/create", controllers.HealthCheck),
	newRoute(http.MethodPut, "/v1/items/update", controllers.HealthCheck),
	newRoute(http.MethodDelete, "/v1/items/remove", controllers.HealthCheck),
	// items categories
	newRoute(http.MethodGet, "/v1/item-category/list", controllers.HealthCheck),
	newRoute(http.MethodGet, "/v1/item-category/get", controllers.HealthCheck),
	newRoute(http.MethodPost, "/v1/item-category/create", controllers.HealthCheck),
	newRoute(http.MethodPut, "/v1/item-category/update", controllers.HealthCheck),
	newRoute(http.MethodDelete, "/v1/item-category/remove", controllers.HealthCheck),
}
