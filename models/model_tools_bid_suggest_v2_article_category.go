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

// ToolsBidSuggestV2ArticleCategory
type ToolsBidSuggestV2ArticleCategory string

// List of tools_bid_suggest_v2_article_category
const (
	CARS_ToolsBidSuggestV2ArticleCategory          ToolsBidSuggestV2ArticleCategory = "CARS"
	WORKPLACE_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "WORKPLACE"
	DIGITAL_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "DIGITAL"
	CONSTELLATION_ToolsBidSuggestV2ArticleCategory ToolsBidSuggestV2ArticleCategory = "CONSTELLATION"
	CULTURE_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "CULTURE"
	LOCAL_ToolsBidSuggestV2ArticleCategory         ToolsBidSuggestV2ArticleCategory = "LOCAL"
	LAWS_ToolsBidSuggestV2ArticleCategory          ToolsBidSuggestV2ArticleCategory = "LAWS"
	REGIMEN_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "REGIMEN"
	MOVIE_ToolsBidSuggestV2ArticleCategory         ToolsBidSuggestV2ArticleCategory = "MOVIE"
	SOCIETY_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "SOCIETY"
	PHOTOGRAPHY_ToolsBidSuggestV2ArticleCategory   ToolsBidSuggestV2ArticleCategory = "PHOTOGRAPHY"
	COLLECTION_ToolsBidSuggestV2ArticleCategory    ToolsBidSuggestV2ArticleCategory = "COLLECTION"
	DESIGN_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "DESIGN"
	GRADUATES_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "GRADUATES"
	SPORTS_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "SPORTS"
	HISTORY_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "HISTORY"
	GOVERNMENT_ToolsBidSuggestV2ArticleCategory    ToolsBidSuggestV2ArticleCategory = "GOVERNMENT"
	COMICS_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "COMICS"
	HOME_ToolsBidSuggestV2ArticleCategory          ToolsBidSuggestV2ArticleCategory = "HOME"
	WEIGHT_LOSING_ToolsBidSuggestV2ArticleCategory ToolsBidSuggestV2ArticleCategory = "WEIGHT_LOSING"
	FINANCE_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "FINANCE"
	SCIENCE_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "SCIENCE"
	EXPLORE_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "EXPLORE"
	HEALTH_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "HEALTH"
	EDUCATION_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "EDUCATION"
	FASHION_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "FASHION"
	INTERNATIONAL_ToolsBidSuggestV2ArticleCategory ToolsBidSuggestV2ArticleCategory = "INTERNATIONAL"
	ESTATE_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "ESTATE"
	GAMES_ToolsBidSuggestV2ArticleCategory         ToolsBidSuggestV2ArticleCategory = "GAMES"
	RUMOR_CRACKER_ToolsBidSuggestV2ArticleCategory ToolsBidSuggestV2ArticleCategory = "RUMOR_CRACKER"
	TRAVEL_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "TRAVEL"
	MILITARY_ToolsBidSuggestV2ArticleCategory      ToolsBidSuggestV2ArticleCategory = "MILITARY"
	PETS_ToolsBidSuggestV2ArticleCategory          ToolsBidSuggestV2ArticleCategory = "PETS"
	FREAK_ToolsBidSuggestV2ArticleCategory         ToolsBidSuggestV2ArticleCategory = "FREAK"
	ESSAY_ToolsBidSuggestV2ArticleCategory         ToolsBidSuggestV2ArticleCategory = "ESSAY"
	BUSINESS_ToolsBidSuggestV2ArticleCategory      ToolsBidSuggestV2ArticleCategory = "BUSINESS"
	GOURMET_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "GOURMET"
	TIPS_ToolsBidSuggestV2ArticleCategory          ToolsBidSuggestV2ArticleCategory = "TIPS"
	PARENTING_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "PARENTING"
	ANIMATION_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "ANIMATION"
	ENTERTAINMENT_ToolsBidSuggestV2ArticleCategory ToolsBidSuggestV2ArticleCategory = "ENTERTAINMENT"
	EMOTION_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "EMOTION"
	TECHNOLOGY_ToolsBidSuggestV2ArticleCategory    ToolsBidSuggestV2ArticleCategory = "TECHNOLOGY"
	AMUSEMENT_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "AMUSEMENT"
	STORIES_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "STORIES"
)

// All allowed values of ToolsBidSuggestV2ArticleCategory enum
var AllowedToolsBidSuggestV2ArticleCategoryEnumValues = []ToolsBidSuggestV2ArticleCategory{
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

// NewToolsBidSuggestV2ArticleCategoryFromValue returns a pointer to a valid ToolsBidSuggestV2ArticleCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2ArticleCategoryFromValue(v string) (*ToolsBidSuggestV2ArticleCategory, error) {
	ev := ToolsBidSuggestV2ArticleCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2ArticleCategory: valid values are %v", v, AllowedToolsBidSuggestV2ArticleCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2ArticleCategory) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2ArticleCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_article_category value
func (v ToolsBidSuggestV2ArticleCategory) Ptr() *ToolsBidSuggestV2ArticleCategory {
	return &v
}
