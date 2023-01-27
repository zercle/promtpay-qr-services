package mocks

import (
	mock "github.com/stretchr/testify/mock"
	"github.com/zercle/promtpay-qr-services/pkg/models"
)

type QrPaymentUsecase struct {
	mock.Mock
}

func (_m *QrPaymentUsecase) Qr30(req models.QrRequest) (resp models.QrResponse, err error) {
	ret := _m.Called(req)

	var r0 models.QrResponse
	if rf, ok := ret.Get(0).(func(models.QrRequest) models.QrResponse); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(models.QrResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.QrRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *QrPaymentUsecase) QrCs(req models.QrRequest) (resp models.QrResponse, err error) {
	ret := _m.Called(req)

	var r0 models.QrResponse
	if rf, ok := ret.Get(0).(func(models.QrRequest) models.QrResponse); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(models.QrResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.QrRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *QrPaymentUsecase) MyPromptQr(req models.QrRequest) (resp models.QrResponse, err error) {
	ret := _m.Called(req)

	var r0 models.QrResponse
	if rf, ok := ret.Get(0).(func(models.QrRequest) models.QrResponse); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(models.QrResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.QrRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
