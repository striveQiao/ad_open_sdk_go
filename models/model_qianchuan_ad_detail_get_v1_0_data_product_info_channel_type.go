/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdDetailGetV10DataProductInfoChannelType
type QianchuanAdDetailGetV10DataProductInfoChannelType string

// List of qianchuan_ad_detail_get_v1.0_data_product_info_channel_type
const (
	SHOP_SELL_QianchuanAdDetailGetV10DataProductInfoChannelType QianchuanAdDetailGetV10DataProductInfoChannelType = "SHOP_SELL"
	STAR_SELL_QianchuanAdDetailGetV10DataProductInfoChannelType QianchuanAdDetailGetV10DataProductInfoChannelType = "STAR_SELL"
)

// All allowed values of QianchuanAdDetailGetV10DataProductInfoChannelType enum
var AllowedQianchuanAdDetailGetV10DataProductInfoChannelTypeEnumValues = []QianchuanAdDetailGetV10DataProductInfoChannelType{
	"SHOP_SELL",
	"STAR_SELL",
}

// NewQianchuanAdDetailGetV10DataProductInfoChannelTypeFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataProductInfoChannelType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataProductInfoChannelTypeFromValue(v string) (*QianchuanAdDetailGetV10DataProductInfoChannelType, error) {
	ev := QianchuanAdDetailGetV10DataProductInfoChannelType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataProductInfoChannelType: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataProductInfoChannelTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataProductInfoChannelType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataProductInfoChannelTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_product_info_channel_type value
func (v QianchuanAdDetailGetV10DataProductInfoChannelType) Ptr() *QianchuanAdDetailGetV10DataProductInfoChannelType {
	return &v
}
