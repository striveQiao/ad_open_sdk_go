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

// CreativeCustomCreativeCreateV2AdDataIsCommentDisable
type CreativeCustomCreativeCreateV2AdDataIsCommentDisable int64

// List of creative_custom_creative_create_v2_ad_data_is_comment_disable
const (
	Enum_0_CreativeCustomCreativeCreateV2AdDataIsCommentDisable CreativeCustomCreativeCreateV2AdDataIsCommentDisable = 0
	Enum_1_CreativeCustomCreativeCreateV2AdDataIsCommentDisable CreativeCustomCreativeCreateV2AdDataIsCommentDisable = 1
)

// All allowed values of CreativeCustomCreativeCreateV2AdDataIsCommentDisable enum
var AllowedCreativeCustomCreativeCreateV2AdDataIsCommentDisableEnumValues = []CreativeCustomCreativeCreateV2AdDataIsCommentDisable{
	0,
	1,
}

// NewCreativeCustomCreativeCreateV2AdDataIsCommentDisableFromValue returns a pointer to a valid CreativeCustomCreativeCreateV2AdDataIsCommentDisable
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeCreateV2AdDataIsCommentDisableFromValue(v int64) (*CreativeCustomCreativeCreateV2AdDataIsCommentDisable, error) {
	ev := CreativeCustomCreativeCreateV2AdDataIsCommentDisable(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeCreateV2AdDataIsCommentDisable: valid values are %v", v, AllowedCreativeCustomCreativeCreateV2AdDataIsCommentDisableEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeCreateV2AdDataIsCommentDisable) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeCreateV2AdDataIsCommentDisableEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_create_v2_ad_data_is_comment_disable value
func (v CreativeCustomCreativeCreateV2AdDataIsCommentDisable) Ptr() *CreativeCustomCreativeCreateV2AdDataIsCommentDisable {
	return &v
}
