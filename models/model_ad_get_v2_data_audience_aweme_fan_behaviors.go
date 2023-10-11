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

// AdGetV2DataAudienceAwemeFanBehaviors
type AdGetV2DataAudienceAwemeFanBehaviors string

// List of ad_get_v2_data_audience_aweme_fan_behaviors
const (
	FOLLOWED_USER_AdGetV2DataAudienceAwemeFanBehaviors        AdGetV2DataAudienceAwemeFanBehaviors = "FOLLOWED_USER"
	GOODS_CARTS_ORDER_AdGetV2DataAudienceAwemeFanBehaviors    AdGetV2DataAudienceAwemeFanBehaviors = "GOODS_CARTS_ORDER"
	GOODS_CARTS_CLICK_AdGetV2DataAudienceAwemeFanBehaviors    AdGetV2DataAudienceAwemeFanBehaviors = "GOODS_CARTS_CLICK"
	LIVE_EXCEPTIONAL_AdGetV2DataAudienceAwemeFanBehaviors     AdGetV2DataAudienceAwemeFanBehaviors = "LIVE_EXCEPTIONAL"
	LIVE_COMMENT_AdGetV2DataAudienceAwemeFanBehaviors         AdGetV2DataAudienceAwemeFanBehaviors = "LIVE_COMMENT"
	LIVE_GOODS_ORDER_AdGetV2DataAudienceAwemeFanBehaviors     AdGetV2DataAudienceAwemeFanBehaviors = "LIVE_GOODS_ORDER"
	COMMENTED_USER_AdGetV2DataAudienceAwemeFanBehaviors       AdGetV2DataAudienceAwemeFanBehaviors = "COMMENTED_USER"
	LIVE_GOODS_CLICK_AdGetV2DataAudienceAwemeFanBehaviors     AdGetV2DataAudienceAwemeFanBehaviors = "LIVE_GOODS_CLICK"
	LIVE_WATCH_AdGetV2DataAudienceAwemeFanBehaviors           AdGetV2DataAudienceAwemeFanBehaviors = "LIVE_WATCH"
	LIKED_USER_AdGetV2DataAudienceAwemeFanBehaviors           AdGetV2DataAudienceAwemeFanBehaviors = "LIKED_USER"
	SHARED_USER_AdGetV2DataAudienceAwemeFanBehaviors          AdGetV2DataAudienceAwemeFanBehaviors = "SHARED_USER"
	LIVE_EFFECTIVE_WATCH_AdGetV2DataAudienceAwemeFanBehaviors AdGetV2DataAudienceAwemeFanBehaviors = "LIVE_EFFECTIVE_WATCH"
)

// All allowed values of AdGetV2DataAudienceAwemeFanBehaviors enum
var AllowedAdGetV2DataAudienceAwemeFanBehaviorsEnumValues = []AdGetV2DataAudienceAwemeFanBehaviors{
	"FOLLOWED_USER",
	"GOODS_CARTS_ORDER",
	"GOODS_CARTS_CLICK",
	"LIVE_EXCEPTIONAL",
	"LIVE_COMMENT",
	"LIVE_GOODS_ORDER",
	"COMMENTED_USER",
	"LIVE_GOODS_CLICK",
	"LIVE_WATCH",
	"LIKED_USER",
	"SHARED_USER",
	"LIVE_EFFECTIVE_WATCH",
}

// NewAdGetV2DataAudienceAwemeFanBehaviorsFromValue returns a pointer to a valid AdGetV2DataAudienceAwemeFanBehaviors
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAudienceAwemeFanBehaviorsFromValue(v string) (*AdGetV2DataAudienceAwemeFanBehaviors, error) {
	ev := AdGetV2DataAudienceAwemeFanBehaviors(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAudienceAwemeFanBehaviors: valid values are %v", v, AllowedAdGetV2DataAudienceAwemeFanBehaviorsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAudienceAwemeFanBehaviors) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAudienceAwemeFanBehaviorsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_audience_aweme_fan_behaviors value
func (v AdGetV2DataAudienceAwemeFanBehaviors) Ptr() *AdGetV2DataAudienceAwemeFanBehaviors {
	return &v
}
