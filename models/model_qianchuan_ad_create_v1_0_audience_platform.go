/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10AudiencePlatform
type QianchuanAdCreateV10AudiencePlatform string

// List of qianchuan_ad_create_v1.0_audience_platform
const (
	ANDROID_QianchuanAdCreateV10AudiencePlatform QianchuanAdCreateV10AudiencePlatform = "ANDROID"
	IOS_QianchuanAdCreateV10AudiencePlatform     QianchuanAdCreateV10AudiencePlatform = "IOS"
)

// All allowed values of QianchuanAdCreateV10AudiencePlatform enum
var AllowedQianchuanAdCreateV10AudiencePlatformEnumValues = []QianchuanAdCreateV10AudiencePlatform{
	"ANDROID",
	"IOS",
}

// NewQianchuanAdCreateV10AudiencePlatformFromValue returns a pointer to a valid QianchuanAdCreateV10AudiencePlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10AudiencePlatformFromValue(v string) (*QianchuanAdCreateV10AudiencePlatform, error) {
	ev := QianchuanAdCreateV10AudiencePlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10AudiencePlatform: valid values are %v", v, AllowedQianchuanAdCreateV10AudiencePlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10AudiencePlatform) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10AudiencePlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_audience_platform value
func (v QianchuanAdCreateV10AudiencePlatform) Ptr() *QianchuanAdCreateV10AudiencePlatform {
	return &v
}
