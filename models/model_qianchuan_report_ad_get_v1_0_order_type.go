/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportAdGetV10OrderType
type QianchuanReportAdGetV10OrderType string

// List of qianchuan_report_ad_get_v1.0_order_type
const (
	ASC_QianchuanReportAdGetV10OrderType  QianchuanReportAdGetV10OrderType = "ASC"
	DESC_QianchuanReportAdGetV10OrderType QianchuanReportAdGetV10OrderType = "DESC"
)

// All allowed values of QianchuanReportAdGetV10OrderType enum
var AllowedQianchuanReportAdGetV10OrderTypeEnumValues = []QianchuanReportAdGetV10OrderType{
	"ASC",
	"DESC",
}

// NewQianchuanReportAdGetV10OrderTypeFromValue returns a pointer to a valid QianchuanReportAdGetV10OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportAdGetV10OrderTypeFromValue(v string) (*QianchuanReportAdGetV10OrderType, error) {
	ev := QianchuanReportAdGetV10OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportAdGetV10OrderType: valid values are %v", v, AllowedQianchuanReportAdGetV10OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportAdGetV10OrderType) IsValid() bool {
	for _, existing := range AllowedQianchuanReportAdGetV10OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_ad_get_v1.0_order_type value
func (v QianchuanReportAdGetV10OrderType) Ptr() *QianchuanReportAdGetV10OrderType {
	return &v
}
