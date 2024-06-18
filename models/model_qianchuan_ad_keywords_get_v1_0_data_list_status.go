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

// QianchuanAdKeywordsGetV10DataListStatus
type QianchuanAdKeywordsGetV10DataListStatus string

// List of qianchuan_ad_keywords_get_v1.0_data_list_status
const (
	AUDIT_QianchuanAdKeywordsGetV10DataListStatus   QianchuanAdKeywordsGetV10DataListStatus = "AUDIT"
	CONFIRM_QianchuanAdKeywordsGetV10DataListStatus QianchuanAdKeywordsGetV10DataListStatus = "CONFIRM"
	DELETE_QianchuanAdKeywordsGetV10DataListStatus  QianchuanAdKeywordsGetV10DataListStatus = "DELETE"
	PAUSED_QianchuanAdKeywordsGetV10DataListStatus  QianchuanAdKeywordsGetV10DataListStatus = "PAUSED"
	REJECT_QianchuanAdKeywordsGetV10DataListStatus  QianchuanAdKeywordsGetV10DataListStatus = "REJECT"
)

// All allowed values of QianchuanAdKeywordsGetV10DataListStatus enum
var AllowedQianchuanAdKeywordsGetV10DataListStatusEnumValues = []QianchuanAdKeywordsGetV10DataListStatus{
	"AUDIT",
	"CONFIRM",
	"DELETE",
	"PAUSED",
	"REJECT",
}

// NewQianchuanAdKeywordsGetV10DataListStatusFromValue returns a pointer to a valid QianchuanAdKeywordsGetV10DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdKeywordsGetV10DataListStatusFromValue(v string) (*QianchuanAdKeywordsGetV10DataListStatus, error) {
	ev := QianchuanAdKeywordsGetV10DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdKeywordsGetV10DataListStatus: valid values are %v", v, AllowedQianchuanAdKeywordsGetV10DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdKeywordsGetV10DataListStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanAdKeywordsGetV10DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_keywords_get_v1.0_data_list_status value
func (v QianchuanAdKeywordsGetV10DataListStatus) Ptr() *QianchuanAdKeywordsGetV10DataListStatus {
	return &v
}
