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

// AdlabGroupCreateV30AdInfoDownloadType
type AdlabGroupCreateV30AdInfoDownloadType string

// List of adlab_group_create_v3.0_ad_info_download_type
const (
	DOWNLOAD_URL_AdlabGroupCreateV30AdInfoDownloadType AdlabGroupCreateV30AdInfoDownloadType = "DOWNLOAD_URL"
	EXTERNAL_URL_AdlabGroupCreateV30AdInfoDownloadType AdlabGroupCreateV30AdInfoDownloadType = "EXTERNAL_URL"
)

// All allowed values of AdlabGroupCreateV30AdInfoDownloadType enum
var AllowedAdlabGroupCreateV30AdInfoDownloadTypeEnumValues = []AdlabGroupCreateV30AdInfoDownloadType{
	"DOWNLOAD_URL",
	"EXTERNAL_URL",
}

// NewAdlabGroupCreateV30AdInfoDownloadTypeFromValue returns a pointer to a valid AdlabGroupCreateV30AdInfoDownloadType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupCreateV30AdInfoDownloadTypeFromValue(v string) (*AdlabGroupCreateV30AdInfoDownloadType, error) {
	ev := AdlabGroupCreateV30AdInfoDownloadType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupCreateV30AdInfoDownloadType: valid values are %v", v, AllowedAdlabGroupCreateV30AdInfoDownloadTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupCreateV30AdInfoDownloadType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupCreateV30AdInfoDownloadTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_create_v3.0_ad_info_download_type value
func (v AdlabGroupCreateV30AdInfoDownloadType) Ptr() *AdlabGroupCreateV30AdInfoDownloadType {
	return &v
}
