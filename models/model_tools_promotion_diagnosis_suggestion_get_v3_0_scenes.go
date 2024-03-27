/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsPromotionDiagnosisSuggestionGetV30Scenes
type ToolsPromotionDiagnosisSuggestionGetV30Scenes string

// List of tools_promotion_diagnosis_suggestion_get_v3.0_scenes
const (
	CLEAN_ToolsPromotionDiagnosisSuggestionGetV30Scenes     ToolsPromotionDiagnosisSuggestionGetV30Scenes = "CLEAN"
	POTENTIAL_ToolsPromotionDiagnosisSuggestionGetV30Scenes ToolsPromotionDiagnosisSuggestionGetV30Scenes = "POTENTIAL"
	ZOMBIE_ToolsPromotionDiagnosisSuggestionGetV30Scenes    ToolsPromotionDiagnosisSuggestionGetV30Scenes = "ZOMBIE"
)

// All allowed values of ToolsPromotionDiagnosisSuggestionGetV30Scenes enum
var AllowedToolsPromotionDiagnosisSuggestionGetV30ScenesEnumValues = []ToolsPromotionDiagnosisSuggestionGetV30Scenes{
	"CLEAN",
	"POTENTIAL",
	"ZOMBIE",
}

// NewToolsPromotionDiagnosisSuggestionGetV30ScenesFromValue returns a pointer to a valid ToolsPromotionDiagnosisSuggestionGetV30Scenes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPromotionDiagnosisSuggestionGetV30ScenesFromValue(v string) (*ToolsPromotionDiagnosisSuggestionGetV30Scenes, error) {
	ev := ToolsPromotionDiagnosisSuggestionGetV30Scenes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPromotionDiagnosisSuggestionGetV30Scenes: valid values are %v", v, AllowedToolsPromotionDiagnosisSuggestionGetV30ScenesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPromotionDiagnosisSuggestionGetV30Scenes) IsValid() bool {
	for _, existing := range AllowedToolsPromotionDiagnosisSuggestionGetV30ScenesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_promotion_diagnosis_suggestion_get_v3.0_scenes value
func (v ToolsPromotionDiagnosisSuggestionGetV30Scenes) Ptr() *ToolsPromotionDiagnosisSuggestionGetV30Scenes {
	return &v
}
