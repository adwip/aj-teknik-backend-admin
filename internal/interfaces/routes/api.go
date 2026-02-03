package routes

import (
	"github.com/adwip/aj-teknik-backend-admin/internal/handlers/rest"
	"github.com/labstack/echo/v5"
)

type apiRoutes struct {
	product *rest.ProductHandler
	brand   *rest.AdmHandler
}

func SetupApiRoutes(product *rest.ProductHandler, brand *rest.AdmHandler) *apiRoutes {
	return &apiRoutes{
		product: product,
		brand:   brand,
	}
}

func (a *apiRoutes) RegisterRoutes(g *echo.Group) {
	// products
	g.GET("/products", a.product.GetProducts)
	g.GET("/product/:id", a.product.GetProductById)
	g.POST("/products", a.product.CreateProduct)
	// g.PUT("/products/:id", a.rest.UpdateProduct)
	// g.DELETE("/products/:id", a.rest.DeleteProduct)

	// brands
	g.GET("/brands", a.brand.GetBrands)
	g.POST("/brands", a.brand.AddBrand)
	g.PUT("/brands/:id", a.brand.UpdateBrand)
	g.DELETE("/brands/:id", a.brand.DeleteBrand)

	// categories
	g.GET("/categories/tree", a.brand.GetCategoryTree)
	g.GET("/categories", a.brand.GetCategories)
	g.POST("/categories", a.brand.AddCategory)
	g.PUT("/categories/:id", a.brand.UpdateCategory)
	g.DELETE("/categories/:id", a.brand.DeleteCategory)
}
