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

// AdlabGroupUpdateV30AdInfoDownloadMode
type AdlabGroupUpdateV30AdInfoDownloadMode string

// List of adlab_group_update_v3.0_ad_info_download_mode
const (
	APP_STORE_DELIVERY_AdlabGroupUpdateV30AdInfoDownloadMode AdlabGroupUpdateV30AdInfoDownloadMode = "APP_STORE_DELIVERY"
	DEFAULT_AdlabGroupUpdateV30AdInfoDownloadMode            AdlabGroupUpdateV30AdInfoDownloadMode = "DEFAULT"
)

// All allowed values of AdlabGroupUpdateV30AdInfoDownloadMode enum
var AllowedAdlabGroupUpdateV30AdInfoDownloadModeEnumValues = []AdlabGroupUpdateV30AdInfoDownloadMode{
	"APP_STORE_DELIVERY",
	"DEFAULT",
}

// NewAdlabGroupUpdateV30AdInfoDownloadModeFromValue returns a pointer to a valid AdlabGroupUpdateV30AdInfoDownloadMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupUpdateV30AdInfoDownloadModeFromValue(v string) (*AdlabGroupUpdateV30AdInfoDownloadMode, error) {
	ev := AdlabGroupUpdateV30AdInfoDownloadMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupUpdateV30AdInfoDownloadMode: valid values are %v", v, AllowedAdlabGroupUpdateV30AdInfoDownloadModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupUpdateV30AdInfoDownloadMode) IsValid() bool {
	for _, existing := range AllowedAdlabGroupUpdateV30AdInfoDownloadModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_update_v3.0_ad_info_download_mode value
func (v AdlabGroupUpdateV30AdInfoDownloadMode) Ptr() *AdlabGroupUpdateV30AdInfoDownloadMode {
	return &v
}
