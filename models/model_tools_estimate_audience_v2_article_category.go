/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsEstimateAudienceV2ArticleCategory
type ToolsEstimateAudienceV2ArticleCategory string

// List of tools_estimate_audience_v2_article_category
const (
	CARS_ToolsEstimateAudienceV2ArticleCategory          ToolsEstimateAudienceV2ArticleCategory = "CARS"
	WORKPLACE_ToolsEstimateAudienceV2ArticleCategory     ToolsEstimateAudienceV2ArticleCategory = "WORKPLACE"
	DIGITAL_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "DIGITAL"
	CONSTELLATION_ToolsEstimateAudienceV2ArticleCategory ToolsEstimateAudienceV2ArticleCategory = "CONSTELLATION"
	CULTURE_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "CULTURE"
	LOCAL_ToolsEstimateAudienceV2ArticleCategory         ToolsEstimateAudienceV2ArticleCategory = "LOCAL"
	LAWS_ToolsEstimateAudienceV2ArticleCategory          ToolsEstimateAudienceV2ArticleCategory = "LAWS"
	REGIMEN_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "REGIMEN"
	MOVIE_ToolsEstimateAudienceV2ArticleCategory         ToolsEstimateAudienceV2ArticleCategory = "MOVIE"
	SOCIETY_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "SOCIETY"
	PHOTOGRAPHY_ToolsEstimateAudienceV2ArticleCategory   ToolsEstimateAudienceV2ArticleCategory = "PHOTOGRAPHY"
	COLLECTION_ToolsEstimateAudienceV2ArticleCategory    ToolsEstimateAudienceV2ArticleCategory = "COLLECTION"
	DESIGN_ToolsEstimateAudienceV2ArticleCategory        ToolsEstimateAudienceV2ArticleCategory = "DESIGN"
	GRADUATES_ToolsEstimateAudienceV2ArticleCategory     ToolsEstimateAudienceV2ArticleCategory = "GRADUATES"
	SPORTS_ToolsEstimateAudienceV2ArticleCategory        ToolsEstimateAudienceV2ArticleCategory = "SPORTS"
	HISTORY_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "HISTORY"
	GOVERNMENT_ToolsEstimateAudienceV2ArticleCategory    ToolsEstimateAudienceV2ArticleCategory = "GOVERNMENT"
	COMICS_ToolsEstimateAudienceV2ArticleCategory        ToolsEstimateAudienceV2ArticleCategory = "COMICS"
	HOME_ToolsEstimateAudienceV2ArticleCategory          ToolsEstimateAudienceV2ArticleCategory = "HOME"
	WEIGHT_LOSING_ToolsEstimateAudienceV2ArticleCategory ToolsEstimateAudienceV2ArticleCategory = "WEIGHT_LOSING"
	FINANCE_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "FINANCE"
	SCIENCE_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "SCIENCE"
	EXPLORE_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "EXPLORE"
	HEALTH_ToolsEstimateAudienceV2ArticleCategory        ToolsEstimateAudienceV2ArticleCategory = "HEALTH"
	EDUCATION_ToolsEstimateAudienceV2ArticleCategory     ToolsEstimateAudienceV2ArticleCategory = "EDUCATION"
	FASHION_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "FASHION"
	INTERNATIONAL_ToolsEstimateAudienceV2ArticleCategory ToolsEstimateAudienceV2ArticleCategory = "INTERNATIONAL"
	ESTATE_ToolsEstimateAudienceV2ArticleCategory        ToolsEstimateAudienceV2ArticleCategory = "ESTATE"
	GAMES_ToolsEstimateAudienceV2ArticleCategory         ToolsEstimateAudienceV2ArticleCategory = "GAMES"
	RUMOR_CRACKER_ToolsEstimateAudienceV2ArticleCategory ToolsEstimateAudienceV2ArticleCategory = "RUMOR_CRACKER"
	TRAVEL_ToolsEstimateAudienceV2ArticleCategory        ToolsEstimateAudienceV2ArticleCategory = "TRAVEL"
	MILITARY_ToolsEstimateAudienceV2ArticleCategory      ToolsEstimateAudienceV2ArticleCategory = "MILITARY"
	PETS_ToolsEstimateAudienceV2ArticleCategory          ToolsEstimateAudienceV2ArticleCategory = "PETS"
	FREAK_ToolsEstimateAudienceV2ArticleCategory         ToolsEstimateAudienceV2ArticleCategory = "FREAK"
	ESSAY_ToolsEstimateAudienceV2ArticleCategory         ToolsEstimateAudienceV2ArticleCategory = "ESSAY"
	BUSINESS_ToolsEstimateAudienceV2ArticleCategory      ToolsEstimateAudienceV2ArticleCategory = "BUSINESS"
	GOURMET_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "GOURMET"
	TIPS_ToolsEstimateAudienceV2ArticleCategory          ToolsEstimateAudienceV2ArticleCategory = "TIPS"
	PARENTING_ToolsEstimateAudienceV2ArticleCategory     ToolsEstimateAudienceV2ArticleCategory = "PARENTING"
	ANIMATION_ToolsEstimateAudienceV2ArticleCategory     ToolsEstimateAudienceV2ArticleCategory = "ANIMATION"
	ENTERTAINMENT_ToolsEstimateAudienceV2ArticleCategory ToolsEstimateAudienceV2ArticleCategory = "ENTERTAINMENT"
	EMOTION_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "EMOTION"
	TECHNOLOGY_ToolsEstimateAudienceV2ArticleCategory    ToolsEstimateAudienceV2ArticleCategory = "TECHNOLOGY"
	AMUSEMENT_ToolsEstimateAudienceV2ArticleCategory     ToolsEstimateAudienceV2ArticleCategory = "AMUSEMENT"
	STORIES_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "STORIES"
)

