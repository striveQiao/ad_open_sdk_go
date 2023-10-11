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

// CustomerCenterAdvertiserListV2AccountSource
type CustomerCenterAdvertiserListV2AccountSource string

// List of customer_center_advertiser_list_v2_account_source
const (
	ENTERPRISE_CustomerCenterAdvertiserListV2AccountSource CustomerCenterAdvertiserListV2AccountSource = "ENTERPRISE"
	AD_CustomerCenterAdvertiserListV2AccountSource         CustomerCenterAdvertiserListV2AccountSource = "AD"
)

// All allowed values of CustomerCenterAdvertiserListV2AccountSource enum
var AllowedCustomerCenterAdvertiserListV2AccountSourceEnumValues = []CustomerCenterAdvertiserListV2AccountSource{
	"ENTERPRISE",
	"AD",
}

// NewCustomerCenterAdvertiserListV2AccountSourceFromValue returns a pointer to a valid CustomerCenterAdvertiserListV2AccountSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomerCenterAdvertiserListV2AccountSourceFromValue(v string) (*CustomerCenterAdvertiserListV2AccountSource, error) {
	ev := CustomerCenterAdvertiserListV2AccountSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomerCenterAdvertiserListV2AccountSource: valid values are %v", v, AllowedCustomerCenterAdvertiserListV2AccountSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomerCenterAdvertiserListV2AccountSource) IsValid() bool {
	for _, existing := range AllowedCustomerCenterAdvertiserListV2AccountSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to customer_center_advertiser_list_v2_account_source value
func (v CustomerCenterAdvertiserListV2AccountSource) Ptr() *CustomerCenterAdvertiserListV2AccountSource {
	return &v
}
