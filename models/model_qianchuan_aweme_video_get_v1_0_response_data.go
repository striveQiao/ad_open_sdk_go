/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeVideoGetV10ResponseData
type QianchuanAwemeVideoGetV10ResponseData struct {
	PageInfo *QianchuanAwemeVideoGetV10ResponseDataPageInfo `json:"page_info,omitempty"`
	//
	VideoList []*QianchuanAwemeVideoGetV10ResponseDataVideoListInner `json:"video_list,omitempty"`
}
