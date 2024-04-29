/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// FileVideoMaterialClearTaskGetV2DataListClearTaskParamsMaterialSource
type FileVideoMaterialClearTaskGetV2DataListClearTaskParamsMaterialSource string

// List of file_video_material_clear_task_get_v2_data_list_clear_task_params_material_source
const (
	AD_FileVideoMaterialClearTaskGetV2DataListClearTaskParamsMaterialSource        FileVideoMaterialClearTaskGetV2DataListClearTaskParamsMaterialSource = "AD"
	QIANCHUAN_FileVideoMaterialClearTaskGetV2DataListClearTaskParamsMaterialSource FileVideoMaterialClearTaskGetV2DataListClearTaskParamsMaterialSource = "QIANCHUAN"
)

// All allowed values of FileVideoMaterialClearTaskGetV2DataListClearTaskParamsMaterialSource enum
var AllowedFileVideoMaterialClearTaskGetV2DataListClearTaskParamsMaterialSourceEnumValues = []FileVideoMaterialClearTaskGetV2DataListClearTaskParamsMaterialSource{
	"AD",
	"QIANCHUAN",
}

// NewFileVideoMaterialClearTaskGetV2DataListClearTaskParamsMaterialSourceFromValue returns a pointer to a valid FileVideoMaterialClearTaskGetV2DataListClearTaskParamsMaterialSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileVideoMaterialClearTaskGetV2DataListClearTaskParamsMaterialSourceFromValue(v string) (*FileVideoMaterialClearTaskGetV2DataListClearTaskParamsMaterialSource, error) {
	ev := FileVideoMaterialClearTaskGetV2DataListClearTaskParamsMaterialSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileVideoMaterialClearTaskGetV2DataListClearTaskParamsMaterialSource: valid values are %v", v, AllowedFileVideoMaterialClearTaskGetV2DataListClearTaskParamsMaterialSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileVideoMaterialClearTaskGetV2DataListClearTaskParamsMaterialSource) IsValid() bool {
	for _, existing := range AllowedFileVideoMaterialClearTaskGetV2DataListClearTaskParamsMaterialSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_video_material_clear_task_get_v2_data_list_clear_task_params_material_source value
func (v FileVideoMaterialClearTaskGetV2DataListClearTaskParamsMaterialSource) Ptr() *FileVideoMaterialClearTaskGetV2DataListClearTaskParamsMaterialSource {
	return &v
}
