/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectBudgetUpdateV30RequestDataInner struct for ProjectBudgetUpdateV30RequestDataInner
type ProjectBudgetUpdateV30RequestDataInner struct {
	//
	Budget     *float64                             `json:"budget,omitempty"`
	BudgetMode ProjectBudgetUpdateV30DataBudgetMode `json:"budget_mode"`
	//
	ProjectId int64 `json:"project_id"`
}
