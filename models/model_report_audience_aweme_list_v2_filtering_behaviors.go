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

// ReportAudienceAwemeListV2FilteringBehaviors
type ReportAudienceAwemeListV2FilteringBehaviors string

// List of report_audience_aweme_list_v2_filtering_behaviors
const (
	FOLLOWED_USER_ReportAudienceAwemeListV2FilteringBehaviors        ReportAudienceAwemeListV2FilteringBehaviors = "FOLLOWED_USER"
	GOODS_CARTS_ORDER_ReportAudienceAwemeListV2FilteringBehaviors    ReportAudienceAwemeListV2FilteringBehaviors = "GOODS_CARTS_ORDER"
	GOODS_CARTS_CLICK_ReportAudienceAwemeListV2FilteringBehaviors    ReportAudienceAwemeListV2FilteringBehaviors = "GOODS_CARTS_CLICK"
	LIVE_EXCEPTIONAL_ReportAudienceAwemeListV2FilteringBehaviors     ReportAudienceAwemeListV2FilteringBehaviors = "LIVE_EXCEPTIONAL"
	LIVE_COMMENT_ReportAudienceAwemeListV2FilteringBehaviors         ReportAudienceAwemeListV2FilteringBehaviors = "LIVE_COMMENT"
	LIVE_GOODS_ORDER_ReportAudienceAwemeListV2FilteringBehaviors     ReportAudienceAwemeListV2FilteringBehaviors = "LIVE_GOODS_ORDER"
	COMMENTED_USER_ReportAudienceAwemeListV2FilteringBehaviors       ReportAudienceAwemeListV2FilteringBehaviors = "COMMENTED_USER"
	LIVE_GOODS_CLICK_ReportAudienceAwemeListV2FilteringBehaviors     ReportAudienceAwemeListV2FilteringBehaviors = "LIVE_GOODS_CLICK"
	LIVE_WATCH_ReportAudienceAwemeListV2FilteringBehaviors           ReportAudienceAwemeListV2FilteringBehaviors = "LIVE_WATCH"
	LIKED_USER_ReportAudienceAwemeListV2FilteringBehaviors           ReportAudienceAwemeListV2FilteringBehaviors = "LIKED_USER"
	SHARED_USER_ReportAudienceAwemeListV2FilteringBehaviors          ReportAudienceAwemeListV2FilteringBehaviors = "SHARED_USER"
	LIVE_EFFECTIVE_WATCH_ReportAudienceAwemeListV2FilteringBehaviors ReportAudienceAwemeListV2FilteringBehaviors = "LIVE_EFFECTIVE_WATCH"
)

// All allowed values of ReportAudienceAwemeListV2FilteringBehaviors enum
var AllowedReportAudienceAwemeListV2FilteringBehaviorsEnumValues = []ReportAudienceAwemeListV2FilteringBehaviors{
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

// NewReportAudienceAwemeListV2FilteringBehaviorsFromValue returns a pointer to a valid ReportAudienceAwemeListV2FilteringBehaviors
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAudienceAwemeListV2FilteringBehaviorsFromValue(v string) (*ReportAudienceAwemeListV2FilteringBehaviors, error) {
	ev := ReportAudienceAwemeListV2FilteringBehaviors(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAudienceAwemeListV2FilteringBehaviors: valid values are %v", v, AllowedReportAudienceAwemeListV2FilteringBehaviorsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAudienceAwemeListV2FilteringBehaviors) IsValid() bool {
	for _, existing := range AllowedReportAudienceAwemeListV2FilteringBehaviorsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_audience_aweme_list_v2_filtering_behaviors value
func (v ReportAudienceAwemeListV2FilteringBehaviors) Ptr() *ReportAudienceAwemeListV2FilteringBehaviors {
	return &v
}
