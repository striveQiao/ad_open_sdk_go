/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
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
	DESIGN_ToolsEstimateAudienceV2ArticleCategory        ToolsEstimateAudienceV2ArticleCategory = "DESIGN"
	ESSAY_ToolsEstimateAudienceV2ArticleCategory         ToolsEstimateAudienceV2ArticleCategory = "ESSAY"
	EDUCATION_ToolsEstimateAudienceV2ArticleCategory     ToolsEstimateAudienceV2ArticleCategory = "EDUCATION"
	FREAK_ToolsEstimateAudienceV2ArticleCategory         ToolsEstimateAudienceV2ArticleCategory = "FREAK"
	CARS_ToolsEstimateAudienceV2ArticleCategory          ToolsEstimateAudienceV2ArticleCategory = "CARS"
	MOVIE_ToolsEstimateAudienceV2ArticleCategory         ToolsEstimateAudienceV2ArticleCategory = "MOVIE"
	HISTORY_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "HISTORY"
	CONSTELLATION_ToolsEstimateAudienceV2ArticleCategory ToolsEstimateAudienceV2ArticleCategory = "CONSTELLATION"
	TECHNOLOGY_ToolsEstimateAudienceV2ArticleCategory    ToolsEstimateAudienceV2ArticleCategory = "TECHNOLOGY"
	GAMES_ToolsEstimateAudienceV2ArticleCategory         ToolsEstimateAudienceV2ArticleCategory = "GAMES"
	HOME_ToolsEstimateAudienceV2ArticleCategory          ToolsEstimateAudienceV2ArticleCategory = "HOME"
	ANIMATION_ToolsEstimateAudienceV2ArticleCategory     ToolsEstimateAudienceV2ArticleCategory = "ANIMATION"
	COLLECTION_ToolsEstimateAudienceV2ArticleCategory    ToolsEstimateAudienceV2ArticleCategory = "COLLECTION"
	BUSINESS_ToolsEstimateAudienceV2ArticleCategory      ToolsEstimateAudienceV2ArticleCategory = "BUSINESS"
	WORKPLACE_ToolsEstimateAudienceV2ArticleCategory     ToolsEstimateAudienceV2ArticleCategory = "WORKPLACE"
	ENTERTAINMENT_ToolsEstimateAudienceV2ArticleCategory ToolsEstimateAudienceV2ArticleCategory = "ENTERTAINMENT"
	HEALTH_ToolsEstimateAudienceV2ArticleCategory        ToolsEstimateAudienceV2ArticleCategory = "HEALTH"
	EXPLORE_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "EXPLORE"
	PARENTING_ToolsEstimateAudienceV2ArticleCategory     ToolsEstimateAudienceV2ArticleCategory = "PARENTING"
	STORIES_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "STORIES"
	RUMOR_CRACKER_ToolsEstimateAudienceV2ArticleCategory ToolsEstimateAudienceV2ArticleCategory = "RUMOR_CRACKER"
	AMUSEMENT_ToolsEstimateAudienceV2ArticleCategory     ToolsEstimateAudienceV2ArticleCategory = "AMUSEMENT"
	TIPS_ToolsEstimateAudienceV2ArticleCategory          ToolsEstimateAudienceV2ArticleCategory = "TIPS"
	LOCAL_ToolsEstimateAudienceV2ArticleCategory         ToolsEstimateAudienceV2ArticleCategory = "LOCAL"
	CULTURE_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "CULTURE"
	GOURMET_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "GOURMET"
	COMICS_ToolsEstimateAudienceV2ArticleCategory        ToolsEstimateAudienceV2ArticleCategory = "COMICS"
	INTERNATIONAL_ToolsEstimateAudienceV2ArticleCategory ToolsEstimateAudienceV2ArticleCategory = "INTERNATIONAL"
	SOCIETY_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "SOCIETY"
	EMOTION_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "EMOTION"
	REGIMEN_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "REGIMEN"
	ESTATE_ToolsEstimateAudienceV2ArticleCategory        ToolsEstimateAudienceV2ArticleCategory = "ESTATE"
	FASHION_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "FASHION"
	GOVERNMENT_ToolsEstimateAudienceV2ArticleCategory    ToolsEstimateAudienceV2ArticleCategory = "GOVERNMENT"
	LAWS_ToolsEstimateAudienceV2ArticleCategory          ToolsEstimateAudienceV2ArticleCategory = "LAWS"
	SPORTS_ToolsEstimateAudienceV2ArticleCategory        ToolsEstimateAudienceV2ArticleCategory = "SPORTS"
	FINANCE_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "FINANCE"
	TRAVEL_ToolsEstimateAudienceV2ArticleCategory        ToolsEstimateAudienceV2ArticleCategory = "TRAVEL"
	PHOTOGRAPHY_ToolsEstimateAudienceV2ArticleCategory   ToolsEstimateAudienceV2ArticleCategory = "PHOTOGRAPHY"
	DIGITAL_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "DIGITAL"
	PETS_ToolsEstimateAudienceV2ArticleCategory          ToolsEstimateAudienceV2ArticleCategory = "PETS"
	WEIGHT_LOSING_ToolsEstimateAudienceV2ArticleCategory ToolsEstimateAudienceV2ArticleCategory = "WEIGHT_LOSING"
	SCIENCE_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "SCIENCE"
	GRADUATES_ToolsEstimateAudienceV2ArticleCategory     ToolsEstimateAudienceV2ArticleCategory = "GRADUATES"
	MILITARY_ToolsEstimateAudienceV2ArticleCategory      ToolsEstimateAudienceV2ArticleCategory = "MILITARY"
)

// All allowed values of ToolsEstimateAudienceV2ArticleCategory enum
var AllowedToolsEstimateAudienceV2ArticleCategoryEnumValues = []ToolsEstimateAudienceV2ArticleCategory{
	"DESIGN",
	"ESSAY",
	"EDUCATION",
	"FREAK",
	"CARS",
	"MOVIE",
	"HISTORY",
	"CONSTELLATION",
	"TECHNOLOGY",
	"GAMES",
	"HOME",
	"ANIMATION",
	"COLLECTION",
	"BUSINESS",
	"WORKPLACE",
	"ENTERTAINMENT",
	"HEALTH",
	"EXPLORE",
	"PARENTING",
	"STORIES",
	"RUMOR_CRACKER",
	"AMUSEMENT",
	"TIPS",
	"LOCAL",
	"CULTURE",
	"GOURMET",
	"COMICS",
	"INTERNATIONAL",
	"SOCIETY",
	"EMOTION",
	"REGIMEN",
	"ESTATE",
	"FASHION",
	"GOVERNMENT",
	"LAWS",
	"SPORTS",
	"FINANCE",
	"TRAVEL",
	"PHOTOGRAPHY",
	"DIGITAL",
	"PETS",
	"WEIGHT_LOSING",
	"SCIENCE",
	"GRADUATES",
	"MILITARY",
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
