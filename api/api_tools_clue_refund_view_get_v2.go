/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
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

// ToolsClueRefundViewGetV2ApiService ToolsClueRefundViewGetV2Api service
type ToolsClueRefundViewGetV2ApiService service

type ApiOpenApi2ToolsClueRefundViewGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsClueRefundViewGetV2ApiService
	advertiserId *int64
	clueId       *int64
}

// 广告主id
func (r *ApiOpenApi2ToolsClueRefundViewGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsClueRefundViewGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 线索id
func (r *ApiOpenApi2ToolsClueRefundViewGetGetRequest) ClueId(clueId int64) *ApiOpenApi2ToolsClueRefundViewGetGetRequest {
	r.clueId = &clueId
	return r
}

func (r *ApiOpenApi2ToolsClueRefundViewGetGetRequest) Execute() (*ToolsClueRefundViewGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsClueRefundViewGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsClueRefundViewGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsClueRefundViewGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsClueRefundViewGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsClueRefundViewGetGet Method for OpenApi2ToolsClueRefundViewGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsClueRefundViewGetGetRequest
*/
func (a *ToolsClueRefundViewGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsClueRefundViewGetGetRequest {
	return &ApiOpenApi2ToolsClueRefundViewGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsClueRefundViewGetV2Response
func (a *ToolsClueRefundViewGetV2ApiService) getExecute(r *ApiOpenApi2ToolsClueRefundViewGetGetRequest) (*ToolsClueRefundViewGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsClueRefundViewGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/clue/refund_view/get/"

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

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "clue_id", r.clueId)
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
