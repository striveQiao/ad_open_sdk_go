/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentChildAgentSelectV2ResponseData
type AgentChildAgentSelectV2ResponseData struct {
	//
	ChildAgentIds []int64 `json:"child_agent_ids,omitempty"`
	//
	List     []int64                                      `json:"list,omitempty"`
	PageInfo *AgentChildAgentSelectV2ResponseDataPageInfo `json:"page_info,omitempty"`
}
