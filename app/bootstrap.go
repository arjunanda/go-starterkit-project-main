package app

import (
	"go-starterkit-project/modules/auth"
	"go-starterkit-project/modules/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

/**
This function is used to register all modules,
this registration is the last process to register
all modules
*/
func Bootstrap(app *fiber.App) {
	// Monitor app
	app.Get("/monitor", monitor.New())

	// Register module user
	user.RegisterModule(app)

	// Register module auth
	auth.RegisterModule(app)
}
