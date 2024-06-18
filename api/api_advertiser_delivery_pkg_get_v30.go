/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
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

// AdvertiserDeliveryPkgGetV30ApiService AdvertiserDeliveryPkgGetV30Api service
type AdvertiserDeliveryPkgGetV30ApiService service

type ApiOpenApiV30AdvertiserDeliveryPkgGetGetRequest struct {
	ctx          context.Context
	ApiService   *AdvertiserDeliveryPkgGetV30ApiService
	advertiserId *int64
	pkgId        *int64
}

// 广告主账户ID
func (r *ApiOpenApiV30AdvertiserDeliveryPkgGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30AdvertiserDeliveryPkgGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 推广产品组id，是推广产品的组标识
func (r *ApiOpenApiV30AdvertiserDeliveryPkgGetGetRequest) PkgId(pkgId int64) *ApiOpenApiV30AdvertiserDeliveryPkgGetGetRequest {
	r.pkgId = &pkgId
	return r
}

func (r *ApiOpenApiV30AdvertiserDeliveryPkgGetGetRequest) Execute() (*AdvertiserDeliveryPkgGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30AdvertiserDeliveryPkgGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30AdvertiserDeliveryPkgGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30AdvertiserDeliveryPkgGetGetRequest) WithLog(enable bool) *ApiOpenApiV30AdvertiserDeliveryPkgGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30AdvertiserDeliveryPkgGetGet Method for OpenApiV30AdvertiserDeliveryPkgGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30AdvertiserDeliveryPkgGetGetRequest
*/
func (a *AdvertiserDeliveryPkgGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30AdvertiserDeliveryPkgGetGetRequest {
	return &ApiOpenApiV30AdvertiserDeliveryPkgGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AdvertiserDeliveryPkgGetV30Response
func (a *AdvertiserDeliveryPkgGetV30ApiService) getExecute(r *ApiOpenApiV30AdvertiserDeliveryPkgGetGetRequest) (*AdvertiserDeliveryPkgGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AdvertiserDeliveryPkgGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/advertiser/delivery_pkg/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.pkgId == nil {
		return localVarReturnValue, nil, ReportError("pkgId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "pkg_id", r.pkgId)
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
