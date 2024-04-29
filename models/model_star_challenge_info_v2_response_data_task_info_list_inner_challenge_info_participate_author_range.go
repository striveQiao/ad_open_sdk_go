/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarChallengeInfoV2ResponseDataTaskInfoListInnerChallengeInfoParticipateAuthorRange
type StarChallengeInfoV2ResponseDataTaskInfoListInnerChallengeInfoParticipateAuthorRange struct {
	//
	AuthorUids       []int64                                                                                              `json:"author_uids,omitempty"`
	AuthorWatcherTag *StarChallengeInfoV2ResponseDataTaskInfoListInnerChallengeInfoParticipateAuthorRangeAuthorWatcherTag `json:"author_watcher_tag,omitempty"`
	//
	ContentTags []int64 `json:"content_tags,omitempty"`
	//
	Gender []string `json:"gender,omitempty"`
	//
	MaxFollower *int64 `json:"max_follower,omitempty"`
	//
	MinFollower *int64 `json:"min_follower,omitempty"`
	//
	Provinces []string `json:"provinces,omitempty"`
}
