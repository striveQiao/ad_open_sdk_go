/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeProceduralCreativeUpdateV2RequestAdDataSupplementsInner struct for CreativeProceduralCreativeUpdateV2RequestAdDataSupplementsInner
type CreativeProceduralCreativeUpdateV2RequestAdDataSupplementsInner struct {
	//
	Games          []*CreativeProceduralCreativeUpdateV2RequestAdDataSupplementsInnerGamesInner `json:"games,omitempty"`
	SupplementType *CreativeProceduralCreativeUpdateV2AdDataSupplementsSupplementType           `json:"supplement_type,omitempty"`
}
