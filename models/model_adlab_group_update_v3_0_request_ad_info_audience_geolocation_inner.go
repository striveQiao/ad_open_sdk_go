/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupUpdateV30RequestAdInfoAudienceGeolocationInner struct for AdlabGroupUpdateV30RequestAdInfoAudienceGeolocationInner
type AdlabGroupUpdateV30RequestAdInfoAudienceGeolocationInner struct {
	// 纬度
	Lat *float64 `json:"lat,omitempty"`
	// 经度
	Long *float64 `json:"long,omitempty"`
	// 地点名称
	Name *string `json:"name,omitempty"`
	// 半径，单位为m，允许值为：3000-15000m
	Radius *int64 `json:"radius,omitempty"`
}
