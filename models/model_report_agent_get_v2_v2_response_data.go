/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAgentGetV2V2ResponseData
type ReportAgentGetV2V2ResponseData struct {
	CursorInfo *ReportAgentGetV2V2ResponseDataCursorInfo `json:"cursor_info,omitempty"`
	//
	List     []*ReportAgentGetV2V2ResponseDataListInner   `json:"list,omitempty"`
	PageInfo *AgentAdvertiserSelectV2ResponseDataPageInfo `json:"page_info,omitempty"`
}
