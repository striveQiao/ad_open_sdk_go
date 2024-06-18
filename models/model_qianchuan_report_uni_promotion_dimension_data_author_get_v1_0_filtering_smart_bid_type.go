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

// QianchuanReportUniPromotionDimensionDataAuthorGetV10FilteringSmartBidType
type QianchuanReportUniPromotionDimensionDataAuthorGetV10FilteringSmartBidType string

// List of qianchuan_report_uni_promotion_dimension_data_author_get_v1.0_filtering_smart_bid_type
const (
	SMART_BID_CONSERVATIVE_QianchuanReportUniPromotionDimensionDataAuthorGetV10FilteringSmartBidType QianchuanReportUniPromotionDimensionDataAuthorGetV10FilteringSmartBidType = "SMART_BID_CONSERVATIVE"
	SMART_BID_CUSTOM_QianchuanReportUniPromotionDimensionDataAuthorGetV10FilteringSmartBidType       QianchuanReportUniPromotionDimensionDataAuthorGetV10FilteringSmartBidType = "SMART_BID_CUSTOM"
)

// All allowed values of QianchuanReportUniPromotionDimensionDataAuthorGetV10FilteringSmartBidType enum
var AllowedQianchuanReportUniPromotionDimensionDataAuthorGetV10FilteringSmartBidTypeEnumValues = []QianchuanReportUniPromotionDimensionDataAuthorGetV10FilteringSmartBidType{
	"SMART_BID_CONSERVATIVE",
	"SMART_BID_CUSTOM",
}

// NewQianchuanReportUniPromotionDimensionDataAuthorGetV10FilteringSmartBidTypeFromValue returns a pointer to a valid QianchuanReportUniPromotionDimensionDataAuthorGetV10FilteringSmartBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportUniPromotionDimensionDataAuthorGetV10FilteringSmartBidTypeFromValue(v string) (*QianchuanReportUniPromotionDimensionDataAuthorGetV10FilteringSmartBidType, error) {
	ev := QianchuanReportUniPromotionDimensionDataAuthorGetV10FilteringSmartBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportUniPromotionDimensionDataAuthorGetV10FilteringSmartBidType: valid values are %v", v, AllowedQianchuanReportUniPromotionDimensionDataAuthorGetV10FilteringSmartBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportUniPromotionDimensionDataAuthorGetV10FilteringSmartBidType) IsValid() bool {
	for _, existing := range AllowedQianchuanReportUniPromotionDimensionDataAuthorGetV10FilteringSmartBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_uni_promotion_dimension_data_author_get_v1.0_filtering_smart_bid_type value
func (v QianchuanReportUniPromotionDimensionDataAuthorGetV10FilteringSmartBidType) Ptr() *QianchuanReportUniPromotionDimensionDataAuthorGetV10FilteringSmartBidType {
	return &v
}
