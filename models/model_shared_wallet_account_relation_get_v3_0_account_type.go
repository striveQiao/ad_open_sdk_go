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

// SharedWalletAccountRelationGetV30AccountType
type SharedWalletAccountRelationGetV30AccountType string

// List of shared_wallet_account_relation_get_v3.0_account_type
const (
	AD_SharedWalletAccountRelationGetV30AccountType        SharedWalletAccountRelationGetV30AccountType = "AD"
	AGENT_SharedWalletAccountRelationGetV30AccountType     SharedWalletAccountRelationGetV30AccountType = "AGENT"
	LOCAL_SharedWalletAccountRelationGetV30AccountType     SharedWalletAccountRelationGetV30AccountType = "LOCAL"
	QIANCHUAN_SharedWalletAccountRelationGetV30AccountType SharedWalletAccountRelationGetV30AccountType = "QIANCHUAN"
)

// All allowed values of SharedWalletAccountRelationGetV30AccountType enum
var AllowedSharedWalletAccountRelationGetV30AccountTypeEnumValues = []SharedWalletAccountRelationGetV30AccountType{
	"AD",
	"AGENT",
	"LOCAL",
	"QIANCHUAN",
}

// NewSharedWalletAccountRelationGetV30AccountTypeFromValue returns a pointer to a valid SharedWalletAccountRelationGetV30AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSharedWalletAccountRelationGetV30AccountTypeFromValue(v string) (*SharedWalletAccountRelationGetV30AccountType, error) {
	ev := SharedWalletAccountRelationGetV30AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SharedWalletAccountRelationGetV30AccountType: valid values are %v", v, AllowedSharedWalletAccountRelationGetV30AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SharedWalletAccountRelationGetV30AccountType) IsValid() bool {
	for _, existing := range AllowedSharedWalletAccountRelationGetV30AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to shared_wallet_account_relation_get_v3.0_account_type value
func (v SharedWalletAccountRelationGetV30AccountType) Ptr() *SharedWalletAccountRelationGetV30AccountType {
	return &v
}
