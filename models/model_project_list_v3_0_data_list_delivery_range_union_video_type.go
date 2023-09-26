/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListDeliveryRangeUnionVideoType
type ProjectListV30DataListDeliveryRangeUnionVideoType string

// List of project_list_v3.0_data_list_delivery_range_union_video_type
const (
	ORIGINAL_VIDEO_ProjectListV30DataListDeliveryRangeUnionVideoType ProjectListV30DataListDeliveryRangeUnionVideoType = "ORIGINAL_VIDEO"
	REWARDED_VIDEO_ProjectListV30DataListDeliveryRangeUnionVideoType ProjectListV30DataListDeliveryRangeUnionVideoType = "REWARDED_VIDEO"
	SPLASH_VIDEO_ProjectListV30DataListDeliveryRangeUnionVideoType   ProjectListV30DataListDeliveryRangeUnionVideoType = "SPLASH_VIDEO"
)

// All allowed values of ProjectListV30DataListDeliveryRangeUnionVideoType enum
var AllowedProjectListV30DataListDeliveryRangeUnionVideoTypeEnumValues = []ProjectListV30DataListDeliveryRangeUnionVideoType{
	"ORIGINAL_VIDEO",
	"REWARDED_VIDEO",
	"SPLASH_VIDEO",
}

// NewProjectListV30DataListDeliveryRangeUnionVideoTypeFromValue returns a pointer to a valid ProjectListV30DataListDeliveryRangeUnionVideoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListDeliveryRangeUnionVideoTypeFromValue(v string) (*ProjectListV30DataListDeliveryRangeUnionVideoType, error) {
	ev := ProjectListV30DataListDeliveryRangeUnionVideoType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListDeliveryRangeUnionVideoType: valid values are %v", v, AllowedProjectListV30DataListDeliveryRangeUnionVideoTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListDeliveryRangeUnionVideoType) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListDeliveryRangeUnionVideoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_delivery_range_union_video_type value
func (v ProjectListV30DataListDeliveryRangeUnionVideoType) Ptr() *ProjectListV30DataListDeliveryRangeUnionVideoType {
	return &v
}
