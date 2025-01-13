package router

import (
	"client/controller"
	"fmt"

	//_ "api-gateway/docs"
	middleware "client/middleware"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
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
	e.GET("/books/update-status", bc.UpdateBookStatus)
	c := cron.New()
	_, err := c.AddFunc("@every 5s", func() {
		_, err := http.Get("http://localhost:8080/books/update-status")
		if err != nil {
			fmt.Println("Error triggering update-status endpoint:", err)
		}
	})
	if err != nil {
		// Handle error
		fmt.Println("Error setting up cron job:", err)
	}
	c.Start()

	// c := cron.New()
    // _, err := c.AddFunc("@every 5s", func() {
    //     _, err := http.Get("http://localhost:8080/books/update-status")
    //     if err != nil {
    //         fmt.Println("Error triggering update-status endpoint:", err)
    //     }
    // })
    // if err != nil {
    //     // Handle error
    //     fmt.Println("Error setting up cron job:", err)
    // }
    // c.Start()
	
	// payment := e.Group("/payment")
	// {
	// 	payment.POST("/xendit/update", oc.PaymentUpdate)
	// }
}