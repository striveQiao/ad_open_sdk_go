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

// FileRebateMaterialDownloadCreateTaskV2FilteringOperatorTag
type FileRebateMaterialDownloadCreateTaskV2FilteringOperatorTag string

// List of file_rebate_material_download_create_task_v2_filtering_operator_tag
const (
	SHOULIANG_FileRebateMaterialDownloadCreateTaskV2FilteringOperatorTag FileRebateMaterialDownloadCreateTaskV2FilteringOperatorTag = "SHOULIANG"
	ZOULIANG_FileRebateMaterialDownloadCreateTaskV2FilteringOperatorTag  FileRebateMaterialDownloadCreateTaskV2FilteringOperatorTag = "ZOULIANG"
	ZIYUNYING_FileRebateMaterialDownloadCreateTaskV2FilteringOperatorTag FileRebateMaterialDownloadCreateTaskV2FilteringOperatorTag = "ZIYUNYING"
)

// All allowed values of FileRebateMaterialDownloadCreateTaskV2FilteringOperatorTag enum
var AllowedFileRebateMaterialDownloadCreateTaskV2FilteringOperatorTagEnumValues = []FileRebateMaterialDownloadCreateTaskV2FilteringOperatorTag{
	"SHOULIANG",
	"ZOULIANG",
	"ZIYUNYING",
}

// NewFileRebateMaterialDownloadCreateTaskV2FilteringOperatorTagFromValue returns a pointer to a valid FileRebateMaterialDownloadCreateTaskV2FilteringOperatorTag
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileRebateMaterialDownloadCreateTaskV2FilteringOperatorTagFromValue(v string) (*FileRebateMaterialDownloadCreateTaskV2FilteringOperatorTag, error) {
	ev := FileRebateMaterialDownloadCreateTaskV2FilteringOperatorTag(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileRebateMaterialDownloadCreateTaskV2FilteringOperatorTag: valid values are %v", v, AllowedFileRebateMaterialDownloadCreateTaskV2FilteringOperatorTagEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileRebateMaterialDownloadCreateTaskV2FilteringOperatorTag) IsValid() bool {
	for _, existing := range AllowedFileRebateMaterialDownloadCreateTaskV2FilteringOperatorTagEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_rebate_material_download_create_task_v2_filtering_operator_tag value
func (v FileRebateMaterialDownloadCreateTaskV2FilteringOperatorTag) Ptr() *FileRebateMaterialDownloadCreateTaskV2FilteringOperatorTag {
	return &v
}
