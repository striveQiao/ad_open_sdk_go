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

// FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes
type FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes string

// List of file_video_material_clear_task_create_v2_clear_task_params_clear_material_types
const (
	INEFFICIENT_MATERIAL_FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes   FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes = "INEFFICIENT_MATERIAL"
	SIMILAR_MATERIAL_FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes       FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes = "SIMILAR_MATERIAL"
	SIMILAR_QUEUE_MATERIAL_FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes = "SIMILAR_QUEUE_MATERIAL"
)

// All allowed values of FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes enum
var AllowedFileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypesEnumValues = []FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes{
	"INEFFICIENT_MATERIAL",
	"SIMILAR_MATERIAL",
	"SIMILAR_QUEUE_MATERIAL",
}

// NewFileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypesFromValue returns a pointer to a valid FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypesFromValue(v string) (*FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes, error) {
	ev := FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes: valid values are %v", v, AllowedFileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes) IsValid() bool {
	for _, existing := range AllowedFileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_video_material_clear_task_create_v2_clear_task_params_clear_material_types value
func (v FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes) Ptr() *FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes {
	return &v
}
