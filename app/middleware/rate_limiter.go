package middleware

import (
	"go-starterkit-project/utils"
	"time"

	respModel "go-starterkit-project/domain/data_models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

/**
This function is used as middleware for rate limit request
*/
func RateLimiter(max int, duration time.Duration) func(ctx *fiber.Ctx) error {
	return limiter.New(limiter.Config{
		LimitReached: func(ctx *fiber.Ctx) error {
			return utils.ApiResponseError(ctx, "Rate limit", fiber.StatusRequestEntityTooLarge, &respModel.ApiErrorResponse{
				StatusCode: fiber.StatusTooManyRequests,
				Message:    "Request too many request",
			})
		},
		Max:        max,
		Expiration: duration * time.Second,
	})
}
