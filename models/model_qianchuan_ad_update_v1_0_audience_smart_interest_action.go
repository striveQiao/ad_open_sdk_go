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

// QianchuanAdUpdateV10AudienceSmartInterestAction
type QianchuanAdUpdateV10AudienceSmartInterestAction string

// List of qianchuan_ad_update_v1.0_audience_smart_interest_action
const (
	CUSTOM_QianchuanAdUpdateV10AudienceSmartInterestAction    QianchuanAdUpdateV10AudienceSmartInterestAction = "CUSTOM"
	RECOMMEND_QianchuanAdUpdateV10AudienceSmartInterestAction QianchuanAdUpdateV10AudienceSmartInterestAction = "RECOMMEND"
)

// All allowed values of QianchuanAdUpdateV10AudienceSmartInterestAction enum
var AllowedQianchuanAdUpdateV10AudienceSmartInterestActionEnumValues = []QianchuanAdUpdateV10AudienceSmartInterestAction{
	"CUSTOM",
	"RECOMMEND",
}

// NewQianchuanAdUpdateV10AudienceSmartInterestActionFromValue returns a pointer to a valid QianchuanAdUpdateV10AudienceSmartInterestAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10AudienceSmartInterestActionFromValue(v string) (*QianchuanAdUpdateV10AudienceSmartInterestAction, error) {
	ev := QianchuanAdUpdateV10AudienceSmartInterestAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10AudienceSmartInterestAction: valid values are %v", v, AllowedQianchuanAdUpdateV10AudienceSmartInterestActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10AudienceSmartInterestAction) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10AudienceSmartInterestActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_audience_smart_interest_action value
func (v QianchuanAdUpdateV10AudienceSmartInterestAction) Ptr() *QianchuanAdUpdateV10AudienceSmartInterestAction {
	return &v
}
