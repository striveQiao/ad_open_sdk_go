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

// QianchuanAdvertiserTypeGetV10DataListEcpType
type QianchuanAdvertiserTypeGetV10DataListEcpType string

// List of qianchuan_advertiser_type_get_v1.0_data_list_ecp_type
const (
	SHOP_QianchuanAdvertiserTypeGetV10DataListEcpType        QianchuanAdvertiserTypeGetV10DataListEcpType = "SHOP"
	SHOP_STAR_QianchuanAdvertiserTypeGetV10DataListEcpType   QianchuanAdvertiserTypeGetV10DataListEcpType = "SHOP_STAR"
	COMMON_STAR_QianchuanAdvertiserTypeGetV10DataListEcpType QianchuanAdvertiserTypeGetV10DataListEcpType = "COMMON_STAR"
	AGENT_QianchuanAdvertiserTypeGetV10DataListEcpType       QianchuanAdvertiserTypeGetV10DataListEcpType = "AGENT"
)

// All allowed values of QianchuanAdvertiserTypeGetV10DataListEcpType enum
var AllowedQianchuanAdvertiserTypeGetV10DataListEcpTypeEnumValues = []QianchuanAdvertiserTypeGetV10DataListEcpType{
	"SHOP",
	"SHOP_STAR",
	"COMMON_STAR",
	"AGENT",
}

// NewQianchuanAdvertiserTypeGetV10DataListEcpTypeFromValue returns a pointer to a valid QianchuanAdvertiserTypeGetV10DataListEcpType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdvertiserTypeGetV10DataListEcpTypeFromValue(v string) (*QianchuanAdvertiserTypeGetV10DataListEcpType, error) {
	ev := QianchuanAdvertiserTypeGetV10DataListEcpType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdvertiserTypeGetV10DataListEcpType: valid values are %v", v, AllowedQianchuanAdvertiserTypeGetV10DataListEcpTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdvertiserTypeGetV10DataListEcpType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdvertiserTypeGetV10DataListEcpTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_advertiser_type_get_v1.0_data_list_ecp_type value
func (v QianchuanAdvertiserTypeGetV10DataListEcpType) Ptr() *QianchuanAdvertiserTypeGetV10DataListEcpType {
	return &v
}
