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

// QianchuanProductAvailableGetV10DataProductListChannelType
type QianchuanProductAvailableGetV10DataProductListChannelType string

// List of qianchuan_product_available_get_v1.0_data_product_list_channel_type
const (
	SHOP_SELL_QianchuanProductAvailableGetV10DataProductListChannelType QianchuanProductAvailableGetV10DataProductListChannelType = "SHOP_SELL"
	STAR_SELL_QianchuanProductAvailableGetV10DataProductListChannelType QianchuanProductAvailableGetV10DataProductListChannelType = "STAR_SELL"
)

// All allowed values of QianchuanProductAvailableGetV10DataProductListChannelType enum
var AllowedQianchuanProductAvailableGetV10DataProductListChannelTypeEnumValues = []QianchuanProductAvailableGetV10DataProductListChannelType{
	"SHOP_SELL",
	"STAR_SELL",
}

// NewQianchuanProductAvailableGetV10DataProductListChannelTypeFromValue returns a pointer to a valid QianchuanProductAvailableGetV10DataProductListChannelType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanProductAvailableGetV10DataProductListChannelTypeFromValue(v string) (*QianchuanProductAvailableGetV10DataProductListChannelType, error) {
	ev := QianchuanProductAvailableGetV10DataProductListChannelType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanProductAvailableGetV10DataProductListChannelType: valid values are %v", v, AllowedQianchuanProductAvailableGetV10DataProductListChannelTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanProductAvailableGetV10DataProductListChannelType) IsValid() bool {
	for _, existing := range AllowedQianchuanProductAvailableGetV10DataProductListChannelTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_product_available_get_v1.0_data_product_list_channel_type value
func (v QianchuanProductAvailableGetV10DataProductListChannelType) Ptr() *QianchuanProductAvailableGetV10DataProductListChannelType {
	return &v
}
