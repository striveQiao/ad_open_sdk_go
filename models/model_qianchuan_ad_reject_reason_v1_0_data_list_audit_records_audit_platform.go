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

// QianchuanAdRejectReasonV10DataListAuditRecordsAuditPlatform
type QianchuanAdRejectReasonV10DataListAuditRecordsAuditPlatform string

// List of qianchuan_ad_reject_reason_v1.0_data_list_audit_records_audit_platform
const (
	CONTENT_QianchuanAdRejectReasonV10DataListAuditRecordsAuditPlatform QianchuanAdRejectReasonV10DataListAuditRecordsAuditPlatform = "CONTENT"
	UNKNOWN_QianchuanAdRejectReasonV10DataListAuditRecordsAuditPlatform QianchuanAdRejectReasonV10DataListAuditRecordsAuditPlatform = "UNKNOWN"
	AD_QianchuanAdRejectReasonV10DataListAuditRecordsAuditPlatform      QianchuanAdRejectReasonV10DataListAuditRecordsAuditPlatform = "AD"
)

// All allowed values of QianchuanAdRejectReasonV10DataListAuditRecordsAuditPlatform enum
var AllowedQianchuanAdRejectReasonV10DataListAuditRecordsAuditPlatformEnumValues = []QianchuanAdRejectReasonV10DataListAuditRecordsAuditPlatform{
	"CONTENT",
	"UNKNOWN",
	"AD",
}

// NewQianchuanAdRejectReasonV10DataListAuditRecordsAuditPlatformFromValue returns a pointer to a valid QianchuanAdRejectReasonV10DataListAuditRecordsAuditPlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdRejectReasonV10DataListAuditRecordsAuditPlatformFromValue(v string) (*QianchuanAdRejectReasonV10DataListAuditRecordsAuditPlatform, error) {
	ev := QianchuanAdRejectReasonV10DataListAuditRecordsAuditPlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdRejectReasonV10DataListAuditRecordsAuditPlatform: valid values are %v", v, AllowedQianchuanAdRejectReasonV10DataListAuditRecordsAuditPlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdRejectReasonV10DataListAuditRecordsAuditPlatform) IsValid() bool {
	for _, existing := range AllowedQianchuanAdRejectReasonV10DataListAuditRecordsAuditPlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_reject_reason_v1.0_data_list_audit_records_audit_platform value
func (v QianchuanAdRejectReasonV10DataListAuditRecordsAuditPlatform) Ptr() *QianchuanAdRejectReasonV10DataListAuditRecordsAuditPlatform {
	return &v
}
