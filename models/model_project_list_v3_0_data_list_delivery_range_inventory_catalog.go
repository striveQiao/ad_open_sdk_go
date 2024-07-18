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

// ProjectListV30DataListDeliveryRangeInventoryCatalog
type ProjectListV30DataListDeliveryRangeInventoryCatalog string

// List of project_list_v3.0_data_list_delivery_range_inventory_catalog
const (
	MANUAL_ProjectListV30DataListDeliveryRangeInventoryCatalog          ProjectListV30DataListDeliveryRangeInventoryCatalog = "MANUAL"
	SCENE_ProjectListV30DataListDeliveryRangeInventoryCatalog           ProjectListV30DataListDeliveryRangeInventoryCatalog = "SCENE"
	SMART_ProjectListV30DataListDeliveryRangeInventoryCatalog           ProjectListV30DataListDeliveryRangeInventoryCatalog = "SMART"
	UNIVERSAL_ProjectListV30DataListDeliveryRangeInventoryCatalog       ProjectListV30DataListDeliveryRangeInventoryCatalog = "UNIVERSAL"
	UNIVERSAL_SMART_ProjectListV30DataListDeliveryRangeInventoryCatalog ProjectListV30DataListDeliveryRangeInventoryCatalog = "UNIVERSAL_SMART"
)

// All allowed values of ProjectListV30DataListDeliveryRangeInventoryCatalog enum
var AllowedProjectListV30DataListDeliveryRangeInventoryCatalogEnumValues = []ProjectListV30DataListDeliveryRangeInventoryCatalog{
	"MANUAL",
	"SCENE",
	"SMART",
	"UNIVERSAL",
	"UNIVERSAL_SMART",
}

// NewProjectListV30DataListDeliveryRangeInventoryCatalogFromValue returns a pointer to a valid ProjectListV30DataListDeliveryRangeInventoryCatalog
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListDeliveryRangeInventoryCatalogFromValue(v string) (*ProjectListV30DataListDeliveryRangeInventoryCatalog, error) {
	ev := ProjectListV30DataListDeliveryRangeInventoryCatalog(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListDeliveryRangeInventoryCatalog: valid values are %v", v, AllowedProjectListV30DataListDeliveryRangeInventoryCatalogEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListDeliveryRangeInventoryCatalog) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListDeliveryRangeInventoryCatalogEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_delivery_range_inventory_catalog value
func (v ProjectListV30DataListDeliveryRangeInventoryCatalog) Ptr() *ProjectListV30DataListDeliveryRangeInventoryCatalog {
	return &v
}
