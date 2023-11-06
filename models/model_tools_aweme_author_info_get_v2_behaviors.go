/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAwemeAuthorInfoGetV2Behaviors
type ToolsAwemeAuthorInfoGetV2Behaviors string

// List of tools_aweme_author_info_get_v2_behaviors
const (
	LIVE_WATCH_ToolsAwemeAuthorInfoGetV2Behaviors           ToolsAwemeAuthorInfoGetV2Behaviors = "LIVE_WATCH"
	LIVE_GOODS_CLICK_ToolsAwemeAuthorInfoGetV2Behaviors     ToolsAwemeAuthorInfoGetV2Behaviors = "LIVE_GOODS_CLICK"
	GOODS_CARTS_CLICK_ToolsAwemeAuthorInfoGetV2Behaviors    ToolsAwemeAuthorInfoGetV2Behaviors = "GOODS_CARTS_CLICK"
	SHARED_USER_ToolsAwemeAuthorInfoGetV2Behaviors          ToolsAwemeAuthorInfoGetV2Behaviors = "SHARED_USER"
	LIVE_EFFECTIVE_WATCH_ToolsAwemeAuthorInfoGetV2Behaviors ToolsAwemeAuthorInfoGetV2Behaviors = "LIVE_EFFECTIVE_WATCH"
	LIVE_GOODS_ORDER_ToolsAwemeAuthorInfoGetV2Behaviors     ToolsAwemeAuthorInfoGetV2Behaviors = "LIVE_GOODS_ORDER"
	GOODS_CARTS_ORDER_ToolsAwemeAuthorInfoGetV2Behaviors    ToolsAwemeAuthorInfoGetV2Behaviors = "GOODS_CARTS_ORDER"
	COMMENTED_USER_ToolsAwemeAuthorInfoGetV2Behaviors       ToolsAwemeAuthorInfoGetV2Behaviors = "COMMENTED_USER"
	LIVE_COMMENT_ToolsAwemeAuthorInfoGetV2Behaviors         ToolsAwemeAuthorInfoGetV2Behaviors = "LIVE_COMMENT"
	LIVE_EXCEPTIONAL_ToolsAwemeAuthorInfoGetV2Behaviors     ToolsAwemeAuthorInfoGetV2Behaviors = "LIVE_EXCEPTIONAL"
	FOLLOWED_USER_ToolsAwemeAuthorInfoGetV2Behaviors        ToolsAwemeAuthorInfoGetV2Behaviors = "FOLLOWED_USER"
	LIKED_USER_ToolsAwemeAuthorInfoGetV2Behaviors           ToolsAwemeAuthorInfoGetV2Behaviors = "LIKED_USER"
)

// All allowed values of ToolsAwemeAuthorInfoGetV2Behaviors enum
var AllowedToolsAwemeAuthorInfoGetV2BehaviorsEnumValues = []ToolsAwemeAuthorInfoGetV2Behaviors{
	"LIVE_WATCH",
	"LIVE_GOODS_CLICK",
	"GOODS_CARTS_CLICK",
	"SHARED_USER",
	"LIVE_EFFECTIVE_WATCH",
	"LIVE_GOODS_ORDER",
	"GOODS_CARTS_ORDER",
	"COMMENTED_USER",
	"LIVE_COMMENT",
	"LIVE_EXCEPTIONAL",
	"FOLLOWED_USER",
	"LIKED_USER",
}

// NewToolsAwemeAuthorInfoGetV2BehaviorsFromValue returns a pointer to a valid ToolsAwemeAuthorInfoGetV2Behaviors
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAwemeAuthorInfoGetV2BehaviorsFromValue(v string) (*ToolsAwemeAuthorInfoGetV2Behaviors, error) {
	ev := ToolsAwemeAuthorInfoGetV2Behaviors(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAwemeAuthorInfoGetV2Behaviors: valid values are %v", v, AllowedToolsAwemeAuthorInfoGetV2BehaviorsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAwemeAuthorInfoGetV2Behaviors) IsValid() bool {
	for _, existing := range AllowedToolsAwemeAuthorInfoGetV2BehaviorsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_aweme_author_info_get_v2_behaviors value
func (v ToolsAwemeAuthorInfoGetV2Behaviors) Ptr() *ToolsAwemeAuthorInfoGetV2Behaviors {
	return &v
}
