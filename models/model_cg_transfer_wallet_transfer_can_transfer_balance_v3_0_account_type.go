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

// CgTransferWalletTransferCanTransferBalanceV30AccountType
type CgTransferWalletTransferCanTransferBalanceV30AccountType string

// List of cg_transfer_wallet_transfer_can_transfer_balance_v3.0_account_type
const (
	AD_CgTransferWalletTransferCanTransferBalanceV30AccountType        CgTransferWalletTransferCanTransferBalanceV30AccountType = "AD"
	AGENT_CgTransferWalletTransferCanTransferBalanceV30AccountType     CgTransferWalletTransferCanTransferBalanceV30AccountType = "AGENT"
	LOCAL_CgTransferWalletTransferCanTransferBalanceV30AccountType     CgTransferWalletTransferCanTransferBalanceV30AccountType = "LOCAL"
	QIANCHUAN_CgTransferWalletTransferCanTransferBalanceV30AccountType CgTransferWalletTransferCanTransferBalanceV30AccountType = "QIANCHUAN"
)

// All allowed values of CgTransferWalletTransferCanTransferBalanceV30AccountType enum
var AllowedCgTransferWalletTransferCanTransferBalanceV30AccountTypeEnumValues = []CgTransferWalletTransferCanTransferBalanceV30AccountType{
	"AD",
	"AGENT",
	"LOCAL",
	"QIANCHUAN",
}

// NewCgTransferWalletTransferCanTransferBalanceV30AccountTypeFromValue returns a pointer to a valid CgTransferWalletTransferCanTransferBalanceV30AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCgTransferWalletTransferCanTransferBalanceV30AccountTypeFromValue(v string) (*CgTransferWalletTransferCanTransferBalanceV30AccountType, error) {
	ev := CgTransferWalletTransferCanTransferBalanceV30AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CgTransferWalletTransferCanTransferBalanceV30AccountType: valid values are %v", v, AllowedCgTransferWalletTransferCanTransferBalanceV30AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CgTransferWalletTransferCanTransferBalanceV30AccountType) IsValid() bool {
	for _, existing := range AllowedCgTransferWalletTransferCanTransferBalanceV30AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cg_transfer_wallet_transfer_can_transfer_balance_v3.0_account_type value
func (v CgTransferWalletTransferCanTransferBalanceV30AccountType) Ptr() *CgTransferWalletTransferCanTransferBalanceV30AccountType {
	return &v
}
