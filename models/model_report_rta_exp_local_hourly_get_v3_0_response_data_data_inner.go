/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportRtaExpLocalHourlyGetV30ResponseDataDataInner struct for ReportRtaExpLocalHourlyGetV30ResponseDataDataInner
type ReportRtaExpLocalHourlyGetV30ResponseDataDataInner struct {
	//
	Click int64 `json:"click"`
	//
	Convert int64 `json:"convert"`
	//
	Cost float64 `json:"cost"`
	//
	CusVid *int64 `json:"cus_vid,omitempty"`
	//
	Date string `json:"date"`
	//
	Show int64 `json:"show"`
	//
	Vid int64 `json:"vid"`
}
