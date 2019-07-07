package handler

import (
	"net/http"
	"strconv"

	"github.com/echo-restful-crud-api-example/db"
	"github.com/echo-restful-crud-api-example/types"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// GetProducts is func get all product
func GetProducts(c echo.Context) error {
	data, err := db.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusNotFound, types.ParseStatus("NOT_FOUND", err.Error()))
	}
	return c.JSON(http.StatusOK, data)
}

// CreateProduct is func create new product
func CreateProduct(c echo.Context) error {
	var objRequest types.Product
	if err := c.Bind(&objRequest); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, types.ParseStatus("REQ_ERR", "Có lỗi xảy ra, vui lòng kiểm tra lại thông tin"))
	}
	if err := c.Validate(&objRequest); err != nil {
		return c.JSON(http.StatusBadRequest, types.ParseStatus("REQ_INVALID", err.Error()))
	}

	data, err := db.CreateNewProduct(&objRequest)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, types.ParseStatus("NOT_ACCEPTED", err.Error()))
	}
	return c.JSON(http.StatusCreated, data)
}

// GetProduct is func get one product
func GetProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ParseStatus("REQ_INVALID", "ID invalid"))
	}
	data, err := db.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, types.ParseStatus("NOT_FOUND", err.Error()))
	}
	return c.JSON(http.StatusOK, data)
}

// UpdateProduct is func update one product
func UpdateProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ParseStatus("REQ_INVALID", "ID invalid"))
	}
	var objRequest types.ProductUpdate
	if err := c.Bind(&objRequest); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, types.ParseStatus("REQ_ERR", "Có lỗi xảy ra, vui lòng kiểm tra lại thông tin"))
	}
	if err := c.Validate(&objRequest); err != nil {
		return c.JSON(http.StatusBadRequest, types.ParseStatus("REQ_INVALID", err.Error()))
	}

	data, err := db.UpdateProduct(id, &objRequest)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, types.ParseStatus("NOT_ACCEPTED", err.Error()))
	}
	return c.JSON(http.StatusOK, data)
}

// DeleteProduct is func delete one product
func DeleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ParseStatus("REQ_INVALID", "ID invalid"))
	}
	data, err := db.DeleteAtProduct(id)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, types.ParseStatus("NOT_ACCEPTED", err.Error()))
	}
	return c.JSON(http.StatusOK, data)
}
