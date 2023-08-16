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

// ProjectCreateV30Classify
type ProjectCreateV30Classify string

// List of project_create_v3.0_classify
const (
	CLASSIFY_APPORTION_ProjectCreateV30Classify  ProjectCreateV30Classify = "CLASSIFY_APPORTION"
	CLASSIFY_EXCHANGE_ProjectCreateV30Classify   ProjectCreateV30Classify = "CLASSIFY_EXCHANGE"
	CLASSIFY_INTERNAL_ProjectCreateV30Classify   ProjectCreateV30Classify = "CLASSIFY_INTERNAL"
	CLASSIFY_SALE_ProjectCreateV30Classify       ProjectCreateV30Classify = "CLASSIFY_SALE"
	CLASSIFY_SUPPLEMENT_ProjectCreateV30Classify ProjectCreateV30Classify = "CLASSIFY_SUPPLEMENT"
)

// All allowed values of ProjectCreateV30Classify enum
var AllowedProjectCreateV30ClassifyEnumValues = []ProjectCreateV30Classify{
	"CLASSIFY_APPORTION",
	"CLASSIFY_EXCHANGE",
	"CLASSIFY_INTERNAL",
	"CLASSIFY_SALE",
	"CLASSIFY_SUPPLEMENT",
}

// NewProjectCreateV30ClassifyFromValue returns a pointer to a valid ProjectCreateV30Classify
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30ClassifyFromValue(v string) (*ProjectCreateV30Classify, error) {
	ev := ProjectCreateV30Classify(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30Classify: valid values are %v", v, AllowedProjectCreateV30ClassifyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30Classify) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30ClassifyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_classify value
func (v ProjectCreateV30Classify) Ptr() *ProjectCreateV30Classify {
	return &v
}
