/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectListV30ResponseDataListInnerInternalAdvertiserInfo
type ProjectListV30ResponseDataListInnerInternalAdvertiserInfo struct {
	//
	AdvertisingVolume *int64                                                `json:"advertising_volume,omitempty"`
	Classify          *ProjectListV30DataListInternalAdvertiserInfoClassify `json:"classify,omitempty"`
	//
	ItemNumber *string `json:"item_number,omitempty"`
}
