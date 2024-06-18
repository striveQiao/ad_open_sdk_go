/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapital
type CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapital string

// List of cg_transfer_wallet_transfer_list_v3.0_data_record_list_transfer_capital_record_list_transfer_capital
const (
	CREDIT_BIDDING_CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapital CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapital = "CREDIT_BIDDING"
	CREDIT_BRAND_CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapital   CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapital = "CREDIT_BRAND"
	CREDIT_GENERAL_CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapital CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapital = "CREDIT_GENERAL"
	PREPAY_BIDDING_CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapital CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapital = "PREPAY_BIDDING"
	PREPAY_BRAND_CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapital   CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapital = "PREPAY_BRAND"
	PREPAY_GENERAL_CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapital CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapital = "PREPAY_GENERAL"
)

// All allowed values of CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapital enum
var AllowedCgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapitalEnumValues = []CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapital{
	"CREDIT_BIDDING",
	"CREDIT_BRAND",
	"CREDIT_GENERAL",
	"PREPAY_BIDDING",
	"PREPAY_BRAND",
	"PREPAY_GENERAL",
}

// NewCgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapitalFromValue returns a pointer to a valid CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapital
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapitalFromValue(v string) (*CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapital, error) {
	ev := CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapital(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapital: valid values are %v", v, AllowedCgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapitalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapital) IsValid() bool {
	for _, existing := range AllowedCgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapitalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cg_transfer_wallet_transfer_list_v3.0_data_record_list_transfer_capital_record_list_transfer_capital value
func (v CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapital) Ptr() *CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapital {
	return &v
}
