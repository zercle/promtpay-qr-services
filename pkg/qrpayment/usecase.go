package qrpayment

import (
	"encoding/base64"

	qrcode "github.com/skip2/go-qrcode"

	"github.com/zercle/promtpay-qr-services/pkg/domain"
	"github.com/zercle/promtpay-qr-services/pkg/models"
	"github.com/zercle/promtpay-qr-services/pkg/utils"
)

type promptpayUsecase struct {
}

func InitQrPaymentUsecase() domain.QrPaymentUsecase {
	return &promptpayUsecase{}
}

// Bill Bayment BR
// This is to support merchant-presented mode (C scan B) QR codes, where the customer scans the merchant’s QR code and pays using either the current or savings account as the source of funds. This type of QR payment is supported by most of the major Thai banks.
func (u *promptpayUsecase) BillpaymentQr(req models.QrRequest) (resp models.QrResponse, imageQr []byte, err error) {

	err = resp.FromReq(req)
	if err != nil {
		return
	}

	resp.Data, err = utils.BillPayStr(
		req.TaxId,
		req.Suffix,
		req.Ref1,
		req.Ref2,
		req.Amount,
	)
	if err != nil {
		return
	}

	imageQr, err = qrcode.Encode(resp.Data, qrcode.High, 512)
	if err != nil {
		return
	}

	resp.Image = "data:image/png;base64, " + base64.StdEncoding.EncodeToString(imageQr)
	return
}

// Thai QR Code Tag 30 (QR 30)
// This is to support merchant-presented mode (C scan B) QR codes, where the customer scans the merchant’s QR code and pays using either the current or savings account as the source of funds. This type of QR payment is supported by most of the major Thai banks.
func (u *promptpayUsecase) Qr30(req models.QrRequest) (resp models.QrResponse, imageQr []byte, err error) {

	err = resp.FromReq(req)
	if err != nil {
		return
	}

	resp.Data, err = utils.Tag30Str(
		req.QrType,
		req.TaxId,
		req.Suffix,
		req.Ref1,
		req.Ref2,
		req.Ref3,
		req.Amount,
	)
	if err != nil {
		return
	}

	imageQr, err = qrcode.Encode(resp.Data, qrcode.High, 512)
	if err != nil {
		return
	}

	resp.Image = "data:image/png;base64, " + base64.StdEncoding.EncodeToString(imageQr)

	return
}

// QR Card Scheme (QR CS)
// This is to support merchant-presented mode (C scan B) QR codes, where the customer scans the merchant’s QR code and pays using credit cards as the source of funds. Currently, QR CS is supported by VISA and Mastercard, making it compatible internationally.
func (u *promptpayUsecase) QrCs(req models.QrRequest) (resp models.QrResponse, imageQr []byte, err error) {

	err = resp.FromReq(req)
	if err != nil {
		return
	}

	resp.Data, err = utils.Tag30CSStr(
		req.QrType,
		req.TaxId,
		req.Suffix,
		req.Ref1,
		req.Ref2,
		req.Ref3,
		req.Amount,
		req.MerchantName,
		req.Invoice,
	)
	if err != nil {
		return
	}

	imageQr, err = qrcode.Encode(resp.Data, qrcode.High, 512)
	if err != nil {
		return
	}

	resp.Image = "data:image/png;base64, " + base64.StdEncoding.EncodeToString(imageQr)

	return
}

// MyPrompt QR
// This is to support consumer-presented mode (B Scan C) QR codes, where the merchant scans the customer’s QR Code and deducts the payment amount from their current or savings account. Major Thai banks will start launching the support of MyPrompt QR in the near future.
func (u *promptpayUsecase) MyPromptQr(req models.QrRequest) (resp models.QrResponse, imageQr []byte, err error) {

	err = resp.FromReq(req)
	if err != nil {
		return
	}

	resp.Data, err = utils.Tag30CSStr(
		req.QrType,
		req.TaxId,
		req.Suffix,
		req.Ref1,
		req.Ref2,
		req.Ref3,
		req.Amount,
		req.MerchantName,
		req.Invoice,
	)
	if err != nil {
		return
	}

	imageQr, err = qrcode.Encode(resp.Data, qrcode.High, 512)
	if err != nil {
		return
	}

	resp.Image = "data:image/png;base64, " + base64.StdEncoding.EncodeToString(imageQr)
	return
}
