/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportRtaGetV2ResponseData
type ReportRtaGetV2ResponseData struct {
	//
	List         []*ReportRtaGetV2ResponseDataListInner  `json:"list,omitempty"`
	PageInfo     *ReportRtaGetV2ResponseDataPageInfo     `json:"page_info,omitempty"`
	TotalMetrics *ReportRtaGetV2ResponseDataTotalMetrics `json:"total_metrics,omitempty"`
}
