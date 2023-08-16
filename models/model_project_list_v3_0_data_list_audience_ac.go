/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListAudienceAc
type ProjectListV30DataListAudienceAc string

// List of project_list_v3.0_data_list_audience_ac
const (
	Enum_2_G_ProjectListV30DataListAudienceAc ProjectListV30DataListAudienceAc = "2G"
	Enum_3_G_ProjectListV30DataListAudienceAc ProjectListV30DataListAudienceAc = "3G"
	Enum_4_G_ProjectListV30DataListAudienceAc ProjectListV30DataListAudienceAc = "4G"
	Enum_5_G_ProjectListV30DataListAudienceAc ProjectListV30DataListAudienceAc = "5G"
	WIFI_ProjectListV30DataListAudienceAc     ProjectListV30DataListAudienceAc = "WIFI"
)

// All allowed values of ProjectListV30DataListAudienceAc enum
var AllowedProjectListV30DataListAudienceAcEnumValues = []ProjectListV30DataListAudienceAc{
	"2G",
	"3G",
	"4G",
	"5G",
	"WIFI",
}

// NewProjectListV30DataListAudienceAcFromValue returns a pointer to a valid ProjectListV30DataListAudienceAc
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceAcFromValue(v string) (*ProjectListV30DataListAudienceAc, error) {
	ev := ProjectListV30DataListAudienceAc(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceAc: valid values are %v", v, AllowedProjectListV30DataListAudienceAcEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceAc) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceAcEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_ac value
func (v ProjectListV30DataListAudienceAc) Ptr() *ProjectListV30DataListAudienceAc {
	return &v
}
