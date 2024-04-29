/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageCreateV2AwemeFanBehaviors
type AudiencePackageCreateV2AwemeFanBehaviors string

// List of audience_package_create_v2_aweme_fan_behaviors
const (
	COMMENTED_USER_AudiencePackageCreateV2AwemeFanBehaviors       AudiencePackageCreateV2AwemeFanBehaviors = "COMMENTED_USER"
	LIVE_GOODS_CLICK_AudiencePackageCreateV2AwemeFanBehaviors     AudiencePackageCreateV2AwemeFanBehaviors = "LIVE_GOODS_CLICK"
	GOODS_CARTS_CLICK_AudiencePackageCreateV2AwemeFanBehaviors    AudiencePackageCreateV2AwemeFanBehaviors = "GOODS_CARTS_CLICK"
	SHARED_USER_AudiencePackageCreateV2AwemeFanBehaviors          AudiencePackageCreateV2AwemeFanBehaviors = "SHARED_USER"
	LIVE_COMMENT_AudiencePackageCreateV2AwemeFanBehaviors         AudiencePackageCreateV2AwemeFanBehaviors = "LIVE_COMMENT"
	LIVE_EFFECTIVE_WATCH_AudiencePackageCreateV2AwemeFanBehaviors AudiencePackageCreateV2AwemeFanBehaviors = "LIVE_EFFECTIVE_WATCH"
	LIVE_GOODS_ORDER_AudiencePackageCreateV2AwemeFanBehaviors     AudiencePackageCreateV2AwemeFanBehaviors = "LIVE_GOODS_ORDER"
	FOLLOWED_USER_AudiencePackageCreateV2AwemeFanBehaviors        AudiencePackageCreateV2AwemeFanBehaviors = "FOLLOWED_USER"
	LIVE_EXCEPTIONAL_AudiencePackageCreateV2AwemeFanBehaviors     AudiencePackageCreateV2AwemeFanBehaviors = "LIVE_EXCEPTIONAL"
	GOODS_CARTS_ORDER_AudiencePackageCreateV2AwemeFanBehaviors    AudiencePackageCreateV2AwemeFanBehaviors = "GOODS_CARTS_ORDER"
	LIKED_USER_AudiencePackageCreateV2AwemeFanBehaviors           AudiencePackageCreateV2AwemeFanBehaviors = "LIKED_USER"
	LIVE_WATCH_AudiencePackageCreateV2AwemeFanBehaviors           AudiencePackageCreateV2AwemeFanBehaviors = "LIVE_WATCH"
)

// All allowed values of AudiencePackageCreateV2AwemeFanBehaviors enum
var AllowedAudiencePackageCreateV2AwemeFanBehaviorsEnumValues = []AudiencePackageCreateV2AwemeFanBehaviors{
	"COMMENTED_USER",
	"LIVE_GOODS_CLICK",
	"GOODS_CARTS_CLICK",
	"SHARED_USER",
	"LIVE_COMMENT",
	"LIVE_EFFECTIVE_WATCH",
	"LIVE_GOODS_ORDER",
	"FOLLOWED_USER",
	"LIVE_EXCEPTIONAL",
	"GOODS_CARTS_ORDER",
	"LIKED_USER",
	"LIVE_WATCH",
}

// NewAudiencePackageCreateV2AwemeFanBehaviorsFromValue returns a pointer to a valid AudiencePackageCreateV2AwemeFanBehaviors
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2AwemeFanBehaviorsFromValue(v string) (*AudiencePackageCreateV2AwemeFanBehaviors, error) {
	ev := AudiencePackageCreateV2AwemeFanBehaviors(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2AwemeFanBehaviors: valid values are %v", v, AllowedAudiencePackageCreateV2AwemeFanBehaviorsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2AwemeFanBehaviors) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2AwemeFanBehaviorsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_aweme_fan_behaviors value
func (v AudiencePackageCreateV2AwemeFanBehaviors) Ptr() *AudiencePackageCreateV2AwemeFanBehaviors {
	return &v
}
