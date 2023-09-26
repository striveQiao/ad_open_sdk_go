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

// AdlabGroupDetailV30DataDataAdInfoDownloadType
type AdlabGroupDetailV30DataDataAdInfoDownloadType string

// List of adlab_group_detail_v3.0_data_data_ad_info_download_type
const (
	DOWNLOAD_URL_AdlabGroupDetailV30DataDataAdInfoDownloadType AdlabGroupDetailV30DataDataAdInfoDownloadType = "DOWNLOAD_URL"
	EXTERNAL_URL_AdlabGroupDetailV30DataDataAdInfoDownloadType AdlabGroupDetailV30DataDataAdInfoDownloadType = "EXTERNAL_URL"
)

// All allowed values of AdlabGroupDetailV30DataDataAdInfoDownloadType enum
var AllowedAdlabGroupDetailV30DataDataAdInfoDownloadTypeEnumValues = []AdlabGroupDetailV30DataDataAdInfoDownloadType{
	"DOWNLOAD_URL",
	"EXTERNAL_URL",
}

// NewAdlabGroupDetailV30DataDataAdInfoDownloadTypeFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataAdInfoDownloadType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataAdInfoDownloadTypeFromValue(v string) (*AdlabGroupDetailV30DataDataAdInfoDownloadType, error) {
	ev := AdlabGroupDetailV30DataDataAdInfoDownloadType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataAdInfoDownloadType: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataAdInfoDownloadTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataAdInfoDownloadType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataAdInfoDownloadTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_ad_info_download_type value
func (v AdlabGroupDetailV30DataDataAdInfoDownloadType) Ptr() *AdlabGroupDetailV30DataDataAdInfoDownloadType {
	return &v
}
