/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AgentTransferTransactionRecordV2FilteringRemitterType
type AgentTransferTransactionRecordV2FilteringRemitterType string

// List of agent_transfer_transaction_record_v2_filtering_remitter_type
const (
	ROLE_ADVERTISER_AgentTransferTransactionRecordV2FilteringRemitterType                    AgentTransferTransactionRecordV2FilteringRemitterType = "ROLE_ADVERTISER"
	ROLE_AGENT_AgentTransferTransactionRecordV2FilteringRemitterType                         AgentTransferTransactionRecordV2FilteringRemitterType = "ROLE_AGENT"
	ROLE_CHILD_AGENT_AgentTransferTransactionRecordV2FilteringRemitterType                   AgentTransferTransactionRecordV2FilteringRemitterType = "ROLE_CHILD_AGENT"
	ROLE_ECP_VIRTUAL_ADVERTISER_AgentTransferTransactionRecordV2FilteringRemitterType        AgentTransferTransactionRecordV2FilteringRemitterType = "ROLE_ECP_VIRTUAL_ADVERTISER"
	ROLE_LOCAL_LIFE_VIRTUAL_ADVERTISER_AgentTransferTransactionRecordV2FilteringRemitterType AgentTransferTransactionRecordV2FilteringRemitterType = "ROLE_LOCAL_LIFE_VIRTUAL_ADVERTISER"
	ROLE_VIRTAUL_ADVERTISER_AgentTransferTransactionRecordV2FilteringRemitterType            AgentTransferTransactionRecordV2FilteringRemitterType = "ROLE_VIRTAUL_ADVERTISER"
)

// All allowed values of AgentTransferTransactionRecordV2FilteringRemitterType enum
var AllowedAgentTransferTransactionRecordV2FilteringRemitterTypeEnumValues = []AgentTransferTransactionRecordV2FilteringRemitterType{
	"ROLE_ADVERTISER",
	"ROLE_AGENT",
	"ROLE_CHILD_AGENT",
	"ROLE_ECP_VIRTUAL_ADVERTISER",
	"ROLE_LOCAL_LIFE_VIRTUAL_ADVERTISER",
	"ROLE_VIRTAUL_ADVERTISER",
}

// NewAgentTransferTransactionRecordV2FilteringRemitterTypeFromValue returns a pointer to a valid AgentTransferTransactionRecordV2FilteringRemitterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentTransferTransactionRecordV2FilteringRemitterTypeFromValue(v string) (*AgentTransferTransactionRecordV2FilteringRemitterType, error) {
	ev := AgentTransferTransactionRecordV2FilteringRemitterType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentTransferTransactionRecordV2FilteringRemitterType: valid values are %v", v, AllowedAgentTransferTransactionRecordV2FilteringRemitterTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentTransferTransactionRecordV2FilteringRemitterType) IsValid() bool {
	for _, existing := range AllowedAgentTransferTransactionRecordV2FilteringRemitterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_transfer_transaction_record_v2_filtering_remitter_type value
func (v AgentTransferTransactionRecordV2FilteringRemitterType) Ptr() *AgentTransferTransactionRecordV2FilteringRemitterType {
	return &v
}
