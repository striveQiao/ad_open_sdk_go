/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCustomConfigGetV30ResponseDataListInner struct for ReportCustomConfigGetV30ResponseDataListInner
type ReportCustomConfigGetV30ResponseDataListInner struct {
	DataTopic *ReportCustomConfigGetV30DataListDataTopic `json:"data_topic,omitempty"`
	//
	Dimensions []*ReportCustomConfigGetV30ResponseDataListInnerDimensionsInner `json:"dimensions,omitempty"`
	//
	Metrics []*ReportCustomConfigGetV30ResponseDataListInnerMetricsInner `json:"metrics,omitempty"`
}
