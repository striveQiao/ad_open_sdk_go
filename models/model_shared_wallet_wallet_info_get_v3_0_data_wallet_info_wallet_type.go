/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// SharedWalletWalletInfoGetV30DataWalletInfoWalletType
type SharedWalletWalletInfoGetV30DataWalletInfoWalletType string

// List of shared_wallet_wallet_info_get_v3.0_data_wallet_info_wallet_type
const (
	MAIN_WALLET_SharedWalletWalletInfoGetV30DataWalletInfoWalletType        SharedWalletWalletInfoGetV30DataWalletInfoWalletType = "MAIN_WALLET"
	SUB_CONSUME_WALLET_SharedWalletWalletInfoGetV30DataWalletInfoWalletType SharedWalletWalletInfoGetV30DataWalletInfoWalletType = "SUB_CONSUME_WALLET"
	SUB_MANAGE_WALLET_SharedWalletWalletInfoGetV30DataWalletInfoWalletType  SharedWalletWalletInfoGetV30DataWalletInfoWalletType = "SUB_MANAGE_WALLET"
)

// All allowed values of SharedWalletWalletInfoGetV30DataWalletInfoWalletType enum
var AllowedSharedWalletWalletInfoGetV30DataWalletInfoWalletTypeEnumValues = []SharedWalletWalletInfoGetV30DataWalletInfoWalletType{
	"MAIN_WALLET",
	"SUB_CONSUME_WALLET",
	"SUB_MANAGE_WALLET",
}

// NewSharedWalletWalletInfoGetV30DataWalletInfoWalletTypeFromValue returns a pointer to a valid SharedWalletWalletInfoGetV30DataWalletInfoWalletType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSharedWalletWalletInfoGetV30DataWalletInfoWalletTypeFromValue(v string) (*SharedWalletWalletInfoGetV30DataWalletInfoWalletType, error) {
	ev := SharedWalletWalletInfoGetV30DataWalletInfoWalletType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SharedWalletWalletInfoGetV30DataWalletInfoWalletType: valid values are %v", v, AllowedSharedWalletWalletInfoGetV30DataWalletInfoWalletTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SharedWalletWalletInfoGetV30DataWalletInfoWalletType) IsValid() bool {
	for _, existing := range AllowedSharedWalletWalletInfoGetV30DataWalletInfoWalletTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to shared_wallet_wallet_info_get_v3.0_data_wallet_info_wallet_type value
func (v SharedWalletWalletInfoGetV30DataWalletInfoWalletType) Ptr() *SharedWalletWalletInfoGetV30DataWalletInfoWalletType {
	return &v
}
