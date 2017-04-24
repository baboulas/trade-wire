package main

import (
	jwt "github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/middleware/logger"

	"bitbucket.com/gregtandiono_/trade-wire/adaptors"
	controller "bitbucket.com/gregtandiono_/trade-wire/controllers"
)

func main() {

	db := adaptors.DBConnector()

	app := iris.New()
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())

	customLogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
	})

	// auth middleware auth
	myJwtMiddleware := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("supersecretkey"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	app.Use(customLogger)

	app.Post("/login", controller.NewUserController(db).Login)
	app.Put("/user/:id", myJwtMiddleware.Serve, controller.NewUserController(db).Update)
	app.Delete("/user/:id", controller.NewUserController(db).Delete)
	app.Get("/fetch", myJwtMiddleware.Serve, controller.NewUserController(db).FetchAll)
	app.Post("/register", controller.NewUserController(db).Register)

	app.Listen(":8080")

}
