package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sigurn/crc16"
)

func init() {
}

// AID Application ID
var AID = map[string]string{
	"26":         "A000000677010101",
	"27":         "A000000677010101",
	"29":         "A000000677010111",
	"29OTA":      "A000000677010114",
	"30":         "A000000677010112",
	"31":         "A000000677010113",
	"31API":      "A000000677012004",
	"39":         "A000000677030100",
	"mastercard": "A0000000041010",
}

func leftPad(input string, pad string, length ...int) string {
	padLength := 2
	// pad length default 2
	if len(length) != 0 {
		padLength = length[0]
	}
	for len(input) < padLength {
		input = pad + input
	}
	return input
}

func numberLen(number float64) (numLen int) {
	numLen = len(strconv.FormatFloat(number, 'f', -1, 64))
	return
}

// Tag30CSStr QR Card Scheme
func Tag30CSStr(qrType string, taxID string, suffix string, ref1 string, ref2 string, ref3 string, amount float64, merchantName string, invoice string) (qrString string, err error) {

	var refData string
	refData += "00" + leftPad(strconv.Itoa(len(AID["30"])), "0") + AID["30"]
	refData += "01" + leftPad(strconv.Itoa(len(taxID+suffix)), "0") + taxID + suffix
	refData += "02" + leftPad(strconv.Itoa(len(ref1)), "0") + ref1
	refData += "03" + leftPad(strconv.Itoa(len(ref2)), "0") + ref2

	var addData string
	if len(invoice) != 0 {
		addData += "01" + leftPad(strconv.Itoa(len(invoice)), "0") + invoice
		addData += "05" + leftPad(strconv.Itoa(len(invoice)), "0") + invoice
	}
	if len(ref3) != 0 {
		addData += "07" + leftPad(strconv.Itoa(len(ref3)), "0") + ref3
	}

	// Payload Format
	qrString = "000201"

	// Init Method
	if qrType == "static" {
		qrString += "010211"
	} else {
		qrString += "010212"
	}

	// Merchant Acc
	// qrString += "02164929842783446615"

	// Tag30 data
	qrString += "30" + leftPad(strconv.Itoa(len(refData)), "0") + refData
	// Merchant Category
	// qrString += "52047011"

	// Transaction Currency
	qrString += "5303764"

	// Transaction Amount
	qrString += "54" + leftPad(strconv.Itoa(numberLen(amount)), "0") + strconv.FormatFloat(amount, 'f', -1, 64)

	// Country
	qrString += "5802TH"

	// Merchant Name
	qrString += "59" + leftPad(strconv.Itoa(len(merchantName)), "0") + merchantName

	// Merchant City
	// qrString += "6007BANGKOK"

	// Additionan Data
	if len(addData) != 0 {
		qrString += "62" + leftPad(strconv.Itoa(len(addData)), "0") + addData
	}

	// CRC
	table := crc16.MakeTable(crc16.CRC16_CCITT_FALSE)
	qrString += "6304"
	qrString += strings.ToUpper(
		fmt.Sprintf("%04X", crc16.Checksum([]byte(qrString), table)),
	)

	return
}

// Tag30Str Thai QR Code Tag 30
func Tag30Str(qrType string, taxID string, suffix string, ref1 string, ref2 string, ref3 string, amount float64) (qrString string, err error) {

	var refData string
	refData += "00" + leftPad(strconv.Itoa(len(AID["30"])), "0") + AID["30"]
	refData += "01" + leftPad(strconv.Itoa(len(taxID+suffix)), "0") + taxID + suffix
	refData += "02" + leftPad(strconv.Itoa(len(ref1)), "0") + ref1
	refData += "03" + leftPad(strconv.Itoa(len(ref2)), "0") + ref2

	var addData string
	if len(ref3) != 0 {
		addData += "07" + leftPad(strconv.Itoa(len(ref3)), "0") + ref3
	}

	// Payload Format
	qrString = "000201"

	// Init Method
	if qrType == "static" {
		qrString += "010211"
	} else {
		qrString += "010212"
	}

	// Merchant Acc
	// qrString += "02164929842783446615"

	// Tag30 data
	qrString += "30" + leftPad(strconv.Itoa(len(refData)), "0") + refData
	// Merchant Category
	// qrString += "52047011"

	// Transaction Currency
	qrString += "5303764"

	// Transaction Amount
	qrString += "54" + leftPad(strconv.Itoa(numberLen(amount)), "0") + strconv.FormatFloat(amount, 'f', -1, 64)

	// Country
	qrString += "5802TH"

	// Additionan Data
	if len(addData) != 0 {
		qrString += "62" + leftPad(strconv.Itoa(len(addData)), "0") + addData
	}

	// CRC
	table := crc16.MakeTable(crc16.CRC16_CCITT_FALSE)
	qrString += "6304"
	qrString += strings.ToUpper(
		fmt.Sprintf("%04X", crc16.Checksum([]byte(qrString), table)),
	)

	return
}

// BillPayStr bill payment
func BillPayStr(taxID string, suffix string, ref1 string, ref2 string, amount float64) (qrString string, err error) {

	qrString = "|"
	qrString += taxID + suffix + "\r"
	qrString += ref1 + "\r"
	qrString += ref2 + "\r"
	amountStr := fmt.Sprintf("%.2f", amount)
	qrString += strings.ReplaceAll(amountStr, ".", "")

	return
}
