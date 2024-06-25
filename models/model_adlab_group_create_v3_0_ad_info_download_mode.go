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

// AdlabGroupCreateV30AdInfoDownloadMode
type AdlabGroupCreateV30AdInfoDownloadMode string

// List of adlab_group_create_v3.0_ad_info_download_mode
const (
	APP_STORE_DELIVERY_AdlabGroupCreateV30AdInfoDownloadMode AdlabGroupCreateV30AdInfoDownloadMode = "APP_STORE_DELIVERY"
	DEFAULT_AdlabGroupCreateV30AdInfoDownloadMode            AdlabGroupCreateV30AdInfoDownloadMode = "DEFAULT"
)

// All allowed values of AdlabGroupCreateV30AdInfoDownloadMode enum
var AllowedAdlabGroupCreateV30AdInfoDownloadModeEnumValues = []AdlabGroupCreateV30AdInfoDownloadMode{
	"APP_STORE_DELIVERY",
	"DEFAULT",
}

// NewAdlabGroupCreateV30AdInfoDownloadModeFromValue returns a pointer to a valid AdlabGroupCreateV30AdInfoDownloadMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupCreateV30AdInfoDownloadModeFromValue(v string) (*AdlabGroupCreateV30AdInfoDownloadMode, error) {
	ev := AdlabGroupCreateV30AdInfoDownloadMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupCreateV30AdInfoDownloadMode: valid values are %v", v, AllowedAdlabGroupCreateV30AdInfoDownloadModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupCreateV30AdInfoDownloadMode) IsValid() bool {
	for _, existing := range AllowedAdlabGroupCreateV30AdInfoDownloadModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_create_v3.0_ad_info_download_mode value
func (v AdlabGroupCreateV30AdInfoDownloadMode) Ptr() *AdlabGroupCreateV30AdInfoDownloadMode {
	return &v
}
