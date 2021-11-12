package routes

import (
	"bookstoreapi/app/controllers"
	"bookstoreapi/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v1")
	route.Post("/book", middleware.JWTProtected(), controllers.CreateBook)
	route.Put("/book", middleware.JWTProtected(), controllers.UpdateBook)
}