// All allowed values of ToolsEstimateAudienceV2ArticleCategory enum
var AllowedToolsEstimateAudienceV2ArticleCategoryEnumValues = []ToolsEstimateAudienceV2ArticleCategory{
	"CARS",
	"WORKPLACE",
	"DIGITAL",
	"CONSTELLATION",
	"CULTURE",
	"LOCAL",
	"LAWS",
	"REGIMEN",
	"MOVIE",
	"SOCIETY",
	"PHOTOGRAPHY",
	"COLLECTION",
	"DESIGN",
	"GRADUATES",
	"SPORTS",
	"HISTORY",
	"GOVERNMENT",
	"COMICS",
	"HOME",
	"WEIGHT_LOSING",
	"FINANCE",
	"SCIENCE",
	"EXPLORE",
	"HEALTH",
	"EDUCATION",
	"FASHION",
	"INTERNATIONAL",
	"ESTATE",
	"GAMES",
	"RUMOR_CRACKER",
	"TRAVEL",
	"MILITARY",
	"PETS",
	"FREAK",
	"ESSAY",
	"BUSINESS",
	"GOURMET",
	"TIPS",
	"PARENTING",
	"ANIMATION",
	"ENTERTAINMENT",
	"EMOTION",
	"TECHNOLOGY",
	"AMUSEMENT",
	"STORIES",
}

// NewToolsEstimateAudienceV2ArticleCategoryFromValue returns a pointer to a valid ToolsEstimateAudienceV2ArticleCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2ArticleCategoryFromValue(v string) (*ToolsEstimateAudienceV2ArticleCategory, error) {
	ev := ToolsEstimateAudienceV2ArticleCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2ArticleCategory: valid values are %v", v, AllowedToolsEstimateAudienceV2ArticleCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2ArticleCategory) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2ArticleCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_article_category value
func (v ToolsEstimateAudienceV2ArticleCategory) Ptr() *ToolsEstimateAudienceV2ArticleCategory {
	return &v
}
