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

// ProjectCreateV30BlueFlowPackageBlueFlowPackageSetting
type ProjectCreateV30BlueFlowPackageBlueFlowPackageSetting string

// List of project_create_v3.0_blue_flow_package_blue_flow_package_setting
const (
	OFF_ProjectCreateV30BlueFlowPackageBlueFlowPackageSetting ProjectCreateV30BlueFlowPackageBlueFlowPackageSetting = "OFF"
	ON_ProjectCreateV30BlueFlowPackageBlueFlowPackageSetting  ProjectCreateV30BlueFlowPackageBlueFlowPackageSetting = "ON"
)

// All allowed values of ProjectCreateV30BlueFlowPackageBlueFlowPackageSetting enum
var AllowedProjectCreateV30BlueFlowPackageBlueFlowPackageSettingEnumValues = []ProjectCreateV30BlueFlowPackageBlueFlowPackageSetting{
	"OFF",
	"ON",
}

// NewProjectCreateV30BlueFlowPackageBlueFlowPackageSettingFromValue returns a pointer to a valid ProjectCreateV30BlueFlowPackageBlueFlowPackageSetting
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30BlueFlowPackageBlueFlowPackageSettingFromValue(v string) (*ProjectCreateV30BlueFlowPackageBlueFlowPackageSetting, error) {
	ev := ProjectCreateV30BlueFlowPackageBlueFlowPackageSetting(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30BlueFlowPackageBlueFlowPackageSetting: valid values are %v", v, AllowedProjectCreateV30BlueFlowPackageBlueFlowPackageSettingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30BlueFlowPackageBlueFlowPackageSetting) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30BlueFlowPackageBlueFlowPackageSettingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_blue_flow_package_blue_flow_package_setting value
func (v ProjectCreateV30BlueFlowPackageBlueFlowPackageSetting) Ptr() *ProjectCreateV30BlueFlowPackageBlueFlowPackageSetting {
	return &v
}
