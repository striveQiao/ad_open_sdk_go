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

// FileImageAdvertiserV2UploadType
type FileImageAdvertiserV2UploadType string

// List of file_image_advertiser_v2_upload_type
const (
	UPLOAD_BY_FILE_FileImageAdvertiserV2UploadType FileImageAdvertiserV2UploadType = "UPLOAD_BY_FILE"
	UPLOAD_BY_URL_FileImageAdvertiserV2UploadType  FileImageAdvertiserV2UploadType = "UPLOAD_BY_URL"
)

// All allowed values of FileImageAdvertiserV2UploadType enum
var AllowedFileImageAdvertiserV2UploadTypeEnumValues = []FileImageAdvertiserV2UploadType{
	"UPLOAD_BY_FILE",
	"UPLOAD_BY_URL",
}

// NewFileImageAdvertiserV2UploadTypeFromValue returns a pointer to a valid FileImageAdvertiserV2UploadType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileImageAdvertiserV2UploadTypeFromValue(v string) (*FileImageAdvertiserV2UploadType, error) {
	ev := FileImageAdvertiserV2UploadType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileImageAdvertiserV2UploadType: valid values are %v", v, AllowedFileImageAdvertiserV2UploadTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileImageAdvertiserV2UploadType) IsValid() bool {
	for _, existing := range AllowedFileImageAdvertiserV2UploadTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_image_advertiser_v2_upload_type value
func (v FileImageAdvertiserV2UploadType) Ptr() *FileImageAdvertiserV2UploadType {
	return &v
}
