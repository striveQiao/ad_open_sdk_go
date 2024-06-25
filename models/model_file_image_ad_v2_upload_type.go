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

// FileImageAdV2UploadType
type FileImageAdV2UploadType string

// List of file_image_ad_v2_upload_type
const (
	UPLOAD_BY_FILE_FileImageAdV2UploadType FileImageAdV2UploadType = "UPLOAD_BY_FILE"
	UPLOAD_BY_URL_FileImageAdV2UploadType  FileImageAdV2UploadType = "UPLOAD_BY_URL"
)

// All allowed values of FileImageAdV2UploadType enum
var AllowedFileImageAdV2UploadTypeEnumValues = []FileImageAdV2UploadType{
	"UPLOAD_BY_FILE",
	"UPLOAD_BY_URL",
}

// NewFileImageAdV2UploadTypeFromValue returns a pointer to a valid FileImageAdV2UploadType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileImageAdV2UploadTypeFromValue(v string) (*FileImageAdV2UploadType, error) {
	ev := FileImageAdV2UploadType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileImageAdV2UploadType: valid values are %v", v, AllowedFileImageAdV2UploadTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileImageAdV2UploadType) IsValid() bool {
	for _, existing := range AllowedFileImageAdV2UploadTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_image_ad_v2_upload_type value
func (v FileImageAdV2UploadType) Ptr() *FileImageAdV2UploadType {
	return &v
}
