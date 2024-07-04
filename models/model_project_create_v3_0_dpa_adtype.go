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

// ProjectCreateV30DpaAdtype
type ProjectCreateV30DpaAdtype string

// List of project_create_v3.0_dpa_adtype
const (
	DPA_APP_ProjectCreateV30DpaAdtype  ProjectCreateV30DpaAdtype = "DPA_APP"
	DPA_LINK_ProjectCreateV30DpaAdtype ProjectCreateV30DpaAdtype = "DPA_LINK"
)

// All allowed values of ProjectCreateV30DpaAdtype enum
var AllowedProjectCreateV30DpaAdtypeEnumValues = []ProjectCreateV30DpaAdtype{
	"DPA_APP",
	"DPA_LINK",
}

// NewProjectCreateV30DpaAdtypeFromValue returns a pointer to a valid ProjectCreateV30DpaAdtype
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30DpaAdtypeFromValue(v string) (*ProjectCreateV30DpaAdtype, error) {
	ev := ProjectCreateV30DpaAdtype(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30DpaAdtype: valid values are %v", v, AllowedProjectCreateV30DpaAdtypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30DpaAdtype) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30DpaAdtypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_dpa_adtype value
func (v ProjectCreateV30DpaAdtype) Ptr() *ProjectCreateV30DpaAdtype {
	return &v
}
