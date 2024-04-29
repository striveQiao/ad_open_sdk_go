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

// QianchuanEcpAwemeAdGetV10FilteringOrderStatus
type QianchuanEcpAwemeAdGetV10FilteringOrderStatus string

// List of qianchuan_ecp_aweme_ad_get_v1.0_filtering_order_status
const (
	AUDITING_QianchuanEcpAwemeAdGetV10FilteringOrderStatus     QianchuanEcpAwemeAdGetV10FilteringOrderStatus = "AUDITING"
	BOOK_QianchuanEcpAwemeAdGetV10FilteringOrderStatus         QianchuanEcpAwemeAdGetV10FilteringOrderStatus = "BOOK"
	CREATING_QianchuanEcpAwemeAdGetV10FilteringOrderStatus     QianchuanEcpAwemeAdGetV10FilteringOrderStatus = "CREATING"
	DELIVERING_QianchuanEcpAwemeAdGetV10FilteringOrderStatus   QianchuanEcpAwemeAdGetV10FilteringOrderStatus = "DELIVERING"
	FAILED_QianchuanEcpAwemeAdGetV10FilteringOrderStatus       QianchuanEcpAwemeAdGetV10FilteringOrderStatus = "FAILED"
	FINISHED_QianchuanEcpAwemeAdGetV10FilteringOrderStatus     QianchuanEcpAwemeAdGetV10FilteringOrderStatus = "FINISHED"
	OVER_QianchuanEcpAwemeAdGetV10FilteringOrderStatus         QianchuanEcpAwemeAdGetV10FilteringOrderStatus = "OVER"
	REJECT_QianchuanEcpAwemeAdGetV10FilteringOrderStatus       QianchuanEcpAwemeAdGetV10FilteringOrderStatus = "REJECT"
	UNDELIVERING_QianchuanEcpAwemeAdGetV10FilteringOrderStatus QianchuanEcpAwemeAdGetV10FilteringOrderStatus = "UNDELIVERING"
	UNPAID_QianchuanEcpAwemeAdGetV10FilteringOrderStatus       QianchuanEcpAwemeAdGetV10FilteringOrderStatus = "UNPAID"
)

// All allowed values of QianchuanEcpAwemeAdGetV10FilteringOrderStatus enum
var AllowedQianchuanEcpAwemeAdGetV10FilteringOrderStatusEnumValues = []QianchuanEcpAwemeAdGetV10FilteringOrderStatus{
	"AUDITING",
	"BOOK",
	"CREATING",
	"DELIVERING",
	"FAILED",
	"FINISHED",
	"OVER",
	"REJECT",
	"UNDELIVERING",
	"UNPAID",
}

// NewQianchuanEcpAwemeAdGetV10FilteringOrderStatusFromValue returns a pointer to a valid QianchuanEcpAwemeAdGetV10FilteringOrderStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanEcpAwemeAdGetV10FilteringOrderStatusFromValue(v string) (*QianchuanEcpAwemeAdGetV10FilteringOrderStatus, error) {
	ev := QianchuanEcpAwemeAdGetV10FilteringOrderStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanEcpAwemeAdGetV10FilteringOrderStatus: valid values are %v", v, AllowedQianchuanEcpAwemeAdGetV10FilteringOrderStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanEcpAwemeAdGetV10FilteringOrderStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanEcpAwemeAdGetV10FilteringOrderStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ecp_aweme_ad_get_v1.0_filtering_order_status value
func (v QianchuanEcpAwemeAdGetV10FilteringOrderStatus) Ptr() *QianchuanEcpAwemeAdGetV10FilteringOrderStatus {
	return &v
}
