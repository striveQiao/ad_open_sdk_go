/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalType
type CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalType string

// List of cg_transfer_wallet_transfer_detail_v3.0_data_transfer_wallet_record_list_transfer_capital_record_list_capital_type
const (
	CREDIT_BIDDING_CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalType CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalType = "CREDIT_BIDDING"
	CREDIT_BRAND_CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalType   CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalType = "CREDIT_BRAND"
	CREDIT_GENERAL_CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalType CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalType = "CREDIT_GENERAL"
	PREPAY_BIDDING_CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalType CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalType = "PREPAY_BIDDING"
	PREPAY_BRAND_CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalType   CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalType = "PREPAY_BRAND"
	PREPAY_GENERAL_CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalType CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalType = "PREPAY_GENERAL"
)

// All allowed values of CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalType enum
var AllowedCgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalTypeEnumValues = []CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalType{
	"CREDIT_BIDDING",
	"CREDIT_BRAND",
	"CREDIT_GENERAL",
	"PREPAY_BIDDING",
	"PREPAY_BRAND",
	"PREPAY_GENERAL",
}

// NewCgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalTypeFromValue returns a pointer to a valid CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalTypeFromValue(v string) (*CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalType, error) {
	ev := CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalType: valid values are %v", v, AllowedCgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalType) IsValid() bool {
	for _, existing := range AllowedCgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cg_transfer_wallet_transfer_detail_v3.0_data_transfer_wallet_record_list_transfer_capital_record_list_capital_type value
func (v CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalType) Ptr() *CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalType {
	return &v
}
