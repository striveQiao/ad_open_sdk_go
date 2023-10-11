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

// QianchuanAwemeSuggestBidV10AudienceAudienceMode
type QianchuanAwemeSuggestBidV10AudienceAudienceMode string

// List of qianchuan_aweme_suggest_bid_v1.0_audience_audience_mode
const (
	ATUO_QianchuanAwemeSuggestBidV10AudienceAudienceMode     QianchuanAwemeSuggestBidV10AudienceAudienceMode = "ATUO"
	CUSTOM_QianchuanAwemeSuggestBidV10AudienceAudienceMode   QianchuanAwemeSuggestBidV10AudienceAudienceMode = "CUSTOM"
	FANS_QianchuanAwemeSuggestBidV10AudienceAudienceMode     QianchuanAwemeSuggestBidV10AudienceAudienceMode = "FANS"
	LIVEFANS_QianchuanAwemeSuggestBidV10AudienceAudienceMode QianchuanAwemeSuggestBidV10AudienceAudienceMode = "LIVEFANS"
)

// All allowed values of QianchuanAwemeSuggestBidV10AudienceAudienceMode enum
var AllowedQianchuanAwemeSuggestBidV10AudienceAudienceModeEnumValues = []QianchuanAwemeSuggestBidV10AudienceAudienceMode{
	"ATUO",
	"CUSTOM",
	"FANS",
	"LIVEFANS",
}

// NewQianchuanAwemeSuggestBidV10AudienceAudienceModeFromValue returns a pointer to a valid QianchuanAwemeSuggestBidV10AudienceAudienceMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeSuggestBidV10AudienceAudienceModeFromValue(v string) (*QianchuanAwemeSuggestBidV10AudienceAudienceMode, error) {
	ev := QianchuanAwemeSuggestBidV10AudienceAudienceMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeSuggestBidV10AudienceAudienceMode: valid values are %v", v, AllowedQianchuanAwemeSuggestBidV10AudienceAudienceModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeSuggestBidV10AudienceAudienceMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeSuggestBidV10AudienceAudienceModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_suggest_bid_v1.0_audience_audience_mode value
func (v QianchuanAwemeSuggestBidV10AudienceAudienceMode) Ptr() *QianchuanAwemeSuggestBidV10AudienceAudienceMode {
	return &v
}
