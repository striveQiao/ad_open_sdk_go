/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportUniPromotionDimensionDataAuthorGetV10OrderType
type QianchuanReportUniPromotionDimensionDataAuthorGetV10OrderType string

// List of qianchuan_report_uni_promotion_dimension_data_author_get_v1.0_order_type
const (
	ASC_QianchuanReportUniPromotionDimensionDataAuthorGetV10OrderType  QianchuanReportUniPromotionDimensionDataAuthorGetV10OrderType = "ASC"
	DESC_QianchuanReportUniPromotionDimensionDataAuthorGetV10OrderType QianchuanReportUniPromotionDimensionDataAuthorGetV10OrderType = "DESC"
)

// All allowed values of QianchuanReportUniPromotionDimensionDataAuthorGetV10OrderType enum
var AllowedQianchuanReportUniPromotionDimensionDataAuthorGetV10OrderTypeEnumValues = []QianchuanReportUniPromotionDimensionDataAuthorGetV10OrderType{
	"ASC",
	"DESC",
}

// NewQianchuanReportUniPromotionDimensionDataAuthorGetV10OrderTypeFromValue returns a pointer to a valid QianchuanReportUniPromotionDimensionDataAuthorGetV10OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportUniPromotionDimensionDataAuthorGetV10OrderTypeFromValue(v string) (*QianchuanReportUniPromotionDimensionDataAuthorGetV10OrderType, error) {
	ev := QianchuanReportUniPromotionDimensionDataAuthorGetV10OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportUniPromotionDimensionDataAuthorGetV10OrderType: valid values are %v", v, AllowedQianchuanReportUniPromotionDimensionDataAuthorGetV10OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportUniPromotionDimensionDataAuthorGetV10OrderType) IsValid() bool {
	for _, existing := range AllowedQianchuanReportUniPromotionDimensionDataAuthorGetV10OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_uni_promotion_dimension_data_author_get_v1.0_order_type value
func (v QianchuanReportUniPromotionDimensionDataAuthorGetV10OrderType) Ptr() *QianchuanReportUniPromotionDimensionDataAuthorGetV10OrderType {
	return &v
}
