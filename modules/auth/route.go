package auth

import (
	"go-starterkit-project/app/middleware"
	"go-starterkit-project/modules/auth/handlers"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
)

type ApiRoute struct {
	AuthHandler handlers.AuthHandler
}

func (handler *ApiRoute) Route(app fiber.Router) {
	const endpointGroup string = "/auth"

	user := app.Group(utils.SetupApiGroup() + endpointGroup)

	user.Post("/", middleware.RateLimiter(5, 120), handler.AuthHandler.Authentication)

	user.Get("/me", middleware.Authenticate(), middleware.RateLimiter(5, 30), handler.AuthHandler.GetProfile)

	user.Get("/refresh-token", middleware.Authenticate(), middleware.RateLimiter(5, 30), handler.AuthHandler.RefreshToken)

	user.Put("/update-user", middleware.Authenticate(), middleware.RateLimiter(5, 30), handler.AuthHandler.UpdateProfile)

	// app.Static("/public/", config.ProjectRootPath+"/public")

}
