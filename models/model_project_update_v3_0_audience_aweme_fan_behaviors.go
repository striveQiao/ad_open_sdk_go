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

// ProjectUpdateV30AudienceAwemeFanBehaviors
type ProjectUpdateV30AudienceAwemeFanBehaviors string

// List of project_update_v3.0_audience_aweme_fan_behaviors
const (
	COMMENTED_USER_ProjectUpdateV30AudienceAwemeFanBehaviors       ProjectUpdateV30AudienceAwemeFanBehaviors = "COMMENTED_USER"
	FOLLOWED_USER_ProjectUpdateV30AudienceAwemeFanBehaviors        ProjectUpdateV30AudienceAwemeFanBehaviors = "FOLLOWED_USER"
	GOODS_CARTS_CLICK_ProjectUpdateV30AudienceAwemeFanBehaviors    ProjectUpdateV30AudienceAwemeFanBehaviors = "GOODS_CARTS_CLICK"
	GOODS_CARTS_ORDER_ProjectUpdateV30AudienceAwemeFanBehaviors    ProjectUpdateV30AudienceAwemeFanBehaviors = "GOODS_CARTS_ORDER"
	LIKED_USER_ProjectUpdateV30AudienceAwemeFanBehaviors           ProjectUpdateV30AudienceAwemeFanBehaviors = "LIKED_USER"
	LIVE_COMMENT_ProjectUpdateV30AudienceAwemeFanBehaviors         ProjectUpdateV30AudienceAwemeFanBehaviors = "LIVE_COMMENT"
	LIVE_EFFECTIVE_WATCH_ProjectUpdateV30AudienceAwemeFanBehaviors ProjectUpdateV30AudienceAwemeFanBehaviors = "LIVE_EFFECTIVE_WATCH"
	LIVE_EXCEPTIONAL_ProjectUpdateV30AudienceAwemeFanBehaviors     ProjectUpdateV30AudienceAwemeFanBehaviors = "LIVE_EXCEPTIONAL"
	LIVE_GOODS_CLICK_ProjectUpdateV30AudienceAwemeFanBehaviors     ProjectUpdateV30AudienceAwemeFanBehaviors = "LIVE_GOODS_CLICK"
	LIVE_GOODS_ORDER_ProjectUpdateV30AudienceAwemeFanBehaviors     ProjectUpdateV30AudienceAwemeFanBehaviors = "LIVE_GOODS_ORDER"
	LIVE_WATCH_ProjectUpdateV30AudienceAwemeFanBehaviors           ProjectUpdateV30AudienceAwemeFanBehaviors = "LIVE_WATCH"
	SHARED_USER_ProjectUpdateV30AudienceAwemeFanBehaviors          ProjectUpdateV30AudienceAwemeFanBehaviors = "SHARED_USER"
)

// All allowed values of ProjectUpdateV30AudienceAwemeFanBehaviors enum
var AllowedProjectUpdateV30AudienceAwemeFanBehaviorsEnumValues = []ProjectUpdateV30AudienceAwemeFanBehaviors{
	"COMMENTED_USER",
	"FOLLOWED_USER",
	"GOODS_CARTS_CLICK",
	"GOODS_CARTS_ORDER",
	"LIKED_USER",
	"LIVE_COMMENT",
	"LIVE_EFFECTIVE_WATCH",
	"LIVE_EXCEPTIONAL",
	"LIVE_GOODS_CLICK",
	"LIVE_GOODS_ORDER",
	"LIVE_WATCH",
	"SHARED_USER",
}

// NewProjectUpdateV30AudienceAwemeFanBehaviorsFromValue returns a pointer to a valid ProjectUpdateV30AudienceAwemeFanBehaviors
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudienceAwemeFanBehaviorsFromValue(v string) (*ProjectUpdateV30AudienceAwemeFanBehaviors, error) {
	ev := ProjectUpdateV30AudienceAwemeFanBehaviors(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudienceAwemeFanBehaviors: valid values are %v", v, AllowedProjectUpdateV30AudienceAwemeFanBehaviorsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudienceAwemeFanBehaviors) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudienceAwemeFanBehaviorsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_aweme_fan_behaviors value
func (v ProjectUpdateV30AudienceAwemeFanBehaviors) Ptr() *ProjectUpdateV30AudienceAwemeFanBehaviors {
	return &v
}
