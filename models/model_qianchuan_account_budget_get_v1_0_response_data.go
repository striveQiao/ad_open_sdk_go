/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAccountBudgetGetV10ResponseData
type QianchuanAccountBudgetGetV10ResponseData struct {
	//
	Budget     *float64                                    `json:"budget,omitempty"`
	BudgetMode *QianchuanAccountBudgetGetV10DataBudgetMode `json:"budget_mode,omitempty"`
}
