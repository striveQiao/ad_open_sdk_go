/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// FileMediaGetV2FilteringStatus
type FileMediaGetV2FilteringStatus string

// List of file_media_get_v2_filtering_status
const (
	FRIENDS_VISIBLE_FileMediaGetV2FilteringStatus FileMediaGetV2FilteringStatus = "FRIENDS_VISIBLE"
	OPEN_VISIBLE_FileMediaGetV2FilteringStatus    FileMediaGetV2FilteringStatus = "OPEN_VISIBLE"
	SELF_VISIBLE_FileMediaGetV2FilteringStatus    FileMediaGetV2FilteringStatus = "SELF_VISIBLE"
)

// All allowed values of FileMediaGetV2FilteringStatus enum
var AllowedFileMediaGetV2FilteringStatusEnumValues = []FileMediaGetV2FilteringStatus{
	"FRIENDS_VISIBLE",
	"OPEN_VISIBLE",
	"SELF_VISIBLE",
}

// NewFileMediaGetV2FilteringStatusFromValue returns a pointer to a valid FileMediaGetV2FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileMediaGetV2FilteringStatusFromValue(v string) (*FileMediaGetV2FilteringStatus, error) {
	ev := FileMediaGetV2FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileMediaGetV2FilteringStatus: valid values are %v", v, AllowedFileMediaGetV2FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileMediaGetV2FilteringStatus) IsValid() bool {
	for _, existing := range AllowedFileMediaGetV2FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_media_get_v2_filtering_status value
func (v FileMediaGetV2FilteringStatus) Ptr() *FileMediaGetV2FilteringStatus {
	return &v
}
