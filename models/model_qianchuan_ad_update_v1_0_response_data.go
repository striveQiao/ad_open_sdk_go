/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdUpdateV10ResponseData
type QianchuanAdUpdateV10ResponseData struct {
	//
	AdId *int64 `json:"ad_id,omitempty"`
	//
	ErrorList []*QianchuanAdUpdateV10ResponseDataErrorListInner `json:"error_list,omitempty"`
	//
	NoticeInfos []*QianchuanAdUpdateV10ResponseDataNoticeInfosInner `json:"notice_infos,omitempty"`
}
