/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportStardeliveryTaskDataGetV30DataListStarTaskExternalAction
type ReportStardeliveryTaskDataGetV30DataListStarTaskExternalAction string

// List of report_stardelivery_task_data_get_v3.0_data_list_star_task_external_action
const (
	AD_CONVERT_TYPE_ACTIVE_ReportStardeliveryTaskDataGetV30DataListStarTaskExternalAction          ReportStardeliveryTaskDataGetV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_ACTIVE"
	AD_CONVERT_TYPE_ACTIVE_REGISTER_ReportStardeliveryTaskDataGetV30DataListStarTaskExternalAction ReportStardeliveryTaskDataGetV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_ACTIVE_REGISTER"
	AD_CONVERT_TYPE_PAY_ReportStardeliveryTaskDataGetV30DataListStarTaskExternalAction             ReportStardeliveryTaskDataGetV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_PAY"
)

// All allowed values of ReportStardeliveryTaskDataGetV30DataListStarTaskExternalAction enum
var AllowedReportStardeliveryTaskDataGetV30DataListStarTaskExternalActionEnumValues = []ReportStardeliveryTaskDataGetV30DataListStarTaskExternalAction{
	"AD_CONVERT_TYPE_ACTIVE",
	"AD_CONVERT_TYPE_ACTIVE_REGISTER",
	"AD_CONVERT_TYPE_PAY",
}

// NewReportStardeliveryTaskDataGetV30DataListStarTaskExternalActionFromValue returns a pointer to a valid ReportStardeliveryTaskDataGetV30DataListStarTaskExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportStardeliveryTaskDataGetV30DataListStarTaskExternalActionFromValue(v string) (*ReportStardeliveryTaskDataGetV30DataListStarTaskExternalAction, error) {
	ev := ReportStardeliveryTaskDataGetV30DataListStarTaskExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportStardeliveryTaskDataGetV30DataListStarTaskExternalAction: valid values are %v", v, AllowedReportStardeliveryTaskDataGetV30DataListStarTaskExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportStardeliveryTaskDataGetV30DataListStarTaskExternalAction) IsValid() bool {
	for _, existing := range AllowedReportStardeliveryTaskDataGetV30DataListStarTaskExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_stardelivery_task_data_get_v3.0_data_list_star_task_external_action value
func (v ReportStardeliveryTaskDataGetV30DataListStarTaskExternalAction) Ptr() *ReportStardeliveryTaskDataGetV30DataListStarTaskExternalAction {
	return &v
}
