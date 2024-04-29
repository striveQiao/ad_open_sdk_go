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

// ReportAudienceInterestActionListV2FilteringActionScene
type ReportAudienceInterestActionListV2FilteringActionScene string

// List of report_audience_interest_action_list_v2_filtering_action_scene
const (
	E_COMMERCE_ReportAudienceInterestActionListV2FilteringActionScene ReportAudienceInterestActionListV2FilteringActionScene = "E-COMMERCE"
	NEWS_ReportAudienceInterestActionListV2FilteringActionScene       ReportAudienceInterestActionListV2FilteringActionScene = "NEWS"
	APP_ReportAudienceInterestActionListV2FilteringActionScene        ReportAudienceInterestActionListV2FilteringActionScene = "APP"
	SEARCH_ReportAudienceInterestActionListV2FilteringActionScene     ReportAudienceInterestActionListV2FilteringActionScene = "SEARCH"
	AD_ReportAudienceInterestActionListV2FilteringActionScene         ReportAudienceInterestActionListV2FilteringActionScene = "AD"
)

// All allowed values of ReportAudienceInterestActionListV2FilteringActionScene enum
var AllowedReportAudienceInterestActionListV2FilteringActionSceneEnumValues = []ReportAudienceInterestActionListV2FilteringActionScene{
	"E-COMMERCE",
	"NEWS",
	"APP",
	"SEARCH",
	"AD",
}

// NewReportAudienceInterestActionListV2FilteringActionSceneFromValue returns a pointer to a valid ReportAudienceInterestActionListV2FilteringActionScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAudienceInterestActionListV2FilteringActionSceneFromValue(v string) (*ReportAudienceInterestActionListV2FilteringActionScene, error) {
	ev := ReportAudienceInterestActionListV2FilteringActionScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAudienceInterestActionListV2FilteringActionScene: valid values are %v", v, AllowedReportAudienceInterestActionListV2FilteringActionSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAudienceInterestActionListV2FilteringActionScene) IsValid() bool {
	for _, existing := range AllowedReportAudienceInterestActionListV2FilteringActionSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_audience_interest_action_list_v2_filtering_action_scene value
func (v ReportAudienceInterestActionListV2FilteringActionScene) Ptr() *ReportAudienceInterestActionListV2FilteringActionScene {
	return &v
}
