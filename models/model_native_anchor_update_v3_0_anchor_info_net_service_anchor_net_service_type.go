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

// NativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceType
type NativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceType string

// List of native_anchor_update_v3.0_anchor_info_net_service_anchor_net_service_type
const (
	GENERAL_NativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceType             NativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceType = "GENERAL"
	MICRO_APP_NativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceType           NativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceType = "MICRO_APP"
	QUICK_APP_NativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceType           NativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceType = "QUICK_APP"
	WECHAT_PACKAGE_NativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceType      NativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceType = "WECHAT_PACKAGE"
	WECOM_PACKAGE_NativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceType       NativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceType = "WECOM_PACKAGE"
	WECHAT_EXTERNAL_URL_NativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceType NativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceType = "WECHAT_EXTERNAL_URL"
)

// All allowed values of NativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceType enum
var AllowedNativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceTypeEnumValues = []NativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceType{
	"GENERAL",
	"MICRO_APP",
	"QUICK_APP",
	"WECHAT_PACKAGE",
	"WECOM_PACKAGE",
	"WECHAT_EXTERNAL_URL",
}

// NewNativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceTypeFromValue returns a pointer to a valid NativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceTypeFromValue(v string) (*NativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceType, error) {
	ev := NativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceType: valid values are %v", v, AllowedNativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceType) IsValid() bool {
	for _, existing := range AllowedNativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_update_v3.0_anchor_info_net_service_anchor_net_service_type value
func (v NativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceType) Ptr() *NativeAnchorUpdateV30AnchorInfoNetServiceAnchorNetServiceType {
	return &v
}
