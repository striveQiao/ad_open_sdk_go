/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsUnionFlowPackageGetV2ResponseDataListInner struct for ToolsUnionFlowPackageGetV2ResponseDataListInner
type ToolsUnionFlowPackageGetV2ResponseDataListInner struct {
	//
	FlowPackageId   *int64                                             `json:"flow_package_id,omitempty"`
	FlowPackageType *ToolsUnionFlowPackageGetV2DataListFlowPackageType `json:"flow_package_type,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	Rit    []int64                                   `json:"rit,omitempty"`
	Status *ToolsUnionFlowPackageGetV2DataListStatus `json:"status,omitempty"`
}
