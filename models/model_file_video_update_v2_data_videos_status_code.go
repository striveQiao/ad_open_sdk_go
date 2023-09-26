/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// FileVideoUpdateV2DataVideosStatusCode
type FileVideoUpdateV2DataVideosStatusCode string

// List of file_video_update_v2_data_videos_status_code
const (
	SUCCESS_FileVideoUpdateV2DataVideosStatusCode               FileVideoUpdateV2DataVideosStatusCode = "SUCCESS"
	FAILURE_FileVideoUpdateV2DataVideosStatusCode               FileVideoUpdateV2DataVideosStatusCode = "FAILURE"
	PARTIAL_FAILURE_FileVideoUpdateV2DataVideosStatusCode       FileVideoUpdateV2DataVideosStatusCode = "PARTIAL_FAILURE"
	IMAGE_BINDING_EXISTED_FileVideoUpdateV2DataVideosStatusCode FileVideoUpdateV2DataVideosStatusCode = "IMAGE_BINDING_EXISTED"
	VIDEO_BINDING_EXISTED_FileVideoUpdateV2DataVideosStatusCode FileVideoUpdateV2DataVideosStatusCode = "VIDEO_BINDING_EXISTED"
)

// All allowed values of FileVideoUpdateV2DataVideosStatusCode enum
var AllowedFileVideoUpdateV2DataVideosStatusCodeEnumValues = []FileVideoUpdateV2DataVideosStatusCode{
	"SUCCESS",
	"FAILURE",
	"PARTIAL_FAILURE",
	"IMAGE_BINDING_EXISTED",
	"VIDEO_BINDING_EXISTED",
}

// NewFileVideoUpdateV2DataVideosStatusCodeFromValue returns a pointer to a valid FileVideoUpdateV2DataVideosStatusCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileVideoUpdateV2DataVideosStatusCodeFromValue(v string) (*FileVideoUpdateV2DataVideosStatusCode, error) {
	ev := FileVideoUpdateV2DataVideosStatusCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileVideoUpdateV2DataVideosStatusCode: valid values are %v", v, AllowedFileVideoUpdateV2DataVideosStatusCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileVideoUpdateV2DataVideosStatusCode) IsValid() bool {
	for _, existing := range AllowedFileVideoUpdateV2DataVideosStatusCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_video_update_v2_data_videos_status_code value
func (v FileVideoUpdateV2DataVideosStatusCode) Ptr() *FileVideoUpdateV2DataVideosStatusCode {
	return &v
}
