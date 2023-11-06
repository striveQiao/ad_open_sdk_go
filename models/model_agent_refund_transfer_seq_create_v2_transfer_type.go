/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AgentRefundTransferSeqCreateV2TransferType
type AgentRefundTransferSeqCreateV2TransferType string

// List of agent_refund_transfer_seq_create_v2_transfer_type
const (
	CASH_DEFAULT_AgentRefundTransferSeqCreateV2TransferType   AgentRefundTransferSeqCreateV2TransferType = "CASH_DEFAULT"
	CREDIT_BIDDING_AgentRefundTransferSeqCreateV2TransferType AgentRefundTransferSeqCreateV2TransferType = "CREDIT_BIDDING"
	CREDIT_BRAND_AgentRefundTransferSeqCreateV2TransferType   AgentRefundTransferSeqCreateV2TransferType = "CREDIT_BRAND"
	CREDIT_GENERAL_AgentRefundTransferSeqCreateV2TransferType AgentRefundTransferSeqCreateV2TransferType = "CREDIT_GENERAL"
	GRANT_GENERAL_AgentRefundTransferSeqCreateV2TransferType  AgentRefundTransferSeqCreateV2TransferType = "GRANT_GENERAL"
	PREPAY_BIDDING_AgentRefundTransferSeqCreateV2TransferType AgentRefundTransferSeqCreateV2TransferType = "PREPAY_BIDDING"
	PREPAY_BRAND_AgentRefundTransferSeqCreateV2TransferType   AgentRefundTransferSeqCreateV2TransferType = "PREPAY_BRAND"
	PREPAY_GENERAL_AgentRefundTransferSeqCreateV2TransferType AgentRefundTransferSeqCreateV2TransferType = "PREPAY_GENERAL"
)

// All allowed values of AgentRefundTransferSeqCreateV2TransferType enum
var AllowedAgentRefundTransferSeqCreateV2TransferTypeEnumValues = []AgentRefundTransferSeqCreateV2TransferType{
	"CASH_DEFAULT",
	"CREDIT_BIDDING",
	"CREDIT_BRAND",
	"CREDIT_GENERAL",
	"GRANT_GENERAL",
	"PREPAY_BIDDING",
	"PREPAY_BRAND",
	"PREPAY_GENERAL",
}

// NewAgentRefundTransferSeqCreateV2TransferTypeFromValue returns a pointer to a valid AgentRefundTransferSeqCreateV2TransferType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentRefundTransferSeqCreateV2TransferTypeFromValue(v string) (*AgentRefundTransferSeqCreateV2TransferType, error) {
	ev := AgentRefundTransferSeqCreateV2TransferType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentRefundTransferSeqCreateV2TransferType: valid values are %v", v, AllowedAgentRefundTransferSeqCreateV2TransferTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentRefundTransferSeqCreateV2TransferType) IsValid() bool {
	for _, existing := range AllowedAgentRefundTransferSeqCreateV2TransferTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_refund_transfer_seq_create_v2_transfer_type value
func (v AgentRefundTransferSeqCreateV2TransferType) Ptr() *AgentRefundTransferSeqCreateV2TransferType {
	return &v
}
