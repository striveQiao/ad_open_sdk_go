/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FilePreauditGetV30ResponseDataListInner struct for FilePreauditGetV30ResponseDataListInner
type FilePreauditGetV30ResponseDataListInner struct {
	//
	Errmsg          string                                    `json:"errmsg"`
	MaterialChannel FilePreauditGetV30DataListMaterialChannel `json:"material_channel"`
	MaterialType    FilePreauditGetV30DataListMaterialType    `json:"material_type"`
	//
	PreauditId int64                            `json:"preaudit_id"`
	Status     FilePreauditGetV30DataListStatus `json:"status"`
	//
	VideoUrl string `json:"video_url"`
}
