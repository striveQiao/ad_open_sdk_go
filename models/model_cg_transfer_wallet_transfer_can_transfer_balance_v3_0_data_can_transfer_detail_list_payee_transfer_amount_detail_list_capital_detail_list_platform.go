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

// CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatform
type CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatform string

// List of cg_transfer_wallet_transfer_can_transfer_balance_v3.0_data_can_transfer_detail_list_payee_transfer_amount_detail_list_capital_detail_list_platform
const (
	AD_CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatform        CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatform = "AD"
	AD_ALL_CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatform    CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatform = "AD_ALL"
	BENDITUI_CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatform  CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatform = "BENDITUI"
	QIANCHUAN_CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatform CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatform = "QIANCHUAN"
	STAR_CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatform      CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatform = "STAR"
)

// All allowed values of CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatform enum
var AllowedCgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatformEnumValues = []CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatform{
	"AD",
	"AD_ALL",
	"BENDITUI",
	"QIANCHUAN",
	"STAR",
}

// NewCgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatformFromValue returns a pointer to a valid CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatformFromValue(v string) (*CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatform, error) {
	ev := CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatform: valid values are %v", v, AllowedCgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatform) IsValid() bool {
	for _, existing := range AllowedCgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cg_transfer_wallet_transfer_can_transfer_balance_v3.0_data_can_transfer_detail_list_payee_transfer_amount_detail_list_capital_detail_list_platform value
func (v CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatform) Ptr() *CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListPayeeTransferAmountDetailListCapitalDetailListPlatform {
	return &v
}
