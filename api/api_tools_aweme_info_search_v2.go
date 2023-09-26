/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
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

// ToolsAwemeInfoSearchV2ApiService ToolsAwemeInfoSearchV2Api service
type ToolsAwemeInfoSearchV2ApiService service

type ApiOpenApi2ToolsAwemeInfoSearchGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsAwemeInfoSearchV2ApiService
	advertiserId *int64
	queryWord    *string
	behaviors    *[]*ToolsAwemeInfoSearchV2Behaviors
}

func (r *ApiOpenApi2ToolsAwemeInfoSearchGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsAwemeInfoSearchGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsAwemeInfoSearchGetRequest) QueryWord(queryWord string) *ApiOpenApi2ToolsAwemeInfoSearchGetRequest {
	r.queryWord = &queryWord
	return r
}

func (r *ApiOpenApi2ToolsAwemeInfoSearchGetRequest) Behaviors(behaviors []*ToolsAwemeInfoSearchV2Behaviors) *ApiOpenApi2ToolsAwemeInfoSearchGetRequest {
	r.behaviors = &behaviors
	return r
}

func (r *ApiOpenApi2ToolsAwemeInfoSearchGetRequest) Execute() (*ToolsAwemeInfoSearchV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsAwemeInfoSearchGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAwemeInfoSearchGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAwemeInfoSearchGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsAwemeInfoSearchGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAwemeInfoSearchGet Method for OpenApi2ToolsAwemeInfoSearchGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAwemeInfoSearchGetRequest
*/
func (a *ToolsAwemeInfoSearchV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsAwemeInfoSearchGetRequest {
	return &ApiOpenApi2ToolsAwemeInfoSearchGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAwemeInfoSearchV2Response
func (a *ToolsAwemeInfoSearchV2ApiService) getExecute(r *ApiOpenApi2ToolsAwemeInfoSearchGetRequest) (*ToolsAwemeInfoSearchV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsAwemeInfoSearchV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/aweme_info_search/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}
	if r.queryWord == nil {
		return localVarReturnValue, nil, ReportError("queryWord is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "query_word", r.queryWord)
	if r.behaviors != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "behaviors", r.behaviors)
	}
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
