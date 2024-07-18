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

// DiagnosisTaskAgentCreateV2DiagnoseConfigPricingType
type DiagnosisTaskAgentCreateV2DiagnoseConfigPricingType string

// List of diagnosis_task_agent_create_v2_diagnose_config_pricing_type
const (
	OCPC_DiagnosisTaskAgentCreateV2DiagnoseConfigPricingType DiagnosisTaskAgentCreateV2DiagnoseConfigPricingType = "OCPC"
	CPA_DiagnosisTaskAgentCreateV2DiagnoseConfigPricingType  DiagnosisTaskAgentCreateV2DiagnoseConfigPricingType = "CPA"
	OCPM_DiagnosisTaskAgentCreateV2DiagnoseConfigPricingType DiagnosisTaskAgentCreateV2DiagnoseConfigPricingType = "OCPM"
)

// All allowed values of DiagnosisTaskAgentCreateV2DiagnoseConfigPricingType enum
var AllowedDiagnosisTaskAgentCreateV2DiagnoseConfigPricingTypeEnumValues = []DiagnosisTaskAgentCreateV2DiagnoseConfigPricingType{
	"OCPC",
	"CPA",
	"OCPM",
}

// NewDiagnosisTaskAgentCreateV2DiagnoseConfigPricingTypeFromValue returns a pointer to a valid DiagnosisTaskAgentCreateV2DiagnoseConfigPricingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDiagnosisTaskAgentCreateV2DiagnoseConfigPricingTypeFromValue(v string) (*DiagnosisTaskAgentCreateV2DiagnoseConfigPricingType, error) {
	ev := DiagnosisTaskAgentCreateV2DiagnoseConfigPricingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DiagnosisTaskAgentCreateV2DiagnoseConfigPricingType: valid values are %v", v, AllowedDiagnosisTaskAgentCreateV2DiagnoseConfigPricingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DiagnosisTaskAgentCreateV2DiagnoseConfigPricingType) IsValid() bool {
	for _, existing := range AllowedDiagnosisTaskAgentCreateV2DiagnoseConfigPricingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to diagnosis_task_agent_create_v2_diagnose_config_pricing_type value
func (v DiagnosisTaskAgentCreateV2DiagnoseConfigPricingType) Ptr() *DiagnosisTaskAgentCreateV2DiagnoseConfigPricingType {
	return &v
}
