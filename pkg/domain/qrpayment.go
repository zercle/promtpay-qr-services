package domain

import "github.com/zercle/promtpay-qr-services/pkg/models"

type QrPaymentUsecase interface {
	BillpaymentQr(req models.QrRequest) (resp models.QrResponse, imageQr []byte, err error)

	Qr30(req models.QrRequest) (resp models.QrResponse, imageQr []byte, err error)

	QrCs(req models.QrRequest) (resp models.QrResponse, imageQr []byte, err error)

	MyPromptQr(req models.QrRequest) (resp models.QrResponse, imageQr []byte, err error)
}
