/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectBudgetUpdateV30DataBudgetMode
type ProjectBudgetUpdateV30DataBudgetMode string

// List of project_budget_update_v3.0_data_budget_mode
const (
	BUDGET_MODE_DAY_ProjectBudgetUpdateV30DataBudgetMode      ProjectBudgetUpdateV30DataBudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_INFINITE_ProjectBudgetUpdateV30DataBudgetMode ProjectBudgetUpdateV30DataBudgetMode = "BUDGET_MODE_INFINITE"
)

// All allowed values of ProjectBudgetUpdateV30DataBudgetMode enum
var AllowedProjectBudgetUpdateV30DataBudgetModeEnumValues = []ProjectBudgetUpdateV30DataBudgetMode{
	"BUDGET_MODE_DAY",
	"BUDGET_MODE_INFINITE",
}

// NewProjectBudgetUpdateV30DataBudgetModeFromValue returns a pointer to a valid ProjectBudgetUpdateV30DataBudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectBudgetUpdateV30DataBudgetModeFromValue(v string) (*ProjectBudgetUpdateV30DataBudgetMode, error) {
	ev := ProjectBudgetUpdateV30DataBudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectBudgetUpdateV30DataBudgetMode: valid values are %v", v, AllowedProjectBudgetUpdateV30DataBudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectBudgetUpdateV30DataBudgetMode) IsValid() bool {
	for _, existing := range AllowedProjectBudgetUpdateV30DataBudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_budget_update_v3.0_data_budget_mode value
func (v ProjectBudgetUpdateV30DataBudgetMode) Ptr() *ProjectBudgetUpdateV30DataBudgetMode {
	return &v
}
