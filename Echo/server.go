package main

import (
	"net/http"
	// "fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// practising binding
type UserRequest struct {
	Name  string  `query:"name"`
	ID  string  `param:"id"`
}

// json binding
type LoginDTO struct {
	Email  string `json:"email"`
	Password  string `json:"password"`
}

func main() {
	e := echo.New() // creates new wcho server

	e.Use(middleware.RequestLogger())

	// custom logger configuration
	// e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
	// 	LogMethod:  true,
	// 	LogURI:     true,
	// 	LogStatus:  true,
	// 	LogLatency: true,
	// 	LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
	// 		fmt.Printf("METHOD=%s URI=%s STATUS=%d LATENCY=%v\n",
	// 			v.Method,
	// 			v.URI,
	// 			v.Status,
	// 			v.Latency,
	// 		)
	// 		return nil
	// 	},
	// }))


	// cors configuration
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			echo.GET,
			echo.POST,
		},

		AllowHeaders: []string{
			echo.HeaderContentType,
		},


	}))

	// redirecting to https
	// e.Pre(middleware.HTTPSRedirect())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// cors middleware
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	// hello from query string
	e.GET("/hello", func(c echo.Context) error {
		var req UserRequest
		req.Name = c.QueryParam("name")
		if req.Name == "" {
			req.Name = "World"
		}

		return c.String(http.StatusOK, "Hello, " + req.Name + "!")
	})

	// user from path parameter
	e.GET("/user/:id", func(c echo.Context) error {
		var req UserRequest
		req.ID = c.Param("id")

		if req.ID == "" {
			req.ID = "unknown"
		}
		return c.String(http.StatusOK, "User ID: " + req.ID)
	})

	// login with json body
	e.POST("/login", func(c echo.Context) error {
		var dto LoginDTO
		
		if err := c.Bind(&dto); err != nil {
			return c.String(http.StatusBadRequest, "Invalid request body")
		}

		response := map[string]string{
			"message": "Login successful",
			"email": dto.Email,
		}
		return c.JSON(http.StatusOK, response)
	})

	// start the server
	if err := e.Start(":1323"); err != nil {
		e.Logger.Fatal(err)
	}
}
