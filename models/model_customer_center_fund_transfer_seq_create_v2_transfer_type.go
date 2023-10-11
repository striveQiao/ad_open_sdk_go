/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CustomerCenterFundTransferSeqCreateV2TransferType
type CustomerCenterFundTransferSeqCreateV2TransferType string

// List of customer_center_fund_transfer_seq_create_v2_transfer_type
const (
	CREDIT_BRAND_CustomerCenterFundTransferSeqCreateV2TransferType     CustomerCenterFundTransferSeqCreateV2TransferType = "CREDIT_BRAND"
	GRANT_CustomerCenterFundTransferSeqCreateV2TransferType            CustomerCenterFundTransferSeqCreateV2TransferType = "GRANT"
	PREPAY_BID_CustomerCenterFundTransferSeqCreateV2TransferType       CustomerCenterFundTransferSeqCreateV2TransferType = "PREPAY_BID"
	CREDIT_UNIVERSAL_CustomerCenterFundTransferSeqCreateV2TransferType CustomerCenterFundTransferSeqCreateV2TransferType = "CREDIT_UNIVERSAL"
	PREPAY_BRAND_CustomerCenterFundTransferSeqCreateV2TransferType     CustomerCenterFundTransferSeqCreateV2TransferType = "PREPAY_BRAND"
	PREPAY_UNIVERSAL_CustomerCenterFundTransferSeqCreateV2TransferType CustomerCenterFundTransferSeqCreateV2TransferType = "PREPAY_UNIVERSAL"
	CREDIT_BID_CustomerCenterFundTransferSeqCreateV2TransferType       CustomerCenterFundTransferSeqCreateV2TransferType = "CREDIT_BID"
)

// All allowed values of CustomerCenterFundTransferSeqCreateV2TransferType enum
var AllowedCustomerCenterFundTransferSeqCreateV2TransferTypeEnumValues = []CustomerCenterFundTransferSeqCreateV2TransferType{
	"CREDIT_BRAND",
	"GRANT",
	"PREPAY_BID",
	"CREDIT_UNIVERSAL",
	"PREPAY_BRAND",
	"PREPAY_UNIVERSAL",
	"CREDIT_BID",
}

// NewCustomerCenterFundTransferSeqCreateV2TransferTypeFromValue returns a pointer to a valid CustomerCenterFundTransferSeqCreateV2TransferType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomerCenterFundTransferSeqCreateV2TransferTypeFromValue(v string) (*CustomerCenterFundTransferSeqCreateV2TransferType, error) {
	ev := CustomerCenterFundTransferSeqCreateV2TransferType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomerCenterFundTransferSeqCreateV2TransferType: valid values are %v", v, AllowedCustomerCenterFundTransferSeqCreateV2TransferTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomerCenterFundTransferSeqCreateV2TransferType) IsValid() bool {
	for _, existing := range AllowedCustomerCenterFundTransferSeqCreateV2TransferTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to customer_center_fund_transfer_seq_create_v2_transfer_type value
func (v CustomerCenterFundTransferSeqCreateV2TransferType) Ptr() *CustomerCenterFundTransferSeqCreateV2TransferType {
	return &v
}
