/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataAudienceArticleCategory
type AdGetV2DataAudienceArticleCategory string

// List of ad_get_v2_data_audience_article_category
const (
	MILITARY_AdGetV2DataAudienceArticleCategory      AdGetV2DataAudienceArticleCategory = "MILITARY"
	INTERNATIONAL_AdGetV2DataAudienceArticleCategory AdGetV2DataAudienceArticleCategory = "INTERNATIONAL"
	DESIGN_AdGetV2DataAudienceArticleCategory        AdGetV2DataAudienceArticleCategory = "DESIGN"
	EMOTION_AdGetV2DataAudienceArticleCategory       AdGetV2DataAudienceArticleCategory = "EMOTION"
	TIPS_AdGetV2DataAudienceArticleCategory          AdGetV2DataAudienceArticleCategory = "TIPS"
	RUMOR_CRACKER_AdGetV2DataAudienceArticleCategory AdGetV2DataAudienceArticleCategory = "RUMOR_CRACKER"
	HOME_AdGetV2DataAudienceArticleCategory          AdGetV2DataAudienceArticleCategory = "HOME"
	EDUCATION_AdGetV2DataAudienceArticleCategory     AdGetV2DataAudienceArticleCategory = "EDUCATION"
	REGIMEN_AdGetV2DataAudienceArticleCategory       AdGetV2DataAudienceArticleCategory = "REGIMEN"
	WEIGHT_LOSING_AdGetV2DataAudienceArticleCategory AdGetV2DataAudienceArticleCategory = "WEIGHT_LOSING"
	GRADUATES_AdGetV2DataAudienceArticleCategory     AdGetV2DataAudienceArticleCategory = "GRADUATES"
	FINANCE_AdGetV2DataAudienceArticleCategory       AdGetV2DataAudienceArticleCategory = "FINANCE"
	CONSTELLATION_AdGetV2DataAudienceArticleCategory AdGetV2DataAudienceArticleCategory = "CONSTELLATION"
	SPORTS_AdGetV2DataAudienceArticleCategory        AdGetV2DataAudienceArticleCategory = "SPORTS"
	DIGITAL_AdGetV2DataAudienceArticleCategory       AdGetV2DataAudienceArticleCategory = "DIGITAL"
	BUSINESS_AdGetV2DataAudienceArticleCategory      AdGetV2DataAudienceArticleCategory = "BUSINESS"
	PETS_AdGetV2DataAudienceArticleCategory          AdGetV2DataAudienceArticleCategory = "PETS"
	AMUSEMENT_AdGetV2DataAudienceArticleCategory     AdGetV2DataAudienceArticleCategory = "AMUSEMENT"
	STORIES_AdGetV2DataAudienceArticleCategory       AdGetV2DataAudienceArticleCategory = "STORIES"
	HISTORY_AdGetV2DataAudienceArticleCategory       AdGetV2DataAudienceArticleCategory = "HISTORY"
	SCIENCE_AdGetV2DataAudienceArticleCategory       AdGetV2DataAudienceArticleCategory = "SCIENCE"
	COMICS_AdGetV2DataAudienceArticleCategory        AdGetV2DataAudienceArticleCategory = "COMICS"
	TRAVEL_AdGetV2DataAudienceArticleCategory        AdGetV2DataAudienceArticleCategory = "TRAVEL"
	TECHNOLOGY_AdGetV2DataAudienceArticleCategory    AdGetV2DataAudienceArticleCategory = "TECHNOLOGY"
	CULTURE_AdGetV2DataAudienceArticleCategory       AdGetV2DataAudienceArticleCategory = "CULTURE"
	PARENTING_AdGetV2DataAudienceArticleCategory     AdGetV2DataAudienceArticleCategory = "PARENTING"
	ESSAY_AdGetV2DataAudienceArticleCategory         AdGetV2DataAudienceArticleCategory = "ESSAY"
	CARS_AdGetV2DataAudienceArticleCategory          AdGetV2DataAudienceArticleCategory = "CARS"
	FREAK_AdGetV2DataAudienceArticleCategory         AdGetV2DataAudienceArticleCategory = "FREAK"
	PHOTOGRAPHY_AdGetV2DataAudienceArticleCategory   AdGetV2DataAudienceArticleCategory = "PHOTOGRAPHY"
	GOVERNMENT_AdGetV2DataAudienceArticleCategory    AdGetV2DataAudienceArticleCategory = "GOVERNMENT"
	ENTERTAINMENT_AdGetV2DataAudienceArticleCategory AdGetV2DataAudienceArticleCategory = "ENTERTAINMENT"
	GAMES_AdGetV2DataAudienceArticleCategory         AdGetV2DataAudienceArticleCategory = "GAMES"
	EXPLORE_AdGetV2DataAudienceArticleCategory       AdGetV2DataAudienceArticleCategory = "EXPLORE"
	ESTATE_AdGetV2DataAudienceArticleCategory        AdGetV2DataAudienceArticleCategory = "ESTATE"
	SOCIETY_AdGetV2DataAudienceArticleCategory       AdGetV2DataAudienceArticleCategory = "SOCIETY"
	GOURMET_AdGetV2DataAudienceArticleCategory       AdGetV2DataAudienceArticleCategory = "GOURMET"
	FASHION_AdGetV2DataAudienceArticleCategory       AdGetV2DataAudienceArticleCategory = "FASHION"
	LOCAL_AdGetV2DataAudienceArticleCategory         AdGetV2DataAudienceArticleCategory = "LOCAL"
	HEALTH_AdGetV2DataAudienceArticleCategory        AdGetV2DataAudienceArticleCategory = "HEALTH"
	COLLECTION_AdGetV2DataAudienceArticleCategory    AdGetV2DataAudienceArticleCategory = "COLLECTION"
	WORKPLACE_AdGetV2DataAudienceArticleCategory     AdGetV2DataAudienceArticleCategory = "WORKPLACE"
	MOVIE_AdGetV2DataAudienceArticleCategory         AdGetV2DataAudienceArticleCategory = "MOVIE"
	ANIMATION_AdGetV2DataAudienceArticleCategory     AdGetV2DataAudienceArticleCategory = "ANIMATION"
	LAWS_AdGetV2DataAudienceArticleCategory          AdGetV2DataAudienceArticleCategory = "LAWS"
)

