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

// QianchuanReportUniPromotionDimensionDataRoomGetV10FilteringSmartBidType
type QianchuanReportUniPromotionDimensionDataRoomGetV10FilteringSmartBidType string

// List of qianchuan_report_uni_promotion_dimension_data_room_get_v1.0_filtering_smart_bid_type
const (
	SMART_BID_CONSERVATIVE_QianchuanReportUniPromotionDimensionDataRoomGetV10FilteringSmartBidType QianchuanReportUniPromotionDimensionDataRoomGetV10FilteringSmartBidType = "SMART_BID_CONSERVATIVE"
	SMART_BID_CUSTOM_QianchuanReportUniPromotionDimensionDataRoomGetV10FilteringSmartBidType       QianchuanReportUniPromotionDimensionDataRoomGetV10FilteringSmartBidType = "SMART_BID_CUSTOM"
)

// All allowed values of QianchuanReportUniPromotionDimensionDataRoomGetV10FilteringSmartBidType enum
var AllowedQianchuanReportUniPromotionDimensionDataRoomGetV10FilteringSmartBidTypeEnumValues = []QianchuanReportUniPromotionDimensionDataRoomGetV10FilteringSmartBidType{
	"SMART_BID_CONSERVATIVE",
	"SMART_BID_CUSTOM",
}

// NewQianchuanReportUniPromotionDimensionDataRoomGetV10FilteringSmartBidTypeFromValue returns a pointer to a valid QianchuanReportUniPromotionDimensionDataRoomGetV10FilteringSmartBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportUniPromotionDimensionDataRoomGetV10FilteringSmartBidTypeFromValue(v string) (*QianchuanReportUniPromotionDimensionDataRoomGetV10FilteringSmartBidType, error) {
	ev := QianchuanReportUniPromotionDimensionDataRoomGetV10FilteringSmartBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportUniPromotionDimensionDataRoomGetV10FilteringSmartBidType: valid values are %v", v, AllowedQianchuanReportUniPromotionDimensionDataRoomGetV10FilteringSmartBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportUniPromotionDimensionDataRoomGetV10FilteringSmartBidType) IsValid() bool {
	for _, existing := range AllowedQianchuanReportUniPromotionDimensionDataRoomGetV10FilteringSmartBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_uni_promotion_dimension_data_room_get_v1.0_filtering_smart_bid_type value
func (v QianchuanReportUniPromotionDimensionDataRoomGetV10FilteringSmartBidType) Ptr() *QianchuanReportUniPromotionDimensionDataRoomGetV10FilteringSmartBidType {
	return &v
}
