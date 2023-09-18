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

// QianchuanAdDetailGetV10DataKeywordsStatus
type QianchuanAdDetailGetV10DataKeywordsStatus string

// List of qianchuan_ad_detail_get_v1.0_data_keywords_status
const (
	AUDIT_QianchuanAdDetailGetV10DataKeywordsStatus   QianchuanAdDetailGetV10DataKeywordsStatus = "AUDIT"
	CONFIRM_QianchuanAdDetailGetV10DataKeywordsStatus QianchuanAdDetailGetV10DataKeywordsStatus = "CONFIRM"
	DELETE_QianchuanAdDetailGetV10DataKeywordsStatus  QianchuanAdDetailGetV10DataKeywordsStatus = "DELETE"
	PAUSED_QianchuanAdDetailGetV10DataKeywordsStatus  QianchuanAdDetailGetV10DataKeywordsStatus = "PAUSED"
	REJECT_QianchuanAdDetailGetV10DataKeywordsStatus  QianchuanAdDetailGetV10DataKeywordsStatus = "REJECT"
)

// All allowed values of QianchuanAdDetailGetV10DataKeywordsStatus enum
var AllowedQianchuanAdDetailGetV10DataKeywordsStatusEnumValues = []QianchuanAdDetailGetV10DataKeywordsStatus{
	"AUDIT",
	"CONFIRM",
	"DELETE",
	"PAUSED",
	"REJECT",
}

// NewQianchuanAdDetailGetV10DataKeywordsStatusFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataKeywordsStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataKeywordsStatusFromValue(v string) (*QianchuanAdDetailGetV10DataKeywordsStatus, error) {
	ev := QianchuanAdDetailGetV10DataKeywordsStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataKeywordsStatus: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataKeywordsStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataKeywordsStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataKeywordsStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_keywords_status value
func (v QianchuanAdDetailGetV10DataKeywordsStatus) Ptr() *QianchuanAdDetailGetV10DataKeywordsStatus {
	return &v
}
