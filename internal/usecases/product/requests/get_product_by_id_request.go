package requests

import (
	"github.com/labstack/echo/v5"
)

type GetProductByIdRequest struct {
	SecureID string `param:"id"`
}

func NewGetProductByIdRequest(c *echo.Context) (GetProductByIdRequest, error) {
	var req GetProductByIdRequest
	if err := c.Bind(&req); err != nil {
		return req, err
	}
	return req, nil
}
