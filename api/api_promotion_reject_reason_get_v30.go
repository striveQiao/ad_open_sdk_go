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

// PromotionRejectReasonGetV30ApiService PromotionRejectReasonGetV30Api service
type PromotionRejectReasonGetV30ApiService service

type ApiOpenApiV30PromotionRejectReasonGetGetRequest struct {
	ctx          context.Context
	ApiService   *PromotionRejectReasonGetV30ApiService
	advertiserId *int64
	promotionIds *[]int64
	deliveryMode *PromotionRejectReasonGetV30DeliveryMode
}

// 广告主ID
func (r *ApiOpenApiV30PromotionRejectReasonGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30PromotionRejectReasonGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 广告id列表
func (r *ApiOpenApiV30PromotionRejectReasonGetGetRequest) PromotionIds(promotionIds []int64) *ApiOpenApiV30PromotionRejectReasonGetGetRequest {
	r.promotionIds = &promotionIds
	return r
}

// 投放模式，允许值： MANUAL手动投放（默认值） PROCEDURAL自动投放
func (r *ApiOpenApiV30PromotionRejectReasonGetGetRequest) DeliveryMode(deliveryMode PromotionRejectReasonGetV30DeliveryMode) *ApiOpenApiV30PromotionRejectReasonGetGetRequest {
	r.deliveryMode = &deliveryMode
	return r
}

func (r *ApiOpenApiV30PromotionRejectReasonGetGetRequest) Execute() (*PromotionRejectReasonGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30PromotionRejectReasonGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30PromotionRejectReasonGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30PromotionRejectReasonGetGetRequest) WithLog(enable bool) *ApiOpenApiV30PromotionRejectReasonGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30PromotionRejectReasonGetGet Method for OpenApiV30PromotionRejectReasonGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30PromotionRejectReasonGetGetRequest
*/
func (a *PromotionRejectReasonGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30PromotionRejectReasonGetGetRequest {
	return &ApiOpenApiV30PromotionRejectReasonGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PromotionRejectReasonGetV30Response
func (a *PromotionRejectReasonGetV30ApiService) getExecute(r *ApiOpenApiV30PromotionRejectReasonGetGetRequest) (*PromotionRejectReasonGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *PromotionRejectReasonGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/promotion/reject_reason/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}
	if r.promotionIds == nil {
		return localVarReturnValue, nil, ReportError("promotionIds is required and must be specified")
	}
	if len(*r.promotionIds) < 1 {
		return localVarReturnValue, nil, ReportError("promotionIds must have at least 1 elements")
	}
	if len(*r.promotionIds) > 10 {
		return localVarReturnValue, nil, ReportError("promotionIds must have less than 10 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "promotion_ids", r.promotionIds)
	if r.deliveryMode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delivery_mode", r.deliveryMode)
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
