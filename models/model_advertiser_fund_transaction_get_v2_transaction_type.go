/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdvertiserFundTransactionGetV2TransactionType
type AdvertiserFundTransactionGetV2TransactionType string

// List of advertiser_fund_transaction_get_v2_transaction_type
const (
	RECHARGE_AdvertiserFundTransactionGetV2TransactionType AdvertiserFundTransactionGetV2TransactionType = "RECHARGE"
	TRANSFER_AdvertiserFundTransactionGetV2TransactionType AdvertiserFundTransactionGetV2TransactionType = "TRANSFER"
)

// All allowed values of AdvertiserFundTransactionGetV2TransactionType enum
var AllowedAdvertiserFundTransactionGetV2TransactionTypeEnumValues = []AdvertiserFundTransactionGetV2TransactionType{
	"RECHARGE",
	"TRANSFER",
}

// NewAdvertiserFundTransactionGetV2TransactionTypeFromValue returns a pointer to a valid AdvertiserFundTransactionGetV2TransactionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserFundTransactionGetV2TransactionTypeFromValue(v string) (*AdvertiserFundTransactionGetV2TransactionType, error) {
	ev := AdvertiserFundTransactionGetV2TransactionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserFundTransactionGetV2TransactionType: valid values are %v", v, AllowedAdvertiserFundTransactionGetV2TransactionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserFundTransactionGetV2TransactionType) IsValid() bool {
	for _, existing := range AllowedAdvertiserFundTransactionGetV2TransactionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_fund_transaction_get_v2_transaction_type value
func (v AdvertiserFundTransactionGetV2TransactionType) Ptr() *AdvertiserFundTransactionGetV2TransactionType {
	return &v
}
