package infrastructure

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/zercle/promtpay-qr-services/pkg/qrpayment"

	_ "github.com/zercle/promtpay-qr-services/assets/docs"
)

// SetupRoutes is the Router for GoFiber App
func (s *Server) SetupRoutes(app *fiber.App) {

	// App Services
	qrpaymentUsecase := qrpayment.InitQrPaymentUsecase()

	// App Routes
	qrpayment.NewQrPaymentHandler(app.Group("/api/qr/v1"), qrpaymentUsecase)

	app.Get("/*", swagger.HandlerDefault)
}
