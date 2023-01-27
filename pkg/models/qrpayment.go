package models

import "encoding/json"

type QrRequest struct {
	QrType       string  `form:"qr_type" json:"qr_type" validate:"optional" enums:"static,dynamic"`
	TaxId        string  `form:"tax_id" json:"tax_id" validate:"required" minLength:"10" maxLength:"13"`
	Suffix       string  `form:"suffix" json:"suffix" validate:"optional" maxLength:"2"`
	Ref1         string  `form:"ref1" json:"ref1" validate:"required" maxLength:"20"`
	Ref2         string  `form:"ref2" json:"ref2" validate:"optional" maxLength:"20"`
	Ref3         string  `form:"ref3" json:"ref3" validate:"optional" maxLength:"20"`
	Amount       float64 `form:"amount" json:"amount,omitempty" validate:"required" minimum:"0" maximum:"99999999999.99"`
	MerchantName string  `form:"merchan_name" json:"merchan_name,omitempty" validate:"optional" maxLength:"20"`
	Invoice      string  `form:"invoice" json:"invoice" validate:"optional" maxLength:"20"`
}

type QrResponse struct {
	QrType       string  `json:"qr_type" enums:"static,dynamic"`
	TaxId        string  `json:"tax_id"`
	Suffix       string  `json:"suffix"`
	Ref1         string  `json:"ref1"`
	Ref2         string  `json:"ref2"`
	Ref3         string  `json:"ref3"`
	Amount       float64 `json:"amount"`
	MerchantName string  `json:"merchan_name,omitempty"`
	Invoice      string  `json:"invoice,omitempty"`
	Data         string  `json:"data"`
	Image        string  `json:"image"`
}

func (c *QrResponse) FromReq(req QrRequest) (err error) {
	buff, err := json.Marshal(req)
	if err != nil {
		return
	}
	err = json.Unmarshal(buff, c)
	if err != nil {
		return
	}
	return
}
