/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DpaProductUpdateV2ProductInfoStatus
type DpaProductUpdateV2ProductInfoStatus int64

// List of dpa_product_update_v2_product_info_status
const (
	Enum_0_DpaProductUpdateV2ProductInfoStatus DpaProductUpdateV2ProductInfoStatus = 0
	Enum_1_DpaProductUpdateV2ProductInfoStatus DpaProductUpdateV2ProductInfoStatus = 1
)

// All allowed values of DpaProductUpdateV2ProductInfoStatus enum
var AllowedDpaProductUpdateV2ProductInfoStatusEnumValues = []DpaProductUpdateV2ProductInfoStatus{
	0,
	1,
}

// NewDpaProductUpdateV2ProductInfoStatusFromValue returns a pointer to a valid DpaProductUpdateV2ProductInfoStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaProductUpdateV2ProductInfoStatusFromValue(v int64) (*DpaProductUpdateV2ProductInfoStatus, error) {
	ev := DpaProductUpdateV2ProductInfoStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaProductUpdateV2ProductInfoStatus: valid values are %v", v, AllowedDpaProductUpdateV2ProductInfoStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaProductUpdateV2ProductInfoStatus) IsValid() bool {
	for _, existing := range AllowedDpaProductUpdateV2ProductInfoStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_product_update_v2_product_info_status value
func (v DpaProductUpdateV2ProductInfoStatus) Ptr() *DpaProductUpdateV2ProductInfoStatus {
	return &v
}
