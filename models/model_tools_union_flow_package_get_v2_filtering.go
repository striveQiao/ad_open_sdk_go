/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsUnionFlowPackageGetV2Filtering
type ToolsUnionFlowPackageGetV2Filtering struct {
	//
	FlowPackageIds  []int64                                             `json:"flow_package_ids,omitempty"`
	FlowPackageType *ToolsUnionFlowPackageGetV2FilteringFlowPackageType `json:"flow_package_type,omitempty"`
}
