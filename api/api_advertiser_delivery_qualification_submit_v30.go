/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
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

// AdvertiserDeliveryQualificationSubmitV30ApiService AdvertiserDeliveryQualificationSubmitV30Api service
type AdvertiserDeliveryQualificationSubmitV30ApiService service

type ApiOpenApiV30AdvertiserDeliveryQualificationSubmitPostRequest struct {
	ctx                                             context.Context
	ApiService                                      *AdvertiserDeliveryQualificationSubmitV30ApiService
	advertiserDeliveryQualificationSubmitV30Request *AdvertiserDeliveryQualificationSubmitV30Request
}

func (r *ApiOpenApiV30AdvertiserDeliveryQualificationSubmitPostRequest) AdvertiserDeliveryQualificationSubmitV30Request(advertiserDeliveryQualificationSubmitV30Request AdvertiserDeliveryQualificationSubmitV30Request) *ApiOpenApiV30AdvertiserDeliveryQualificationSubmitPostRequest {
	r.advertiserDeliveryQualificationSubmitV30Request = &advertiserDeliveryQualificationSubmitV30Request
	return r
}

func (r *ApiOpenApiV30AdvertiserDeliveryQualificationSubmitPostRequest) Execute() (*AdvertiserDeliveryQualificationSubmitV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30AdvertiserDeliveryQualificationSubmitPostRequest) AccessToken(accessToken string) *ApiOpenApiV30AdvertiserDeliveryQualificationSubmitPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30AdvertiserDeliveryQualificationSubmitPostRequest) WithLog(enable bool) *ApiOpenApiV30AdvertiserDeliveryQualificationSubmitPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30AdvertiserDeliveryQualificationSubmitPost Method for OpenApiV30AdvertiserDeliveryQualificationSubmitPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30AdvertiserDeliveryQualificationSubmitPostRequest
*/
func (a *AdvertiserDeliveryQualificationSubmitV30ApiService) Post(ctx context.Context) *ApiOpenApiV30AdvertiserDeliveryQualificationSubmitPostRequest {
	return &ApiOpenApiV30AdvertiserDeliveryQualificationSubmitPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AdvertiserDeliveryQualificationSubmitV30Response
func (a *AdvertiserDeliveryQualificationSubmitV30ApiService) postExecute(r *ApiOpenApiV30AdvertiserDeliveryQualificationSubmitPostRequest) (*AdvertiserDeliveryQualificationSubmitV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *AdvertiserDeliveryQualificationSubmitV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/advertiser/delivery_qualification/submit/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.advertiserDeliveryQualificationSubmitV30Request
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
