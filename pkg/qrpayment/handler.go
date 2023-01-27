package qrpayment

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/promtpay-qr-services/pkg/domain"
	"github.com/zercle/promtpay-qr-services/pkg/models"
)

type QrPaymentHandler struct {
	QrPaymentUsecase domain.QrPaymentUsecase
}

func NewQrPaymentHandler(appRoute fiber.Router, qrPaymentUsecase domain.QrPaymentUsecase) {
	handler := &QrPaymentHandler{
		QrPaymentUsecase: qrPaymentUsecase,
	}

	appRoute.Get("/billpayment", handler.GetBillpaymentQr())
	appRoute.Post("/billpayment", handler.PostBillpaymentQr())

	appRoute.Get("/qr30", handler.GetQr30())
	appRoute.Post("/qr30", handler.PostQr30())

	appRoute.Get("/qrcs", handler.GetQrCs())
	appRoute.Post("/qrcs", handler.PostQrCs())

	appRoute.Get("/mypromptqr", handler.GetMyPromptQr())
	appRoute.Post("/mypromptqr", handler.PostMyPromptQr())
}

// Bill Payment QR
func (h *QrPaymentHandler) billpaymentQr(c *fiber.Ctx) (err error) {
	var req models.QrRequest
	var resp helpers.ResponseForm

	respCode := http.StatusInternalServerError

	if err = c.BodyParser(&req); errors.Is(err, fiber.ErrUnprocessableEntity) {
		err = c.QueryParser(&req)
	}

	if err != nil {
		return
	}

	resp.Result, err = h.QrPaymentUsecase.BillpaymentQr(req)

	if err == nil {
		resp.Success = true
		respCode = http.StatusOK

	}
	return c.Status(respCode).JSON(resp)
}

// @Param request formData models.QrRequest true "transaction detail"
// @Success 200 {object} helpers.ResponseForm
// @Router /api/v1/qr/billpayment [get]
func (h *QrPaymentHandler) GetBillpaymentQr() fiber.Handler {
	return h.billpaymentQr
}

// @Param request body models.QrRequest true "transaction detail"
// @Success 200 {object} helpers.ResponseForm
// @Router /api/v1/qr/billpayment [post]
func (h *QrPaymentHandler) PostBillpaymentQr() fiber.Handler {
	return h.billpaymentQr
}

// Thai QR Code Tag 30 (QR 30)
func (h *QrPaymentHandler) qr30(c *fiber.Ctx) (err error) {
	var req models.QrRequest
	var resp helpers.ResponseForm

	respCode := http.StatusInternalServerError

	if err = c.BodyParser(&req); errors.Is(err, fiber.ErrUnprocessableEntity) {
		err = c.QueryParser(&req)
	}

	if err != nil {
		return
	}

	resp.Result, err = h.QrPaymentUsecase.Qr30(req)

	if err == nil {
		resp.Success = true
		respCode = http.StatusOK

	}
	return c.Status(respCode).JSON(resp)
}

// @Param request formData models.QrRequest true "transaction detail"
// @Success 200 {object} helpers.ResponseForm
// @Router /api/v1/qr/qr30 [get]
func (h *QrPaymentHandler) GetQr30() fiber.Handler {
	return h.qr30
}

// @Param request body models.QrRequest true "transaction detail"
// @Success 200 {object} helpers.ResponseForm
// @Router /api/v1/qr/qr30 [post]
func (h *QrPaymentHandler) PostQr30() fiber.Handler {
	return h.qr30
}

// QR Card Scheme (QR CS)
func (h *QrPaymentHandler) qrCs(c *fiber.Ctx) (err error) {
	var req models.QrRequest
	var resp helpers.ResponseForm

	respCode := http.StatusInternalServerError

	if err = c.BodyParser(&req); errors.Is(err, fiber.ErrUnprocessableEntity) {
		err = c.QueryParser(&req)
	}

	if err != nil {
		return
	}

	resp.Result, err = h.QrPaymentUsecase.QrCs(req)

	if err == nil {
		resp.Success = true
		respCode = http.StatusOK

	}
	return c.Status(respCode).JSON(resp)
}

// @Param request formData models.QrRequest true "transaction detail"
// @Success 200 {object} helpers.ResponseForm
// @Router /api/v1/qr/qrcs [get]
func (h *QrPaymentHandler) GetQrCs() fiber.Handler {
	return h.qrCs
}

// @Param request body models.QrRequest true "transaction detail"
// @Success 200 {object} helpers.ResponseForm
// @Router /api/v1/qr/qrcs [post]
func (h *QrPaymentHandler) PostQrCs() fiber.Handler {
	return h.qrCs
}

// MyPrompt QR
func (h *QrPaymentHandler) myPromptQr(c *fiber.Ctx) (err error) {
	var req models.QrRequest
	var resp helpers.ResponseForm

	respCode := http.StatusInternalServerError

	if err = c.BodyParser(&req); errors.Is(err, fiber.ErrUnprocessableEntity) {
		err = c.QueryParser(&req)
	}

	if err != nil {
		return
	}

	resp.Result, err = h.QrPaymentUsecase.MyPromptQr(req)

	if err == nil {
		resp.Success = true
		respCode = http.StatusOK

	}
	return c.Status(respCode).JSON(resp)
}

// @Param get_request formData models.QrRequest true "transaction detail"
// @Success 200 {object} helpers.ResponseForm
// @Router /api/v1/qr/mypromptqr [get]
func (h *QrPaymentHandler) GetMyPromptQr() fiber.Handler {
	return h.myPromptQr
}

// @Param post_request body models.QrRequest true "transaction detail"
// @Success 200 {object} helpers.ResponseForm
// @Router /api/v1/qr/mypromptqr [post]
func (h *QrPaymentHandler) PostMyPromptQr() fiber.Handler {
	return h.myPromptQr
}
