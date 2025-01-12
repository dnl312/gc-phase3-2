package helpers

import (
	"client/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetClaims(c echo.Context) (model.Claims, error) {
	claimsTmp := c.Get("user")
	//utils.ErrUnauthorized.EchoFormatDetails("Failed to fetch user claims from JWT")
	if claimsTmp == nil {
		return model.Claims{}, echo.NewHTTPError(echo.ErrUnauthorized.Code, "Failed to fetch user claims from JWT")
	}

	claims := claimsTmp.(jwt.MapClaims)
	return model.Claims{
		Exp: claims["exp"].(float64),
		UserID:  claims["userid"].(string),
	}, nil
}