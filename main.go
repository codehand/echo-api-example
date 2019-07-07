package main

import (
	"github.com/echo-restful-crud-api-example/config"
	"github.com/echo-restful-crud-api-example/db"
	"github.com/echo-restful-crud-api-example/handler"
	"github.com/echo-restful-crud-api-example/middlewares"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	e := echo.New()
	e.Validator = middlewares.InitCustomValidator()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())

	api := e.Group("/api/v1", serverHeader)
	api.GET("/products", handler.GetProducts)          // Returns all resources of this product
	api.POST("/products", handler.CreateProduct)       // Creates a resource of this product and stores the data you posted, then returns the ID
	api.GET("/products/:id", handler.GetProduct)       // Returns the resource of this product with that ID
	api.PUT("/products/:id", handler.UpdateProduct)    // Updates the resource of this product with that ID
	api.DELETE("/products/:id", handler.DeleteProduct) // Deletes the resource of this product with that ID

	apiV2 := e.Group("/api/v2", serverHeaderVersion2)
	apiV2.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == config.Config.API.Token, nil
	}))
	apiV2.GET("/products", handler.GetProducts)          // Returns all resources of this product
	apiV2.POST("/products", handler.CreateProduct)       // Creates a resource of this product and stores the data you posted, then returns the ID
	apiV2.GET("/products/:id", handler.GetProduct)       // Returns the resource of this product with that ID
	apiV2.PUT("/products/:id", handler.UpdateProduct)    // Updates the resource of this product with that ID
	apiV2.DELETE("/products/:id", handler.DeleteProduct) // Deletes the resource of this product with that ID

	err := db.Ping()
	if err != nil {
		logrus.Fatal(err)
	}

	// service start at port :9090
	err = e.Start(":9090")
	if err != nil {
		logrus.Fatal(err)
	}
}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("x-version", "Test/v1.0")
		return next(c)
	}
}

func serverHeaderVersion2(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("x-version", "Test/v2.0")
		return next(c)
	}
}
