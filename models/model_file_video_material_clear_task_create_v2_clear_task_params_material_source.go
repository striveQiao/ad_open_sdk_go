/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// FileVideoMaterialClearTaskCreateV2ClearTaskParamsMaterialSource
type FileVideoMaterialClearTaskCreateV2ClearTaskParamsMaterialSource string

// List of file_video_material_clear_task_create_v2_clear_task_params_material_source
const (
	AD_FileVideoMaterialClearTaskCreateV2ClearTaskParamsMaterialSource        FileVideoMaterialClearTaskCreateV2ClearTaskParamsMaterialSource = "AD"
	QIANCHUAN_FileVideoMaterialClearTaskCreateV2ClearTaskParamsMaterialSource FileVideoMaterialClearTaskCreateV2ClearTaskParamsMaterialSource = "QIANCHUAN"
)

// All allowed values of FileVideoMaterialClearTaskCreateV2ClearTaskParamsMaterialSource enum
var AllowedFileVideoMaterialClearTaskCreateV2ClearTaskParamsMaterialSourceEnumValues = []FileVideoMaterialClearTaskCreateV2ClearTaskParamsMaterialSource{
	"AD",
	"QIANCHUAN",
}

// NewFileVideoMaterialClearTaskCreateV2ClearTaskParamsMaterialSourceFromValue returns a pointer to a valid FileVideoMaterialClearTaskCreateV2ClearTaskParamsMaterialSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileVideoMaterialClearTaskCreateV2ClearTaskParamsMaterialSourceFromValue(v string) (*FileVideoMaterialClearTaskCreateV2ClearTaskParamsMaterialSource, error) {
	ev := FileVideoMaterialClearTaskCreateV2ClearTaskParamsMaterialSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileVideoMaterialClearTaskCreateV2ClearTaskParamsMaterialSource: valid values are %v", v, AllowedFileVideoMaterialClearTaskCreateV2ClearTaskParamsMaterialSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileVideoMaterialClearTaskCreateV2ClearTaskParamsMaterialSource) IsValid() bool {
	for _, existing := range AllowedFileVideoMaterialClearTaskCreateV2ClearTaskParamsMaterialSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_video_material_clear_task_create_v2_clear_task_params_material_source value
func (v FileVideoMaterialClearTaskCreateV2ClearTaskParamsMaterialSource) Ptr() *FileVideoMaterialClearTaskCreateV2ClearTaskParamsMaterialSource {
	return &v
}
