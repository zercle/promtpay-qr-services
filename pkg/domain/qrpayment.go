package domain

import "github.com/zercle/promtpay-qr-services/pkg/models"

type QrPaymentUsecase interface {
	BillpaymentQr(req models.QrRequest) (resp models.QrResponse, err error)

	Qr30(req models.QrRequest) (resp models.QrResponse, err error)

	QrCs(req models.QrRequest) (resp models.QrResponse, err error)

	MyPromptQr(req models.QrRequest) (resp models.QrResponse, err error)
}
