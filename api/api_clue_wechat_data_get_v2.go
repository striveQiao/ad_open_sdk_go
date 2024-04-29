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

// ClueWechatDataGetV2ApiService ClueWechatDataGetV2Api service
type ClueWechatDataGetV2ApiService service

type ApiOpenApi2ClueWechatDataGetGetRequest struct {
	ctx          context.Context
	ApiService   *ClueWechatDataGetV2ApiService
	advertiserId *int64
	unionId      *string
	state        *string
}

// 广告账户id
func (r *ApiOpenApi2ClueWechatDataGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ClueWechatDataGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 微信union_id
func (r *ApiOpenApi2ClueWechatDataGetGetRequest) UnionId(unionId string) *ApiOpenApi2ClueWechatDataGetGetRequest {
	r.unionId = &unionId
	return r
}

// 企业微信「联系我」state参数
func (r *ApiOpenApi2ClueWechatDataGetGetRequest) State(state string) *ApiOpenApi2ClueWechatDataGetGetRequest {
	r.state = &state
	return r
}

func (r *ApiOpenApi2ClueWechatDataGetGetRequest) Execute() (*ClueWechatDataGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ClueWechatDataGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ClueWechatDataGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ClueWechatDataGetGetRequest) WithLog(enable bool) *ApiOpenApi2ClueWechatDataGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ClueWechatDataGetGet Method for OpenApi2ClueWechatDataGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ClueWechatDataGetGetRequest
*/
func (a *ClueWechatDataGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ClueWechatDataGetGetRequest {
	return &ApiOpenApi2ClueWechatDataGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ClueWechatDataGetV2Response
func (a *ClueWechatDataGetV2ApiService) getExecute(r *ApiOpenApi2ClueWechatDataGetGetRequest) (*ClueWechatDataGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ClueWechatDataGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/clue/wechat_data/get/"

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
	if r.unionId == nil {
		return localVarReturnValue, nil, ReportError("unionId is required and must be specified")
	}
	if strlen(*r.unionId) < 1 {
		return localVarReturnValue, nil, ReportError("unionId must have at least 1 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "union_id", r.unionId)
	if r.state != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "state", r.state)
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
