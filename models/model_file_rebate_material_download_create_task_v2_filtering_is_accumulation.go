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

// FileRebateMaterialDownloadCreateTaskV2FilteringIsAccumulation
type FileRebateMaterialDownloadCreateTaskV2FilteringIsAccumulation string

// List of file_rebate_material_download_create_task_v2_filtering_is_accumulation
const (
	NO_FileRebateMaterialDownloadCreateTaskV2FilteringIsAccumulation  FileRebateMaterialDownloadCreateTaskV2FilteringIsAccumulation = "NO"
	YES_FileRebateMaterialDownloadCreateTaskV2FilteringIsAccumulation FileRebateMaterialDownloadCreateTaskV2FilteringIsAccumulation = "YES"
)

// All allowed values of FileRebateMaterialDownloadCreateTaskV2FilteringIsAccumulation enum
var AllowedFileRebateMaterialDownloadCreateTaskV2FilteringIsAccumulationEnumValues = []FileRebateMaterialDownloadCreateTaskV2FilteringIsAccumulation{
	"NO",
	"YES",
}

// NewFileRebateMaterialDownloadCreateTaskV2FilteringIsAccumulationFromValue returns a pointer to a valid FileRebateMaterialDownloadCreateTaskV2FilteringIsAccumulation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileRebateMaterialDownloadCreateTaskV2FilteringIsAccumulationFromValue(v string) (*FileRebateMaterialDownloadCreateTaskV2FilteringIsAccumulation, error) {
	ev := FileRebateMaterialDownloadCreateTaskV2FilteringIsAccumulation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileRebateMaterialDownloadCreateTaskV2FilteringIsAccumulation: valid values are %v", v, AllowedFileRebateMaterialDownloadCreateTaskV2FilteringIsAccumulationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileRebateMaterialDownloadCreateTaskV2FilteringIsAccumulation) IsValid() bool {
	for _, existing := range AllowedFileRebateMaterialDownloadCreateTaskV2FilteringIsAccumulationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_rebate_material_download_create_task_v2_filtering_is_accumulation value
func (v FileRebateMaterialDownloadCreateTaskV2FilteringIsAccumulation) Ptr() *FileRebateMaterialDownloadCreateTaskV2FilteringIsAccumulation {
	return &v
}
