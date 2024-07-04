/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorGetV30Filtering
type NativeAnchorGetV30Filtering struct {
	// 锚点创建结束日期，格式：yyyy-MM-dd
	AnchorEndTime *string `json:"anchor_end_time,omitempty"`
	//
	AnchorId *string `json:"anchor_id,omitempty"`
	//
	AnchorName *string `json:"anchor_name,omitempty"`
	// 锚点创建开始日期，格式：yyyy-MM-dd
	AnchorStartTime *string `json:"anchor_start_time,omitempty"`
	// 锚点类型，允许值： - 应用下载-游戏：APP_GAME - 应用下载-网服：APP_INTERNET_SERVICE - 应用下载-电商：APP_SHOP - 高级在线预约：ONLINE_SUBSCRIBE
	AnchorType []*NativeAnchorGetV30FilteringAnchorType `json:"anchor_type,omitempty"`
	//
	AndroidPackageName *string                                    `json:"android_package_name,omitempty"`
	ExternalAction     *NativeAnchorGetV30FilteringExternalAction `json:"external_action,omitempty"`
	//
	IosPackageName *string                                 `json:"ios_package_name,omitempty"`
	LandingType    *NativeAnchorGetV30FilteringLandingType `json:"landing_type,omitempty"`
	//
	Source []*NativeAnchorGetV30FilteringSource `json:"source,omitempty"`
	//
	Status []*NativeAnchorGetV30FilteringStatus `json:"status,omitempty"`
}
