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

// QianchuanAwemeOrderGetV10DataListStatus
type QianchuanAwemeOrderGetV10DataListStatus string

// List of qianchuan_aweme_order_get_v1.0_data_list_status
const (
	AUDIT_QianchuanAwemeOrderGetV10DataListStatus         QianchuanAwemeOrderGetV10DataListStatus = "AUDIT"
	BOOK_QianchuanAwemeOrderGetV10DataListStatus          QianchuanAwemeOrderGetV10DataListStatus = "BOOK"
	CREATING_QianchuanAwemeOrderGetV10DataListStatus      QianchuanAwemeOrderGetV10DataListStatus = "CREATING"
	DELIVERY_OK_QianchuanAwemeOrderGetV10DataListStatus   QianchuanAwemeOrderGetV10DataListStatus = "DELIVERY_OK"
	FAILED_QianchuanAwemeOrderGetV10DataListStatus        QianchuanAwemeOrderGetV10DataListStatus = "FAILED"
	FINISHED_QianchuanAwemeOrderGetV10DataListStatus      QianchuanAwemeOrderGetV10DataListStatus = "FINISHED"
	FROZEN_QianchuanAwemeOrderGetV10DataListStatus        QianchuanAwemeOrderGetV10DataListStatus = "FROZEN"
	OFFLINE_AUDIT_QianchuanAwemeOrderGetV10DataListStatus QianchuanAwemeOrderGetV10DataListStatus = "OFFLINE_AUDIT"
	OVER_QianchuanAwemeOrderGetV10DataListStatus          QianchuanAwemeOrderGetV10DataListStatus = "OVER"
	TIMEOUT_QianchuanAwemeOrderGetV10DataListStatus       QianchuanAwemeOrderGetV10DataListStatus = "TIMEOUT"
	UNPAID_QianchuanAwemeOrderGetV10DataListStatus        QianchuanAwemeOrderGetV10DataListStatus = "UNPAID"
	UNPAIDCANCEL_QianchuanAwemeOrderGetV10DataListStatus  QianchuanAwemeOrderGetV10DataListStatus = "UNPAIDCANCEL"
)

// All allowed values of QianchuanAwemeOrderGetV10DataListStatus enum
var AllowedQianchuanAwemeOrderGetV10DataListStatusEnumValues = []QianchuanAwemeOrderGetV10DataListStatus{
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

// NewQianchuanAwemeOrderGetV10DataListStatusFromValue returns a pointer to a valid QianchuanAwemeOrderGetV10DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderGetV10DataListStatusFromValue(v string) (*QianchuanAwemeOrderGetV10DataListStatus, error) {
	ev := QianchuanAwemeOrderGetV10DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderGetV10DataListStatus: valid values are %v", v, AllowedQianchuanAwemeOrderGetV10DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderGetV10DataListStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderGetV10DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_get_v1.0_data_list_status value
func (v QianchuanAwemeOrderGetV10DataListStatus) Ptr() *QianchuanAwemeOrderGetV10DataListStatus {
	return &v
}
