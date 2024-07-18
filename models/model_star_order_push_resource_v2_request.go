/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderPushResourceV2Request struct for StarOrderPushResourceV2Request
type StarOrderPushResourceV2Request struct {
	// 任务ID
	OrderId int64 `json:"order_id"`
	// 推送账户ID
	PushAccountIds []int64 `json:"push_account_ids"`
	// 推送目标平台 仅限以下 (1)：巨量引擎
	PushPlatform int64 `json:"push_platform"`
	// 星图客户ID
	StarId int64 `json:"star_id"`
}
