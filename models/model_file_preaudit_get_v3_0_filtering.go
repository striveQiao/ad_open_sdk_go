/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FilePreauditGetV30Filtering
type FilePreauditGetV30Filtering struct {
	MaterialChannel *FilePreauditGetV30FilteringMaterialChannel `json:"material_channel,omitempty"`
	MaterialType    *FilePreauditGetV30FilteringMaterialType    `json:"material_type,omitempty"`
	//
	PreauditIds []int64                            `json:"preaudit_ids,omitempty"`
	Status      *FilePreauditGetV30FilteringStatus `json:"status,omitempty"`
}
