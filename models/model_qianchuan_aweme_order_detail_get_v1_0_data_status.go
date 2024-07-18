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

// QianchuanAwemeOrderDetailGetV10DataStatus
type QianchuanAwemeOrderDetailGetV10DataStatus string

// List of qianchuan_aweme_order_detail_get_v1.0_data_status
const (
	AUDIT_QianchuanAwemeOrderDetailGetV10DataStatus         QianchuanAwemeOrderDetailGetV10DataStatus = "AUDIT"
	BOOK_QianchuanAwemeOrderDetailGetV10DataStatus          QianchuanAwemeOrderDetailGetV10DataStatus = "BOOK"
	CREATING_QianchuanAwemeOrderDetailGetV10DataStatus      QianchuanAwemeOrderDetailGetV10DataStatus = "CREATING"
	DELIVERY_OK_QianchuanAwemeOrderDetailGetV10DataStatus   QianchuanAwemeOrderDetailGetV10DataStatus = "DELIVERY_OK"
	FAILED_QianchuanAwemeOrderDetailGetV10DataStatus        QianchuanAwemeOrderDetailGetV10DataStatus = "FAILED"
	FINISHED_QianchuanAwemeOrderDetailGetV10DataStatus      QianchuanAwemeOrderDetailGetV10DataStatus = "FINISHED"
	FROZEN_QianchuanAwemeOrderDetailGetV10DataStatus        QianchuanAwemeOrderDetailGetV10DataStatus = "FROZEN"
	OFFLINE_AUDIT_QianchuanAwemeOrderDetailGetV10DataStatus QianchuanAwemeOrderDetailGetV10DataStatus = "OFFLINE_AUDIT"
	OVER_QianchuanAwemeOrderDetailGetV10DataStatus          QianchuanAwemeOrderDetailGetV10DataStatus = "OVER"
	TIMEOUT_QianchuanAwemeOrderDetailGetV10DataStatus       QianchuanAwemeOrderDetailGetV10DataStatus = "TIMEOUT"
	UNPAID_QianchuanAwemeOrderDetailGetV10DataStatus        QianchuanAwemeOrderDetailGetV10DataStatus = "UNPAID"
	UNPAIDCANCEL_QianchuanAwemeOrderDetailGetV10DataStatus  QianchuanAwemeOrderDetailGetV10DataStatus = "UNPAIDCANCEL"
)

// All allowed values of QianchuanAwemeOrderDetailGetV10DataStatus enum
var AllowedQianchuanAwemeOrderDetailGetV10DataStatusEnumValues = []QianchuanAwemeOrderDetailGetV10DataStatus{
	"AUDIT",
	"BOOK",
	"CREATING",
	"DELIVERY_OK",
	"FAILED",
	"FINISHED",
	"FROZEN",
	"OFFLINE_AUDIT",
	"OVER",
	"TIMEOUT",
	"UNPAID",
	"UNPAIDCANCEL",
}

// NewQianchuanAwemeOrderDetailGetV10DataStatusFromValue returns a pointer to a valid QianchuanAwemeOrderDetailGetV10DataStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderDetailGetV10DataStatusFromValue(v string) (*QianchuanAwemeOrderDetailGetV10DataStatus, error) {
	ev := QianchuanAwemeOrderDetailGetV10DataStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderDetailGetV10DataStatus: valid values are %v", v, AllowedQianchuanAwemeOrderDetailGetV10DataStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderDetailGetV10DataStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderDetailGetV10DataStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_detail_get_v1.0_data_status value
func (v QianchuanAwemeOrderDetailGetV10DataStatus) Ptr() *QianchuanAwemeOrderDetailGetV10DataStatus {
	return &v
}
