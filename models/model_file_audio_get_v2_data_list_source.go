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

// FileAudioGetV2DataListSource
type FileAudioGetV2DataListSource string

// List of file_audio_get_v2_data_list_source
const (
	AD_SITE_FileAudioGetV2DataListSource           FileAudioGetV2DataListSource = "AD_SITE"
	OPEN_API_FileAudioGetV2DataListSource          FileAudioGetV2DataListSource = "OPEN_API"
	TTS_TEXT_TO_AUDIO_FileAudioGetV2DataListSource FileAudioGetV2DataListSource = "TTS_TEXT_TO_AUDIO"
)

// All allowed values of FileAudioGetV2DataListSource enum
var AllowedFileAudioGetV2DataListSourceEnumValues = []FileAudioGetV2DataListSource{
	"AD_SITE",
	"OPEN_API",
	"TTS_TEXT_TO_AUDIO",
}

// NewFileAudioGetV2DataListSourceFromValue returns a pointer to a valid FileAudioGetV2DataListSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileAudioGetV2DataListSourceFromValue(v string) (*FileAudioGetV2DataListSource, error) {
	ev := FileAudioGetV2DataListSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileAudioGetV2DataListSource: valid values are %v", v, AllowedFileAudioGetV2DataListSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileAudioGetV2DataListSource) IsValid() bool {
	for _, existing := range AllowedFileAudioGetV2DataListSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_audio_get_v2_data_list_source value
func (v FileAudioGetV2DataListSource) Ptr() *FileAudioGetV2DataListSource {
	return &v
}
