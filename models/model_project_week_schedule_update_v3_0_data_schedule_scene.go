/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectWeekScheduleUpdateV30DataScheduleScene
type ProjectWeekScheduleUpdateV30DataScheduleScene string

// List of project_week_schedule_update_v3.0_data_schedule_scene
const (
	NEXT_DAY_ProjectWeekScheduleUpdateV30DataScheduleScene ProjectWeekScheduleUpdateV30DataScheduleScene = "NEXT_DAY"
	REALTIME_ProjectWeekScheduleUpdateV30DataScheduleScene ProjectWeekScheduleUpdateV30DataScheduleScene = "REALTIME"
)

// All allowed values of ProjectWeekScheduleUpdateV30DataScheduleScene enum
var AllowedProjectWeekScheduleUpdateV30DataScheduleSceneEnumValues = []ProjectWeekScheduleUpdateV30DataScheduleScene{
	"NEXT_DAY",
	"REALTIME",
}

// NewProjectWeekScheduleUpdateV30DataScheduleSceneFromValue returns a pointer to a valid ProjectWeekScheduleUpdateV30DataScheduleScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectWeekScheduleUpdateV30DataScheduleSceneFromValue(v string) (*ProjectWeekScheduleUpdateV30DataScheduleScene, error) {
	ev := ProjectWeekScheduleUpdateV30DataScheduleScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectWeekScheduleUpdateV30DataScheduleScene: valid values are %v", v, AllowedProjectWeekScheduleUpdateV30DataScheduleSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectWeekScheduleUpdateV30DataScheduleScene) IsValid() bool {
	for _, existing := range AllowedProjectWeekScheduleUpdateV30DataScheduleSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_week_schedule_update_v3.0_data_schedule_scene value
func (v ProjectWeekScheduleUpdateV30DataScheduleScene) Ptr() *ProjectWeekScheduleUpdateV30DataScheduleScene {
	return &v
}
