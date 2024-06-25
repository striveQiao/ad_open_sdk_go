/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// StardeliveryTaskAuthorVideoDetailV30DataValidAuthorSubmittedVideosStarVideoStatus
type StardeliveryTaskAuthorVideoDetailV30DataValidAuthorSubmittedVideosStarVideoStatus string

// List of stardelivery_task_author_video_detail_v3.0_data_valid_author_submitted_videos_star_video_status
const (
	INVISIBLE_StardeliveryTaskAuthorVideoDetailV30DataValidAuthorSubmittedVideosStarVideoStatus StardeliveryTaskAuthorVideoDetailV30DataValidAuthorSubmittedVideosStarVideoStatus = "INVISIBLE"
	VISIBLE_StardeliveryTaskAuthorVideoDetailV30DataValidAuthorSubmittedVideosStarVideoStatus   StardeliveryTaskAuthorVideoDetailV30DataValidAuthorSubmittedVideosStarVideoStatus = "VISIBLE"
)

// All allowed values of StardeliveryTaskAuthorVideoDetailV30DataValidAuthorSubmittedVideosStarVideoStatus enum
var AllowedStardeliveryTaskAuthorVideoDetailV30DataValidAuthorSubmittedVideosStarVideoStatusEnumValues = []StardeliveryTaskAuthorVideoDetailV30DataValidAuthorSubmittedVideosStarVideoStatus{
	"INVISIBLE",
	"VISIBLE",
}

// NewStardeliveryTaskAuthorVideoDetailV30DataValidAuthorSubmittedVideosStarVideoStatusFromValue returns a pointer to a valid StardeliveryTaskAuthorVideoDetailV30DataValidAuthorSubmittedVideosStarVideoStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStardeliveryTaskAuthorVideoDetailV30DataValidAuthorSubmittedVideosStarVideoStatusFromValue(v string) (*StardeliveryTaskAuthorVideoDetailV30DataValidAuthorSubmittedVideosStarVideoStatus, error) {
	ev := StardeliveryTaskAuthorVideoDetailV30DataValidAuthorSubmittedVideosStarVideoStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StardeliveryTaskAuthorVideoDetailV30DataValidAuthorSubmittedVideosStarVideoStatus: valid values are %v", v, AllowedStardeliveryTaskAuthorVideoDetailV30DataValidAuthorSubmittedVideosStarVideoStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StardeliveryTaskAuthorVideoDetailV30DataValidAuthorSubmittedVideosStarVideoStatus) IsValid() bool {
	for _, existing := range AllowedStardeliveryTaskAuthorVideoDetailV30DataValidAuthorSubmittedVideosStarVideoStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to stardelivery_task_author_video_detail_v3.0_data_valid_author_submitted_videos_star_video_status value
func (v StardeliveryTaskAuthorVideoDetailV30DataValidAuthorSubmittedVideosStarVideoStatus) Ptr() *StardeliveryTaskAuthorVideoDetailV30DataValidAuthorSubmittedVideosStarVideoStatus {
	return &v
}
