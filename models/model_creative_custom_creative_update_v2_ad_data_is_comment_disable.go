/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeCustomCreativeUpdateV2AdDataIsCommentDisable
type CreativeCustomCreativeUpdateV2AdDataIsCommentDisable int64

// List of creative_custom_creative_update_v2_ad_data_is_comment_disable
const (
	Enum_0_CreativeCustomCreativeUpdateV2AdDataIsCommentDisable CreativeCustomCreativeUpdateV2AdDataIsCommentDisable = 0
	Enum_1_CreativeCustomCreativeUpdateV2AdDataIsCommentDisable CreativeCustomCreativeUpdateV2AdDataIsCommentDisable = 1
)

// All allowed values of CreativeCustomCreativeUpdateV2AdDataIsCommentDisable enum
var AllowedCreativeCustomCreativeUpdateV2AdDataIsCommentDisableEnumValues = []CreativeCustomCreativeUpdateV2AdDataIsCommentDisable{
	0,
	1,
}

// NewCreativeCustomCreativeUpdateV2AdDataIsCommentDisableFromValue returns a pointer to a valid CreativeCustomCreativeUpdateV2AdDataIsCommentDisable
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeUpdateV2AdDataIsCommentDisableFromValue(v int64) (*CreativeCustomCreativeUpdateV2AdDataIsCommentDisable, error) {
	ev := CreativeCustomCreativeUpdateV2AdDataIsCommentDisable(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeUpdateV2AdDataIsCommentDisable: valid values are %v", v, AllowedCreativeCustomCreativeUpdateV2AdDataIsCommentDisableEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeUpdateV2AdDataIsCommentDisable) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeUpdateV2AdDataIsCommentDisableEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_update_v2_ad_data_is_comment_disable value
func (v CreativeCustomCreativeUpdateV2AdDataIsCommentDisable) Ptr() *CreativeCustomCreativeUpdateV2AdDataIsCommentDisable {
	return &v
}
