/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupDetailV30ResponseDataDataAdInfoAudienceGeolocationInner struct for AdlabGroupDetailV30ResponseDataDataAdInfoAudienceGeolocationInner
type AdlabGroupDetailV30ResponseDataDataAdInfoAudienceGeolocationInner struct {
	// 纬度
	Lat *float64 `json:"lat,omitempty"`
	// 经度
	Long *float64 `json:"long,omitempty"`
	// 地点名称
	Name *string `json:"name,omitempty"`
	// 半径
	Radius *int64 `json:"radius,omitempty"`
}
