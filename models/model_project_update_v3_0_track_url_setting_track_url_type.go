/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectUpdateV30TrackUrlSettingTrackUrlType
type ProjectUpdateV30TrackUrlSettingTrackUrlType string

// List of project_update_v3.0_track_url_setting_track_url_type
const (
	CUSTOM_ProjectUpdateV30TrackUrlSettingTrackUrlType   ProjectUpdateV30TrackUrlSettingTrackUrlType = "CUSTOM"
	GROUP_ID_ProjectUpdateV30TrackUrlSettingTrackUrlType ProjectUpdateV30TrackUrlSettingTrackUrlType = "GROUP_ID"
)

// All allowed values of ProjectUpdateV30TrackUrlSettingTrackUrlType enum
var AllowedProjectUpdateV30TrackUrlSettingTrackUrlTypeEnumValues = []ProjectUpdateV30TrackUrlSettingTrackUrlType{
	"CUSTOM",
	"GROUP_ID",
}

// NewProjectUpdateV30TrackUrlSettingTrackUrlTypeFromValue returns a pointer to a valid ProjectUpdateV30TrackUrlSettingTrackUrlType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30TrackUrlSettingTrackUrlTypeFromValue(v string) (*ProjectUpdateV30TrackUrlSettingTrackUrlType, error) {
	ev := ProjectUpdateV30TrackUrlSettingTrackUrlType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30TrackUrlSettingTrackUrlType: valid values are %v", v, AllowedProjectUpdateV30TrackUrlSettingTrackUrlTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30TrackUrlSettingTrackUrlType) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30TrackUrlSettingTrackUrlTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_track_url_setting_track_url_type value
func (v ProjectUpdateV30TrackUrlSettingTrackUrlType) Ptr() *ProjectUpdateV30TrackUrlSettingTrackUrlType {
	return &v
}
