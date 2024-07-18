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

// FilePreauditGetV30DataListStatus
type FilePreauditGetV30DataListStatus string

// List of file_preaudit_get_v3.0_data_list_status
const (
	AUDITING_FilePreauditGetV30DataListStatus       FilePreauditGetV30DataListStatus = "AUDITING"
	AUDIT_ACCEPTED_FilePreauditGetV30DataListStatus FilePreauditGetV30DataListStatus = "AUDIT_ACCEPTED"
	AUDIT_FAILED_FilePreauditGetV30DataListStatus   FilePreauditGetV30DataListStatus = "AUDIT_FAILED"
	AUDIT_REJECT_FilePreauditGetV30DataListStatus   FilePreauditGetV30DataListStatus = "AUDIT_REJECT"
	AUDIT_SUBMIT_FilePreauditGetV30DataListStatus   FilePreauditGetV30DataListStatus = "AUDIT_SUBMIT"
	AUDIT_TIMEOUT_FilePreauditGetV30DataListStatus  FilePreauditGetV30DataListStatus = "AUDIT_TIMEOUT"
)

// All allowed values of FilePreauditGetV30DataListStatus enum
var AllowedFilePreauditGetV30DataListStatusEnumValues = []FilePreauditGetV30DataListStatus{
	"AUDITING",
	"AUDIT_ACCEPTED",
	"AUDIT_FAILED",
	"AUDIT_REJECT",
	"AUDIT_SUBMIT",
	"AUDIT_TIMEOUT",
}

// NewFilePreauditGetV30DataListStatusFromValue returns a pointer to a valid FilePreauditGetV30DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFilePreauditGetV30DataListStatusFromValue(v string) (*FilePreauditGetV30DataListStatus, error) {
	ev := FilePreauditGetV30DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FilePreauditGetV30DataListStatus: valid values are %v", v, AllowedFilePreauditGetV30DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FilePreauditGetV30DataListStatus) IsValid() bool {
	for _, existing := range AllowedFilePreauditGetV30DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_preaudit_get_v3.0_data_list_status value
func (v FilePreauditGetV30DataListStatus) Ptr() *FilePreauditGetV30DataListStatus {
	return &v
}
