/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAudienceAwemeListV2ResponseDataListInner struct for ReportAudienceAwemeListV2ResponseDataListInner
type ReportAudienceAwemeListV2ResponseDataListInner struct {
	//
	AudienceLevel *string `json:"audience_level,omitempty"`
	//
	LabelName *string `json:"label_name,omitempty"`
	//
	Metrics map[string]interface{} `json:"metrics,omitempty"`
	//
	SuperiorLabelName *string `json:"superior_label_name,omitempty"`
}
