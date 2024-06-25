/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// FileVideoUploadTaskListV2AccountType
type FileVideoUploadTaskListV2AccountType string

// List of file_video_upload_task_list_v2_account_type
const (
	ADVERTISER_FileVideoUploadTaskListV2AccountType FileVideoUploadTaskListV2AccountType = "ADVERTISER"
	AGENT_FileVideoUploadTaskListV2AccountType      FileVideoUploadTaskListV2AccountType = "AGENT"
)

// All allowed values of FileVideoUploadTaskListV2AccountType enum
var AllowedFileVideoUploadTaskListV2AccountTypeEnumValues = []FileVideoUploadTaskListV2AccountType{
	"ADVERTISER",
	"AGENT",
}

// NewFileVideoUploadTaskListV2AccountTypeFromValue returns a pointer to a valid FileVideoUploadTaskListV2AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileVideoUploadTaskListV2AccountTypeFromValue(v string) (*FileVideoUploadTaskListV2AccountType, error) {
	ev := FileVideoUploadTaskListV2AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileVideoUploadTaskListV2AccountType: valid values are %v", v, AllowedFileVideoUploadTaskListV2AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileVideoUploadTaskListV2AccountType) IsValid() bool {
	for _, existing := range AllowedFileVideoUploadTaskListV2AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_video_upload_task_list_v2_account_type value
func (v FileVideoUploadTaskListV2AccountType) Ptr() *FileVideoUploadTaskListV2AccountType {
	return &v
}
