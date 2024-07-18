/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// ToolsDiagnosisSuggestionGetV2ApiService ToolsDiagnosisSuggestionGetV2Api service
type ToolsDiagnosisSuggestionGetV2ApiService service

type ApiOpenApi2ToolsDiagnosisSuggestionGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsDiagnosisSuggestionGetV2ApiService
	adIds        *[]int64
	scenes       *[]*ToolsDiagnosisSuggestionGetV2Scenes
	advertiserId *int64
}

func (r *ApiOpenApi2ToolsDiagnosisSuggestionGetGetRequest) AdIds(adIds []int64) *ApiOpenApi2ToolsDiagnosisSuggestionGetGetRequest {
	r.adIds = &adIds
	return r
}

func (r *ApiOpenApi2ToolsDiagnosisSuggestionGetGetRequest) Scenes(scenes []*ToolsDiagnosisSuggestionGetV2Scenes) *ApiOpenApi2ToolsDiagnosisSuggestionGetGetRequest {
	r.scenes = &scenes
	return r
}

func (r *ApiOpenApi2ToolsDiagnosisSuggestionGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsDiagnosisSuggestionGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsDiagnosisSuggestionGetGetRequest) Execute() (*ToolsDiagnosisSuggestionGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsDiagnosisSuggestionGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsDiagnosisSuggestionGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsDiagnosisSuggestionGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsDiagnosisSuggestionGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsDiagnosisSuggestionGetGet Method for OpenApi2ToolsDiagnosisSuggestionGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsDiagnosisSuggestionGetGetRequest
*/
func (a *ToolsDiagnosisSuggestionGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsDiagnosisSuggestionGetGetRequest {
	return &ApiOpenApi2ToolsDiagnosisSuggestionGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsDiagnosisSuggestionGetV2Response
func (a *ToolsDiagnosisSuggestionGetV2ApiService) getExecute(r *ApiOpenApi2ToolsDiagnosisSuggestionGetGetRequest) (*ToolsDiagnosisSuggestionGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsDiagnosisSuggestionGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/diagnosis/suggestion/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.adIds == nil {
		return localVarReturnValue, nil, ReportError("adIds is required and must be specified")
	}
	if len(*r.adIds) < 1 {
		return localVarReturnValue, nil, ReportError("adIds must have at least 1 elements")
	}
	if len(*r.adIds) > 100 {
		return localVarReturnValue, nil, ReportError("adIds must have less than 100 elements")
	}
	if r.scenes == nil {
		return localVarReturnValue, nil, ReportError("scenes is required and must be specified")
	}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}
	if *r.advertiserId > 9223372036854775807 {
		return localVarReturnValue, nil, ReportError("advertiserId must be less than 9223372036854775807")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "ad_ids", r.adIds)
	parameterAddToHeaderOrQuery(localVarQueryParams, "scenes", r.scenes)
	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
