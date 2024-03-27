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

// QianchuanReportLongTransferOrderGetV10DataListPayType
type QianchuanReportLongTransferOrderGetV10DataListPayType string

// List of qianchuan_report_long_transfer_order_get_v1.0_data_list_pay_type
const (
	DIRECT_QianchuanReportLongTransferOrderGetV10DataListPayType   QianchuanReportLongTransferOrderGetV10DataListPayType = "DIRECT"
	INDIRECT_QianchuanReportLongTransferOrderGetV10DataListPayType QianchuanReportLongTransferOrderGetV10DataListPayType = "INDIRECT"
)

// All allowed values of QianchuanReportLongTransferOrderGetV10DataListPayType enum
var AllowedQianchuanReportLongTransferOrderGetV10DataListPayTypeEnumValues = []QianchuanReportLongTransferOrderGetV10DataListPayType{
	"DIRECT",
	"INDIRECT",
}

// NewQianchuanReportLongTransferOrderGetV10DataListPayTypeFromValue returns a pointer to a valid QianchuanReportLongTransferOrderGetV10DataListPayType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportLongTransferOrderGetV10DataListPayTypeFromValue(v string) (*QianchuanReportLongTransferOrderGetV10DataListPayType, error) {
	ev := QianchuanReportLongTransferOrderGetV10DataListPayType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportLongTransferOrderGetV10DataListPayType: valid values are %v", v, AllowedQianchuanReportLongTransferOrderGetV10DataListPayTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportLongTransferOrderGetV10DataListPayType) IsValid() bool {
	for _, existing := range AllowedQianchuanReportLongTransferOrderGetV10DataListPayTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_long_transfer_order_get_v1.0_data_list_pay_type value
func (v QianchuanReportLongTransferOrderGetV10DataListPayType) Ptr() *QianchuanReportLongTransferOrderGetV10DataListPayType {
	return &v
}
