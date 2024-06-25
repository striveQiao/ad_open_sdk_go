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

// AdlabGroupDetailV30DataDataAdInfoDownloadMode
type AdlabGroupDetailV30DataDataAdInfoDownloadMode string

// List of adlab_group_detail_v3.0_data_data_ad_info_download_mode
const (
	APP_STORE_DELIVERY_AdlabGroupDetailV30DataDataAdInfoDownloadMode AdlabGroupDetailV30DataDataAdInfoDownloadMode = "APP_STORE_DELIVERY"
	DEFAULT_AdlabGroupDetailV30DataDataAdInfoDownloadMode            AdlabGroupDetailV30DataDataAdInfoDownloadMode = "DEFAULT"
)

// All allowed values of AdlabGroupDetailV30DataDataAdInfoDownloadMode enum
var AllowedAdlabGroupDetailV30DataDataAdInfoDownloadModeEnumValues = []AdlabGroupDetailV30DataDataAdInfoDownloadMode{
	"APP_STORE_DELIVERY",
	"DEFAULT",
}

// NewAdlabGroupDetailV30DataDataAdInfoDownloadModeFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataAdInfoDownloadMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataAdInfoDownloadModeFromValue(v string) (*AdlabGroupDetailV30DataDataAdInfoDownloadMode, error) {
	ev := AdlabGroupDetailV30DataDataAdInfoDownloadMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataAdInfoDownloadMode: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataAdInfoDownloadModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataAdInfoDownloadMode) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataAdInfoDownloadModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_ad_info_download_mode value
func (v AdlabGroupDetailV30DataDataAdInfoDownloadMode) Ptr() *AdlabGroupDetailV30DataDataAdInfoDownloadMode {
	return &v
}
