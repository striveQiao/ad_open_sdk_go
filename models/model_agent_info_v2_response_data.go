/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentInfoV2ResponseData struct for AgentInfoV2ResponseData
type AgentInfoV2ResponseData struct {
	AccountStatus *AgentInfoV2DataAccountStatus `json:"account_status,omitempty"`
	//
	AgentId *int64 `json:"agent_id,omitempty"`
	//
	AgentName *string `json:"agent_name,omitempty"`
	//
	CompanyId *int64 `json:"company_id,omitempty"`
	//
	CompanyName *string `json:"company_name,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	CustomerId *int64 `json:"customer_id,omitempty"`
	//
	CustomerName *string              `json:"customer_name,omitempty"`
	Role         *AgentInfoV2DataRole `json:"role,omitempty"`
}
