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

// FileImageAdvertiserV2UploadTo
type FileImageAdvertiserV2UploadTo string

// List of file_image_advertiser_v2_upload_to
const (
	AD_FileImageAdvertiserV2UploadTo     FileImageAdvertiserV2UploadTo = "AD"
	AVATAR_FileImageAdvertiserV2UploadTo FileImageAdvertiserV2UploadTo = "AVATAR"
	CG_FileImageAdvertiserV2UploadTo     FileImageAdvertiserV2UploadTo = "CG"
)

// All allowed values of FileImageAdvertiserV2UploadTo enum
var AllowedFileImageAdvertiserV2UploadToEnumValues = []FileImageAdvertiserV2UploadTo{
	"AD",
	"AVATAR",
	"CG",
}

// NewFileImageAdvertiserV2UploadToFromValue returns a pointer to a valid FileImageAdvertiserV2UploadTo
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileImageAdvertiserV2UploadToFromValue(v string) (*FileImageAdvertiserV2UploadTo, error) {
	ev := FileImageAdvertiserV2UploadTo(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileImageAdvertiserV2UploadTo: valid values are %v", v, AllowedFileImageAdvertiserV2UploadToEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileImageAdvertiserV2UploadTo) IsValid() bool {
	for _, existing := range AllowedFileImageAdvertiserV2UploadToEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_image_advertiser_v2_upload_to value
func (v FileImageAdvertiserV2UploadTo) Ptr() *FileImageAdvertiserV2UploadTo {
	return &v
}
