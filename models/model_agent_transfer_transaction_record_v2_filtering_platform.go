/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AgentTransferTransactionRecordV2FilteringPlatform
type AgentTransferTransactionRecordV2FilteringPlatform string

// List of agent_transfer_transaction_record_v2_filtering_platform
const (
	AD_AgentTransferTransactionRecordV2FilteringPlatform                    AgentTransferTransactionRecordV2FilteringPlatform = "AD"
	EFFECT_AgentTransferTransactionRecordV2FilteringPlatform                AgentTransferTransactionRecordV2FilteringPlatform = "EFFECT"
	STAR_AgentTransferTransactionRecordV2FilteringPlatform                  AgentTransferTransactionRecordV2FilteringPlatform = "STAR"
	TRUSTEESHIP_MARKETING_AgentTransferTransactionRecordV2FilteringPlatform AgentTransferTransactionRecordV2FilteringPlatform = "TRUSTEESHIP_MARKETING"
	ZTT_AgentTransferTransactionRecordV2FilteringPlatform                   AgentTransferTransactionRecordV2FilteringPlatform = "ZTT"
)

// All allowed values of AgentTransferTransactionRecordV2FilteringPlatform enum
var AllowedAgentTransferTransactionRecordV2FilteringPlatformEnumValues = []AgentTransferTransactionRecordV2FilteringPlatform{
	"AD",
	"EFFECT",
	"STAR",
	"TRUSTEESHIP_MARKETING",
	"ZTT",
}

// NewAgentTransferTransactionRecordV2FilteringPlatformFromValue returns a pointer to a valid AgentTransferTransactionRecordV2FilteringPlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentTransferTransactionRecordV2FilteringPlatformFromValue(v string) (*AgentTransferTransactionRecordV2FilteringPlatform, error) {
	ev := AgentTransferTransactionRecordV2FilteringPlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentTransferTransactionRecordV2FilteringPlatform: valid values are %v", v, AllowedAgentTransferTransactionRecordV2FilteringPlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentTransferTransactionRecordV2FilteringPlatform) IsValid() bool {
	for _, existing := range AllowedAgentTransferTransactionRecordV2FilteringPlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_transfer_transaction_record_v2_filtering_platform value
func (v AgentTransferTransactionRecordV2FilteringPlatform) Ptr() *AgentTransferTransactionRecordV2FilteringPlatform {
	return &v
}
