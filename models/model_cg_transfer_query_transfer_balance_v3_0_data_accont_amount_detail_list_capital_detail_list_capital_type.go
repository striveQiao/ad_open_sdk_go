/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalType
type CgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalType string

// List of cg_transfer_query_transfer_balance_v3.0_data_accont_amount_detail_list_capital_detail_list_capital_type
const (
	CREDIT_BIDDING_CgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalType CgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalType = "CREDIT_BIDDING"
	CREDIT_BRAND_CgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalType   CgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalType = "CREDIT_BRAND"
	CREDIT_GENERAL_CgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalType CgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalType = "CREDIT_GENERAL"
	PREPAY_BIDDING_CgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalType CgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalType = "PREPAY_BIDDING"
	PREPAY_BRAND_CgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalType   CgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalType = "PREPAY_BRAND"
	PREPAY_GENERAL_CgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalType CgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalType = "PREPAY_GENERAL"
)

// All allowed values of CgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalType enum
var AllowedCgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalTypeEnumValues = []CgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalType{
	"CREDIT_BIDDING",
	"CREDIT_BRAND",
	"CREDIT_GENERAL",
	"PREPAY_BIDDING",
	"PREPAY_BRAND",
	"PREPAY_GENERAL",
}

// NewCgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalTypeFromValue returns a pointer to a valid CgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalTypeFromValue(v string) (*CgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalType, error) {
	ev := CgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalType: valid values are %v", v, AllowedCgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalType) IsValid() bool {
	for _, existing := range AllowedCgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cg_transfer_query_transfer_balance_v3.0_data_accont_amount_detail_list_capital_detail_list_capital_type value
func (v CgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalType) Ptr() *CgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalType {
	return &v
}
