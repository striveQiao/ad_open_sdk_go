/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportVideoFrameGetV2ResponseDataListInnerMetrics
type ReportVideoFrameGetV2ResponseDataListInnerMetrics struct {
	//
	ClickCnt *int64 `json:"click_cnt,omitempty"`
	//
	DyComment *int64 `json:"dy_comment,omitempty"`
	//
	DyFollow *int64 `json:"dy_follow,omitempty"`
	//
	DyLike *int64 `json:"dy_like,omitempty"`
	//
	UserLoseCnt *int64 `json:"user_lose_cnt,omitempty"`
}
