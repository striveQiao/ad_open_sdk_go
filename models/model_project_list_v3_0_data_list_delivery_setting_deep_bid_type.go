/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListDeliverySettingDeepBidType
type ProjectListV30DataListDeliverySettingDeepBidType string

// List of project_list_v3.0_data_list_delivery_setting_deep_bid_type
const (
	ALI_FL_ProjectListV30DataListDeliverySettingDeepBidType                       ProjectListV30DataListDeliverySettingDeepBidType = "ALI_FL"
	AUTO_MIN_SECOND_STAGE_ProjectListV30DataListDeliverySettingDeepBidType        ProjectListV30DataListDeliverySettingDeepBidType = "AUTO_MIN_SECOND_STAGE"
	BID_PER_ACTION_ProjectListV30DataListDeliverySettingDeepBidType               ProjectListV30DataListDeliverySettingDeepBidType = "BID_PER_ACTION"
	DEEP_BID_DEFAULT_ProjectListV30DataListDeliverySettingDeepBidType             ProjectListV30DataListDeliverySettingDeepBidType = "DEEP_BID_DEFAULT"
	DEEP_BID_MIN_ProjectListV30DataListDeliverySettingDeepBidType                 ProjectListV30DataListDeliverySettingDeepBidType = "DEEP_BID_MIN"
	DEEP_BID_PACING_ProjectListV30DataListDeliverySettingDeepBidType              ProjectListV30DataListDeliverySettingDeepBidType = "DEEP_BID_PACING"
	DEEP_BID_TYPE_RETENTION_DAYS_ProjectListV30DataListDeliverySettingDeepBidType ProjectListV30DataListDeliverySettingDeepBidType = "DEEP_BID_TYPE_RETENTION_DAYS"
	FIRST_AND_SEVEN_PAY_ROI_ProjectListV30DataListDeliverySettingDeepBidType      ProjectListV30DataListDeliverySettingDeepBidType = "FIRST_AND_SEVEN_PAY_ROI"
	FORM_BID_ProjectListV30DataListDeliverySettingDeepBidType                     ProjectListV30DataListDeliverySettingDeepBidType = "FORM_BID"
	MIN_SECOND_STAGE_ProjectListV30DataListDeliverySettingDeepBidType             ProjectListV30DataListDeliverySettingDeepBidType = "MIN_SECOND_STAGE"
	PACING_SECOND_STAGE_ProjectListV30DataListDeliverySettingDeepBidType          ProjectListV30DataListDeliverySettingDeepBidType = "PACING_SECOND_STAGE"
	PER_AND_SEVEN_PAY_ROI_ProjectListV30DataListDeliverySettingDeepBidType        ProjectListV30DataListDeliverySettingDeepBidType = "PER_AND_SEVEN_PAY_ROI"
	PHONE_CONNECT_BID_ProjectListV30DataListDeliverySettingDeepBidType            ProjectListV30DataListDeliverySettingDeepBidType = "PHONE_CONNECT_BID"
	ROI_COEFFICIENT_ProjectListV30DataListDeliverySettingDeepBidType              ProjectListV30DataListDeliverySettingDeepBidType = "ROI_COEFFICIENT"
	ROI_DIRECT_MAIL_ProjectListV30DataListDeliverySettingDeepBidType              ProjectListV30DataListDeliverySettingDeepBidType = "ROI_DIRECT_MAIL"
	ROI_PACING_ProjectListV30DataListDeliverySettingDeepBidType                   ProjectListV30DataListDeliverySettingDeepBidType = "ROI_PACING"
	SMARTBID_ProjectListV30DataListDeliverySettingDeepBidType                     ProjectListV30DataListDeliverySettingDeepBidType = "SMARTBID"
	SOCIAL_ROI_ProjectListV30DataListDeliverySettingDeepBidType                   ProjectListV30DataListDeliverySettingDeepBidType = "SOCIAL_ROI"
)

// All allowed values of ProjectListV30DataListDeliverySettingDeepBidType enum
var AllowedProjectListV30DataListDeliverySettingDeepBidTypeEnumValues = []ProjectListV30DataListDeliverySettingDeepBidType{
	"ALI_FL",
	"AUTO_MIN_SECOND_STAGE",
	"BID_PER_ACTION",
	"DEEP_BID_DEFAULT",
	"DEEP_BID_MIN",
	"DEEP_BID_PACING",
	"DEEP_BID_TYPE_RETENTION_DAYS",
	"FIRST_AND_SEVEN_PAY_ROI",
	"FORM_BID",
	"MIN_SECOND_STAGE",
	"PACING_SECOND_STAGE",
	"PER_AND_SEVEN_PAY_ROI",
	"PHONE_CONNECT_BID",
	"ROI_COEFFICIENT",
	"ROI_DIRECT_MAIL",
	"ROI_PACING",
	"SMARTBID",
	"SOCIAL_ROI",
}

// NewProjectListV30DataListDeliverySettingDeepBidTypeFromValue returns a pointer to a valid ProjectListV30DataListDeliverySettingDeepBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListDeliverySettingDeepBidTypeFromValue(v string) (*ProjectListV30DataListDeliverySettingDeepBidType, error) {
	ev := ProjectListV30DataListDeliverySettingDeepBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListDeliverySettingDeepBidType: valid values are %v", v, AllowedProjectListV30DataListDeliverySettingDeepBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListDeliverySettingDeepBidType) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListDeliverySettingDeepBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_delivery_setting_deep_bid_type value
func (v ProjectListV30DataListDeliverySettingDeepBidType) Ptr() *ProjectListV30DataListDeliverySettingDeepBidType {
	return &v
}
