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

// AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlSendType
type AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlSendType string

// List of adlab_group_create_v3.0_ad_info_track_url_setting_track_url_send_type
const (
	CLIENT_SEND_AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlSendType AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlSendType = "CLIENT_SEND"
	SERVER_SEND_AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlSendType AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlSendType = "SERVER_SEND"
)

// All allowed values of AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlSendType enum
var AllowedAdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlSendTypeEnumValues = []AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlSendType{
	"CLIENT_SEND",
	"SERVER_SEND",
}

// NewAdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlSendTypeFromValue returns a pointer to a valid AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlSendType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlSendTypeFromValue(v string) (*AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlSendType, error) {
	ev := AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlSendType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlSendType: valid values are %v", v, AllowedAdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlSendTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlSendType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlSendTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_create_v3.0_ad_info_track_url_setting_track_url_send_type value
func (v AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlSendType) Ptr() *AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlSendType {
	return &v
}
