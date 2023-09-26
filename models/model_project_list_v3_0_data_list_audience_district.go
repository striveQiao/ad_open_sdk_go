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

// ProjectListV30DataListAudienceDistrict
type ProjectListV30DataListAudienceDistrict string

// List of project_list_v3.0_data_list_audience_district
const (
	BUSINESS_DISTRICT_ProjectListV30DataListAudienceDistrict ProjectListV30DataListAudienceDistrict = "BUSINESS_DISTRICT"
	CITY_ProjectListV30DataListAudienceDistrict              ProjectListV30DataListAudienceDistrict = "CITY"
	COUNTY_ProjectListV30DataListAudienceDistrict            ProjectListV30DataListAudienceDistrict = "COUNTY"
	NONE_ProjectListV30DataListAudienceDistrict              ProjectListV30DataListAudienceDistrict = "NONE"
	OVERSEA_ProjectListV30DataListAudienceDistrict           ProjectListV30DataListAudienceDistrict = "OVERSEA"
	REGION_ProjectListV30DataListAudienceDistrict            ProjectListV30DataListAudienceDistrict = "REGION"
)

// All allowed values of ProjectListV30DataListAudienceDistrict enum
var AllowedProjectListV30DataListAudienceDistrictEnumValues = []ProjectListV30DataListAudienceDistrict{
	"BUSINESS_DISTRICT",
	"CITY",
	"COUNTY",
	"NONE",
	"OVERSEA",
	"REGION",
}

// NewProjectListV30DataListAudienceDistrictFromValue returns a pointer to a valid ProjectListV30DataListAudienceDistrict
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceDistrictFromValue(v string) (*ProjectListV30DataListAudienceDistrict, error) {
	ev := ProjectListV30DataListAudienceDistrict(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceDistrict: valid values are %v", v, AllowedProjectListV30DataListAudienceDistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceDistrict) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceDistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_district value
func (v ProjectListV30DataListAudienceDistrict) Ptr() *ProjectListV30DataListAudienceDistrict {
	return &v
}
