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

// ToolsClueCallVirtualNumberGetV2ApiService ToolsClueCallVirtualNumberGetV2Api service
type ToolsClueCallVirtualNumberGetV2ApiService service

type ApiOpenApi2ToolsClueCallVirtualNumberGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsClueCallVirtualNumberGetV2ApiService
	advertiserId *int64
	clueId       *int64
	callerNumber *string
	calleeNumber *string
}

// 广告主id
func (r *ApiOpenApi2ToolsClueCallVirtualNumberGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsClueCallVirtualNumberGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 线索id
func (r *ApiOpenApi2ToolsClueCallVirtualNumberGetGetRequest) ClueId(clueId int64) *ApiOpenApi2ToolsClueCallVirtualNumberGetGetRequest {
	r.clueId = &clueId
	return r
}

// 主叫号码，必须为11位手机号码，否则呼叫失败
func (r *ApiOpenApi2ToolsClueCallVirtualNumberGetGetRequest) CallerNumber(callerNumber string) *ApiOpenApi2ToolsClueCallVirtualNumberGetGetRequest {
	r.callerNumber = &callerNumber
	return r
}

// 被叫号码，即线索号码
func (r *ApiOpenApi2ToolsClueCallVirtualNumberGetGetRequest) CalleeNumber(calleeNumber string) *ApiOpenApi2ToolsClueCallVirtualNumberGetGetRequest {
	r.calleeNumber = &calleeNumber
	return r
}

func (r *ApiOpenApi2ToolsClueCallVirtualNumberGetGetRequest) Execute() (*ToolsClueCallVirtualNumberGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsClueCallVirtualNumberGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsClueCallVirtualNumberGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsClueCallVirtualNumberGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsClueCallVirtualNumberGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsClueCallVirtualNumberGetGet Method for OpenApi2ToolsClueCallVirtualNumberGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsClueCallVirtualNumberGetGetRequest
*/
func (a *ToolsClueCallVirtualNumberGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsClueCallVirtualNumberGetGetRequest {
	return &ApiOpenApi2ToolsClueCallVirtualNumberGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsClueCallVirtualNumberGetV2Response
func (a *ToolsClueCallVirtualNumberGetV2ApiService) getExecute(r *ApiOpenApi2ToolsClueCallVirtualNumberGetGetRequest) (*ToolsClueCallVirtualNumberGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsClueCallVirtualNumberGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/clue/call_virtual_number/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}
	if r.clueId == nil {
		return localVarReturnValue, nil, ReportError("clueId is required and must be specified")
	}
	if *r.clueId < 1 {
		return localVarReturnValue, nil, ReportError("clueId must be greater than 1")
	}
	if r.callerNumber == nil {
		return localVarReturnValue, nil, ReportError("callerNumber is required and must be specified")
	}
	if strlen(*r.callerNumber) < 11 {
		return localVarReturnValue, nil, ReportError("callerNumber must have at least 11 elements")
	}
	if strlen(*r.callerNumber) > 11 {
		return localVarReturnValue, nil, ReportError("callerNumber must have less than 11 elements")
	}
	if r.calleeNumber == nil {
		return localVarReturnValue, nil, ReportError("calleeNumber is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "clue_id", r.clueId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "caller_number", r.callerNumber)
	parameterAddToHeaderOrQuery(localVarQueryParams, "callee_number", r.calleeNumber)
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
