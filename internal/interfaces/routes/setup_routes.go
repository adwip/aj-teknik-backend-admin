package routes

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/infrastructure"
	"github.com/adwip/aj-teknik-backend-admin/internal/handlers/rest"
	"github.com/adwip/aj-teknik-backend-admin/internal/handlers/web"
)

func SetupRoutes(productHandler *rest.ProductHandler, brandHandler *rest.AdmHandler, web *web.WebHandler, server infrastructure.HttpServer) {
	apiRoutes := SetupApiRoutes(productHandler, brandHandler)
	webRoutes := SetupWebRoutes(web)

	apiRoutes.RegisterRoutes(server.RouteInit("/api"))
	webRoutes.RegisterRoutes(server.RouteInit("/"))
}
