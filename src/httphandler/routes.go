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
	// test
	newRoute(http.MethodGet, "/test/error", controllers.TestError),
	newRoute(http.MethodGet, "/test/latency", controllers.TestLongLatency),
	// users
	newRoute(http.MethodGet, "/v1/users/me", controllers.GetMyInfoV1),
	newRoute(http.MethodPut, "/v1/users/me", controllers.GetMyInfoV1),
	newRoute(http.MethodPut, "/v1/users/me/deposit", controllers.MyDepositV1),
	newRoute(http.MethodGet, "/v1/users/list", controllers.GetUsersListV1),
	newRoute(http.MethodGet, "/v1/users/list", controllers.GetActiveUsersListV1),
	newRoute(http.MethodGet, "/v1/users/get/([0-9]+)", controllers.GetUserV1),
	newRoute(http.MethodPost, "/v1/users/add-loyalty", controllers.AddLoyaltyV1),
	newRoute(http.MethodDelete, "/v1/users/remove-loyalty", controllers.RemoveLoyaltyV1),
	newRoute(http.MethodPost, "/v1/users/create", controllers.CreateUserV1),
	newRoute(http.MethodPut, "/v1/users/update", controllers.UpdateUserV1),
	newRoute(http.MethodPut, "/v1/users/category-update", controllers.UserCategoryUpdateV1),
	newRoute(http.MethodDelete, "/v1/users/delete/([0-9]+)", controllers.BlockUserV1),
	newRoute(http.MethodPost, "/v1/users/reset-password", controllers.ResetUserPasswordV1),
	newRoute(http.MethodGet, "/v1/users/statistics", controllers.GetStatisticsV1),
	// users categories
	newRoute(http.MethodGet, "/v1/user-category/list", controllers.GetUserCategoriesListV1),
	newRoute(http.MethodGet, "/v1/user-category/get/([0-9]+)", controllers.GetUserCategoryV1),
	newRoute(http.MethodGet, "/v1/user-category/get-by-user/([0-9]+)", controllers.GetCategoryByUserV1),
	newRoute(http.MethodPost, "/v1/user-category/create", controllers.CreateUserCategoryV1),
	newRoute(http.MethodPut, "/v1/user-category/update", controllers.UpdateUserCategoryV1),
	newRoute(http.MethodDelete, "/v1/user-category/remove", controllers.RemoveUserCategoryV1),
	// items
	newRoute(http.MethodGet, "/v1/items/list", controllers.GetItemsListV1),
	newRoute(http.MethodGet, "/v1/items/get/([0-9]+)", controllers.GetItemV1),
	newRoute(http.MethodPost, "/v1/items/create", controllers.CreateItemV1),
	newRoute(http.MethodPut, "/v1/items/update", controllers.UpdateItemV1),
	newRoute(http.MethodDelete, "/v1/items/remove/([0-9]+)", controllers.RemoveItemV1),
	newRoute(http.MethodPost, "/v1/items/buy", controllers.BuyItemV1),
	// items categories
	newRoute(http.MethodGet, "/v1/item-category/list", controllers.GetItemsCatogoryListV1),
	newRoute(http.MethodGet, "/v1/item-category/get", controllers.GetItemCategoryV1),
	newRoute(http.MethodPost, "/v1/item-category/create", controllers.CreateItemCategoryV1),
	newRoute(http.MethodPut, "/v1/item-category/update", controllers.UpdateItemCategoryV1),
	newRoute(http.MethodDelete, "/v1/item-category/remove", controllers.RemoveItemCategoryV1),
}