// All allowed values of AdGetV2DataAudienceArticleCategory enum
var AllowedAdGetV2DataAudienceArticleCategoryEnumValues = []AdGetV2DataAudienceArticleCategory{
	"MILITARY",
	"INTERNATIONAL",
	"DESIGN",
	"EMOTION",
	"TIPS",
	"RUMOR_CRACKER",
	"HOME",
	"EDUCATION",
	"REGIMEN",
	"WEIGHT_LOSING",
	"GRADUATES",
	"FINANCE",
	"CONSTELLATION",
	"SPORTS",
	"DIGITAL",
	"BUSINESS",
	"PETS",
	"AMUSEMENT",
	"STORIES",
	"HISTORY",
	"SCIENCE",
	"COMICS",
	"TRAVEL",
	"TECHNOLOGY",
	"CULTURE",
	"PARENTING",
	"ESSAY",
	"CARS",
	"FREAK",
	"PHOTOGRAPHY",
	"GOVERNMENT",
	"ENTERTAINMENT",
	"GAMES",
	"EXPLORE",
	"ESTATE",
	"SOCIETY",
	"GOURMET",
	"FASHION",
	"LOCAL",
	"HEALTH",
	"COLLECTION",
	"WORKPLACE",
	"MOVIE",
	"ANIMATION",
	"LAWS",
}

// NewAdGetV2DataAudienceArticleCategoryFromValue returns a pointer to a valid AdGetV2DataAudienceArticleCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAudienceArticleCategoryFromValue(v string) (*AdGetV2DataAudienceArticleCategory, error) {
	ev := AdGetV2DataAudienceArticleCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAudienceArticleCategory: valid values are %v", v, AllowedAdGetV2DataAudienceArticleCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAudienceArticleCategory) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAudienceArticleCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_audience_article_category value
func (v AdGetV2DataAudienceArticleCategory) Ptr() *AdGetV2DataAudienceArticleCategory {
	return &v
}
