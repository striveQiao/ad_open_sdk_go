/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.6
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

// PromotionAutoGenerateConfigGetV30ApiService PromotionAutoGenerateConfigGetV30Api service
type PromotionAutoGenerateConfigGetV30ApiService service

type ApiOpenApiV30PromotionAutoGenerateConfigGetGetRequest struct {
	ctx          context.Context
	ApiService   *PromotionAutoGenerateConfigGetV30ApiService
	advertiserId *int64
	configId     *int64
}

// 广告主id
func (r *ApiOpenApiV30PromotionAutoGenerateConfigGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30PromotionAutoGenerateConfigGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 配置id
func (r *ApiOpenApiV30PromotionAutoGenerateConfigGetGetRequest) ConfigId(configId int64) *ApiOpenApiV30PromotionAutoGenerateConfigGetGetRequest {
	r.configId = &configId
	return r
}

func (r *ApiOpenApiV30PromotionAutoGenerateConfigGetGetRequest) Execute() (*PromotionAutoGenerateConfigGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30PromotionAutoGenerateConfigGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30PromotionAutoGenerateConfigGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30PromotionAutoGenerateConfigGetGetRequest) WithLog(enable bool) *ApiOpenApiV30PromotionAutoGenerateConfigGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30PromotionAutoGenerateConfigGetGet Method for OpenApiV30PromotionAutoGenerateConfigGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30PromotionAutoGenerateConfigGetGetRequest
*/
func (a *PromotionAutoGenerateConfigGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30PromotionAutoGenerateConfigGetGetRequest {
	return &ApiOpenApiV30PromotionAutoGenerateConfigGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PromotionAutoGenerateConfigGetV30Response
func (a *PromotionAutoGenerateConfigGetV30ApiService) getExecute(r *ApiOpenApiV30PromotionAutoGenerateConfigGetGetRequest) (*PromotionAutoGenerateConfigGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *PromotionAutoGenerateConfigGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/promotion/auto_generate_config/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.configId == nil {
		return localVarReturnValue, nil, ReportError("configId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "config_id", r.configId)
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
