/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsMicroAppListV30Filtering
type ToolsMicroAppListV30Filtering struct {
	AuditStatus *ToolsMicroAppListV30FilteringAuditStatus `json:"audit_status,omitempty"`
	CreateTime  *ToolsMicroAppListV30FilteringCreateTime  `json:"create_time,omitempty"`
	//
	SearchKey  *string                                  `json:"search_key,omitempty"`
	SearchType *ToolsMicroAppListV30FilteringSearchType `json:"search_type,omitempty"`
}
