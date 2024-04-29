/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
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

// ToolsBidsSuggestV30ApiService ToolsBidsSuggestV30Api service
type ToolsBidsSuggestV30ApiService service

type ApiOpenApiV30ToolsBidsSuggestGetRequest struct {
	ctx                context.Context
	ApiService         *ToolsBidsSuggestV30ApiService
	advertiserId       *int64
	pricing            *ToolsBidsSuggestV30Pricing
	externalAction     *ToolsBidsSuggestV30ExternalAction
	projectId          *int64
	deepExternalAction *ToolsBidsSuggestV30DeepExternalAction
	deepBidType        *ToolsBidsSuggestV30DeepBidType
}

func (r *ApiOpenApiV30ToolsBidsSuggestGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ToolsBidsSuggestGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30ToolsBidsSuggestGetRequest) Pricing(pricing ToolsBidsSuggestV30Pricing) *ApiOpenApiV30ToolsBidsSuggestGetRequest {
	r.pricing = &pricing
	return r
}

func (r *ApiOpenApiV30ToolsBidsSuggestGetRequest) ExternalAction(externalAction ToolsBidsSuggestV30ExternalAction) *ApiOpenApiV30ToolsBidsSuggestGetRequest {
	r.externalAction = &externalAction
	return r
}

func (r *ApiOpenApiV30ToolsBidsSuggestGetRequest) ProjectId(projectId int64) *ApiOpenApiV30ToolsBidsSuggestGetRequest {
	r.projectId = &projectId
	return r
}

func (r *ApiOpenApiV30ToolsBidsSuggestGetRequest) DeepExternalAction(deepExternalAction ToolsBidsSuggestV30DeepExternalAction) *ApiOpenApiV30ToolsBidsSuggestGetRequest {
	r.deepExternalAction = &deepExternalAction
	return r
}

func (r *ApiOpenApiV30ToolsBidsSuggestGetRequest) DeepBidType(deepBidType ToolsBidsSuggestV30DeepBidType) *ApiOpenApiV30ToolsBidsSuggestGetRequest {
	r.deepBidType = &deepBidType
	return r
}

func (r *ApiOpenApiV30ToolsBidsSuggestGetRequest) Execute() (*ToolsBidsSuggestV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ToolsBidsSuggestGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsBidsSuggestGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsBidsSuggestGetRequest) WithLog(enable bool) *ApiOpenApiV30ToolsBidsSuggestGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsBidsSuggestGet Method for OpenApiV30ToolsBidsSuggestGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsBidsSuggestGetRequest
*/
func (a *ToolsBidsSuggestV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ToolsBidsSuggestGetRequest {
	return &ApiOpenApiV30ToolsBidsSuggestGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsBidsSuggestV30Response
func (a *ToolsBidsSuggestV30ApiService) getExecute(r *ApiOpenApiV30ToolsBidsSuggestGetRequest) (*ToolsBidsSuggestV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsBidsSuggestV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/bids/suggest/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.pricing == nil {
		return localVarReturnValue, nil, ReportError("pricing is required and must be specified")
	}
	if r.externalAction == nil {
		return localVarReturnValue, nil, ReportError("externalAction is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.projectId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "project_id", r.projectId)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "pricing", r.pricing)
	parameterAddToHeaderOrQuery(localVarQueryParams, "external_action", r.externalAction)
	if r.deepExternalAction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deep_external_action", r.deepExternalAction)
	}
	if r.deepBidType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deep_bid_type", r.deepBidType)
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
