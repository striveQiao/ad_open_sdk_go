/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsLandingGroupSiteOptStatusUpdateV2RequestDataInner struct for ToolsLandingGroupSiteOptStatusUpdateV2RequestDataInner
type ToolsLandingGroupSiteOptStatusUpdateV2RequestDataInner struct {
	// 成员 ID，即站点在落地页组中的唯一标识
	MemberId int64 `json:"member_id"`
	// 站点启用状态，详见【附录-枚举值】，可选值: SITE_OPT_STATUS_DISABLE,SITE_OPT_STATUS_ENABLE
	SiteOptStatus string `json:"site_opt_status"`
}
