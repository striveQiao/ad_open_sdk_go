/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// KeywordCreateV30DataSuccessListStatus
type KeywordCreateV30DataSuccessListStatus string

// List of keyword_create_v3.0_data_success_list_status
const (
	AUDITING_KeywordCreateV30DataSuccessListStatus       KeywordCreateV30DataSuccessListStatus = "AUDITING"
	AUDIT_ACCEPTED_KeywordCreateV30DataSuccessListStatus KeywordCreateV30DataSuccessListStatus = "AUDIT_ACCEPTED"
	AUDIT_REJECTED_KeywordCreateV30DataSuccessListStatus KeywordCreateV30DataSuccessListStatus = "AUDIT_REJECTED"
	DELETED_KeywordCreateV30DataSuccessListStatus        KeywordCreateV30DataSuccessListStatus = "DELETED"
)

// All allowed values of KeywordCreateV30DataSuccessListStatus enum
var AllowedKeywordCreateV30DataSuccessListStatusEnumValues = []KeywordCreateV30DataSuccessListStatus{
	"AUDITING",
	"AUDIT_ACCEPTED",
	"AUDIT_REJECTED",
	"DELETED",
}

// NewKeywordCreateV30DataSuccessListStatusFromValue returns a pointer to a valid KeywordCreateV30DataSuccessListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordCreateV30DataSuccessListStatusFromValue(v string) (*KeywordCreateV30DataSuccessListStatus, error) {
	ev := KeywordCreateV30DataSuccessListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordCreateV30DataSuccessListStatus: valid values are %v", v, AllowedKeywordCreateV30DataSuccessListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordCreateV30DataSuccessListStatus) IsValid() bool {
	for _, existing := range AllowedKeywordCreateV30DataSuccessListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_create_v3.0_data_success_list_status value
func (v KeywordCreateV30DataSuccessListStatus) Ptr() *KeywordCreateV30DataSuccessListStatus {
	return &v
}
