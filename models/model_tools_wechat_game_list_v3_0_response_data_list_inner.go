/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsWechatGameListV30ResponseDataListInner struct for ToolsWechatGameListV30ResponseDataListInner
type ToolsWechatGameListV30ResponseDataListInner struct {
	//
	AccountId   *int64                                     `json:"account_id,omitempty"`
	AccountType *ToolsWechatGameListV30DataListAccountType `json:"account_type,omitempty"`
	//
	AgeRemindUrl *string `json:"age_remind_url,omitempty"`
	//
	AntiAddictionUrl    *string                                            `json:"anti_addiction_url,omitempty"`
	AuditStatus         *ToolsWechatGameListV30DataListAuditStatus         `json:"audit_status,omitempty"`
	AuthorizationStatus *ToolsWechatGameListV30DataListAuthorizationStatus `json:"authorization_status,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	InstanceId *int64 `json:"instance_id,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	Path *string `json:"path,omitempty"`
	//
	RealNameUrl *string `json:"real_name_url,omitempty"`
	//
	Reason *string `json:"reason,omitempty"`
	//
	ReasonUnauthorize *string `json:"reason_unauthorize,omitempty"`
	//
	ScreenRecordUrl *string `json:"screen_record_url,omitempty"`
	//
	UserName *string `json:"user_name,omitempty"`
}
