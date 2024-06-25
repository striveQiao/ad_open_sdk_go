/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsLandingGroupGetV2ResponseDataListInner struct for ToolsLandingGroupGetV2ResponseDataListInner
type ToolsLandingGroupGetV2ResponseDataListInner struct {
	GroupFlowType *ToolsLandingGroupGetV2DataListGroupFlowType `json:"group_flow_type,omitempty"`
	//
	GroupId     *int64                                     `json:"group_id,omitempty"`
	GroupStatus *ToolsLandingGroupGetV2DataListGroupStatus `json:"group_status,omitempty"`
	//
	GroupTitle *string `json:"group_title,omitempty"`
	//
	GroupUrl *string `json:"group_url,omitempty"`
	//
	Sites []*ToolsLandingGroupGetV2ResponseDataListInnerSitesInner `json:"sites,omitempty"`
}
