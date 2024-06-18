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

// QianchuanAdPivativewordsUpdateV10DataStatus
type QianchuanAdPivativewordsUpdateV10DataStatus string

// List of qianchuan_ad_pivativewords_update_v1.0_data_status
const (
	UNKNOWN_QianchuanAdPivativewordsUpdateV10DataStatus QianchuanAdPivativewordsUpdateV10DataStatus = "UNKNOWN"
	SUCCESS_QianchuanAdPivativewordsUpdateV10DataStatus QianchuanAdPivativewordsUpdateV10DataStatus = "SUCCESS"
	FAIL_QianchuanAdPivativewordsUpdateV10DataStatus    QianchuanAdPivativewordsUpdateV10DataStatus = "FAIL"
)

// All allowed values of QianchuanAdPivativewordsUpdateV10DataStatus enum
var AllowedQianchuanAdPivativewordsUpdateV10DataStatusEnumValues = []QianchuanAdPivativewordsUpdateV10DataStatus{
	"UNKNOWN",
	"SUCCESS",
	"FAIL",
}

// NewQianchuanAdPivativewordsUpdateV10DataStatusFromValue returns a pointer to a valid QianchuanAdPivativewordsUpdateV10DataStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdPivativewordsUpdateV10DataStatusFromValue(v string) (*QianchuanAdPivativewordsUpdateV10DataStatus, error) {
	ev := QianchuanAdPivativewordsUpdateV10DataStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdPivativewordsUpdateV10DataStatus: valid values are %v", v, AllowedQianchuanAdPivativewordsUpdateV10DataStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdPivativewordsUpdateV10DataStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanAdPivativewordsUpdateV10DataStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_pivativewords_update_v1.0_data_status value
func (v QianchuanAdPivativewordsUpdateV10DataStatus) Ptr() *QianchuanAdPivativewordsUpdateV10DataStatus {
	return &v
}
