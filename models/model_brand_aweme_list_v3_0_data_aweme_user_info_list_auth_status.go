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

// BrandAwemeListV30DataAwemeUserInfoListAuthStatus
type BrandAwemeListV30DataAwemeUserInfoListAuthStatus string

// List of brand_aweme_list_v3.0_data_aweme_user_info_list_auth_status
const (
	AUTHING_BrandAwemeListV30DataAwemeUserInfoListAuthStatus BrandAwemeListV30DataAwemeUserInfoListAuthStatus = "AUTHING"
	EXPIRED_BrandAwemeListV30DataAwemeUserInfoListAuthStatus BrandAwemeListV30DataAwemeUserInfoListAuthStatus = "EXPIRED"
	UNBIND_BrandAwemeListV30DataAwemeUserInfoListAuthStatus  BrandAwemeListV30DataAwemeUserInfoListAuthStatus = "UNBIND"
)

// All allowed values of BrandAwemeListV30DataAwemeUserInfoListAuthStatus enum
var AllowedBrandAwemeListV30DataAwemeUserInfoListAuthStatusEnumValues = []BrandAwemeListV30DataAwemeUserInfoListAuthStatus{
	"AUTHING",
	"EXPIRED",
	"UNBIND",
}

// NewBrandAwemeListV30DataAwemeUserInfoListAuthStatusFromValue returns a pointer to a valid BrandAwemeListV30DataAwemeUserInfoListAuthStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAwemeListV30DataAwemeUserInfoListAuthStatusFromValue(v string) (*BrandAwemeListV30DataAwemeUserInfoListAuthStatus, error) {
	ev := BrandAwemeListV30DataAwemeUserInfoListAuthStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAwemeListV30DataAwemeUserInfoListAuthStatus: valid values are %v", v, AllowedBrandAwemeListV30DataAwemeUserInfoListAuthStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAwemeListV30DataAwemeUserInfoListAuthStatus) IsValid() bool {
	for _, existing := range AllowedBrandAwemeListV30DataAwemeUserInfoListAuthStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_aweme_list_v3.0_data_aweme_user_info_list_auth_status value
func (v BrandAwemeListV30DataAwemeUserInfoListAuthStatus) Ptr() *BrandAwemeListV30DataAwemeUserInfoListAuthStatus {
	return &v
}
