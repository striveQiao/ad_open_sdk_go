/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanFileVideoAwemeGetV10ResponseData
type QianchuanFileVideoAwemeGetV10ResponseData struct {
	PageInfo *QianchuanFileVideoAwemeGetV10ResponseDataPageInfo `json:"page_info,omitempty"`
	//
	VideoList []*QianchuanFileVideoAwemeGetV10ResponseDataVideoListInner `json:"video_list,omitempty"`
}
