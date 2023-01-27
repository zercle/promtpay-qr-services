package infrastructure

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/zercle/promtpay-qr-services/internal/handlers"
	"github.com/zercle/promtpay-qr-services/pkg/qrpayment"

	_ "github.com/zercle/promtpay-qr-services/assets/docs"
)

// SetupRoutes is the Router for GoFiber App
func (s *Server) SetupRoutes(app *fiber.App) {

	// API routes group
	groupApiV1 := app.Group("/api/v:version?", handlers.ApiLimiter)
	{
		groupApiV1.Get("/", handlers.Index())
	}

	// App Services
	qrpaymentUsecase := qrpayment.InitQrPaymentUsecase()

	// App Routes
	qrpayment.NewQrPaymentHandler(app.Group("/api/v1/qr"), qrpaymentUsecase)

	app.Get("/*", swagger.HandlerDefault)
}
