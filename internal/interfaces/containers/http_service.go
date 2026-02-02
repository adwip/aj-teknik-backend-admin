package containers

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/infrastructure"
	"github.com/adwip/aj-teknik-backend-admin/common-lib/logger"
	"github.com/adwip/aj-teknik-backend-admin/internal/handlers/rest"
	"github.com/adwip/aj-teknik-backend-admin/internal/handlers/web"
	"github.com/adwip/aj-teknik-backend-admin/internal/interfaces/drivers"
	"github.com/adwip/aj-teknik-backend-admin/internal/interfaces/routes"
	"github.com/adwip/aj-teknik-backend-admin/internal/repositories/mysql"
	"github.com/adwip/aj-teknik-backend-admin/internal/shared/config"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/administration"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/product"
)

func SetupServiceContainer() (err error) {
	log, report, err := logger.SetupLogger("")
	if err != nil {
		return err
	}

	httpServer := infrastructure.SetupHttpServer(log)

	_, err = config.SetupConfig()
	if err != nil {
		return err
	}

	db, err := drivers.SetupDatabase("root:12345678@tcp(localhost:3306)/aj_teknik_db?parseTime=true")
	if err != nil {
		return err
	}

	// setup repo
	_ = mysql.SetupUsersRepository(db)
	productRepo := mysql.SetupProductsRepository(db)
	brandRepo := mysql.SetupBrandRepository(db)

	// setup usecase
	// _ = admin.SetupAdminUsecase()
	productUsecase := product.SetupProductUsecase(report, productRepo, brandRepo)
	brandUsecase := administration.SetupAdministrationUsecase(brandRepo)

	// setup handler
	productHandler := rest.SetupProductHandler(productUsecase)
	brandHandler := rest.SetupAdministrationHandler(brandUsecase)
	web := web.SetupWebHandlers()

	routes.SetupRoutes(productHandler, brandHandler, web, httpServer)

	err = httpServer.StartServer(":8080")
	if err != nil {
		return err
	}
	return nil
}
