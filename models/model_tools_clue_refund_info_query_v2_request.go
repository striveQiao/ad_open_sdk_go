/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueRefundInfoQueryV2Request struct for ToolsClueRefundInfoQueryV2Request
type ToolsClueRefundInfoQueryV2Request struct {
	// 广告主ID
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// 线索ID
	ClueIds []int64 `json:"clue_ids"`
}
