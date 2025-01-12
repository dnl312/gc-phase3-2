package router

import (
	"client/controller"
	//_ "api-gateway/docs"
	middleware "client/middleware"

	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Echo(e *echo.Echo, uc controller.UserController, bc controller.BookController) {
	e.GET("/users/verified", func(c echo.Context) error {
		return c.Render(http.StatusOK, "verification.html", nil)
	})
	
	e.GET("", func(c echo.Context) error {
		return c.Redirect(http.StatusTemporaryRedirect, "/swagger/index.html")
	})
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	users := e.Group("/users")
	{		
		users.POST("/login", uc.LoginUser)
		users.POST("/register", uc.RegisterUser)
		//users.GET("/logout", uc.Logout, middlewares.RequireAuth)
	}

	book := e.Group("books")
	book.Use(middleware.RequireAuth)
	{	
		book.GET("", bc.GetAllBooks)
		book.POST("/borrow", bc.BorrowBook)	
		book.POST("/return", bc.ReturnBook)
	}
	
	// payment := e.Group("/payment")
	// {
	// 	payment.POST("/xendit/update", oc.PaymentUpdate)
	// }
}