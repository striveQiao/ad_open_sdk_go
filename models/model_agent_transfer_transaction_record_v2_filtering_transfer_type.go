/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AgentTransferTransactionRecordV2FilteringTransferType
type AgentTransferTransactionRecordV2FilteringTransferType string

// List of agent_transfer_transaction_record_v2_filtering_transfer_type
const (
	ADD_MONEY_AgentTransferTransactionRecordV2FilteringTransferType       AgentTransferTransactionRecordV2FilteringTransferType = "ADD_MONEY"
	MUTUAL_TRANSFER_AgentTransferTransactionRecordV2FilteringTransferType AgentTransferTransactionRecordV2FilteringTransferType = "MUTUAL_TRANSFER"
	REFUND_MONEY_AgentTransferTransactionRecordV2FilteringTransferType    AgentTransferTransactionRecordV2FilteringTransferType = "REFUND_MONEY"
)

// All allowed values of AgentTransferTransactionRecordV2FilteringTransferType enum
var AllowedAgentTransferTransactionRecordV2FilteringTransferTypeEnumValues = []AgentTransferTransactionRecordV2FilteringTransferType{
	"ADD_MONEY",
	"MUTUAL_TRANSFER",
	"REFUND_MONEY",
}

// NewAgentTransferTransactionRecordV2FilteringTransferTypeFromValue returns a pointer to a valid AgentTransferTransactionRecordV2FilteringTransferType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentTransferTransactionRecordV2FilteringTransferTypeFromValue(v string) (*AgentTransferTransactionRecordV2FilteringTransferType, error) {
	ev := AgentTransferTransactionRecordV2FilteringTransferType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentTransferTransactionRecordV2FilteringTransferType: valid values are %v", v, AllowedAgentTransferTransactionRecordV2FilteringTransferTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentTransferTransactionRecordV2FilteringTransferType) IsValid() bool {
	for _, existing := range AllowedAgentTransferTransactionRecordV2FilteringTransferTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_transfer_transaction_record_v2_filtering_transfer_type value
func (v AgentTransferTransactionRecordV2FilteringTransferType) Ptr() *AgentTransferTransactionRecordV2FilteringTransferType {
	return &v
}
